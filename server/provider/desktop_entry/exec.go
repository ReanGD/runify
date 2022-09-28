package desktop_entry

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func buildArgs(agrsStr string, needTerminal bool, terminal string) ([]string, error) {
	arg := []rune{}
	res := []string{}
	fieldCodeInd := -1
	inEscape := false
	inSingleQuote := false
	inDoubleQuote := false

	if needTerminal {
		res = append(res, terminal, "-e")
	}

	for ind, c := range strings.Replace(agrsStr, "\\\\", "\\", -1) {
		if inEscape {
			inEscape = false
			arg = append(arg, c)
			continue
		}

		switch c {
		case 'u', 'U', 'f', 'F':
			if fieldCodeInd == ind {
				// TODO ignore field codes for the moment
				// strip %-char at the end of the argument
				if len(arg) > 0 {
					arg = arg[:len(arg)-1]
				}
				continue
			}
		case '"':
			if inDoubleQuote {
				inDoubleQuote = false
				res = append(res, string(arg))
				arg = arg[:0]
				continue
			}
			if !inSingleQuote {
				inDoubleQuote = true
				continue
			}

		case '\'':
			if inSingleQuote {
				inSingleQuote = false
				res = append(res, string(arg))
				arg = arg[:0]
				continue
			}
			if !inDoubleQuote {
				inSingleQuote = true
				continue
			}

		case '\\':
			if inDoubleQuote {
				inEscape = true
				continue
			}

		case '%':
			if !(inDoubleQuote || inSingleQuote) {
				fieldCodeInd = ind + 1
			}

		case ' ':
			if !(inDoubleQuote || inSingleQuote) {
				if len(arg) != 0 {
					res = append(res, string(arg))
					arg = arg[:0]
				}
				continue
			}
		}

		arg = append(arg, c)
	}

	if len(arg) != 0 {
		if !(inEscape || inDoubleQuote || inSingleQuote) {
			res = append(res, string(arg))
		} else {
			return nil, fmt.Errorf("Exec value contains an unbalanced number of quote characters: %s", agrsStr)
		}
	}

	return res, nil
}

func execCmd(agrsStr string, needTerminal bool, terminal string) error {
	args, err := buildArgs(agrsStr, needTerminal, terminal)
	if err != nil {
		return err
	}
	if len(args) == 0 {
		return errors.New("empty exec string")
	}

	name := args[0]
	var cmd *exec.Cmd
	if len(args) == 1 {
		cmd = exec.Command(name)
	} else {
		cmd = exec.Command(name, args[1:]...)
	}

	return cmd.Start()
}
