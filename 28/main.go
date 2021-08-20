package main

import (
	"fmt"

	"28/pkg/misc"
	"28/pkg/storage"
	"28/pkg/student"
)

func main() {
	misc.NewLogger()

	var sm storage.StudentMap = make(map[string]*student.Student)
	storage.ReadStudents(sm)

	fmt.Println()
	fmt.Println()
	fmt.Println(sm)
}
