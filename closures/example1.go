package closures

func MakeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}


// You need a function that wraps another function so it runs only once, 
// and after that its result is reused (cached).

func Create () func(int) int {
	isCached  := false
	var val int = num
	return func (num int) int {
		if isCached == false {
			return val
		}
		isCached = true
		return val
	}
} 