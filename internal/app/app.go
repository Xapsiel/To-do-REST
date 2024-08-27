package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"test_case/internal/speller"
	"test_case/internal/task"
	"test_case/internal/user"

	"github.com/gorilla/mux"
)

func Add(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	userID, _ := strconv.Atoi(r.URL.Query().Get("userID"))
	content := r.URL.Query().Get("content")
	ys := speller.NewSpeller()
	newContent, err := ys.CheckText(content, "ru")
	content = newContent

	task := task.NewTask()
	id, err := task.Add(userID, content)
	if err != nil {
		json.NewEncoder(rw).Encode(map[string]interface{}{"result": false, "taskID": id})
		return
	}

	if err != nil {
		json.NewEncoder(rw).Encode(map[string]interface{}{"result": false, "taskID": id})
		return
	}
	err = json.NewEncoder(rw).Encode(map[string]interface{}{"result": true, "taskID": id, "content": content})
	if err != nil {
		log.Fatal(err.Error())
	}
}
func Get(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	value := mux.Vars(r)
	userID, _ := strconv.Atoi(value["userID"])
	task := task.NewTask()
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
	var u *user.User = &user.User{ID: -1}

	u.Login = r.URL.Query().Get("login")
	u.Password = r.URL.Query().Get("password")
	var err error
	u, err = user.SignIn(u.Login, u.Password)
	if err != nil {
		json.NewEncoder(rw).Encode(map[string]interface{}{"result": false, "id": u.ID, "login": u.Login})
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
	var u *user.User = &user.User{ID: -1}

	u.Login = r.URL.Query().Get("login")
	u.Password = r.URL.Query().Get("password")
	var err error
	u, err = user.SignUp(u.Login, u.Password)
	if err != nil {
		json.NewEncoder(rw).Encode(map[string]interface{}{"result": false, "id": u.ID, "login": u.Login})
		return
	}
	err = json.NewEncoder(rw).Encode(map[string]interface{}{"result": true, "id": u.ID, "login": u.Login})

	if err != nil {
		log.Fatal(err.Error())
	}
}
