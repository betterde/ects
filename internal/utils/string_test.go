package utils

import (
	"testing"
)

func TestRandom(t *testing.T) {
	str := Random(64)
	if len(str) != 64 {
		t.Errorf("随机生成字符串长度不匹配")
	}
}
