package utils

// sum even numbers and also return the evens
func SumEven(nums []int) (int, []int) {
	sum := 0
	evens := []int{}
	for _, n := range nums {
		if n%2 == 0 {
			sum += n
			evens = append(evens, n)
		}
	}
	return sum, evens
}

// check if two strings anagram
// ref: https://golangtutorial.dev/problems/string-anagram-go/
func IsAnagram(src, tgt string) bool {
	if len(src) != len(tgt) {
		return false
	}

	var cnt [26]int
	for i := 0; i < len(src); i++ {
		cnt[src[i]-'a']++
		cnt[tgt[i]-'a']--
	}

	for _, v := range cnt {
		if v != 0 {
			return false
		}
	}
	return true
}
