package de

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const (
	FILE0 = "/home/user/file0"
	FILE1 = "/home/user/file1"
	URL0  = "http://example0.com"
	URL1  = "http://example1.com"
)

type LaunchSuite struct {
	suite.Suite

	urls         []string
	files        []string
	needTerminal bool
	terminalPath string
}

func (s *LaunchSuite) SetupTest() {
	s.urls = []string{}
	s.files = []string{}
	s.needTerminal = false
	s.terminalPath = ""
}

func (s *LaunchSuite) check(execStr string, expected []string) {
	t := s.T()

	df := &desktopFile{
		exec:     execStr,
		terminal: s.needTerminal,
	}
	actual, err := df.buildLaunchArgs(s.terminalPath, s.urls, s.files)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func (s *LaunchSuite) TestExecStr() {
	s.needTerminal = false
	s.terminalPath = "bash"

	s.check(`vim`, []string{"vim"})
	s.check(`vim test`, []string{"vim", "test"})

	s.files = []string{FILE0}
	s.check(`vim`, []string{"vim"})
	s.check(`vim test`, []string{"vim", "test"})

	s.files = []string{FILE0, FILE1}
	s.check(`vim`, []string{"vim"})
	s.check(`vim test`, []string{"vim", "test"})

	s.files = []string{}
	s.urls = []string{URL0}
	s.check(`vim`, []string{"vim"})
	s.check(`vim test`, []string{"vim", "test"})

	s.urls = []string{URL0, URL1}
	s.check(`vim`, []string{"vim"})
	s.check(`vim test`, []string{"vim", "test"})
}

func (s *LaunchSuite) TestExecStrWithFileParam() {
	s.needTerminal = false
	s.terminalPath = "bash"

	s.check(`vim %f`, []string{"vim"})
	s.check(`vim test %f`, []string{"vim", "test"})

	s.files = []string{FILE0}
	s.check(`vim %f`, []string{"vim", FILE0})
	s.check(`vim test %f`, []string{"vim", "test", FILE0})

	s.files = []string{FILE0, FILE1}
	s.check(`vim %f`, []string{"vim", FILE0})
	s.check(`vim test %f`, []string{"vim", "test", FILE0})

	s.files = []string{}
	s.urls = []string{URL0}
	s.check(`vim %f`, []string{"vim"})
	s.check(`vim test %f`, []string{"vim", "test"})

	s.urls = []string{URL0, URL1}
	s.check(`vim %f`, []string{"vim"})
	s.check(`vim test %f`, []string{"vim", "test"})
}

func (s *LaunchSuite) TestExecStrWithFilesParam() {
	s.needTerminal = false
	s.terminalPath = "bash"

	s.check(`vim %F`, []string{"vim"})
	s.check(`vim test %F`, []string{"vim", "test"})

	s.files = []string{FILE0}
	s.check(`vim %F`, []string{"vim", FILE0})
	s.check(`vim test %F`, []string{"vim", "test", FILE0})

	s.files = []string{FILE0, FILE1}
	s.check(`vim %F`, []string{"vim", FILE0, FILE1})
	s.check(`vim test %F`, []string{"vim", "test", FILE0, FILE1})

	s.files = []string{}
	s.urls = []string{URL0}
	s.check(`vim %F`, []string{"vim"})
	s.check(`vim test %F`, []string{"vim", "test"})

	s.urls = []string{URL0, URL1}
	s.check(`vim %F`, []string{"vim"})
	s.check(`vim test %F`, []string{"vim", "test"})
}

func (s *LaunchSuite) TestExecStrWithUrlParam() {
	s.needTerminal = false
	s.terminalPath = "bash"

	s.check(`vim %u`, []string{"vim"})
	s.check(`vim test %u`, []string{"vim", "test"})

	s.files = []string{FILE0}
	s.check(`vim %u`, []string{"vim"})
	s.check(`vim test %u`, []string{"vim", "test"})

	s.files = []string{FILE0, FILE1}
	s.check(`vim %u`, []string{"vim"})
	s.check(`vim test %u`, []string{"vim", "test"})

	s.files = []string{}
	s.urls = []string{URL0}
	s.check(`vim %u`, []string{"vim", URL0})
	s.check(`vim test %u`, []string{"vim", "test", URL0})

	s.urls = []string{URL0, URL1}
	s.check(`vim %u`, []string{"vim", URL0})
	s.check(`vim test %u`, []string{"vim", "test", URL0})
}

func (s *LaunchSuite) TestExecStrWithUrlsParam() {
	s.needTerminal = false
	s.terminalPath = "bash"

	s.check(`vim %U`, []string{"vim"})
	s.check(`vim test %U`, []string{"vim", "test"})

	s.files = []string{FILE0}
	s.check(`vim %U`, []string{"vim"})
	s.check(`vim test %U`, []string{"vim", "test"})

	s.files = []string{FILE0, FILE1}
	s.check(`vim %U`, []string{"vim"})
	s.check(`vim test %U`, []string{"vim", "test"})

	s.files = []string{}
	s.urls = []string{URL0}
	s.check(`vim %U`, []string{"vim", URL0})
	s.check(`vim test %U`, []string{"vim", "test", URL0})

	s.urls = []string{URL0, URL1}
	s.check(`vim %U`, []string{"vim", URL0, URL1})
	s.check(`vim test %U`, []string{"vim", "test", URL0, URL1})
}

func (s *LaunchSuite) TestNonValidFieldCodes() {
	s.needTerminal = false
	s.terminalPath = "bash"

	s.check(`vim test %x`, []string{"vim", "test", "%x"})
	s.check(`vim %x test`, []string{"vim", "%x", "test"})
	s.check(`vim %x/.vimrc`, []string{"vim", "%x/.vimrc"})

	s.files = []string{FILE0}
	s.check(`vim test %x`, []string{"vim", "test", "%x"})
	s.check(`vim %x test`, []string{"vim", "%x", "test"})
	s.check(`vim %x/.vimrc`, []string{"vim", "%x/.vimrc"})

	s.files = []string{FILE0, FILE1}
	s.check(`vim test %x`, []string{"vim", "test", "%x"})
	s.check(`vim %x test`, []string{"vim", "%x", "test"})
	s.check(`vim %x/.vimrc`, []string{"vim", "%x/.vimrc"})

	s.files = []string{}
	s.urls = []string{URL0}
	s.check(`vim test %x`, []string{"vim", "test", "%x"})
	s.check(`vim %x test`, []string{"vim", "%x", "test"})
	s.check(`vim %x/.vimrc`, []string{"vim", "%x/.vimrc"})

	s.urls = []string{URL0, URL1}
	s.check(`vim test %x`, []string{"vim", "test", "%x"})
	s.check(`vim %x test`, []string{"vim", "%x", "test"})
	s.check(`vim %x/.vimrc`, []string{"vim", "%x/.vimrc"})
}

func (s *LaunchSuite) TestQuotes() {
	s.needTerminal = false
	s.terminalPath = "bash"

	// TODO: https://specifications.freedesktop.org/desktop-entry-spec/latest/ar01s07.html
	s.check(`"gvim" test`, []string{"gvim", "test"})
	s.check(`"gvim test"`, []string{"gvim test"})
	s.check(`vim ~/.vimrc`, []string{"vim", "~/.vimrc"})
	s.check(`vim '~/.vimrc test'`, []string{"vim", "~/.vimrc test"})
	s.check(`vim '~/.vimrc " test'`, []string{"vim", `~/.vimrc " test`})
}

func (s *LaunchSuite) TestEscapeSequences() {
	s.needTerminal = false
	s.terminalPath = "bash"

	s.check(`"gvim test" test2 "test \\" 3"`, []string{"gvim test", "test2", `test " 3`})
	s.check(`"test \\\\\\\\ \\" moin" test`, []string{`test \\ " moin`, "test"})
	s.check(`"gvim \\\\\\\\ \\`+"`"+`test\\$"`, []string{`gvim \\ ` + "`" + `test$`})
}

func (s *LaunchSuite) TestTerminal() {
	s.terminalPath = "bash"

	s.needTerminal = true
	s.check(`vim ~/.vimrc`, []string{"bash", "-e", "vim", "~/.vimrc"})

	s.needTerminal = false
	s.check(`sh -c 'vim ~/.vimrc " test'`, []string{"sh", "-c", `vim ~/.vimrc " test`})
	s.check(`sh -c 'vim ~/.vimrc " test"'`, []string{"sh", "-c", `vim ~/.vimrc " test"`})
}

func (s *LaunchSuite) TestEscapeValidFieldCodes() {
	s.needTerminal = false
	s.terminalPath = "bash"

	s.check(`vim "\\%u" ~/.vimrc`, []string{"vim", "%u", "~/.vimrc"})
}

func TestLaunchSuite(t *testing.T) {
	suite.Run(t, new(LaunchSuite))
}
