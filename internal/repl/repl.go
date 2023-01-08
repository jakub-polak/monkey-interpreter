package repl

import (
	"bufio"
	"fmt"
	"io"

	"monkey-interpreter/internal/lexer"
	"monkey-interpreter/internal/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			b := []byte(fmt.Sprintf("%+v\n", tok))
			_, _ = out.Write(b)
		}
	}
}
