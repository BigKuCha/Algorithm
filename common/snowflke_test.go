package common

import (
	"testing"
	"time"
)

func TestSnowFlake_GetID(t *testing.T) {
	snow := NewSnowFlake(1)
	m := make(map[int64]bool)
	timeout := time.After(time.Second)
	for {
		select {
		case <-timeout:
			t.Log("1 second")
			goto ForEnd
		default:
			id := snow.GetID()
			if m[id] {
				t.Error("ID重复")
			}
			m[id] = true
		}
	}
ForEnd:
}
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
