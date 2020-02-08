package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/SandeepMultani/restful_api_go/models"
)

type todoController struct {
	todoIDPattern *regexp.Regexp
}

func (tc todoController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/todos" {
		switch r.Method {
		case http.MethodGet:
			tc.getAll(w, r)
		case http.MethodPost:
			tc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := tc.todoIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		switch r.Method {
		case http.MethodGet:
			tc.get(id, w)
		case http.MethodPut:
			tc.put(id, w, r)
		case http.MethodDelete:
			tc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (tc *todoController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetToDos(), w)
}

func (tc *todoController) get(id int, w http.ResponseWriter) {
	t, err := models.GetToDoByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(t, w)
}

func (tc *todoController) post(w http.ResponseWriter, r *http.Request) {
	t, err := tc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse ToDo object"))
		return
	}
	t, err = models.AddTodo(t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(t, w)
}

func (tc *todoController) put(id int, w http.ResponseWriter, r *http.Request) {
	t, err := tc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse ToDo object"))
		return
	}
	if id != t.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted todo must match ID in URL"))
		return
	}
	t, err = models.UpdateToDo(t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(t, w)
}

func (tc *todoController) delete(id int, w http.ResponseWriter) {
	err := models.DeleteToDoByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (tc *todoController) parseRequest(r *http.Request) (models.ToDo, error) {
	dec := json.NewDecoder(r.Body)
	var t models.ToDo
	err := dec.Decode(&t)
	if err != nil {
		return models.ToDo{}, err
	}
	return t, nil
}

func newTodoController() *todoController {
	return &todoController{
		todoIDPattern: regexp.MustCompile(`^/todos/(\d+)/?`),
	}
}
