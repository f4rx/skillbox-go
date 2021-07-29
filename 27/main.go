package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	logger "github.com/f4rx/logger-zap-wrapper"
	"go.uber.org/zap"
)

var slog *zap.SugaredLogger //nolint:gochecknoglobals

var ErrorParseInputLine = fmt.Errorf("ошибка парсинга входящей строки")

type Student struct {
	Name  string
	Age   int
	grade int
}

func (s Student) String() string {
	return fmt.Sprintf("Студент: %s, возвраст: %d, оценка: %d", s.Name, s.Age, s.grade)
}

func newStudent(name string, age int, grade int) *Student {
	student := Student{name, age, grade}
	return &student
}

type studentMap map[string]*Student

func (sm studentMap) put(student *Student) error {
	if _, ok := sm[student.Name]; ok {
		return fmt.Errorf("такой студент уже существует в словаре")
	}
	sm[student.Name] = student
	return nil
}

func (sm studentMap) get(name string) (*Student, error) {
	if student, ok := sm[name]; ok {
		return student, nil
	}
	return nil, fmt.Errorf("такого пользователя нет в словаре")
}

func (sm studentMap) String() string {
	keys := make([]string, 0, len(sm))
	for k := range sm {
		slog.Debug("key:", k)
		keys = append(keys, k)
	}
	slog.Debug(keys)

	sort.Strings(keys)

	var sb strings.Builder

	slog.Debug(keys)

	for _, key := range keys {
		s, err := sm.get(key)
		if err != nil {
			slog.Panic("Ошибка при работе с ключами")
		}
		sb.WriteString(s.String())
		sb.WriteString("\n")
	}
	return sb.String()
}

func parseStringToStudent(line string) (*Student, error) {
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

	student := newStudent(name, age, grade)
	return student, nil
}

func readStudents(sm studentMap) {
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
			slog.Warn("'", line, "' ", err)
			continue
		}
		err = sm.put(s)
		if err != nil {
			slog.Warn("'", s.Name, "' ", err)
			continue
		}
	}
}

func main() {
	slog = logger.NewSugaredLogger()
	slog.Sync() //nolint:errcheck

	var sm studentMap = make(map[string]*Student)
	readStudents(sm)

	fmt.Println()
	fmt.Println()
	fmt.Println(sm)
}
