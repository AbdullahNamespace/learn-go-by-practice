// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 56: sum_of_digits.go
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
// Write a program that reads a positive number from the user and calculates the sum
// of all its digits. For example, if the input is 1234, the sum is 1+2+3+4 = 10.
//
// AR:
// اكتب برنامجًا يقرأ عددًا موجبًا من المستخدم ويحسب مجموع جميع أرقامه. على سبيل
// المثال، إذا كان الإدخال 1234، فإن المجموع هو 1+2+3+4 = 10.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  1234
// Output: Sum Of Digits = 10
//   Why:    1 + 2 + 3 + 4 = 10
//
// Example 2:
// Input:  506
// Output: Sum Of Digits = 11
//   Why:    5 + 0 + 6 = 11
//
// Example 3:
// Input:  7
// Output: Sum Of Digits = 7
//   Why:    Single digit, sum equals the number itself
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Number must be a positive integer (> 0) | الرقم يجب أن يكون عددًا صحيحًا موجبًا
// • Use modulo (%) and division (/) to extract digits | استخدم باقي القسمة والقسمة لاستخراج الأرقام
// • Return the sum as an integer | أرجع المجموع كعدد صحيح
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPositiveNumber(message string) int
// func sumOfDigits(number int) int
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
			printError("Invalid input please enter a  positive number!")
			continue
		}

		return input
	}
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func sumOfDigits(input int) int {
	sum := 0

	for input > 0 {
		sum += input % 10
		input /= 10
	}

	return sum
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(input int) {
	fmt.Printf("Sum Of Digits = %d\n", sumOfDigits(input))
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readPositiveNumber("Please enter number"))
}
