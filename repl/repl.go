package repl

import (
	"bufio"
	"fmt"
	"github.com/petrvelicka/amo/lexer"
	"github.com/petrvelicka/amo/token"
	"io"
)

const PS1 = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PS1)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
