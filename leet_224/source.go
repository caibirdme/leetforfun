package leet_224

import (
	"strconv"
)

func calculate(s string) int {
	interpreter := &Interpreter{
		TokenStreamer: NewTokenStream(s),
	}
	defer func() {
		if r := recover(); nil != r {
		}
	}()
	return interpreter.Expr()
}

type TokenStream struct {
	pos    int
	endPos int
	source []byte
	length int
}

func NewTokenStream(text string) *TokenStream {
	return &TokenStream{
		source: []byte(text),
		length: len(text),
	}
}

func (t *TokenStream) ProbeNextToken() Token {
	for t.pos < t.length && t.source[t.pos] == ' ' {
		t.pos++
	}
	if t.pos >= t.length {
		return eofToken
	}
	switch t.source[t.pos] {
	case '*':
		t.endPos = t.pos + 1
		return mulToken
	case '/':
		t.endPos = t.pos + 1
		return divToken
	case '+':
		t.endPos = t.pos + 1
		return addToken
	case '-':
		t.endPos = t.pos + 1
		return decToken
	case '(':
		t.endPos = t.pos + 1
		return lpToken
	case ')':
		t.endPos = t.pos + 1
		return rpToken
	default:
		for t.endPos = t.pos; t.endPos < t.length && isNum(t.source[t.endPos]); t.endPos++ {
		}
		return Token(string(t.source[t.pos:t.endPos]))
	}
}

func (t *TokenStream) Confirm() {
	t.pos = t.endPos
}

func isNum(b byte) bool {
	return b >= '0' && b <= '9'
}

type Token string

func (t Token) Atoi() int {
	p, _ := strconv.Atoi(string(t))
	return p
}

var (
	mulToken Token = "*"
	divToken Token = "/"
	addToken Token = "+"
	decToken Token = "-"
	lpToken  Token = "("
	rpToken  Token = ")"
	eofToken Token = ""
)

type TokenStreamer interface {
	ProbeNextToken() Token
	Confirm()
}

type Interpreter struct {
	TokenStreamer
}

// expr: term (('+'|'-') term)*
func (i *Interpreter) Expr() int {
	result := i.term()
	token := i.ProbeNextToken()
	for token == addToken || token == decToken {
		i.Confirm()
		switch token {
		case addToken:
			result += i.term()
		case decToken:
			result -= i.term()
		}
		token = i.ProbeNextToken()
	}
	return result
}

// term: factor (('*'|'/') factor)*
func (i *Interpreter) term() int {
	result := i.factor()
	token := i.ProbeNextToken()
	for token == mulToken || token == divToken {
		i.Confirm()
		switch token {
		case mulToken:
			result *= i.factor()
		case decToken:
			result /= i.factor()
		}
		token = i.ProbeNextToken()
	}
	return result
}

// factor: INT | '(' expr ')'
func (i *Interpreter) factor() int {
	token := i.ProbeNextToken()
	if isNumberToken(token) {
		i.Confirm()
		return token.Atoi()
	}
	if token == lpToken {
		i.Confirm()
		result := i.Expr()
		if i.ProbeNextToken() == rpToken {
			i.Confirm()
		} else {
			panic("need )")
		}
		return result
	}
	panic("wtf")
}

func isNumberToken(token Token) bool {
	switch token {
	case mulToken, divToken, addToken, decToken, lpToken, rpToken, eofToken:
		return false
	default:
		return true
	}
}
