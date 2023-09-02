package snip

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
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

func ListSnip(path string) []string {
	var snips []string
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			snips = append(snips, path)
			return nil
		})
	if err != nil {
		return []string{}
	}
	return snips
}
