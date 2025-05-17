package main

func main() {
	findAnagrams("bpaa", "aa")

}

func findAnagrams(s string, p string) []int {
	result := []int{}

	if len(s) < len(p) {
		return result
	}

	pMap := make(map[uint8]int)
	for _, b := range p {
		pMap[uint8(b)]++
	}
	for i := 0; i < len(p); i++ {
		pMap[s[i]]--
	}

	differ := 0
	for _, v := range pMap {
		if v != 0 {
			differ++
		}
	}
	if differ == 0 {
		result = append(result, 0)
	}

	left := len(p)
	for left < len(s) {

		pMap[s[left-len(p)]]++
		if pMap[s[left-len(p)]] == 0 {
			differ--
		} else if pMap[s[left-len(p)]] == 1 {
			differ++
		}
		pMap[s[left]]--
		if pMap[s[left]] == 0 {
			differ--
		} else if pMap[s[left]] == -1 {
			differ++
		}
		left++
		if differ == 0 {
			result = append(result, left-len(p))
		}
	}

	return result
}
