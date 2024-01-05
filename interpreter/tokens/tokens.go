package tokens

type Token struct {
	Raw    string
	Type   Type
	Line   int
	Column int
}

type Type uint8

const (
	UNKNOWN Type = iota
	NUMBER
	PLUS
	MINUS
	ASTERIKS
	SLASH
	BRACE_LEFT
	BRACE_RIGHT
	EOF
)

var TYPE_MAP = map[Type]string{
	UNKNOWN:     "UNKNOWN",
	NUMBER:      "NUMBER",
	PLUS:        "PLUS",
	MINUS:       "MINUS",
	ASTERIKS:    "ASTERIKS",
	SLASH:       "SLASH",
	BRACE_LEFT:  "BRACE_LEFT",
	BRACE_RIGHT: "BRACE_RIGHT",
	EOF:         "EOF",
}
