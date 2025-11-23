package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	DB *sql.DB
}

func NewStore() *Store {
	dsn := "root:Rahul@123@tcp(127.0.0.1:3306)/studentdb?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Database ping error:", err)
	}

	return &Store{DB: db}
}

func (s *Store) createStudent(st Student) (*Student, error) {
	query := `
        INSERT INTO students (first_name, last_name, email, age, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?, ?)
    `
	now := time.Now()
	result, err := s.DB.Exec(query, st.FirstName, st.LastName, st.Email, st.Age, now, now)
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()
	st.ID = int(id)
	st.CreatedAt = now
	st.UpdatedAt = now
	return &st, nil
}

func (s *Store) listStudents() ([]Student, error) {
	rows, err := s.DB.Query(`SELECT id, first_name, last_name, email, age, created_at, updated_at FROM students`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var st Student
		if err := rows.Scan(&st.ID, &st.FirstName, &st.LastName, &st.Email, &st.Age, &st.CreatedAt, &st.UpdatedAt); err != nil {
			return nil, err
		}
		students = append(students, st)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (s *Store) getStudent(id int) (*Student, error) {
	query := `SELECT id, first_name, last_name, email, age, created_at, updated_at FROM students WHERE id = ?`

	var st Student
	err := s.DB.QueryRow(query, id).Scan(
		&st.ID, &st.FirstName, &st.LastName,
		&st.Email, &st.Age, &st.CreatedAt, &st.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &st, nil
}

func (s *Store) updateStudent(id int, st Student) (*Student, error) {
	query := `UPDATE students SET first_name=?, last_name=?, email=?, age=?, updated_at=? WHERE id=?`
	now := time.Now()
	_, err := s.DB.Exec(query, st.FirstName, st.LastName, st.Email, st.Age, now, id)

	if err != nil {
		return nil, err
	}

	return s.getStudent(id)
}

func (s *Store) deleteStudent(id int) error {
	_, err := s.DB.Exec(`DELETE FROM students WHERE id = ?`, id)
	return err
}
