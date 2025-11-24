package student

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"student/student-demo-app/response"
)

type StudentHandlers struct {
	service *StudentService
}

func NewStudentHandlers(service *StudentService) *StudentHandlers {
	return &StudentHandlers{service: service}
}

func (h *StudentHandlers) HandleStudents(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		students, err := h.service.ListStudents()
		if err != nil {
			response.JsonResponse(w, http.StatusInternalServerError, false, nil, err.Error())
			return
		}
		response.JsonResponse(w, http.StatusOK, true, students, "")
		return

	case http.MethodPost:
		var st Student
		if err := json.NewDecoder(r.Body).Decode(&st); err != nil {
			response.JsonResponse(w, http.StatusBadRequest, false, nil, "invalid request body")
			return
		}

		created, err := h.service.CreateStudent(st)
		if err != nil {
			response.JsonResponse(w, http.StatusBadRequest, false, nil, err.Error())
			return
		}
		response.JsonResponse(w, http.StatusCreated, true, created, "")
		return

	default:
		response.JsonResponse(w, http.StatusMethodNotAllowed, false, nil, "method not allowed")
	}
}

func (h *StudentHandlers) HandleStudentByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 2 {
		response.JsonResponse(w, http.StatusNotFound, false, nil, "invalid path")
		return
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		response.JsonResponse(w, http.StatusBadRequest, false, nil, "invalid id")
		return
	}

	switch r.Method {

	case http.MethodGet:
		st, err := h.service.GetStudent(id)
		if err != nil {
			response.JsonResponse(w, http.StatusNotFound, false, nil, err.Error())
			return
		}
		response.JsonResponse(w, http.StatusOK, true, st, "")
		return

	case http.MethodPut:
		var update Student
		if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
			response.JsonResponse(w, http.StatusBadRequest, false, nil, "invalid body")
			return
		}

		st, err := h.service.UpdateStudent(id, update)
		if err != nil {
			response.JsonResponse(w, http.StatusBadRequest, false, nil, err.Error())
			return
		}
		response.JsonResponse(w, http.StatusOK, true, st, "")
		return

	case http.MethodDelete:
		err := h.service.DeleteStudent(id)
		if err != nil {
			response.JsonResponse(w, http.StatusBadRequest, false, nil, err.Error())
			return
		}
		response.JsonResponse(w, http.StatusOK, true, nil, "student deleted")
		return

	default:
		response.JsonResponse(w, http.StatusMethodNotAllowed, false, nil, "method not allowed")
	}
}
