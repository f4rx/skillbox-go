package core

import (
	"28/pkg/student"
	"28/pkg/storage"
	"28/pkg/misc"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var ErrorParseInputLine = fmt.Errorf("ошибка парсинга входящей строки")

func Run(){
	misc.NewLogger()

	// var sm storage.StudentMap = make(map[string]*student.Student)
	sm := storage.NewStorage()
	readStudents(sm)

	fmt.Println()
	fmt.Println()
	fmt.Println(sm)
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

func readStudents(sm storage.Storage) {
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
