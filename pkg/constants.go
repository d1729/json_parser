package pkg

const JsonQuote = "\""

const TrueLen = len("true")
const FalseLen = len("false")
const NullLen = len("null")
const JsonLeftBracket = "["
const JsonRightBracket = "]"
const JsonLeftBrace = "{"
const JsonRightBrace = "}"
const JsonComma = ","
const JsonColon = ":"

var JsonWhiteSpaces = map[rune]bool{
	' ':  true,
	'\t': true,
	'\n': true,
	'\r': true,
}

var JsonSyntax = map[rune]bool{
	'{': true,
	'}': true,
	'[': true,
	']': true,
	':': true,
	',': true,
}
