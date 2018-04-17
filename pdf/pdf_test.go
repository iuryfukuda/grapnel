package pdf_test

import (
	"os"
	"io"
	"log"
	"testing"
	"github.com/iuryfukuda/grapnel/pdf"
)

func loadReaderOf(rpath string) io.Reader {
	lpath := "./test_files"
	f, err := os.Open(lpath + "/" + rpath)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func verboseTextErr(text string, err error) {
	log.Print("TEXT:\n", text)
	log.Print("ERR:\n", err)
}

func TestSuccessConvertPdfToText(t *testing.T) {
	r := loadReaderOf("valid.pdf")
	text, err := pdf.ToText(r)
	if err != nil {
		t.Errorf("Bad Test Success pdf to text %v", text)
	}
	verboseTextErr(text, err)
}

func TestErrorConvertInvalidPdfToText(t *testing.T) {
	r := loadReaderOf("invalid.pdf")
	text, err := pdf.ToText(r)
	if err == nil {
		t.Errorf("Bad Test Error pdf to invalid %v", err)
	}
	verboseTextErr(text, err)
}
