package db0301

import (
	"strings"
)

type Parser struct {
	buf string
	pos int
}

func NewParser(s string) Parser {
	return Parser{buf: s, pos: 0}
}

func isSpace(ch byte) bool {
	switch ch {
	case '\t', '\n', '\v', '\f', '\r', ' ':
		return true
	}
	return false
}
func isAlpha(ch byte) bool {
	return 'a' <= (ch|32) && (ch|32) <= 'z'
}
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
func isNameStart(ch byte) bool {
	return isAlpha(ch) || ch == '_'
}
func isNameContinue(ch byte) bool {
	return isAlpha(ch) || isDigit(ch) || ch == '_'
}
func isSeparator(ch byte) bool {
	return ch < 128 && !isNameContinue(ch)
}

func (p *Parser) skipSpaces() {
	for p.pos < len(p.buf) && isSpace(p.buf[p.pos]) {
		p.pos += 1
	}
}

func (p *Parser) tryKeyword(kw string) bool {
	lastPos := p.pos
	name, ok := p.tryName()
	if !ok {
		return false
	}

	if strings.ToLower(name) == strings.ToLower(kw) {
		return true
	} else {
		p.pos = lastPos
	}

	return false
}

func (p *Parser) tryName() (string, bool) {
	for {
		if p.isEnd() {
			break
		}
		char := p.buf[p.pos]
		if isSpace(char) {
			p.skipSpaces()
			continue
		}
		if !isNameStart(char) {
			break
		}
		start := p.pos
		p.pos++
		for p.pos < len(p.buf) && isNameContinue(p.buf[p.pos]) {
			p.pos++
		}
		end := p.pos
		name := p.buf[start:end]
		return name, true
	}
	return "", false
}

func (p *Parser) isEnd() bool {
	p.skipSpaces()
	return p.pos >= len(p.buf)
}

// QzBQWVJJOUhU https://trialofcode.org/
