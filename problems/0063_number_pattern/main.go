// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 63: number_pattern.go
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
// Write a program that reads a positive number N and prints a number pattern.
// The first line prints the number 1 once, the second line prints the number 2 twice,
// and so on until the Nth line, which prints the number N repeated N times.
//
// العربية:
// اكتب برنامجًا يقرأ عددًا موجبًا N ويطبع نمط أرقام. السطر الأول يطبع الرقم 1 مرة واحدة،
// والسطر الثاني يطبع الرقم 2 مرتين، وهكذا حتى السطر N الذي يطبع الرقم N مكررًا N مرة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  5
// Output:
// 1
// 22
// 333
// 4444
// 55555
//
// Example 2:
// Input:  3
// Output:
// 1
// 22
// 333
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
// func printNumberPattern(number int)
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

func printNumberPattern(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i)
		}
		fmt.Println()
	}
}

// ======================
//         MAIN
// ======================

func main() {
	printNumberPattern(readPositiveNumber("Please enter a number"))
}
