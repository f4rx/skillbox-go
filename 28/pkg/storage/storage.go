package storage

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"28/pkg/misc"
	"28/pkg/student"
)

var ErrorParseInputLine = fmt.Errorf("ошибка парсинга входящей строки")

type Storage interface {
	Put(student *student.Student) error
	Get(name string) (*student.Student, error)
}

type StudentMap map[string]*student.Student

func (sm StudentMap) Put(student *student.Student) error {
	if _, ok := sm[student.GetName()]; ok {
		return fmt.Errorf("такой студент уже существует в словаре")
	}
	sm[student.GetName()] = student
	return nil
}

func (sm StudentMap) Get(name string) (*student.Student, error) {
	if student, ok := sm[name]; ok {
		return student, nil
	}
	return nil, fmt.Errorf("такого пользователя нет в словаре")
}

func (sm StudentMap) String() string {
	keys := make([]string, 0, len(sm))
	for k := range sm {
		misc.Slog.Debug("key:", k)
		keys = append(keys, k)
	}
	misc.Slog.Debug(keys)

	sort.Strings(keys)

	var sb strings.Builder

	misc.Slog.Debug(keys)

	for _, key := range keys {
		s, err := sm.Get(key)
		if err != nil {
			misc.Slog.Panic("Ошибка при работе с ключами")
		}
		sb.WriteString(s.String())
		sb.WriteString("\n")
	}
	return sb.String()
}

func parseStringToStudent(line string) (*student.Student, error) {
	words := strings.Split(line, " ")
	if len(words) != 3 {
		return nil, ErrorParseInputLine
	}
	name := words[0]
	age, err := strconv.Atoi(words[1])
	if err != nil {
		return nil, ErrorParseInputLine
	}
	grade, err := strconv.Atoi(words[2])
	if err != nil {
		return nil, ErrorParseInputLine
	}

	student := student.NewStudent(name, age, grade)
	return student, nil
}

func ReadStudents(sm StudentMap) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Введите через запятую имя студента, возраст, оценку: ")
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		s, err := parseStringToStudent(line)
		if err != nil {
			misc.Slog.Warn("'", line, "' ", err)
			continue
		}
		err = sm.Put(s)
		if err != nil {
			misc.Slog.Warn("'", s.GetName(), "' ", err)
			continue
		}
	}
}
