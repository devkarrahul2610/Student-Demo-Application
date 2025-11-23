package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func (s *Store) handleStudents(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		students := s.listStudents()
		jsonResponse(w, http.StatusOK, true, students, "")

	case http.MethodPost:
		var st Student
		if err := json.NewDecoder(r.Body).Decode(&st); err != nil {
			jsonResponse(w, http.StatusBadRequest, false, nil, "invalid JSON payload")
			return
		}
		created := s.createStudent(st)
		jsonResponse(w, http.StatusCreated, true, created, "")

	default:
		jsonResponse(w, http.StatusMethodNotAllowed, false, nil, "method not allowed")
	}
}

func (s *Store) handleStudentByID(w http.ResponseWriter, r *http.Request) {
	// URL format: /students/{id}
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 2 {
		jsonResponse(w, http.StatusNotFound, false, nil, "invalid path")
		return
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, false, nil, "invalid id")
		return
	}

	switch r.Method {

	case http.MethodGet:
		st, ok := s.getStudent(id)
		if !ok {
			jsonResponse(w, http.StatusNotFound, false, nil, "student not found")
			return
		}
		jsonResponse(w, http.StatusOK, true, st, "")
	case http.MethodPut:
		var update Student
		if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
			jsonResponse(w, http.StatusBadRequest, false, nil, "invalid JSON payload")
			return
		}
		st, ok := s.updateStudent(id, update)
		if !ok {
			jsonResponse(w, http.StatusNotFound, false, nil, "student not found")
			return
		}
		jsonResponse(w, http.StatusOK, true, st, "")

	case http.MethodDelete:
		if !s.deleteStudent(id) {
			jsonResponse(w, http.StatusNotFound, false, nil, "student not found")
			return
		}
		jsonResponse(w, http.StatusNoContent, true, nil, "")

	default:
		jsonResponse(w, http.StatusMethodNotAllowed, false, nil, "method not allowed")
	}
}
