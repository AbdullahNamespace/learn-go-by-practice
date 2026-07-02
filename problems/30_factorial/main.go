// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 30: factorial.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Math & Loops | الرياضيات والحلقات
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that reads a non-negative number N from the user and calculates
// the factorial of N. The factorial of a number is the product of all positive
// integers less than or equal to that number.
// Factorial formula: N! = N × (N-1) × (N-2) × ... × 2 × 1
//
// AR:
// اكتب برنامجاً يقرأ رقماً غير سالب N من المستخدم ويحسب المضروب (factorial) للرقم N.
// المضروب لرقم ما هو حاصل ضرب جميع الأعداد الصحيحة الموجبة الأقل من أو تساوي ذلك الرقم.
// صيغة المضروب: N! = N × (N-1) × (N-2) × ... × 2 × 1
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  5
// Output: 120
//   Why:    5! = 5 × 4 × 3 × 2 × 1 = 120
//
// Example 2:
// Input:  4
// Output: 24
//   Why:    4! = 4 × 3 × 2 × 1 = 24
//
// Example 3 (Edge Case):
// Input:  0
// Output: 1
//   Why:    By definition, 0! = 1
//
// Example 4 (Edge Case):
// Input:  1
// Output: 1
//   Why:    1! = 1
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • N must be a non-negative integer (N >= 0) | يجب أن يكون N عدداً صحيحاً غير سالب
// • Keep prompting until valid input is received | استمر في طلب المدخل حتى يكون صحيحاً
// • Handle edge cases: 0! = 1 and 1! = 1 | تعامل مع الحالات الخاصة
// • Use uint64 to prevent quick overflow | استخدم uint64 لمنع التجاوز السريع للسعة
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPositiveNumber() int
// func factorial(n int) uint64
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

var reader = bufio.NewReader(os.Stdin)

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

		if input < 0 {
			printError("Invalid input please enter a positive number!")
			continue
		}

		return input
	}
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func factorial(n int) uint64 {
	if n <= 1 {
		return 1
	}

	return uint64(n) * factorial(n-1)
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(n int) {
	fmt.Printf("%d\n", factorial(n))
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readPositiveNumber("Enter a positive number"))
}

// ======================
//   UTILITY FUNCTIONS
// ======================

func printError(prompt string) {
	fmt.Printf("❌ Error : %s\n", prompt)
}
