package sum

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/efficientgo/core/errcapture"
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
        -memprofile=output/${ver}.mem.pprof \
        -cpuprofile=output/${ver}.cpu.pprof \
    | tee output/${ver}.txt
*/
func BenchmarkSum(b *testing.B) {
	for _, tcase := range []struct {
		numLines int
	}{
		{numLines: 0},
		{numLines: 1e2},
		{numLines: 1e4},
		{numLines: 1e6},
		{numLines: 2e6},
	} {
		b.Run(fmt.Sprintf("lines-%d", tcase.numLines), func(b *testing.B) {
			b.ReportAllocs()

			fn := lazyCreateTestInput(b, tcase.numLines)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, err := Sum(fn)
				assert.NoError(b, err)
			}
		})
	}
}

func TestBenchSum(t *testing.T) {
	fn := filepath.Join(t.TempDir(), "/test.2M.txt")
	sum, err := createTestInput(fn, 2e6)
	assert.NoError(t, err)
	ret, err := Sum(fn)
	assert.NoError(t, err)
	assert.Equal(t, sum, ret)
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

func lazyCreateTestInput(tb testing.TB, numLines int) (filename string) {
	tb.Helper()

	filename = fmt.Sprintf("testdata/test.%v.txt", numLines)
	_, err := os.Stat(filename)
	if errors.Is(err, os.ErrNotExist) {
		_, err = createTestInput(filename, numLines)
		assert.NoError(tb, err)
	} else {
		assert.NoError(tb, err)
	}

	return filename
}
