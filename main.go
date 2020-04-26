package main

import (
	"os"
	"io"
	"fmt"
	"log"
	"flag"
	"strings"
	"net/http"
	"io/ioutil"
	"github.com/zbioe/grapnel/html"
	"github.com/zbioe/grapnel/pdf"
)

var readerType = flag.String("type", "", "pdf|html, keep empty for detect content")

func main() {
	flag.Parse()
	setup()

	s, err := readerToText(os.Stdin, *readerType)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(os.Stdout, s)
}

func readerToText(r io.Reader, readerType string) (string, error) {
	switch readerType {
	case "":
		return unknowReaderToText(os.Stdin)
	case "pdf":
		return html.ToTextFromReader(os.Stdin)
	case "html":
		return html.ToTextFromReader(os.Stdin)
	}
	return "", fmt.Errorf("unknow type: %s", readerType)
}

func unknowReaderToText(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("Can't read stdin: %s", err)
	}
	var ct = http.DetectContentType(b)
	switch strings.SplitN(ct, ";", -1)[0] {
	case "application/pdf":
		return pdf.ToText(b)
	case "text/html":
		return html.ToText(b)
	case "text/plain":
		return string(b[:]), nil
	}
	return "", fmt.Errorf("Invalid Content Type: %s", ct)
}

func setup() {
	log.SetFlags(0)
	log.SetPrefix("grapnel: ")
}
