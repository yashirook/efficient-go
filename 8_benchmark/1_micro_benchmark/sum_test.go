package sum

import (
	"testing"

	"github.com/efficientgo/core/testutil"
)

func TestSum(t *testing.T) {
	ret, err := Sum("testdata/input.txt")
	testutil.Ok(t, err)
	testutil.Equals(t, 3110800, ret)
}

func BenchmarkSum(b *testing.B) {
	b.ReportAllocs()

	// 【TODO】初期化処理がある場合は追加する

	// 初期化処理の時間をベンチマークに含めたくない場合はタイマーをリセット
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Sum("testdata/test.2M.txt")
	}
}
