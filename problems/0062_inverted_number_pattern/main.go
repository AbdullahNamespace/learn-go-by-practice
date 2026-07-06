// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 62: inverted_number_pattern.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Loops & Patterns | الحلقات والأنماط
// **Difficulty:** Medium | متوسط
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that reads a positive number and prints an inverted number pattern.
// The first line prints the given number repeated as many times as its value,
// the next line prints the number minus one repeated, and so on until it reaches 1.
//
// العربية:
// اكتب برنامجًا يقرأ عددًا موجبًا ويطبع نمط أرقام معكوس. السطر الأول يطبع الرقم المُدخل
// مكررًا بنفس عدد مرات قيمته، والسطر التالي يطبع الرقم ناقص واحدًا مكررًا، وهكذا
// حتى يصل إلى الرقم 1.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  5
// Output:
// 55555
// 4444
// 333
// 22
// 1
//
// Example 2:
// Input:  3
// Output:
// 333
// 22
// 1
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Input must be a positive integer (> 0). | المدخل يجب أن يكون عددًا صحيحًا موجبًا (> 0).
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPositiveNumber(prompt string) int
// func printInvertedNumberPattern(number int)
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

func printInvertedNumberPattern(number int) {
	for i := number; i >= 1; i-- {
		for j := i; j >= 1; j-- {
			fmt.Print(i)
		}
		fmt.Println()
	}
}

// ======================
//         MAIN
// ======================

func main() {
	printInvertedNumberPattern(readPositiveNumber("Please enter a number"))
}
