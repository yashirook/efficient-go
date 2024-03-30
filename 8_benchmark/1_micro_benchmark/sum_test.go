package sum

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/efficientgo/core/errcapture"
	"github.com/efficientgo/core/testutil"
	"github.com/efficientgo/examples/pkg/sum/sumtestutil"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

// `Sum`関数をベンチマークする。
// 推奨の実行オプションは以下
/*
export ver=v1 && \
    go test -run '^$' -bench '^BenchmarkSum$' -benchtime 10s -count 6 \
        -cpu 4 \
        -benchmem \
        -memprofile=${ver}.mem.pprof \
        -cpuprofile=${ver}.cpu.pprof \
    | tee ${ver}.txt
*/
func BenchmarkSum(b *testing.B) {
	benchmarkSum(testutil.NewTB(b))
}
func TestBenchSum(t *testing.T) {
	benchmarkSum(testutil.NewTB(t))
}

func benchmarkSum(tb testutil.TB) {
	fn := filepath.Join(tb.TempDir(), "/test.2M.txt")
	sum, err := createTestInput(fn, 2e6)
	assert.NoError(tb, err)

	tb.ResetTimer()
	for i := 0; i < tb.N(); i++ {
		ret, err := Sum(fn)
		assert.NoError(tb, err)
		if !tb.IsBenchmark() {
			assert.Equal(tb, sum, ret)
		}
	}
}

func createTestInput(fn string, numLen int) (sum int64, err error) {
	return createTestInputWithExpectedResult(fn, numLen)
}

func createTestInputWithExpectedResult(fn string, numLen int) (sum int64, err error) {
	err = os.MkdirAll(filepath.Dir(fn), os.ModePerm)
	if err != nil {
		return 0, err
	}

	f, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return 0, errors.Wrap(err, "open")
	}
	defer errcapture.Do(&err, f.Close, "close file")

	return sumtestutil.CreateTestInputWithExpectedResult(f, numLen)
}
