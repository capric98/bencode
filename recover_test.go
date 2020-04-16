package bencode

import "testing"

func BenchmarkRecover(b *testing.B) {
	ob := []byte{0}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		func() {
			// defer func() { _ = recover() }()
			if len(ob) > 1 {
				ob[1] = 1
			}
		}()
	}
}
