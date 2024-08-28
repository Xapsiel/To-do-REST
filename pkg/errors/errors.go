package errors

type Errors struct {
	stage string
	text  string
}

func New(stage, text string) *Errors {
	return &Errors{stage: stage, text: text}
}

func (err *Errors) Error() string {
	return err.text
}

func (err *Errors) GetStage() string {
	return err.stage
}
