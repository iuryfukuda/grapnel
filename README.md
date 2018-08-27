# Grapnel

Repository with tools to convert for text some content types

## PDF

Receive Pdf in []byte or io.Reader and transform him to text with [pdftotext](https://www.xpdfreader.com/pdftotext-man.html)

### Example

create file `main.go` with content:
```go
package main

import (
	"os"
	"fmt"
	"github.com/iuryfukuda/grapnel/pdf"
)

func main() {
	text, err := pdf.ToTextFromReader(os.Stdin)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Print(text)
}
```
run on command line:
```sh
go run main.go < pdf/test_files/valid.pdf
curl -Ls "http://www.orimi.com/pdf-test.pdf" | go run main.go
```
### Requirements

- [pdftotext](https://www.xpdfreader.com/download.html)


## HTML

Receive html in bytes or reader and transform him to text

### Example

create file `main.go` with content:
```go
package main

import (
	"os"
	"fmt"
	"github.com/iuryfukuda/grapnel/html"
)

func main() {
	text, err := html.ToTextFromReader(os.Stdin)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Print(text)
}
```
run on command line:
```sh
go run main.go < pdf/testdata/valid.html
curl -Ls "https://reddit.com/" | go run main.go
```
