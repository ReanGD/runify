package syn_select_test

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type SelectSuite struct {
	suite.Suite
}

func (s *SelectSuite) SetupSuite() {
}

func (s *SelectSuite) TearDownSuite() {
}

func (s *SelectSuite) TestError() {
	chInt := make(chan int, 10)
	chStr := make(chan string, 10)
	cntCh := 2

	cases := make([]reflect.SelectCase, cntCh)
	cases[0] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(chInt)}
	cases[1] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(chStr)}

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(100 * time.Millisecond)
			chInt <- i
		}
		close(chInt)
	}()

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(100 * time.Millisecond)
			chStr <- "str" + strconv.Itoa(i)
		}
		close(chStr)
	}()

	for i := 0; i < 9; i++ {
		chosen, recv, recvOk := reflect.Select(cases)
		if !recvOk {
			fmt.Println("recv not ok", chosen)
			cases = append(cases[:chosen], cases[chosen+1:]...)
			continue
		}

		switch chosen {
		case 0:
			fmt.Println("chosen Int", recv.Int())
		case 1:
			fmt.Println("chosen Str", recv.String())
		}
	}

	s.Require().Equal(1, 1)
}

func TestSelectSuite(t *testing.T) {
	suite.Run(t, new(SelectSuite))
}
