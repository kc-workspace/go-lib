package random

type StringMode int

const (
	Numeric = iota
	AlphabetLower
	AlphabetUpper
	AlphaNumericLower
	AlphaNumericUpper
	AlphaNumeric
)

var (
	numeric           = []rune("0123456789")
	alphabetLower     = []rune("abcdefghijklmnopqrstuvwxyz")
	alphabetUpper     = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	alphanumericLower = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	alphanumericUpper = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	alphanumeric      = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func selectRune(mode StringMode) []rune {
	switch mode {
	case Numeric:
		return numeric
	case AlphabetLower:
		return alphabetLower
	case AlphabetUpper:
		return alphabetUpper
	case AlphaNumericLower:
		return alphanumericLower
	case AlphaNumericUpper:
		return alphanumericUpper
	case AlphaNumeric:
		return alphanumeric
	default:
		return make([]rune, 0)
	}
}
