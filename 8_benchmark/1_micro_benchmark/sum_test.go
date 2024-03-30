package sum

import (
	"testing"

	"github.com/efficientgo/core/testutil"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	ret, err := Sum("testdata/input.txt")
	assert.NoError(t, err)
	assert.Equal(t, int64(3110800), ret)
}

func BenchmarkSum(b *testing.B) {
	b.ReportAllocs()

	// 【TODO】初期化処理がある場合は追加する

	// 初期化処理の時間をベンチマークに含めたくない場合はタイマーをリセット
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Sum("testdata/test.2M.txt")
		assert.NoError(b, err)
	}
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
