package mathutils

func Factorial(num int) int {
	i := 1
	product := 1
	for i < num+1 {
		product *= i
		i++
	}
	return product
}
