package main

type Calculator struct {
	result int
}

func (c Calculator) Add (num int) {
	c.result = c.result + num
}
func (c Calculator) Sutract (num int) {
	c.result = c.result - num	
}

func (c Calculator) GetResult() int {
	return c.result
}