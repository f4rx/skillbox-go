package student

import (
	"fmt"
)

type Student struct {
	name  string
	age   int
	grade int
}

func (s Student) String() string {
	return fmt.Sprintf("Студент: %s, возвраст: %d, оценка: %d", s.name, s.age, s.grade)
}

func (s Student) GetName() string {
	return s.name
}

func NewStudent(name string, age int, grade int) *Student {
	student := Student{name, age, grade}
	return &student
}
