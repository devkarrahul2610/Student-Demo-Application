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
		students, err := s.listStudents()
		if err != nil {
			jsonResponse(w, http.StatusInternalServerError, false, nil, err.Error())
			return
		}
		jsonResponse(w, http.StatusOK, true, students, "")
		return

	case http.MethodPost:
		var st Student
		if err := json.NewDecoder(r.Body).Decode(&st); err != nil {
			jsonResponse(w, http.StatusBadRequest, false, nil, "invalid request body")
			return
		}

		created, err := s.createStudent(st)
		if err != nil {
			jsonResponse(w, http.StatusInternalServerError, false, nil, err.Error())
			return
		}

		jsonResponse(w, http.StatusCreated, true, created, "")
		return

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
		st, err := s.getStudent(id)
		if err != nil {
			jsonResponse(w, http.StatusNotFound, false, nil, err.Error())
			return
		}
		jsonResponse(w, http.StatusOK, true, st, "")
		return
	case http.MethodPut:
		var update Student
		if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
			jsonResponse(w, http.StatusBadRequest, false, nil, "invalid body")
			return
		}

		st, err := s.updateStudent(id, update)
		if err != nil {
			jsonResponse(w, http.StatusNotFound, false, nil, err.Error())
			return
		}

		jsonResponse(w, http.StatusOK, true, st, "")
		return

	case http.MethodDelete:
		err := s.deleteStudent(id)
		if err != nil {
			jsonResponse(w, http.StatusNotFound, false, nil, err.Error())
			return
		}
		jsonResponse(w, http.StatusOK, true, nil, "student deleted")
		return

	default:
		jsonResponse(w, http.StatusMethodNotAllowed, false, nil, "method not allowed")
	}
}
