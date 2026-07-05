// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 55: print_digits_reverse.go
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
// Write a program that reads a positive number from the user and prints its digits
// in reverse order, one digit per line. For example, if the input is 1234, the
// output should be 4, 3, 2, 1 each on a separate line.
//
// AR:
// اكتب برنامجًا يقرأ عددًا موجبًا من المستخدم ويطبع أرقامه بترتيب عكسي، رقمًا
// واحدًا في كل سطر. على سبيل المثال، إذا كان الإدخال 1234، يجب أن يكون
// المخرج 4، 3، 2، 1 كلٌ في سطر منفصل.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  1234
// Output: 4
//         3
//         2
//         1
//   Why:    Digits extracted from right to left using modulo and division
//
// Example 2:
// Input:  506
// Output: 6
//         0
//         5
//   Why:    506 % 10 = 6, 50 % 10 = 0, 5 % 10 = 5
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Number must be a positive integer (> 0) | الرقم يجب أن يكون عددًا صحيحًا موجبًا
// • Use modulo (%) and division (/) to extract digits | استخدم باقي القسمة والقسمة لاستخراج الأرقام
// • Print each digit on a new line | اطبع كل رقم في سطر جديد
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPositiveNumber(message string) int
// func printDigits(number int)
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
			printError("Invalid input please enter a positive number!")
			continue
		}

		return input
	}
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printDigits(input int) {
	for input > 0 {
		number := input % 10
		fmt.Println(number)
		input /= 10
	}
}

// ======================
//         MAIN
// ======================

func main() {
	printDigits(readPositiveNumber("Please enter a number"))
}
