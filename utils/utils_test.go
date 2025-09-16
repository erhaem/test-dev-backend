package utils

import (
	"reflect"
	"testing"
)

func TestSumEven(t *testing.T) {
	cases := []struct {
		nums []int
		wantSum int``
		wantEvens []int
	}{
		{[]int{15, 18, 3, 9, 6, 2, 12, 14}, 52, []int{18, 6, 2, 12, 14}},
		{[]int{1, 3, 5}, 0, []int{}},
		{[]int{2, 4, 6}, 12, []int{2, 4, 6}},
		{[]int{}, 0, []int{}},
	}

	for _, c := range cases {
		gotSum, gotEvens := SumEven(c.nums)
		if gotSum != c.wantSum {
			t.Errorf("SumEven(%v) sum = %d; want %d", c.nums, gotSum, c.wantSum)
		}
		if !reflect.DeepEqual(gotEvens, c.wantEvens) {
			t.Errorf("SumEven(%v) evens = %v; want %v", c.nums, gotEvens, c.wantEvens)
		}
	}
}

func TestIsAnagram(t *testing.T) {
	cases := []struct {
		src, tgt string
		want bool
	}{
		{"kamu", "muka", true},   // simple anagram
		{"otto", "toto", true},   // repeated letters
		{"max", "milo", false},   // different lengths
		{"listen", "silent", true}, // simple anagram
		{"abc", "def", false},    // no overlap
		{"", "", true},           // empty strings anagram
	}

	for _, c := range cases {
		got := IsAnagram(c.src, c.tgt)
		if got != c.want {
			t.Errorf("IsAnagram(%q, %q) = %v; want %v", c.src, c.tgt, got, c.want)
		}
	}
}
