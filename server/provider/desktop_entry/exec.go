package desktop_entry

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"syscall"

	"github.com/ReanGD/runify/server/paths"
)

type ExecParams struct {
	urls         []string
	files        []string
	needTerminal bool
	terminal     string
}

func NewExecParams(needTerminal bool, terminal string) *ExecParams {
	return &ExecParams{
		urls:         []string{},
		files:        []string{},
		needTerminal: needTerminal,
		terminal:     terminal,
	}
}

func (p *ExecParams) SetUrls(urls ...string) *ExecParams {
	p.urls = urls

	return p
}

func (p *ExecParams) SetFiles(files ...string) *ExecParams {
	p.files = files

	return p
}

func fillFieldCodes(code rune, urls []string, files []string) []string {
	switch code {
	case 'u':
		if len(urls) > 0 {
			return []string{urls[0]}
		}
	case 'U':
		return urls
	case 'f':
		if len(files) > 0 {
			return []string{files[0]}
		}
	case 'F':
		if len(files) > 0 {
			return files
		}
	}

	return []string{}
}

func buildArgs(agrsStr string, ep *ExecParams) ([]string, error) {
	arg := []rune{}
	res := []string{}
	fieldCodeInd := -1
	inEscape := false
	inSingleQuote := false
	inDoubleQuote := false

	if ep.needTerminal {
		res = append(res, ep.terminal, "-e")
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
				if len(arg) > 0 {
					arg = arg[:len(arg)-1]
					res = append(res, fillFieldCodes(c, ep.urls, ep.files)...)
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

func execCmd(agrsStr string, ep *ExecParams) error {
	args, err := buildArgs(agrsStr, ep)
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

	// If the parent process does not exit correctly, then all child processes will also be killed
	// This code cancel this behavior
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true, Pgid: 0}
	cmd.Dir = paths.GetUserHome()

	if err = cmd.Start(); err != nil {
		return err
	}

	go cmd.Wait()

	return nil
}
