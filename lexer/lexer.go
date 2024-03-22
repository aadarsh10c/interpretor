package lexer

import "github.com/aadarsh10c/interpretor/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// returns a new lexer
func New(input string) *Lexer {
	l := Lexer{input: input}
	l.readChar()
	return &l
}

// sets the current charachter and also poinst ot the next char
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// reads char until the EOF or whitespace and return the read literal
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string{
	position := l.position
	for isDigit(l.ch){
		l.readChar()
	}
	return l.input[position:l.position]
}
func (l *Lexer) skipWhitespace(){
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r'{
		l.readChar()
	}
}
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

//returns  the new token each time it is called  
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	//skip unncessary whitespace charachters
	l.skipWhitespace()

	//set the token value based upon the current charahcter or group of charachter
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
	default:// check wheteher the token is a identifier/keyword or an unkonwn charachter
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		}else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	//read the next charahcter after setting the token
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
