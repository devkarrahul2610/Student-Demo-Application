package student

import (
	"fmt"
	"time"

	"student/student-demo-app/internal/cache"
	"student/student-demo-app/response"
)

type StudentService struct {
	repo StudentRepository
}

func NewStudentService(repo StudentRepository) *StudentService {
	return &StudentService{repo: repo}
}

// ListStudents remains DB-only (or optionally you could cache all students if needed)
func (s *StudentService) ListStudents() ([]Student, error) {
	return s.repo.ListStudents()
}

func (s *StudentService) CreateStudent(st Student) (*Student, error) {
	// Basic validation
	if st.FirstName == "" || st.LastName == "" {
		return nil, response.ErrValidation("first name and last name are required")
	}

	created, err := s.repo.CreateStudent(st)
	if err != nil {
		return nil, err
	}

	// Optional: cache newly created student
	cacheKey := fmt.Sprintf("student:%d", created.ID)
	cache.Set(cacheKey, created, 5*time.Minute)

	return created, nil
}

func (s *StudentService) GetStudent(id int) (*Student, error) {
	cacheKey := fmt.Sprintf("student:%d", id)

	// Try cache first
	var st Student
	err := cache.Get(cacheKey, &st)
	if err == nil {
		// Cache hit
		fmt.Printf("[Cache HIT] student:%d\n", id)
		return &st, nil
	}

	// Cache miss → fetch from DB
	stPtr, err := s.repo.GetStudent(id)
	if err != nil {
		return nil, err
	}

	// Save result to cache for next time
	cache.Set(cacheKey, stPtr, 5*time.Minute)

	return stPtr, nil
}

func (s *StudentService) UpdateStudent(id int, st Student) (*Student, error) {
	updated, err := s.repo.UpdateStudent(id, st)
	if err != nil {
		return nil, err
	}

	// Invalidate cache
	cacheKey := fmt.Sprintf("student:%d", id)
	cache.Delete(cacheKey)

	return updated, nil
}

func (s *StudentService) DeleteStudent(id int) error {
	err := s.repo.DeleteStudent(id)
	if err != nil {
		return err
	}

	// Invalidate cache
	cacheKey := fmt.Sprintf("student:%d", id)
	cache.Delete(cacheKey)

	return nil
}
