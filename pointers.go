package main

func Double (n int) int {
	n = n * 2;
	return n;
}

func DoubleUsingPointer (n *int)  {
	*n = *n * 2
}