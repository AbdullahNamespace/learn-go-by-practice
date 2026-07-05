// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 38: prime_number_check.go
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
// Write a program that reads a positive number and checks if it is a prime number
// or not. A prime number is greater than 1 and has no divisors other than 1 and itself.
//
// AR:
// اكتب برنامجًا يقرأ عددًا موجبًا ويتحقق مما إذا كان عددًا أوليًا أم لا.
// العدد الأولي أكبر من 1 وليس له قواسم سوى 1 ونفسه.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  7
// Output: The number is Prime
//   Why:    7 is only divisible by 1 and 7
//
// Example 2:
// Input:  10
// Output: The number is Not Prime
//   Why:    10 is divisible by 2 and 5
//
// Example 3 (Edge Case):
// Input:  1
// Output: The number is Not Prime
//   Why:    1 is not considered a prime number
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Number must be > 0 | الرقم يجب أن يكون أكبر من 0
// • Optimize by checking divisors only up to N/2 | قم بتحسين الخوارزمية بالتحقق فقط حتى N/2
// • Return a custom enum/type for Prime/NotPrime | أرجع نوعًا مخصصًا لأولي / غير أولي
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// type PrimeStatus int
// func readPositiveNumber(prompt string) int
// func checkPrime(number int) PrimeStatus
// func printNumberType(number int)
//
// ────────────────────────────────────────────────────────────────────────────
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ======================
//     INPUT FUNCTIONS
// ======================

func readString(prompt string) (string, error) {
	fmt.Printf("%s : ", prompt)

	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

func readNumber(prompt string) int {
	for {
		input, err := readString(prompt)

		if err != nil {
			printError("Invalid input please enter again!")
			continue
		}

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
			printError("invalid input please enter a positive number!")
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

func printNumberType(number int) {
	if checkPrime(number) == Prime {
		fmt.Printf("\nThe number is prime\n")
		return
	}
	fmt.Printf("\nThe number is not prime\n")
}

// ======================
//         MAIN
// ======================

func main() {
	printNumberType(readPositiveNumber("Please enter number"))
}

// ======================
//   UTILITY FUNCTIONS
// ======================

type PrimeStatus int

const (
	Prime    PrimeStatus = 0
	NotPrime PrimeStatus = 1
)

var reader = bufio.NewReader(os.Stdin)

func printError(prompt string) {
	fmt.Printf("❌ Error : %s", prompt)
}
