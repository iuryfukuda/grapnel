package pdf

import (
	"io"
	"bytes"
	"errors"
	"os/exec"

	"github.com/zbioe/grapnel/text"
)

func ToWords(b []byte) (map[string]int, error) {
	s, err := ToText(b)
	if err != nil {
		return nil, err
	}
	return text.ToWords(s), nil
}

func ToWordsFromReader(r io.Reader) (map[string]int, error) {
	s, err := ToTextFromReader(r)
	if err != nil {
		return nil, err
	}
	return text.ToWords(s), nil
}

func ToText(b [] byte) (string, error) {
	r := bytes.NewReader(b)
	return ToTextFromReader(r)
}

func ToTextFromReader(r io.Reader) (string, error) {
	var bufText bytes.Buffer
	var bufErr bytes.Buffer
	var cmd = exec.Command(
		"pdftotext", "-nopgbrk", "-enc", "UTF-8",
		"-eol", "unix", "-", "-",
	)
	cmd.Stdout = &bufText
	cmd.Stderr = &bufErr
	cmd.Stdin = r
	err := cmd.Start()
	if err != nil {
		return "", err
	}
	cmd.Wait()
	if bufErr.Len() > 0 {
		err = errors.New(bufErr.String())
	}
	if err != nil {
		return "", err
	}
	return bufText.String(), nil
}
