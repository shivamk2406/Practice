package errorpoc

import (
	"fmt"

	"github.com/pkg/errors"
)

const badInput = "abc"

type BadInputError struct {
	input string
}

func (e *BadInputError) Error() string {
	return fmt.Sprintf("%s : Amount Limit Exceeded", e.input)
}

func validateInput(input string) error {
	if input == badInput {
		//return fmt.Errorf("%w validateInput:", &BadInputError{input: input})
		return fmt.Errorf("Amount Limit Excedded %w ", &BadInputError{input: input})
	}
	return nil
}

type AmountLimitEror struct {
	Input string
}

func (e *AmountLimitEror) Error() string {
	return fmt.Sprintf("Amount Limit Exceeded", e.Input)
}

func Start() error {
	err := errors.New("Hello")
	err2 := errors.Wrap(err, "World")
	//err3 := errors.Wrap("err2")
	if errors.As(err2, err) {
		fmt.Println("Error Occurred")
	}

	//bytes:=[]byte{[123, 34, 116 97 115 107 34 58 34 99 114 109 45 100 97 105 108 121 34 125]}
	//bytes=

	return nil
}
