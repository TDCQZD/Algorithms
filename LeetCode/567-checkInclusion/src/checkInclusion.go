package src

func checkInclusion(s1 string, s2 string) bool {
	l1 := len(s1)
	l2 := len(s2)
	if l1 > l2 {
		return false
	}

	var c1 [26]int // s1每个字符出现的次数
	var c2 [26]int // s2每个字符出现的次数
	for k, v := range s1 {
		c1[v-'a']++
		c2[s2[k]-'a']++
	}
	for i := 0; i < l2-l1+1; i++ {
		if c1 == c2 {
			return true
		}
		if i < l2-l1 {
			c2[s2[i]-'a']--
			c2[s2[i+l1]-'a']++
		}
	}

	return false
}

func checkInclusion1(s1 string, s2 string) bool {

	winLen := len(s1)           //窗口大小
	s1Map := make(map[byte]int) //用于表示s1中每个字符的出现次数
	m := make(map[byte][]int)   //用于表示s2中指定字符出现的所有位置索引，每个切片都是升序
	for _, v := range s1 {
		s1Map[byte(v)] += 1
	}
	for i, v := range s2 {
		m[byte(v)] = append(m[byte(v)], i)
	}
	//fmt.Println(len(s1Map),len(m))
	for i := 0; i <= len(s2)-winLen; i++ {
		l, r := i, i+winLen-1
		s1MapCurLenth := 0

		for k, v := range s1Map {

			if len(m[k]) < v {
				return false
			} else {
				arr := m[k]
				tl, tr := -1, -1
				for i, v := range arr {
					if v >= l && (i == 0 || arr[i-1] < l) {
						tl = i
					}
					if v <= r && (i == len(arr)-1 || arr[i+1] > r) {
						tr = i
					}
				}
				if tr-tl+1 != v {
					break
				} else {
					s1MapCurLenth++
				}
			}
		}

		if s1MapCurLenth == len(s1Map) {
			return true
		}
	}
	return false
}
