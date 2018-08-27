package pdf_test

import (
	"os"
	"testing"
	"io/ioutil"
	"github.com/iuryfukuda/grapnel/pdf"
)


var lpath = "./testdata"

func load(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func mustLoad(t *testing.T, fn string) []byte {
	b, err := load(lpath + "/" + fn)
	if err != nil {
		t.Fatalf("can't load file: %s", fn)
	}
	return b
}

func TestSuccessConvertPdfToText(t *testing.T) {
	b := mustLoad(t, "valid.pdf")
	text, err := pdf.ToText(b)
	if err != nil {
		t.Errorf("want nil, got: %s", err)
	}
	if testing.Verbose() {
		t.Log(text)
	}
}

func TestErrorConvertInvalidPdfToText(t *testing.T) {
	b := mustLoad(t, "invalid.pdf")
	text, err := pdf.ToText(b)
	if err == nil {
		t.Errorf("want error, got: %s", err)
	}
	if testing.Verbose() {
		t.Log(text)
	}
}
