package sum

import (
	"testing"

	"github.com/efficientgo/core/testutil"
	"github.com/stretchr/testify/assert"
)

func BenchmarkSum(b *testing.B) {
	benchmarkSum(testutil.NewTB(b))
}
func TestBenchSum(t *testing.T) {
	benchmarkSum(testutil.NewTB(t))
}

func benchmarkSum(tb testutil.TB) {
	for i := 0; i < tb.N(); i++ {
		ret, err := Sum("testdata/test.2M.txt")
		assert.NoError(tb, err)
		if !tb.IsBenchmark() {
			assert.Equal(tb, int64(32773481347), ret)
		}
	}
}
