package tokens

type Token struct {
	Raw    string
	Type   Type
	Line   int
	Column int
}

type Type byte

const (
	UNKNOWN Type = iota

	NUMBER

	PLUS
	MINUS
	ASTERIKS
	SLASH

	EOF
)
