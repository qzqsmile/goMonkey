package lexer

import (
	"goMokeney/src/pkg/token"
)

type Lexer struct {
	input        string
	position     int  //current position
	readPosition int  //position after reading position
	ch           byte //
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhiteSpace()
	l.skipCarriage()
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
	default:
		if IsLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIndent(tok.Literal)
			return tok
		} else if IsNumber(l.ch){
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		}else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for IsLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string{
	position  := l.position
	for IsNumber(l.ch){
		l.readChar()
	}
	return l.input[position:l.position]
}

func IsLetter(s byte) bool {
	return ('a' <= s && s <= 'z') || ('A' <= s && s <= 'Z') || (s == '_')
}

func IsNumber(s byte) bool{
	return '0' <= s && s <= '9'
}

func  (l *Lexer) skipWhiteSpace(){
	for l.ch == ' '{
		l.readChar()
	}
}

func  (l *Lexer) skipCarriage(){
	for l.ch == '\n' || l.ch=='\t'{
		l.readChar()
	}
}