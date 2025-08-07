package token

type Position struct {
	Filename string
	Line     int
	Column   int
}

type Type int

type Token struct {
	Text     string
	Type     Type
	Position Position
}

func (t Token) Valid() bool {
	return t.Type != token_bg
}

const (
	token_bg Type = iota
	COMMON
	ILLEGAL
	EOF
)
