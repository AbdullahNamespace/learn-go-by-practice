// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 52: prime_numbers_1_to_n.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Math & Loops | الرياضيات والحلقات
// **Difficulty:** Medium | متوسط
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that reads a positive number N from the user and prints all prime
// numbers from 1 to N. A prime number is greater than 1 and has no divisors other
// than 1 and itself.
//
// AR:
// اكتب برنامجًا يقرأ عددًا موجبًا N من المستخدم ويطبع جميع الأعداد الأولية من
// 1 إلى N. العدد الأولي أكبر من 1 وليس له قواسم سوى 1 ونفسه.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  20
// Output: Prime Numbers from 1 To 20 are :
//         2
//         3
//         5
//         7
//         11
//         13
//         17
//         19
//
// Example 2:
// Input:  10
// Output: Prime Numbers from 1 To 10 are :
//         2
//         3
//         5
//         7
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • N must be a positive integer (> 0) | N يجب أن يكون عددًا صحيحًا موجبًا
// • Optimize by checking divisors only up to N/2 | تحسين الخوارزمية بالتحقق فقط حتى N/2
// • Use a custom enum/type for Prime/NotPrime status | استخدم نوعًا مخصصًا لحالة أولي/غير أولي
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// type PrimeStatus int
// func readPositiveNumber(prompt string) int
// func checkPrime(number int) PrimeStatus
// func printPrimeNumbersFrom1ToN(number int)
//
// ────────────────────────────────────────────────────────────────────────────
package main

import (
	"fmt"
	"strconv"
)

// ======================
//   UTILITY
// ======================

type PrimeStatus int

const (
	Prime    PrimeStatus = 0
	NotPrime PrimeStatus = 1
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

// ======================
//   PROCESSING FUNCTIONS
// ======================

func checkPrime(number int) PrimeStatus {
	if number <= 1 {
		return NotPrime
	}

	for i := 2; i <= (number / 2); i++ {
		if number%i == 0 {
			return NotPrime
		}
	}
	return Prime
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printPrimeNumbersFrom1ToN(number int) {
	for i := 1; i <= number; i++ {
		if checkPrime(i) == Prime {
			fmt.Println(i)
		}
	}
}

// ======================
//         MAIN
// ======================

func main() {
	printPrimeNumbersFrom1ToN(readPositiveNumber("Please enter number"))
}
