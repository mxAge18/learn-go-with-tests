package letterCombinationsDFS


func LetterCombinations(digits string) []string {
	 // 空则返回空结果
	 if len(digits) == 0 {
        return []string{}
    }
	// 外部定义电话号和string的映射关系，用map[string][string]
	var numsMap map[string]string = map[string]string {
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	// 定义要返回的结果
	var res []string
	res = []string{}
	str := ""
    index := 0
    dfs(numsMap, &res, digits, index, str)
    return res
}

// res 需要传二级指针，因为发生扩容后，函数里的slice新生成的一个数组，
// 所以这里slice的地址就变化了，原函数的指针指的数组还是原数组
func dfs(numsMap map[string]string, res *[]string, digits string, index int, str string) {
	if len(digits) == index {
		*res = append(*res, str)
	} else {
		digit := string(digits[index])
        letters := numsMap[digit] // 对应index的所有字符
		// 遍历字符，连上之前的str后向下一步递归调用dfs
        for i:= 0; i < len(letters); i++ {
            dfs(numsMap, res, digits, index+1, str + string(letters[i]))
        }
	}
}