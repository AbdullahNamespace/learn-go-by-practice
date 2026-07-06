// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 65: letter_pattern.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Loops & Patterns & ASCII | الحلقات والأنماط ورموز ASCII
// **Difficulty:** Medium | متوسط
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that reads a positive number N and prints a letter pattern.
// The pattern starts from 'A'. The first line prints 'A' once, the second line
// prints 'B' twice, and so on until the Nth line, which prints the Nth letter
// repeated N times.
//
// العربية:
// اكتب برنامجًا يقرأ عددًا موجبًا N ويطبع نمط حروف. يبدأ النمط من الحرف 'A'.
// السطر الأول يطبع 'A' مرة واحدة، والسطر الثاني يطبع 'B' مرتين، وهكذا حتى
// السطر N الذي يطبع الحرف ذو الترتيب N مكررًا N مرة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  3
// Output:
// A
// BB
// CCC
//
// Example 2:
// Input:  5
// Output:
// A
// BB
// CCC
// DDDD
// EEEEE
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Input must be a positive integer between 1 and 26. | المدخل يجب أن يكون عددًا صحيحًا موجبًا بين 1 و 26.
// • Pattern uses uppercase English letters only. | النمط يستخدم الحروف الإنجليزية الكبيرة فقط.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func printLetterPattern(number int)
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

func readNumberInRange(prompt string, from, to int) int {
	for {
		input := readNumber(prompt)

		if isNumberInRange(input, from, to) {
			return input
		}

		printError(fmt.Sprintf("Invalid input please enter a number in range (%d-%d)", from, to))
	}
}

func readNumberOfLetter() int {
	return readNumberInRange("please enter number of letter in range (1-26)", 1, 26)
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func isNumberInRange(number, from, to int) bool {
	return number >= from && number <= to
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

// A
// BB
// CCC
// DDDD
// EEEEE

func printLetterPattern(number int) {
	for i := 65; i < 65+number; i++ {
		for j := 65; j <= i; j++ {
			fmt.Print(string(i))
		}
		fmt.Println()
	}
}

// ======================
//         MAIN
// ======================

func main() {
	printLetterPattern(readNumberOfLetter())
}
