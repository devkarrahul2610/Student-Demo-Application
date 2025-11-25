package student

import (
	"database/sql"
	"time"
)

type MySQLStudentRepository struct {
	DB *sql.DB
}

func NewMySQLStudentRepository(db *sql.DB) StudentRepository {
	return &MySQLStudentRepository{DB: db}
}

func (r *MySQLStudentRepository) ListStudents() ([]Student, error) {
	rows, err := r.DB.Query("SELECT id, first_name, last_name, email, age, created_at, updated_at FROM students")
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

	return students, nil
}

func (r *MySQLStudentRepository) CreateStudent(st Student) (*Student, error) {
	now := time.Now()
	result, err := r.DB.Exec(
		"INSERT INTO students (first_name, last_name, email, age, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		st.FirstName, st.LastName, st.Email, st.Age, now, now,
	)
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()
	st.ID = int(id)
	st.CreatedAt = now
	st.UpdatedAt = now

	return &st, nil
}

func (r *MySQLStudentRepository) GetStudent(id int) (*Student, error) {
	var st Student
	err := r.DB.QueryRow(
		"SELECT id, first_name, last_name, email, age, created_at, updated_at FROM students WHERE id=?",
		id,
	).Scan(&st.ID, &st.FirstName, &st.LastName, &st.Email, &st.Age, &st.CreatedAt, &st.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &st, nil
}

func (r *MySQLStudentRepository) UpdateStudent(id int, st Student) (*Student, error) {
	now := time.Now()
	_, err := r.DB.Exec(
		"UPDATE students SET first_name=?, last_name=?, email=?, age=?, updated_at=? WHERE id=?",
		st.FirstName, st.LastName, st.Email, st.Age, now, id,
	)
	if err != nil {
		return nil, err
	}

	st.ID = id
	st.UpdatedAt = now
	return &st, nil
}

func (r *MySQLStudentRepository) DeleteStudent(id int) error {
	_, err := r.DB.Exec("DELETE FROM students WHERE id=?", id)
	return err
}
