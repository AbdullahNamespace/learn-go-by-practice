// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 54: perfect_numbers_1_to_n.go
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
// Write a program that reads a positive number N from the user and prints all
// perfect numbers from 1 to N. A perfect number equals the sum of its proper
// divisors (excluding itself).
//
// AR:
// اكتب برنامجًا يقرأ عددًا موجبًا N من المستخدم ويطبع جميع الأعداد التامة من
// 1 إلى N. العدد التام يساوي مجموع قواسمه الحقيقية (باستثناء نفسه).
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  30
// Output: 6
//         28
//   Why:    6 = 1+2+3, 28 = 1+2+4+7+14
//
// Example 2:
// Input:  500
// Output: 6
//         28
//         496
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • N must be a positive integer (> 0) | N يجب أن يكون عددًا صحيحًا موجبًا
// • Reuse the isPerfectNumber logic from Problem 53 | أعد استخدام منطق isPerfectNumber من المسألة 53
// • Print each perfect number on a new line | اطبع كل عدد تام في سطر جديد
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPositiveNumber(prompt string) int
// func isPerfectNumber(number int) bool
// func printPerfectNumbersFrom1ToN(number int)
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

type PerfectStatus int

const (
	Perfect    PerfectStatus = 0
	NotPerfect PerfectStatus = 1
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

func checkPerfect(number int) PerfectStatus {
	sum := 0

	for i := 1; i < number; i++ {
		if number%i == 0 {
			sum += i
		}
	}

	if sum == number {
		return Perfect
	}

	return NotPerfect
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printPerfectNumbersFrom1ToN(number int) {
	for i := 1; i <= number; i++ {
		if checkPerfect(i) == Perfect {
			fmt.Println(i)
		}
	}
}

// ======================
//         MAIN
// ======================

func main() {
	printPerfectNumbersFrom1ToN(readPositiveNumber("Please enter a number"))
}
