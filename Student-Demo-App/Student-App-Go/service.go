package main

type StudentService struct {
	repo StudentRepository
}

func NewStudentService(repo StudentRepository) *StudentService {
	return &StudentService{repo: repo}
}

func (s *StudentService) ListStudents() ([]Student, error) {
	return s.repo.ListStudents()
}

func (s *StudentService) CreateStudent(st Student) (*Student, error) {
	// Example validation layer (can expand later)
	if st.FirstName == "" || st.LastName == "" {
		return nil, ErrValidation("first name and last name are required")
	}
	return s.repo.CreateStudent(st)
}

func (s *StudentService) GetStudent(id int) (*Student, error) {
	return s.repo.GetStudent(id)
}

func (s *StudentService) UpdateStudent(id int, st Student) (*Student, error) {
	return s.repo.UpdateStudent(id, st)
}

func (s *StudentService) DeleteStudent(id int) error {
	return s.repo.DeleteStudent(id)
}
