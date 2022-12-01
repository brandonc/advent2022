package ui

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mitchellh/colorstring"
)

var answerCount int

func Die(err error) bool {
	if err != nil {
		colorstring.Printf("[red]An unexpected error occurred:[reset] %s\n", err)
		return true
	}
	return false
}

func Debug(message string) {
	if os.Getenv("LOG_LEVEL") != "debug" {
		return
	}
	colorstring.Printf("[dark_gray][DEBUG] %s\n", message)
}

func Debugf(message string, a ...any) {
	Debug(fmt.Sprintf(message, a...))
}

func Answer(answer string, description ...string) {
	answerCount++
	desc := ""
	if len(description) > 0 {
		desc = " (" + strings.Join(description, " ") + ")"
	}

	colorstring.Printf("[cyan][ANSWER #%d] [magenta]%s[yellow]%s\n", answerCount, answer, desc)
}

func AnswerInt(answer int, description ...string) {
	Answer(strconv.FormatInt(int64(answer), 10), description...)
}
