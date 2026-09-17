func lengthOfLongestSubstring(s string) int {
	l, res := 0,0
	ch := make(map[byte]bool)

	for r:=0; r < len(s); r++ {
		for ch[s[r]]{
			delete(ch, s[l])
			l++
		}

		ch[s[r]] = true
		if r - l + 1 > res {
			res = r - l + 1
		}
	}

	return res
}
