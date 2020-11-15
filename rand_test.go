package hgenid

import "testing"

func TestRand32(t *testing.T) {
	t.Log(Rand32())
	t.Log(Rand32())
}
