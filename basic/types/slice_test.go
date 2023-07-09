package basic

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	/// 报错  t.Log(s2[0], s2[1], s2[2], s2[3], s2[4])
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar",
		"Aar", "Mar", "Jar",
		"Jar", "Aar", "Sar",
		"Oar", "Nar", "Dar"}

	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))

	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknown"
	t.Log(Q2)
	t.Log(year)
}

func TestSliceCompare(t *testing.T) {
	//a:=[]int{1,2,3,4}
	//b:=[]int{1,2,3,4}
	//if a==b { 报错
	//
	//}
}

// https://www.luozhiyun.com/archives/797
func TestSliceIn100Mistake(t *testing.T) {
	// s := make([]int, 3, 6)
	// s[4] = 0 // panic

	s1 := make([]int, 3, 6)
	s2 := s1[1:3]
	s1[1] = 2
	s2 = append(s2, 3)
	fmt.Println(s1) // [0 2 0]
	fmt.Println(s2) // [2 0 3]

	s1 = append(s1, 4)
	fmt.Println(s1) // [0 2 0 4]
	fmt.Println(s2) // [2 0 4]

	s2 = append(s2, 5, 6, 7)
	fmt.Println(s1) // [0 2 0 4]
	fmt.Println(s2) // [2 0 4 5 6 7]
}
