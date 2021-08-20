package student

import (
	"fmt"
	"strconv"
	"strings"
)

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

func ParseStringToStudent(line string) (*Student, error) {
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
