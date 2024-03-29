package sum_test

import "testing"

func BenchmarkSum(b *testing.B) {
	b.ReportAllocs()

	// 【TODO】初期化処理がある場合は追加する

	// 初期化処理の時間をベンチマークに含めたくない場合はタイマーをリセット
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// todo
	}
}
