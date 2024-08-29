package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"test_case/internal/task"
	"test_case/internal/user"
	envreader "test_case/pkg/envReader"
	"test_case/pkg/errors"
	"test_case/pkg/speller"
	"time"

	"github.com/gorilla/mux"
)

func Add(rw http.ResponseWriter, r *http.Request) {
	timeout, err := strconv.Atoi(envreader.EnvReader{}.GetEnvOrDefault("TIMEOUT_SECONDS", "60"))
	if err != nil {
		err = errors.New("Add Handler", err.Error(), http.StatusServiceUnavailable)
		if e, ok := err.(*errors.Errors); ok {
			log.Println(e.Print())
			rw.WriteHeader(e.GetCode())
		}
		json.NewEncoder(rw).Encode(map[string]interface{}{"error": err.Error()})

		return
	}
	rw.Header().Set("Content-Type", "application/json")
	userID, err := strconv.Atoi(r.URL.Query().Get("userID"))
	if err != nil {
		err = errors.New("Add Handler", err.Error(), http.StatusServiceUnavailable)
		if e, ok := err.(*errors.Errors); ok {
			log.Println(e.Print())
			rw.WriteHeader(e.GetCode())
		}
		json.NewEncoder(rw).Encode(map[string]interface{}{"error": err.Error()})

		return
	}
	content := r.URL.Query().Get("content")
	ys := speller.NewSpeller(time.Second * time.Duration(timeout))
	newContent, err := ys.CheckText(content, "ru")
	if err != nil {

		if e, ok := err.(*errors.Errors); ok {
			log.Println(e.Print())
			rw.WriteHeader(e.GetCode())
		}
		json.NewEncoder(rw).Encode(map[string]interface{}{"error": err.Error()})

		return

	}
	content = newContent
	task := task.NewTask(userID, content)

	id, err := task.Add()
	if err != nil {
		if e, ok := err.(*errors.Errors); ok {
			log.Println(e.Print())
			rw.WriteHeader(e.GetCode())
		}
		json.NewEncoder(rw).Encode(map[string]interface{}{"error": err.Error()})

		return
	}

	rw.WriteHeader(http.StatusOK)
	err = json.NewEncoder(rw).Encode(map[string]interface{}{"result": true, "taskID": id, "content": content})
	if err != nil {

		err = errors.New("Add Handler", err.Error(), http.StatusServiceUnavailable)
		if e, ok := err.(*errors.Errors); ok {
			log.Println(e.Print())
			rw.WriteHeader(e.GetCode())
		}
		json.NewEncoder(rw).Encode(map[string]interface{}{"error": err.Error()})

		return

	}
}
func Get(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	value := mux.Vars(r)
	userID, _ := strconv.Atoi(value["userID"])
	task := task.NewTask(userID, "")
	res, err := task.Get(userID)
	if err != nil {
		if e, ok := err.(*errors.Errors); ok {
			log.Println(e.Print())
			rw.WriteHeader(e.GetCode())
		}
		json.NewEncoder(rw).Encode(map[string]interface{}{"error": err.Error()})

		return
	}
	rw.WriteHeader(http.StatusOK)

	err = json.NewEncoder(rw).Encode(map[string]interface{}{"result": true, "tasks": res})
	if err != nil {
		if e, ok := err.(*errors.Errors); ok {
			log.Println(e.Print())
			rw.WriteHeader(e.GetCode())
		}
		json.NewEncoder(rw).Encode(map[string]interface{}{"error": err.Error()})

		return
	}

}

func SignIn(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	login := r.URL.Query().Get("login")
	password := r.URL.Query().Get("password")
	var err error
	u := user.New(login, password)
	u, err = u.SignIn()
	if err != nil {
		if e, ok := err.(*errors.Errors); ok {
			log.Println(e.Print())
			rw.WriteHeader(e.GetCode())
		}
		json.NewEncoder(rw).Encode(map[string]interface{}{"error": err.Error()})

		return

	}
	rw.WriteHeader(http.StatusOK)

	err = json.NewEncoder(rw).Encode(map[string]interface{}{"result": true, "id": u.ID, "login": u.Login})
	if err != nil {
		if e, ok := err.(*errors.Errors); ok {
			log.Println(e.Print())
			rw.WriteHeader(e.GetCode())
		}
		json.NewEncoder(rw).Encode(map[string]interface{}{"error": err.Error()})

		return
	}
}
func SignUp(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	login := r.URL.Query().Get("login")
	password := r.URL.Query().Get("password")
	var err error
	u := user.New(login, password)
	u, err = u.SignUp()
	if err != nil {
		if e, ok := err.(*errors.Errors); ok {
			log.Println(e.Print())
			rw.WriteHeader(e.GetCode())
		}
		json.NewEncoder(rw).Encode(map[string]interface{}{"error": err.Error()})

		return
	}
	rw.WriteHeader(http.StatusOK)

	err = json.NewEncoder(rw).Encode(map[string]interface{}{"result": true, "id": u.ID, "login": u.Login})

	if err != nil {
		if e, ok := err.(*errors.Errors); ok {
			log.Println(e.Print())
			rw.WriteHeader(e.GetCode())
		}
		json.NewEncoder(rw).Encode(map[string]interface{}{"error": err.Error()})

		return
	}
}
