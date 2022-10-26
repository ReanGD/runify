package calculator_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/ReanGD/runify/server/system"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func assertEqualDecimal(t *testing.T, expected decimal.Decimal, actual decimal.Decimal, msgAndArgs ...interface{}) {
	t.Helper()
	if !expected.Equal(actual) {
		assert.Fail(t, fmt.Sprintf("Not equal: \n"+
			"expected: %s\n"+
			"actual  : %s", expected.String(), actual.String()), msgAndArgs...)
		t.FailNow()
	}
}

const (
	Add int = iota
	Sub
	Mul
	Div
	Pow1
	Pow2
	UnaryPlus
	UnaryMinus
	Brackets
	Numbers
	LastOperation
)

var opNames = map[int]string{
	Add:        "Add",
	Sub:        "Sub",
	Mul:        "Mul",
	Div:        "Div",
	Pow1:       "Pow1",
	Pow2:       "Pow2",
	UnaryPlus:  "UnaryPlus",
	UnaryMinus: "UnaryMinus",
	Brackets:   "Brackets",
	Numbers:    "Numbers",
}

type testDataStr struct {
	expression string
	result     string
}

type testDataGenerator struct {
	sgen    rand.Source64
	gen     *rand.Rand
	errCode system.Error
	stats   []int
	total   int
	maxSize int
	sumSize int
}

func newTestDataGenerator(seed int64) *testDataGenerator {
	src := rand.NewSource(seed)
	return &testDataGenerator{
		sgen:    src.(rand.Source64),
		gen:     rand.New(src),
		errCode: system.Success,
		stats:   make([]int, LastOperation),
		total:   0,
		maxSize: 0,
		sumSize: 0,
	}
}

const (
	kindN         = 1000
	rngMax        = 1 << 63
	rngMask       = rngMax - 1
	rngMax31      = 1 << 31
	rngMask31     = rngMax31 - 1
	rndMaxKindGen = uint64(rngMax - 1 - rngMax%uint64(kindN))
)

// Return value in hange [0, kindN)
// Analog of s.gen.Intn(kindN)
func (s *testDataGenerator) getKind() uint64 {
	v := s.sgen.Uint64()
	for v > rndMaxKindGen {
		v = s.sgen.Uint64()
	}
	return v % uint64(kindN)
}

func (s *testDataGenerator) getInt31() int32 {
	return int32(s.sgen.Uint64() & rngMask31)
}

func (s *testDataGenerator) showStats(d time.Duration) {
	total := 0
	for _, cnt := range s.stats {
		total += cnt
	}

	fmt.Printf("Expression: Max size - %d, Avg size - %.2f, Processing time - %.2f mcs\n",
		s.maxSize, float64(s.sumSize)/float64(s.total), float64(d.Microseconds())/float64(s.total))

	for id, cnt := range s.stats {
		fmt.Printf("%s: %d (%.2f%%)\n", opNames[id], cnt, float64(cnt)/float64(total)*100)
	}
}

func (s *testDataGenerator) next() (string, decimal.Decimal, system.Error) {
	s.errCode = system.Success
	result, expression := s.genExprAddSub()
	eLen := len(expression)
	s.total++
	s.sumSize += eLen
	if s.maxSize < eLen {
		s.maxSize = eLen
	}
	return expression, result, s.errCode
}

func (s *testDataGenerator) genExprAddSub() (decimal.Decimal, string) {
	kind := s.getKind()
	if kind >= 700 {
		return s.genExprMulDiv()
	}

	res := decimal.Zero
	left, leftStr := s.genExprMulDiv()
	right, rightStr := s.genExprMulDiv()
	if kind >= 350 {
		if s.errCode == system.Success {
			res = left.Add(right)
		}
		s.stats[Add]++
		return res, fmt.Sprintf("%s + %s", leftStr, rightStr)
	}

	if s.errCode == system.Success {
		res = left.Sub(right)
	}
	s.stats[Sub]++
	return res, fmt.Sprintf("%s - %s", leftStr, rightStr)
}

func (s *testDataGenerator) genExprMulDiv() (decimal.Decimal, string) {
	kind := s.getKind()
	if kind >= 400 {
		return s.genExprUnaryPlusMinus()
	}

	res := decimal.Zero
	left, leftStr := s.genExprUnaryPlusMinus()
	right, rightStr := s.genExprUnaryPlusMinus()
	if kind >= 200 {
		if s.errCode == system.Success {
			res = left.Mul(right)
		}
		s.stats[Mul]++
		return res, fmt.Sprintf("%s * %s", leftStr, rightStr)
	}

	if right.Sign() == 0 {
		s.errCode = system.CalculatorDivisionByZero
	}
	if s.errCode == system.Success {
		res = left.Div(right)
	}
	s.stats[Div]++
	return res, fmt.Sprintf("%s / %s", leftStr, rightStr)
}

func (s *testDataGenerator) genExprUnaryPlusMinus() (decimal.Decimal, string) {
	kind := s.getKind()
	if kind >= 100 {
		return s.genExprPow()
	}

	res := decimal.Zero
	right, rightStr := s.genExprPow()
	if kind >= 10 {
		if s.errCode == system.Success {
			res = right.Neg()
		}
		s.stats[UnaryMinus]++
		return res, fmt.Sprintf("-%s", rightStr)
	}

	if s.errCode == system.Success {
		res = right
	}
	s.stats[UnaryPlus]++
	return res, fmt.Sprintf("+%s", rightStr)
}

func (s *testDataGenerator) genExprPow() (decimal.Decimal, string) {
	kind := s.getKind()
	if kind >= 10 {
		return s.genExprBracketsInt()
	}

	res := decimal.Zero
	left, leftStr := s.genExprBracketsInt()
	// TODO: fix pow
	right, rightStr := decimal.NewFromInt(2), "2"
	if s.errCode == system.Success {
		res = left.Pow(right)
	}

	if kind >= 5 {
		s.stats[Pow1]++
		return res, fmt.Sprintf("%s ^ %s", leftStr, rightStr)
	}

	s.stats[Pow2]++
	return res, fmt.Sprintf("%s ** %s", leftStr, rightStr)
}

func (s *testDataGenerator) genExprBracketsInt() (decimal.Decimal, string) {
	kind := s.getKind()
	if kind >= 300 {
		return s.genInt64()
	}

	res := decimal.Zero
	inner, innerStr := s.genExprAddSub()
	if s.errCode == system.Success {
		res = inner
	}
	s.stats[Brackets]++
	return res, fmt.Sprintf("(%s)", innerStr)
}

func (s *testDataGenerator) genInt64() (decimal.Decimal, string) {
	res := decimal.Zero

	val := decimal.NewFromInt32(s.getInt31())
	if s.errCode == system.Success {
		res = val
	}
	s.stats[Numbers]++
	return res, val.String()
}
