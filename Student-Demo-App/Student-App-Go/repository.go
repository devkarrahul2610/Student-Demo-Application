package main

type StudentRepository interface {
	ListStudents() ([]Student, error)
	CreateStudent(st Student) (*Student, error)
	GetStudent(id int) (*Student, error)
	UpdateStudent(id int, st Student) (*Student, error)
	DeleteStudent(id int) error
}
