package errors

import (
	"fmt"
)

type Errors struct {
	code  int
	stage string
	text  string
}

func New(stage, text string, code int) *Errors {
	return &Errors{stage: stage, text: text, code: code}
}

func (err Errors) Error() string {
	return err.text
}

func (err *Errors) GetStage() string {
	return err.stage
}

func (err *Errors) GetCode() int {
	return err.code
}
func (err *Errors) Print() string {
	return fmt.Sprintf("Stage: %s: \n\tError:%s", err.stage, err.text)
}
