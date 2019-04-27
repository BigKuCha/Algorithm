package common

import "testing"

func TestBloomFilter_Has(t *testing.T) {
	testData := []string{"aa", "bb", "cc", "hello"}
	verifyData := []struct {
		in   string
		want bool
	}{{"aa", true}, {"ss", false}}
	bloom := NewDefaultBloomFilter()
	for _, value := range testData {
		bloom.Add(value)
	}

	for _, vv := range verifyData {
		if bloom.Has(vv.in) != vv.want {
			t.Error("过滤失败:", vv.in)
		}
	}
}
