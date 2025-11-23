package main

import (
	"sync"
	"time"
)

type Store struct {
	sync.Mutex
	data map[int]*Student
	next int
}

func NewStore() *Store {
	return &Store{
		data: make(map[int]*Student),
		next: 1,
	}
}

func (s *Store) createStudent(st Student) *Student {
	s.Lock()
	defer s.Unlock()

	st.ID = s.next
	st.CreatedAt = time.Now()
	st.UpdatedAt = st.CreatedAt

	s.data[st.ID] = &st
	s.next++

	return &st
}

func (s *Store) listStudents() []*Student {
	s.Lock()
	defer s.Unlock()

	students := make([]*Student, 0, len(s.data))
	for _, st := range s.data {
		students = append(students, st)
	}
	return students
}

func (s *Store) getStudent(id int) (*Student, bool) {
	s.Lock()
	defer s.Unlock()

	st, ok := s.data[id]
	return st, ok
}

func (s *Store) updateStudent(id int, update Student) (*Student, bool) {
	s.Lock()
	defer s.Unlock()

	st, ok := s.data[id]
	if !ok {
		return nil, false
	}

	if update.FirstName != "" {
		st.FirstName = update.FirstName
	}
	if update.LastName != "" {
		st.LastName = update.LastName
	}
	if update.Email != "" {
		st.Email = update.Email
	}
	if update.Age != 0 {
		st.Age = update.Age
	}

	st.UpdatedAt = time.Now()

	return st, true
}

func (s *Store) deleteStudent(id int) bool {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.data[id]; !ok {
		return false
	}

	delete(s.data, id)
	return true
}
