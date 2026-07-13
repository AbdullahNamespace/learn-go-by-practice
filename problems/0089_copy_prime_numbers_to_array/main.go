// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 89: copy_prime_numbers_to_array.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Arrays & Math | المصفوفات والرياضيات
// **Difficulty:** Medium | متوسط
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that fills a source array with random numbers between 1 and
// 100, then copies only the prime numbers into a new destination slice. Unlike
// a naive check up to N-1, the primality test should only check divisors up to
// round(N/2) for a small efficiency gain, and should return a custom enum type
// (Prime / NotPrime) instead of a plain boolean. Finally, print both the source
// array and the resulting array of prime numbers.
//
// العربية:
// اكتب برنامجًا يملأ مصفوفة مصدر بأرقام عشوائية بين 1 و 100، ثم ينسخ فقط
// الأعداد الأولية إلى شريحة وجهة جديدة. على عكس الفحص الساذج حتى N-1، يجب أن
// يفحص اختبار الأولية القواسم فقط حتى round(N/2) لتحسين بسيط في الأداء، ويجب
// أن يُرجع نوعًا مخصصًا (Prime / NotPrime) بدلاً من قيمة منطقية بسيطة. أخيرًا،
// اطبع المصفوفة المصدرية والمصفوفة الناتجة التي تحتوي على الأعداد الأولية.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Array 1 = [10 7 4 11 6 13]
// Output: Array 2 Prime numbers: 7 11 13
//
// Example 2:
// Input:  Array 1 = [4 6 8 9]
// Output: Array 2 Prime numbers: (empty)
//   Why:    None of the source elements are prime
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Array elements are positive integers (1 to 100) | عناصر المصفوفة أعداد صحيحة موجبة
// • Divisor check should stop at round(number/2) | فحص القواسم يجب أن يتوقف عند round(العدد/2)
// • Use a custom enum type for Prime/NotPrime status | استخدم نوعًا مخصصًا لحالة أولي/غير أولي
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// type PrimeStatus int
// const (Prime PrimeStatus = iota + 1; NotPrime)
// func checkPrime(number int) PrimeStatus
// func randomNumber(from, to int) int
// func fillArrayWithRandomNumbers(arr []int)
// func copyPrimeNumbers(arrSource []int) []int
// func printArray(arr []int)
//
// ────────────────────────────────────────────────────────────────────────────

package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

// ======================
//   UTILITY
// ======================

type PrimeStatus int

const (
	Prime PrimeStatus = iota + 1
	NotPrime
)

func printError(prompt string) {
	fmt.Printf("X Error : %s\n", prompt)
}

// ======================
//     INPUT FUNCTIONS
// ======================

func readString(prompt string) string {
	fmt.Printf("%s : ", prompt)

	var input string

	fmt.Scan(&input)

	return input
}

func readNumber(prompt string) int {
	for {
		input := readString(prompt)

		inputNumber, err := strconv.Atoi(input)

		if err != nil {
			printError("Invalid input please enter a valid number!")
			continue
		}

		return inputNumber
	}
}

func readPositiveNumber(prompt string) int {
	for {
		input := readNumber(prompt)

		if input <= 0 {
			printError("Invalid input please enter a positive number!")
			continue
		}

		return input
	}
}

func fillArrayWithRandomNumbers(lengthOfArray int) []int {
	arr := make([]int, lengthOfArray)

	for key := range lengthOfArray {
		arr[key] = randomNumber(1, 100)
	}

	return arr
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func checkPrime(number int) PrimeStatus {
	if number <= 1 {
		return NotPrime
	}
	if number == 2 {
		return Prime
	}
	if number%2 == 0 {
		return NotPrime
	}
	for i := 3; i <= number/2; i = i + 2 {
		if number%i == 0 {
			return NotPrime
		}
	}

	return Prime
}

func randomNumber(min, max int) int {
	return rand.IntN(max-min+1) + min
}

func copyPrimeNumbers(arrSource []int) []int {
	var arr []int

	for _, value := range arrSource {
		if checkPrime(value) == Prime {
			arr = append(arr, value)
		}
	}

	return arr
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printArray(arr []int) {
	fmt.Printf("Array element is %v\n", arr)
}

// ======================
//         MAIN
// ======================

func main() {
	lengthOfArray := readPositiveNumber("Please enter number of elements")

	arr := fillArrayWithRandomNumbers(lengthOfArray)

	fmt.Println("All number in array")
	printArray(arr)

	primeArr := copyPrimeNumbers(arr)

	fmt.Println("Prime number in array")
	printArray(primeArr)
}
