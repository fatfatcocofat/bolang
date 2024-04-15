package repl

import (
	"bolang/lexer"
	"bolang/parser"
	"os"
	"strings"

	"github.com/chzyer/readline"
)

type Repl struct {
	stdin  *os.File
	stdout *os.File
	stderr *os.File
	rl     *readline.Instance
}

func New(stdin, stdout, stderr *os.File) *Repl {
	cfg := &readline.Config{
		Prompt:            ">> ",
		InterruptPrompt:   "^C",
		EOFPrompt:         " ",
		HistorySearchFold: true,
		Stdin:             stdin,
		Stdout:            stdout,
		Stderr:            stderr,
	}

	rl, err := readline.NewEx(cfg)
	if err != nil {
		panic(err)
	}

	return &Repl{
		stdin:  stdin,
		stdout: stdout,
		stderr: stderr,
		rl:     rl,
	}
}

func (r *Repl) Start() {
	defer r.rl.Close()

	for {
		line, err := r.rl.Readline()
		if err != nil {
			r.stderrWrite("Error reading line: " + err.Error())
			break
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if r.isReplCommand(line) {
			switch line {
			case "!help":
				r.stdoutWrite("BoLang REPL Help\n")
				r.stdoutWrite("  !help - Show this help message\n")
				r.stdoutWrite("  !exit - Exit the REPL\n")
				continue
			case "!exit":
				r.PrintGoodbye()
				return
			default:
				r.stderrWrite("Unknown command: " + line + "\n")
			}
		} else {
			l := lexer.New(line)
			p := parser.New(l)

			program := p.ParseProgram()
			if len(p.Errors()) != 0 {
				r.printParserError(p.Errors())
				continue
			}

			r.stdoutWrite(program.String())
			r.stdoutWrite("\n")

			r.saveHistory(line)
		}
	}
}

func (r *Repl) saveHistory(line string) {
	r.rl.SaveHistory(line)
}

func (r *Repl) stdoutWrite(s string) {
	r.rl.Stdout().Write([]byte(s))
}

func (r *Repl) stderrWrite(s string) {
	r.rl.Stderr().Write([]byte(s))
}

func (r *Repl) printParserError(errors []string) {
	r.stderrWrite("Whoops! Something bad happened:\n")
	for _, e := range errors {
		msg := "  [x] " + e + "\n"
		r.stderrWrite(msg)
	}
}

func (r *Repl) isReplCommand(line string) bool {
	return strings.HasPrefix(line, "!")
}

func (r *Repl) PrintBanner() {
	r.stdoutWrite(r.GetBoLangLogo())
	r.stdoutWrite("\n\n")
	r.stdoutWrite("Welcome to BoLang v0.1.0.\n")
	r.stdoutWrite("Type !help to see the help message.\n")
}

func (r *Repl) PrintGoodbye() {
	r.stdoutWrite("Goodbye! See you later.\n")
}

func (r *Repl) GetBoLangLogo() string {
	return `
   _____    _____       
  /\  __/\  ) ___ (      
  ) )(_ ) )/ /\_/\ \     
 / / __/ // /_/ (_\ \    
 \ \  _\ \\ \ )_/ / /_   
  ) )(__) )\ \/_\/ //_/\ 
  \/____\/  )_____( \_\/ `
}
