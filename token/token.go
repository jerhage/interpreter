package token

type TokenType string

type Token struct {
    Type TokenType
    Literal string

}

const (
    // An unknown token or character
    ILLEGAL = "ILLEGAL"

    // Tells parser it can stop
    EOF = "EOF"

    // Indentifiers and literals
    IDENT = "IDENT"
    INT = "INT"

    // Operators
    ASSIGN = "="
    PLUS = "+"

    // Delimiters
    COMMA = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    // Keywords
    FUNCTION = "FUNCTION"
    LET = "LET"
)

// Map of keywords for the programming language
var keywords = map[string]TokenType {
    "fn": FUNCTION,
    "let": LET,
}

// Check if identifier is a keyword, else return as IDENT
func LookUpIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}
