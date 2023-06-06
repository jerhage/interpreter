package lexer

import "interpreter/token"

type Lexer struct {
    input string
    // points to the ch corresponding to the Lexer's ch byte
    position int
    // always points to the next ch byte
    readPosition int
    ch byte
}

func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar()
    return l
}

// Gives the next ch and advances the position of the Lexer on the input string
func (l *Lexer) readChar() {
    // If EOF, yield ASCII code for "NUL"
    if l.readPosition >= len(l.input) {
        // byte 0 corresponds to NUL
        l.ch = 0
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
    var tok token.Token

    switch l.ch {
    case '=': 
        tok = newToken(token.ASSIGN, l.ch)
    case ';': 
        tok = newToken(token.SEMICOLON, l.ch)
    case '(': 
        tok = newToken(token.LPAREN, l.ch)
    case ')': 
        tok = newToken(token.RPAREN, l.ch)
    case ',': 
        tok = newToken(token.COMMA, l.ch)
    case '+': 
        tok = newToken(token.PLUS, l.ch)
    case '{': 
        tok = newToken(token.LBRACE, l.ch)
    case '}': 
        tok = newToken(token.RBRACE, l.ch)
    case 0:
        tok.Literal = ""
        tok.Type = token.EOF
    }

    l.readChar()
    return tok
}

// Create token from TokenType and a ch
func newToken(tokenType token.TokenType, ch byte) token.Token {
    return token.Token{Type: tokenType, Literal: string(ch)}
}
