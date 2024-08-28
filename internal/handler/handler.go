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
		err = errors.New("Add Handler", err.Error())
		json.NewEncoder(rw).Encode(map[string]interface{}{"error": err.Error()})
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	userID, err := strconv.Atoi(r.URL.Query().Get("userID"))
	if err != nil {
		err = errors.New("Add Handler", err.Error())
		json.NewEncoder(rw).Encode(map[string]interface{}{"error": err.Error()})
		return
	}
	content := r.URL.Query().Get("content")
	ys := speller.NewSpeller(time.Second * time.Duration(timeout))
	newContent, err := ys.CheckText(content, "ru")
	if err != nil {

		json.NewEncoder(rw).Encode(map[string]interface{}{"error": err.Error()})
		return
	}
	content = newContent
	task := task.NewTask(userID, content)

	id, err := task.Add()
	if err != nil {
		json.NewEncoder(rw).Encode(map[string]interface{}{"result": false, "taskID": id})
		return
	}

	err = json.NewEncoder(rw).Encode(map[string]interface{}{"result": true, "taskID": id, "content": content})
	if err != nil {

		err = errors.New("Add Handler", err.Error())
		json.NewEncoder(rw).Encode(map[string]interface{}{"error": err.Error()})
		return
	}
}
func Get(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	value := mux.Vars(r)
	userID, _ := strconv.Atoi(value["userID"])
	task := task.NewTask(userID, "")
	res, err := task.Get(userID)
	if err != nil {
		json.NewEncoder(rw).Encode(map[string]interface{}{"result": false})
		return
	}
	err = json.NewEncoder(rw).Encode(map[string]interface{}{"result": true, "tasks": res})
	if err != nil {
		log.Fatal(err.Error())
	}

}

func SignIn(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	login := r.URL.Query().Get("login")
	password := r.URL.Query().Get("password")
	var err error
	u := user.New(login, password)
	u, err = u.SignIn()
	if err != nil {
		json.NewEncoder(rw).Encode(map[string]interface{}{"result": false, "id": u.ID, "login": u.Login, "error": err.Error()})
		log.Println(err.Error())
		return
	}
	err = json.NewEncoder(rw).Encode(map[string]interface{}{"result": true, "id": u.ID, "login": u.Login})
	if err != nil {
		log.Fatal(err.Error())
	}
}
func SignUp(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	login := r.URL.Query().Get("login")
	password := r.URL.Query().Get("password")
	var err error
	u := user.New(login, password)
	u, err = u.SignUp()
	if err != nil {
		json.NewEncoder(rw).Encode(map[string]interface{}{"result": false, "id": u.ID, "login": u.Login})
		return
	}
	err = json.NewEncoder(rw).Encode(map[string]interface{}{"result": true, "id": u.ID, "login": u.Login})

	if err != nil {
		log.Fatal(err.Error())
	}
}
