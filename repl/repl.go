package repl

import (
	"bufio"
	"fmt"
	"github.com/digital-codex/monkey/evaluator"
	"github.com/digital-codex/monkey/object"
	"github.com/digital-codex/monkey/parser"
	"io"
	"log"
	"os/user"
)

const MONKEY = `
            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   ._/  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`
const PROMPT = ">> "

func Start(in io.Reader, out io.Writer, current *user.User) {
	_, err := io.WriteString(out, MONKEY)
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintf(out, "Hello %s! This is the Monkey programming language!\n", current.Username)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()

	for {
		_, err := fmt.Fprintf(out, PROMPT)
		if err != nil {
			log.Fatal(err)
		}
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		p := parser.New(line)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}
		evaluator.DefineMacros(program, macroEnv)
		expanded := evaluator.ExpandMacros(program, macroEnv)

		evaluated := evaluator.Eval(expanded, env)
		if evaluated != nil {
			_, err := fmt.Fprintf(out, "%s\n", evaluated.Inspect())
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func printParseErrors(out io.Writer, errors []error) {
	_, err := io.WriteString(out, MONKEY)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.WriteString(out, "Whoops! We ran into some monkey business here!\n")
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.WriteString(out, "parser errors:\n")
	if err != nil {
		log.Fatal(err)
	}
	for _, msg := range errors {
		_, err := fmt.Fprintf(out, "\t%s\n", msg)
		if err != nil {
			log.Fatal(err)
		}
	}
}
