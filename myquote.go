package myquote

import "rsc.io/quote"

func TestQuote1() string {
	return quote.Hello() + "\n" + quote.Glass() + "\n" + quote.Opt() + "\n" + quote.Go()
}
