package ui

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mitchellh/colorstring"
)

func Die(err error) bool {
	if err != nil {
		colorstring.Printf("[red]An unexpected error occurred:[reset] %s\n", err)
		return true
	}
	return false
}

func Assert(expr bool, description string) {
	if !expr {
		Die(errors.New(description))
		os.Exit(1)
	}
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

func maxLength(a, b string) int {
	if len(a) > len(b) {
		return len(a)
	}
	return len(b)
}

func leftAlign(v, other string) string {
	if len(v) > len(other) {
		return fmt.Sprintf(" %s ", v)
	} else {
		return fmt.Sprintf(" %s%s ", v, strings.Repeat(" ", len(other)-len(v)))
	}
}

func answerString(first, second string) {
	//  +-----------------+--------+
	//  | Part 1          | Part 2 |
	//  |-----------------+--------|
	//  | 123413419459185 | 12345  |
	//  +-----------------+--------+

	dashesA1 := strings.Repeat("-", maxLength("Part 1", first)+2)
	dashesA2 := strings.Repeat("-", maxLength("Part 2", second)+2)

	colorstring.Printf("[yellow]+%s+%s+\n", dashesA1, dashesA2)
	colorstring.Printf("[yellow]|[green]%s[yellow]|[green]%s[yellow]|\n", leftAlign("Part 1", first), leftAlign("Part 2", second))
	colorstring.Printf("[yellow]|%s+%s|\n", dashesA1, dashesA2)
	colorstring.Printf("[yellow]|[cyan]%s[yellow]|[cyan]%s[yellow]|\n", leftAlign(first, "Part 1"), leftAlign(second, "Part 2"))
	colorstring.Printf("[yellow]+%s+%s+\n", dashesA1, dashesA2)
}

func answerInt(first, second int) {
	a1 := strconv.FormatInt(int64(first), 10)
	a2 := strconv.FormatInt(int64(second), 10)

	answerString(a1, a2)
}

func Answer(first, second interface{}) {
	switch first.(type) {
	case int, int64:
		answerInt(first.(int), second.(int))
	case string:
		answerString(first.(string), second.(string))
	}
}
