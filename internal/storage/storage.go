package storage

import "github.com/HakashiKatake/Students-Rest-Api-Go/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	// GetList() ([]types.Student, error)
	// UpdateStudent(id int64, name string, email string, age int) error
	// DeleteStudent(id int64) error
}
