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
    MINUS = "-"
    BANG = "!"
    ASTERISK = "*"
    SLASH = "/"
    GT = ">"
    LT = "<"
    EQ = "=="
    NOT_EQ = "!="

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
    IF = "IF"
    ELSE = "ELSE"
    RETURN = "RETURN"
    TRUE = "TRUE"
    FALSE = "FALSE"

)

// Map of keywords for the programming language
var keywords = map[string]TokenType {
    "fn": FUNCTION,
    "let": LET,
    "if": IF,
    "else": ELSE,
    "return": RETURN,
    "true": TRUE,
    "false": FALSE,

}

// Check if identifier is a keyword, else return as IDENT
func LookUpIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}
