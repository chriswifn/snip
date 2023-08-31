package snip

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Fill(form string, fields ...string) string {
	for n, v := range fields {
		lk := fmt.Sprintf("{%v}", n+1)
		form = strings.ReplaceAll(form, lk, v)
	}
	return form
}

func FillForm(r io.Reader, fields ...string) string {
	byt, err := io.ReadAll(r)
	if err != nil {
		return ""
	}
	form := string(byt)
	for n, v := range fields {
		lk := fmt.Sprintf("{%v}", n+1)
		form = strings.ReplaceAll(form, lk, v)
	}
	return form
}

func FillIn(fields ...string) string {
	return FillForm(os.Stdin, fields...)
}

func FillFile(file string, fields ...string) string {
	dat, err := os.ReadFile(file)
	if err != nil {
		return ""
	}
	return Fill(string(dat), fields...)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println(FillIn(os.Args[1:]...))
	}
	file := os.Args[1]
	args := os.Args[2:]
	if len(args) == 0 {
		fmt.Println("")
	}
	fmt.Println(FillFile(file, args...))
}
