package db0302

import (
	"errors"
	"fmt"
	"strconv"
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
	p.skipSpaces()
	if !(p.pos+len(kw) <= len(p.buf) && strings.EqualFold(p.buf[p.pos:p.pos+len(kw)], kw)) {
		return false
	}
	if p.pos+len(kw) < len(p.buf) && !isSeparator(p.buf[p.pos+len(kw)]) {
		return false
	}
	p.pos += len(kw)
	return true
}

func (p *Parser) tryName() (string, bool) {
	p.skipSpaces()
	start, cur := p.pos, p.pos
	if !(cur < len(p.buf) && isNameStart(p.buf[cur])) {
		return "", false
	}
	cur++
	for cur < len(p.buf) && isNameContinue(p.buf[cur]) {
		cur++
	}
	p.pos = cur
	return p.buf[start:cur], true
}

func (p *Parser) parseValue(out *Cell) error {
	p.skipSpaces()
	if p.pos >= len(p.buf) {
		return errors.New("expect value")
	}
	ch := p.buf[p.pos]
	if ch == '"' || ch == '\'' {
		return p.parseString(out)
	} else if isDigit(ch) || ch == '-' || ch == '+' {
		return p.parseInt(out)
	} else {
		return errors.New("expect value")
	}
}

func (p *Parser) parseString(out *Cell) error {
	out.Type = TypeStr
	quote := p.buf[p.pos]
	p.pos++
	for {
		if p.isEnd() {
			return fmt.Errorf("missing closing quote")
		}
		if p.buf[p.pos] == '\\' {
			p.pos++
			if !p.isEnd() {
				out.Str = append(out.Str, p.buf[p.pos])
				p.pos++
			} else {
				return fmt.Errorf("missing closing quote")
			}
		} else if p.buf[p.pos] == quote {
			p.pos++
			return nil
		} else {
			out.Str = append(out.Str, p.buf[p.pos])
			p.pos++
		}
	}

	return fmt.Errorf("missing closing quote")
}

func (p *Parser) parseInt(out *Cell) (err error) {
	start := p.pos
	for {
		if p.isEnd() {
			break
		}
		next := p.buf[p.pos+1]
		if isDigit(next) {
			p.pos++
			continue
		}
		p.pos++
		break
	}
	end := p.pos
	digits := p.buf[start:end]
	digitsInt, err := strconv.ParseInt(digits, 10, 64)
	if err != nil {
		return err
	}
	out.Type = TypeI64
	out.I64 = digitsInt
	return nil
}

func (p *Parser) next() byte {
	if !p.isEnd() {
		return p.buf[p.pos+1]
	}
	return 0
}

func (p *Parser) isEnd() bool {
	p.skipSpaces()
	return p.pos >= len(p.buf)
}

// QzBQWVJJOUhU https://trialofcode.org/
