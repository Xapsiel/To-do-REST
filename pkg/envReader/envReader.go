package envreader

import (
	"net/http"
	"os"
	"test_case/pkg/errors"

	"github.com/joho/godotenv"
)

type EnvReader struct {
}

func Init() {
	_, err := New()
	if err != nil {
	}
}
func New(filenames ...string) (*EnvReader, error) {
	err := godotenv.Load(filenames...)
	if err != nil {
		err = godotenv.Load()
		if err != nil {
			return &EnvReader{}, errors.New(".env loading", err.Error(), http.StatusServiceUnavailable)
		}
	}
	return &EnvReader{}, nil
}
func (er EnvReader) GetEnvOrDefault(key, defaultValue string) string {

	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
