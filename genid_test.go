package hgenid

import (
	"log"
	"testing"
	"time"
)

// 单线程测试
func TestGenId(t *testing.T) {
	var mp = make(map[string]bool)
	for i := 0; i < 10*10000; i++ {
		id := GenId()
		if mp[id] {
			t.Log("出现重复了 i=", i)
			break
		}
		mp[id] = true
		t.Log(id)
		t.Log(time.Now().UnixNano())
	}
}

func TestA(t *testing.T) {

	et := time.Duration(1592304944216)
	t.Log(time.Now().Unix())

	tt := time.Now().Add(et * time.Millisecond)
	log.Println(tt.Format(time.RFC3339))
}

func TestE(t *testing.T) {
	nm := make(map[string]int)
	for i := 0; i < 1000; i++ {
		k := GenId()[16:]
		if nm[k] != 0 {
			t.Log("重复了", i)
			break
		}
		nm[k] = i
	}
}
