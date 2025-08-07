package scanner

import (
	"bytes"
	"fmt"
	"github.com/lerity-yao/cztctl/goTemplate/rabbitmqGen/token"
	"github.com/lerity-yao/cztctl/util/pathx"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Scanner struct {
	filename string
	size     int

	data         []rune
	position     int // current position in input (points to current char)
	readPosition int // current reading position in input (after current char)
	ch           rune

	lines []int
}

// MustNewScanner returns a new scanner for the given filename and data.
func MustNewScanner(filename string, src interface{}) *Scanner {
	sc, err := NewScanner(filename, src)
	if err != nil {
		log.Fatalln(err)
	}
	return sc
}

// NewScanner returns a new scanner for the given filename and data.
func NewScanner(filename string, src interface{}) (*Scanner, error) {
	data, err := readData(filename, src)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("filename: %s, missing input", filename)
	}

	var runeList []rune
	for _, r := range string(data) {
		runeList = append(runeList, r)
	}

	filename = filepath.Base(filename)
	s := &Scanner{
		filename:     filename,
		size:         len(runeList),
		data:         runeList,
		lines:        []int{-1},
		readPosition: 0,
	}

	s.readRune()
	return s, nil
}

func (s *Scanner) NextToken() (token.Token, error) {
	s.skipWhiteSpace()
	if s.isIdentifierLetter(s.ch) {
		return s.scanIdent(), nil
	}

	if s.isDigit(s.ch) {
		return s.scanIntOrDuration(), nil
	}

	if s.ch == 0 {
		return token.Token{
			Type: token.EOF,
		}, nil
	}

	tok := token.Token{
		Type:     token.ILLEGAL,
		Text:     string(s.ch),
		Position: s.newPosition(s.position),
	}
	s.readRune()
	return tok, nil
}

func (s *Scanner) skipWhiteSpace() {
	for s.isWhiteSpace(s.ch) {
		s.readRune()
	}
}

func (s *Scanner) isWhiteSpace(b rune) bool {
	if b == '\n' {
		s.lines = append(s.lines, s.position)
	}
	return b == ' ' || b == '\t' || b == '\r' || b == '\f' || b == '\v' || b == '\n'
}

func (s *Scanner) isIdentifierLetter(b rune) bool {
	if s.isLetter(b) {
		return true
	}
	return b == '_'
}

func (s *Scanner) isLetter(b rune) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func (s *Scanner) isDigit(b rune) bool {
	return b >= '0' && b <= '9'
}

func (s *Scanner) scanIdent() token.Token {
	position := s.position
	for s.isIdentifierLetter(s.ch) || s.isDigit(s.ch) {
		s.readRune()
	}

	ident := string(s.data[position:s.position])

	return token.Token{
		Type:     token.COMMON,
		Text:     ident,
		Position: s.newPosition(position),
	}
}

func (s *Scanner) scanIntOrDuration() token.Token {
	position := s.position
	for s.isDigit(s.ch) {
		s.readRune()
	}

	return token.Token{
		Type:     token.COMMON,
		Text:     string(s.data[position:s.position]),
		Position: s.newPosition(position),
	}

}

func (s *Scanner) newPosition(position int) token.Position {
	line := s.lineCount()
	return token.Position{
		Filename: s.filename,
		Line:     line,
		Column:   position - s.lines[line-1],
	}
}

func (s *Scanner) lineCount() int {
	return len(s.lines)
}

func (s *Scanner) readRune() {
	if s.readPosition >= s.size {
		s.ch = 0
	} else {
		s.ch = s.data[s.readPosition]
	}
	s.position = s.readPosition
	s.readPosition += 1
}

func readData(filename string, src interface{}) ([]byte, error) {
	if strings.HasSuffix(filename, ".api") && pathx.FileExists(filename) {
		data, err := os.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	switch v := src.(type) {
	case []byte:
		return v, nil
	case *bytes.Buffer:
		return v.Bytes(), nil
	case string:
		return []byte(v), nil
	default:
		return nil, fmt.Errorf("unsupported type: %T", src)
	}
}
