package pdf

import (
	"bytes"
	"errors"
	"os/exec"
)

func ToText(b [] byte) (text string, err error) {
	var bufText bytes.Buffer
	var bufErr bytes.Buffer
	var cmd = exec.Command(
		"pdftotext", "-nopgbrk", "-enc", "UTF-8",
		"-eol", "unix", "-", "-",
	)
	cmd.Stdout = &bufText
	cmd.Stderr = &bufErr
	cmd.Stdin = bytes.NewBuffer(b)
	err = cmd.Start()
	if err != nil {
		return
	}
	cmd.Wait()
	if bufErr.Len() > 0 {
		err = errors.New(bufErr.String())
	}
	if err != nil {
		return
	}
	text = bufText.String()
	return
}
