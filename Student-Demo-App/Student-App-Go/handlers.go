package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type StudentHandlers struct {
	service *StudentService
}

func NewStudentHandlers(service *StudentService) *StudentHandlers {
	return &StudentHandlers{service: service}
}

func (h *StudentHandlers) handleStudents(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		students, err := h.service.ListStudents()
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

		created, err := h.service.CreateStudent(st)
		if err != nil {
			jsonResponse(w, http.StatusBadRequest, false, nil, err.Error())
			return
		}
		jsonResponse(w, http.StatusCreated, true, created, "")
		return

	default:
		jsonResponse(w, http.StatusMethodNotAllowed, false, nil, "method not allowed")
	}
}

func (h *StudentHandlers) handleStudentByID(w http.ResponseWriter, r *http.Request) {
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
		st, err := h.service.GetStudent(id)
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

		st, err := h.service.UpdateStudent(id, update)
		if err != nil {
			jsonResponse(w, http.StatusBadRequest, false, nil, err.Error())
			return
		}
		jsonResponse(w, http.StatusOK, true, st, "")
		return

	case http.MethodDelete:
		err := h.service.DeleteStudent(id)
		if err != nil {
			jsonResponse(w, http.StatusBadRequest, false, nil, err.Error())
			return
		}
		jsonResponse(w, http.StatusOK, true, nil, "student deleted")
		return

	default:
		jsonResponse(w, http.StatusMethodNotAllowed, false, nil, "method not allowed")
	}
}
