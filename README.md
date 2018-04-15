# Grapnel

Repository with tools for convert body in response to plain text

## PDF

Receive ReadCloser and transform to text with [pdftotext](https://www.xpdfreader.com/pdftotext-man.html)

### Example Usage

create file `main.go`
```go
package main

import (
	"os"
	"fmt"
	"github.com/iuryfukuda/grapnel/pdf"
)

func main() {
	text, err := pdf.ToText(os.Stdin)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Print(text)
}
```
run on command line
```sh
go run main.go < pdf/test_files/valid.pdf
```
### Requirements

- [pdftotext](https://www.xpdfreader.com/download.html)
