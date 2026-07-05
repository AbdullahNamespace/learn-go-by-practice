// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 53: perfect_number_check.go
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
// Write a program that reads a positive number from the user and checks if it is a
// perfect number. A perfect number is a number that equals the sum of its proper
// divisors (excluding itself). For example, 6 = 1 + 2 + 3.
//
// AR:
// اكتب برنامجًا يقرأ عددًا موجبًا من المستخدم ويتحقق مما إذا كان عددًا تامًا.
// العدد التام هو عدد يساوي مجموع قواسمه الحقيقية (باستثناء نفسه). على سبيل
// المثال: 6 = 1 + 2 + 3.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  6
// Output: 6 Is Perfect Number.
//   Why:    1 + 2 + 3 = 6
//
// Example 2:
// Input:  28
// Output: 28 Is Perfect Number.
//   Why:    1 + 2 + 4 + 7 + 14 = 28
//
// Example 3:
// Input:  10
// Output: 10 Is NOT Perfect Number.
//   Why:    1 + 2 + 5 = 8, which is not equal to 10
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Number must be a positive integer (> 0) | الرقم يجب أن يكون عددًا صحيحًا موجبًا
// • Proper divisors exclude the number itself | القواسم الحقيقية تستبعد العدد نفسه
// • Return a boolean for perfect/not perfect | أرجع قيمة منطقية لتام/غير تام
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPositiveNumber(prompt string) int
// func checkPerfect(number int)
// func printResults(number int)
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

func printResults(number int) {
	if checkPerfect(number) == Perfect {
		fmt.Printf("\n%d Is Perfect Number.\n", number)
	} else {
		fmt.Printf("\n%d Is NOT Perfect Number.\n", number)
	}
}

// ======================
//         MAIN
// ======================

func main() {
	printResults(readPositiveNumber("Please enter a number"))
}
