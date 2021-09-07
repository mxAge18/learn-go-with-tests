package iteration
const repeatCount = 5
func Repeat(str string, repeatTimes int) (res string) {
	if repeatTimes <= 0 {
		repeatTimes = repeatCount
	}
	for i:=0; i < repeatTimes; i++ {
		res += str
	}
	return
}