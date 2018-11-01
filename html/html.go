package html

import (
	"io"
	"bytes"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"github.com/iuryfukuda/grapnel/text"
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

func ToText(b []byte) (string, error) {
	r := bytes.NewReader(b)
	return ToTextFromReader(r)
}

func ToTextFromReader(r io.Reader) (string, error) {
	var (
		text = ""
		hasText = false
		z = html.NewTokenizer(r)
	)

Loop:
	for {
		tokenType := z.Next()
		token := z.Token()
		switch tokenType {
		case html.ErrorToken:
			switch err := z.Err(); err {
			case io.EOF:
				break Loop
			default:
				return text, err
			}
		case html.StartTagToken:
			switch token.DataAtom {
			case atom.Script, atom.Style:
				hasText = false
			default:
				hasText = true
			}
		case html.TextToken:
			if hasText {
				text += token.String() + "\n"
			}
		}
	}
	return text, nil
}
