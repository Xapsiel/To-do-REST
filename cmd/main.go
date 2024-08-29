package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"test_case/internal/handler"
	"test_case/internal/repo"
	envreader "test_case/pkg/envReader"
	"test_case/pkg/errors"
	"test_case/pkg/server"

	"github.com/gorilla/mux"
)

var er *envreader.EnvReader

func init() {
	var err error
	er, err = envreader.New()
	if err != nil {
		if e, ok := err.(errors.Errors); ok {
			log.Println(e.Print())
		}
		return
	}
	_ = repo.Repo{}

}

func main() {
	// Чтение конфигурации
	port := er.GetEnvOrDefault("PORT", "8080")
	timeoutSeconds, _ := strconv.Atoi(er.GetEnvOrDefault("TIMEOUT_SECONDS", "60"))

	// Создание нового сервера
	s := server.New(port, time.Duration(timeoutSeconds)*time.Second)

	// Регистрация маршрутов
	registerRoutes(s.Router)

	// Запуск сервера
	log.Printf("Starting server on port %s", port)
	err := s.Run()
	if err != nil {
		if e, ok := err.(errors.Errors); ok {

			log.Println(e.Print())
		}
	}
}

func registerRoutes(router *mux.Router) {
	router.HandleFunc("/add", handler.Add)
	router.HandleFunc("/get/{userID}", handler.Get).Methods("GET")
	router.HandleFunc("/signIn", handler.SignIn).Methods("POST")
	router.HandleFunc("/signUp", handler.SignUp).Methods("POST")
}

func GetEnvOrDefault(key, defaultValue string) string {

	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
