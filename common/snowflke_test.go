package common

import "testing"

func BenchmarkSnowFlake_GetID(b *testing.B) {
	s := NewSnowFlake(1)
	for i := 0; i < b.N; i++ {
		s.GetID()
	}
}

func BenchmarkSnowFlake_GetIDParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		s := NewSnowFlake(1)
		for pb.Next() {
			s.GetID()
		}
	})
}
