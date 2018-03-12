package parse

import "fmt"

type parser struct {
	l        *lexer
	last     item
	rejected bool
}

func (p *parser) scan() item {
	// there is one rejected item
	if p.rejected {
		p.rejected = false
		return p.last
	}

	p.last = p.l.nextItem()
	p.rejected = false
	return p.last
}

func (p *parser) scanIgnoreWhiteSpace() item {
	t := p.scan()
	if t.typ == itemWhiteSpace {
		t = p.scan()
	}
	return t
}

func (p *parser) reject() {
	p.rejected = true
}

func AST(q string) {
	p := &parser{
		l: lex(q),
	}
	fmt.Println(newStatement(p))
}
