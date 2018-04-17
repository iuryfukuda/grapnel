package pdf

import (
	"io"
	"bytes"
	"errors"
	"os/exec"
)

func ToText(r io.Reader) (text string, err error) {
	var bufText bytes.Buffer
	var bufErr bytes.Buffer
	subProcess := exec.Command(
		"pdftotext", "-nopgbrk", "-enc", "UTF-8",
		"-eol", "unix", "-", "-",
	)
	subProcess.Stdout = &bufText
	subProcess.Stderr = &bufErr
	subProcess.Stdin = r
	err = subProcess.Start()
	subProcess.Wait()
	if bufErr.Len() > 0 {
		err = errors.New(bufErr.String())
	}
	text = bufText.String()
	return text, err
}
