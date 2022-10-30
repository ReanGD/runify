package calculator_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/cockroachdb/apd/v3"
	"github.com/stretchr/testify/assert"
)

func assertEqualDecimal(t *testing.T, expected apd.Decimal, actual apd.Decimal, msgAndArgs ...interface{}) {
	t.Helper()
	if expected.Cmp(&actual) != 0 {
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
	dctx    apd.Context
	stats   []int
	cond    apd.Condition
	total   int
	maxSize int
	sumSize int
}

func newTestDataGenerator(seed int64, dctx apd.Context) *testDataGenerator {
	src := rand.NewSource(seed)
	return &testDataGenerator{
		sgen:    src.(rand.Source64),
		gen:     rand.New(src),
		dctx:    dctx,
		cond:    0,
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

func (s *testDataGenerator) next() (string, apd.Decimal, apd.Condition) {
	s.cond = 0
	result, expression := s.genExprAddSub()
	eLen := len(expression)
	s.total++
	s.sumSize += eLen
	if s.maxSize < eLen {
		s.maxSize = eLen
	}
	return expression, result, s.cond
}

func (s *testDataGenerator) genExprAddSub() (apd.Decimal, string) {
	kind := s.getKind()
	if kind >= 700 {
		return s.genExprMulDiv()
	}

	var res apd.Decimal
	left, leftStr := s.genExprMulDiv()
	right, rightStr := s.genExprMulDiv()
	if kind >= 350 {
		if s.cond == 0 {
			cond, _ := s.dctx.Add(&res, &left, &right)
			s.cond = cond & s.dctx.Traps
		}
		s.stats[Add]++
		return res, fmt.Sprintf("%s + %s", leftStr, rightStr)
	}

	if s.cond == 0 {
		cond, _ := s.dctx.Sub(&res, &left, &right)
		s.cond = cond & s.dctx.Traps
	}
	s.stats[Sub]++
	return res, fmt.Sprintf("%s - %s", leftStr, rightStr)
}

func (s *testDataGenerator) genExprMulDiv() (apd.Decimal, string) {
	kind := s.getKind()
	if kind >= 400 {
		return s.genExprUnaryPlusMinus()
	}

	var res apd.Decimal
	left, leftStr := s.genExprUnaryPlusMinus()
	right, rightStr := s.genExprUnaryPlusMinus()
	if kind >= 200 {
		if s.cond == 0 {
			cond, _ := s.dctx.Mul(&res, &left, &right)
			s.cond = cond & s.dctx.Traps
		}
		s.stats[Mul]++
		return res, fmt.Sprintf("%s * %s", leftStr, rightStr)
	}

	if s.cond == 0 {
		cond, _ := s.dctx.Quo(&res, &left, &right)
		s.cond = cond & s.dctx.Traps
	}
	s.stats[Div]++
	return res, fmt.Sprintf("%s / %s", leftStr, rightStr)
}

func (s *testDataGenerator) genExprUnaryPlusMinus() (apd.Decimal, string) {
	kind := s.getKind()
	if kind >= 100 {
		return s.genExprPow()
	}

	var res apd.Decimal
	right, rightStr := s.genExprPow()
	if kind >= 10 {
		if s.cond == 0 {
			cond, _ := s.dctx.Neg(&res, &right)
			s.cond = cond & s.dctx.Traps
		}
		s.stats[UnaryMinus]++
		return res, fmt.Sprintf("-%s", rightStr)
	}

	if s.cond == 0 {
		res = right
	}
	s.stats[UnaryPlus]++
	return res, fmt.Sprintf("+%s", rightStr)
}

func (s *testDataGenerator) genExprPow() (apd.Decimal, string) {
	kind := s.getKind()
	if kind >= 10 {
		return s.genExprBracketsInt()
	}

	var res apd.Decimal
	left, leftStr := s.genExprBracketsInt()
	// TODO: fix pow
	right, rightStr := apd.New(2, 0), "2"
	if s.cond == 0 {
		cond, _ := s.dctx.Pow(&res, &left, right)
		s.cond = cond & s.dctx.Traps
	}

	if kind >= 5 {
		s.stats[Pow1]++
		return res, fmt.Sprintf("%s ^ %s", leftStr, rightStr)
	}

	s.stats[Pow2]++
	return res, fmt.Sprintf("%s ** %s", leftStr, rightStr)
}

func (s *testDataGenerator) genExprBracketsInt() (apd.Decimal, string) {
	kind := s.getKind()
	if kind >= 300 {
		return s.genInt64()
	}

	var res apd.Decimal
	inner, innerStr := s.genExprAddSub()
	if s.cond == 0 {
		res = inner
	}
	s.stats[Brackets]++
	return res, fmt.Sprintf("(%s)", innerStr)
}

func (s *testDataGenerator) genInt64() (apd.Decimal, string) {
	var res apd.Decimal

	rndVal := s.getInt31()
	if s.cond == 0 {
		res = *apd.New(int64(rndVal), 0)
	}
	s.stats[Numbers]++
	return res, fmt.Sprintf("%d", rndVal)
}
