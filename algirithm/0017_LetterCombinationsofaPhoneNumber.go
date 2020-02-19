package algirithm

func letterCombinations(digits string) []string {

	if digits == "" {
		return nil
	}

	dic := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	result := []string{""}

	//让digits中所有的数字都有机会进行迭代
	for i := 0; i < len(digits); i++ {
		temp := []string{}
		//每一个元素都能添加新字母
		for j := 0; j < len(result); j++ {
			//每个数字所包含的字母都会被添加进去
			for k := 0; k < len(dic[digits[i]]); k++ {
				temp = append(temp, result[j]+dic[digits[i]][k])
			}
		}

		result = temp
	}

	return result
}
