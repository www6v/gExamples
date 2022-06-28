package basic

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))
	m3 := make(map[int]int, 10)  /// capacity 10
	t.Logf("len m3=%d", len(m3)) /// len 0
	//t.Logf("len m3=%d", cap(m3))  /// cap(m3) is not available
}

func TestAccessNotExisting(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])

	/// 和其他语言的差异
	/// m1[3] = 0
	if v, ok := m1[3]; ok {
		t.Logf("aa %d", v)
	} else {
		t.Log("key 3 is not existing")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4}
	for k, v := range m1 {
		t.Log(k, v)
	}
}