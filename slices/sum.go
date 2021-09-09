package slices


func Sum(nums []int) (res int) {
	// for i:= 0; i < len; i++ {}
	for _, val := range(nums) {
		res += val
	}
	return
}

func SumAll(numsToSum ...[]int) (sumRes []int) {
	// 切片有容量的概念。如果你有一个容量为 2 的切片，但使用 mySlice[10]=1 进行赋值，会报运行时错误。

	// 方法1 ： 先根据numsToSum make切片长度
	// len := len(numsToSum)
	// sumRes = make([]int, len)
	// for i, val := range(numsToSum) {
	// 	sumRes[i] = Sum(val)
	// } 

	// 方法2 ： append
	for _, val := range(numsToSum) {
		sumRes = append(sumRes, Sum(val))
	}
	return 
}

// 把每个切片的尾部元素相加（尾部的意思就是除去第一个元素以外的其他元素）。
func SumAllTails(tailsToSum ...[]int) (res []int) {
	for _, numbers := range(tailsToSum) {
		if len(numbers) == 0 {
			res = append(res, 0)
		} else {
			tail := numbers[1:]
			res = append(res, Sum(tail))
		}

	}
	return
}