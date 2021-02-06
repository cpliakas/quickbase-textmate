package qbtm

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/cpliakas/cliutil"
)

// Parser parses data from the CSV file.
type Parser interface {
	fmt.Stringer
	Row([]string) error
}

// FunctionParser parses functions from the CSV file.
type FunctionParser struct {
	funcs []string
}

// NewFunctionParser returns a new *FunctionParser.
func NewFunctionParser() *FunctionParser {
	return &FunctionParser{funcs: []string{}}
}

// Row implements Parser.Row by checking whether the first character of the
// value in the second column is an uppercase letter.
func (p *FunctionParser) Row(row []string) error {
	if len(row[1]) > 0 {
		r, _ := utf8.DecodeRuneInString(row[1])
		if unicode.IsLetter(r) && unicode.IsUpper(r) {
			p.funcs = append(p.funcs, row[1])
		}
	}
	return nil
}

func (p *FunctionParser) String() string {
	if len(p.funcs) > 0 {
		return fmt.Sprintf("\\b(%s)\\b", strings.Join(Unique(p.funcs), "|"))
	}
	return ""
}

// OperationParser parses functions from the CSV file.
type OperationParser struct {
	ops []string
}

// NewOperationParser returns a new *OperationParser.
func NewOperationParser() *OperationParser {
	return &OperationParser{ops: []string{}}
}

// Row implements Parser.Row by checking whether the first character of the
// value in the second column is not a letter.
func (p *OperationParser) Row(row []string) error {
	if len(row[1]) > 0 {
		r, _ := utf8.DecodeRuneInString(row[1])
		if !unicode.IsLetter(r) {
			p.ops = append(p.ops, regexp.QuoteMeta(row[1]))
		}
	}
	return nil
}

func (p *OperationParser) String() string {
	if len(p.ops) > 0 {
		return fmt.Sprintf("(%s)", strings.Join(Unique(p.ops), "|"))
	}
	return ""
}

// Snippet models a snippet.
type Snippet struct {
	Prefix      []string `json:"prefix"`
	Body        []string `json:"body"`
	Description string   `json:"description"`
}

// SnippetParser parses functions from the CSV file.
type SnippetParser struct {
	snippets map[string]*Snippet
}

// NewSnippetParser returns a new *SnippetParser.
func NewSnippetParser() *SnippetParser {
	snippets := map[string]*Snippet{}

	for _, t := range _types {
		k := fmt.Sprintf("%s variable", t)

		body := fmt.Sprintf("var %s $1 = ", t)
		if t == "String" {
			body += `"$2";`
		} else {
			body += "$2;"
		}

		snippets[k] = &Snippet{
			Prefix:      []string{"var " + t, t},
			Body:        []string{body},
			Description: fmt.Sprintf("Create a %s variable", t),
		}
	}

	return &SnippetParser{snippets: snippets}
}

// Row implements Parser.Row.
func (p *SnippetParser) Row(row []string) error {
	if len(row[1]) > 0 {
		r, _ := utf8.DecodeRuneInString(row[1])
		if unicode.IsLetter(r) && unicode.IsUpper(r) {
			k := fmt.Sprintf("%s %s", row[1], row[2])
			p.snippets[k] = &Snippet{
				Prefix:      []string{row[1]},
				Body:        []string{row[1] + ParseArgs(row[2])},
				Description: strings.TrimSpace(row[4]),
			}
		}
	}
	return nil
}

func (p *SnippetParser) String() string {
	s, _ := cliutil.FormatJSON(p.snippets)
	return s
}

// ParseArgs parses (Text t, Number p, Text d) into ($1, $2, $3).
func ParseArgs(args string) string {
	a := strings.Split(strings.Trim(args, "()"), ",")

	v := []string{}
	for idx, a := range a {
		if strings.Contains(a, "..") || a == "" {
			break
		}
		v = append(v, fmt.Sprintf("$%v", idx+1))
	}

	return fmt.Sprintf("(%s)", strings.Join(v, ", "))
}
