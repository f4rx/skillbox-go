package storage

import (
	"fmt"
	"sort"
	"strings"
	"28/pkg/misc"
	"28/pkg/student"
)

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
