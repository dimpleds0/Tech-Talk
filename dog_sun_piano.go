package techtalk

import (
	"fmt"
	"math/rand"
)

// TechTalk is a function that 
// takes in a string and prints it
func TechTalk(inputStr string) {
	fmt.Println("Welcome to the Tech Talk -",inputStr)
}

// TimesFunc is a function that 
// takes two integers and runs the 
// provided function those many times
func TimesFunc(from int, to int, fn func()) {
	for i := from; i < to; i++ {
		fn()
	}
}

// Randomize is a function that 
// takes in two integers and 
// randomly selects a number 
// between them
func Randomize(min int, max int) int {
	return rand.Intn(max - min) + min 
}

// Greeting is a function that 
// takes in a string and prints 
// a greeting with the string
func Greeting(name string) {
	fmt.Println("Welcome", name)
}

// Sum is a function that takes 
// in two integers and returns 
// the sum of them
func Sum(num1 int, num2 int) int {
	return num1 + num2
}

// Average is a function that 
// takes in an array of integers 
// and returns the average
func Average(arr []int) float64 {
	var sum int
	for _, val := range arr {
		sum += val
	}

	return float64(sum) / float64(len(arr))
}

// Fibonacci is a function that 
// takes in an integer and returns 
// the corresponding Fibonacci number
func Fibonacci(num int) int {
	if num == 0 {
		return 0	
	} else if num == 1 {
		return 1
	}
	
	return Fibonacci(num-1) + Fibonacci(num-2)
}

// Smallest is a function that 
// takes in an array of integers 
// and returns the smallest one
func Smallest(arr []int) int {
	var smallest = arr[0]
	for _, val := range arr {
		if val < smallest {
			smallest = val
		}
	}
	return smallest 
}

// Largest is a function that 
// takes in an array of integers 
// and returns the largest one
func Largest(arr []int) int {
	var largest = arr[0]
	for _, val := range arr {
		if val > largest {
			largest = val
		}
	}
	return largest 
}

// Reverse is a function that 
// takes in a string and reverses it
func Reverse(str string) string {
	var reversed string
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	return reversed
}

// CountLetters is a function that 
// takes in a string and returns 
// the number of letters in it
func CountLetters(str string) int {
	var count int
	for _,char := range str {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			count++
		}
	}
	return count
}

// Contains is a function that 
// takes in two strings and returns 
// true if the second string is contained in the first
func Contains(str1 string, str2 string) bool {
	return strings.Contains(str1, str2)
}

// ToUpper is a function that 
// takes in a string and returns it in all uppercase
func ToUpper(str string) string {
	return strings.ToUpper(str)
}

// ToLower is a function that 
// takes in a string and returns it in all lowercase
func ToLower(str string) string {
	return strings.ToLower(str)
}

// IsPalindrome is a function that 
// takes in a string and returns true 
// if the string is a palindrome
func IsPalindrome(str string) bool {
	var reversed string
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	return reversed == str
}

// IsPangram is a function that 
// takes in a string and returns true 
// if the string is a pangram
func IsPangram(str string) bool {
	var set = make(map[rune]bool)
	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			set[char] = true
		} else if char >= 'A' && char <= 'Z' {
			char = rune(char - 'A' + 'a')
			set[char] = true
		}
	}
	return len(set) == 26
}

// Capitalize is a function that 
// takes in a string and capitalizes 
// the first letter of each word
func Capitalize(str string) string {
	var words []string
	for _, word := range strings.Split(str, " ") {
		words = append(words, strings.Title(word))
	}
	return strings.Join(words, " ")
}

// IsPrime is a function that 
// takes in an integer and returns 
// true if the number is prime
func IsPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}

// IsEven is a function that 
// takes in an integer and returns 
// true if the number is even
func IsEven(num int) bool {
	return num % 2 == 0
}

// IsOdd is a function that 
// takes in an integer and returns 
// true if the number is odd
func IsOdd(num int) bool {
	return num % 2 != 0
}

// SwapCase is a function that 
// takes in a string and swaps 
// the case of each character
func SwapCase(str string) string {
	var swapped string
	for _, char := range str {
		if char >= 'A' && char <= 'Z' {
			swapped += string(char + 'a' - 'A')
		} else if char >= 'a' && char <= 'z' {
			swapped += string(char - 'a' + 'A')
		} else {
			swapped += string(char)
		}
	}
	return swapped
}

// IndexOf is a function that 
// takes in a string and a character 
// and returns the index of the 
// first occurrence of the character
func IndexOf(str string, char rune) int {
	for i, r := range str {
		if r == char {
			return i
		}
	}
	return -1
}

// LastIndexOf is a function that 
// takes in a string and a character 
// and returns the index of the 
// last occurrence of the character
func LastIndexOf(str string, char rune) int {
	for i := len(str) - 1; i >= 0; i-- {
		if rune(str[i]) == char {
			return i
		}
	}
	return -1
}

// IsAnagram is a function that 
// takes in two strings and returns 
// true if they are anagrams
func IsAnagram(str1 string, str2 string) bool {
	var counts1 [26]int
	var counts2 [26]int
	for _, char := range str1 {
		counts1[char - 'a']++
	}
	for _, char := range str2 {
		counts2[char - 'a']++
	}
	return counts1 == counts2
}

// IsPalindromeNumber is a function 
// that takes in an integer and returns 
// true if the number is a palindrome
func IsPalindromeNumber(num int) bool {
	var reversedNum int
	temp := num
	for temp > 0 {
		reversedNum = reversedNum*10 + temp%10
		temp /= 10
	}
	return reversedNum == num
}

// NumFactors is a function that 
// takes in an integer and returns 
// the number of factors it has
func NumFactors(num int) int {
	var count int
	for i := 1; i <= num; i++ {
		if num % i == 0 {
			count++
		}
	}
	return count
}

// RandomString is a function that 
// takes in an integer and returns 
// a random string of that length
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// Factorial is a function that 
// takes in an integer and returns 
// the factorial of the number
func Factorial(num int) int {
	if num <= 1 {
		return 1
	}
	return num * Factorial(num - 1)
}

// SquareRoot is a function that 
// takes in an integer and returns 
// the square root of the number
func SquareRoot(num int) float64 {
	start := 0.0
	end := float64(num)
	for start < end {
		mid := (start + end) / 2
		if mid*mid == float64(num) {
			return mid 
		} else if mid*mid < float64(num) {
			start = mid 
		} else {
			end = mid
		}
	}
	return start
}

// IsDivisible is a function that 
// takes in two integers and returns 
// true if the first is divisible by the second
func IsDivisible(num1 int, num2 int) bool {
	return num1 % num2 == 0
}

// PrimeFactors is a function that 
// takes in an integer and returns 
// an array of the prime factors
func PrimeFactors(num int) []int {
	var factors []int
	for i := 2; i <= num; i++ {
		for num % i == 0 {
			factors = append(factors, i)
			num /= i
		}
	}
	return factors
}

// HCF is a function that 
// takes in two integers and 
// returns the highest common factor
func HCF(a int, b int) int {
	max := a
	if b > a {
		max = b
	}
	for i := max; i > 0; i-- {
		if a % i == 0 && b % i == 0 {
			return i
		}
	}
	return 1
}

// LCM is a function that 
// takes in two integers and 
// returns the lowest common multiple
func LCM(a int, b int) int {
	min := a
	if b < a {
		min = b
	}
	for i := min; i <= a*b; i++ {
		if i % a == 0 && i % b == 0 {
			return i
		}
	}
	return 0
}

// Trim is a function that 
// takes in a string and returns 
// the string with the whitespace 
// removed from the start and end
func Trim(str string) string {
	return strings.TrimSpace(str)
}

// Search is a function that 
// takes in a string and returns 
// the index of the first occurrence 
// of the provided substring
func Search(str string, substr string) int {
	return strings.Index(str, substr)
}

// StrsToInts is a function that 
// takes in an array of strings 
// and returns an array of integers
func StrsToInts(strs []string) []int {
	ints := make([]int, len(strs))
	for i, str := range strs {
		ints[i], _ = strconv.Atoi(str)
	}
	return ints
}

// GetIndex is a function that 
// takes in an array of integers 
// and returns the index of the 
// first occurrence of the provided
// integer
func GetIndex(arr []int, val int) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}

// Diff is a function that 
// takes in two arrays of integers 
// and returns an array containing 
// the elements in the first array 
// that are not present in the second
func Diff(arr1 []int, arr2 []int) []int {
	diff := make([]int, 0)
	for _, val1 := range arr1 {
		found := false
		for _, val2 := range arr2 {
			if val1 == val2 {
				found = true
				break
			}
		}
		if !found {
			diff = append(diff, val1)
		}
	}
	return diff
}

// Intersect is a function that 
// takes in two arrays of integers 
// and returns an array containing 
// the elements that are present in 
// both the arrays
func Intersect(arr1 []int, arr2 []int) []int {
	inters := make([]int, 0)
	for _, val1 := range arr1 {
		for _, val2 := range arr2 {
			if val1 == val2 {
				inters = append(inters, val1)
				break
			}
		}
	}
	return inters
}

// Shuffle is a function that 
// takes