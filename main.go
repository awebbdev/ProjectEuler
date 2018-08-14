package main

import (
	"sort"
	"fmt"
	"math"
)



func main(){
	//largestPrimeFactor()
	//largestPalindromeProduct()
	smallestDivisible()
}

func threeFiveMults(){
	mults := make([]int, 0)

	for i := 1; i < 1000; i++ {
		if math.Mod(float64(i), 3.0) == 0 || math.Mod(float64(i), 5.0) == 0 {
			mults = append(mults, i)
		}
	}
	val := 0
	for _, n := range mults{
		val += n
	}
	fmt.Println(val)
}

func fibonacci(){
	i := 0
	f := fib()
	fibs := make([]int,0)
	for {
		fibs = append(fibs, f())
		if fibs[i] > 4000000{
			fibs = fibs[:i]
			break
		}
		i++
	}
	sum := 0
	for _, fi := range fibs{
		if math.Mod(float64(fi), 2.0) == 0 {
			sum += fi
		}
	}
	fmt.Println(fibs)
	fmt.Println(sum)
}

func fib() func() int {
	a, b := 1, 2
	return func() int{
		a,b = b, a + b
		return a
	}

}

func largestPrimeFactor(){
	primes := make([]int, 0)
	x := 600851475143
	for x%2 == 0 {
		primes = append(primes, 2)
		x /= 2
	}
	for i := 3 ; i*i <= x  ; i = i + 2 {
		for x%i == 0 {
			primes = append(primes, i)
			x /= i
		}
	}

	if x > 2 {
		primes = append(primes, x)
	}
	fmt.Println(primes)
}

func largestPalindromeProduct(){
	pals := make([]int, 0)
	for i := 999; i > 99; i-- {
		for j:= 999; j > 99; j-- {
			if checkPalindrome(i*j){
				pals = append(pals, (i*j))
			}
		}
	}
	sort.Ints(pals)
	fmt.Println(pals[len(pals)-1])
}

func checkPalindrome(num int) bool {
	var remainder, temp int
	reverse := 0
	temp = num
	for{
		remainder = num%10
		reverse = reverse*10 + remainder
		num /= 10
		if num == 0 {
			break
		}
	}
	if temp == reverse {
		return true
	}
	return false
}

func smallestDivisible(){
	var remainder int
	num := 2520
	for{
		for i := 1; i <= 20; i++ {
			remainder += num%i
		}
		if remainder == 0 {
			fmt.Println(num)
			break
		}
		remainder = 0
		num++
	}
}