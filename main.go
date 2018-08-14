package main

import (
	"strconv"
	"math/big"
	"sort"
	"fmt"
	"math"
)



func main(){
	//largestPrimeFactor()
	//largestPalindromeProduct()
	//smallestDivisible()
	//sumSquaresDifference()
	//findPrime(10001)
	//largestProductinSeries(13)
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

func sumSquaresDifference(){
	a := 0
	b := 0
	for i := 1; i <= 100; i++ {
		a += i*i
		b += i
	}
	b = b*b
	fmt.Printf("%v - %v = ", b,a)
	fmt.Println(b-a)
}

func findPrime(i int){
	x := 2
	for{
		if big.NewInt(int64(x)).ProbablyPrime(0) {
			i--
			if i == 0 {
				fmt.Println(x)
				return
			}
		}
		x++
	}
}

func largestProductinSeries(i int){
	s := "7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450"
	x := []rune(s)
	product := 0
	for a := 0 ; a < len(x) - i ; a++ {
		c, err := strconv.Atoi(string(x[a]))
		if err != nil{
			panic(err)
		}
		for b := 1 ; b < i; b++ {
			d, err := strconv.Atoi(string(x[a + b]))
			if err != nil{
				panic(err)
			}
			c *= d
		}
		if c > product {
			product = c
			fmt.Println(product)
		}
	}
	fmt.Println(product)
}