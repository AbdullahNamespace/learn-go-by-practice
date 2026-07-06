// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 64: inverted_letter_pattern.go
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
// Write a program that reads a positive number N and prints an inverted letter pattern.
// The pattern starts from the Nth uppercase letter (e.g., if N=3, it starts with 'C')
// and goes down to 'A'. The first letter is printed N times, the next letter N-1 times,
// and so on, until 'A' is printed once.
//
// العربية:
// اكتب برنامجًا يقرأ عددًا موجبًا N ويطبع نمط حروف معكوس. يبدأ النمط من الحرف
// الأبجدي الكبير ذو الترتيب N (مثلاً إذا N=3 يبدأ بـ 'C') وينزل حتى 'A'.
// يُطبع الحرف الأول N مرة، والحرف التالي N-1 مرة، وهكذا حتى يُطبع الحرف 'A' مرة واحدة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  3
// Output:
// CCC
// BB
// A
//
// Example 2:
// Input:  5
// Output:
// EEEEE
// DDDD
// CCC
// BB
// A
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
// func printInvertedLetterPattern(number int)
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
			printError("Invalid input please enter a number!")
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
		printError(fmt.Sprintf("Invalid input please enter number in range (%d-%d)", from, to))
	}
}

func readNumberOfLetter() int {
	return readNumberInRange("please enter number of letter in range (1-26)", 1, 26)
}

// ======================
//  PROCESSING FUNCTIONS
// ======================

func isNumberInRange(number, from, to int) bool {
	return number >= from && number <= to
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printInvertedLetterPattern(number int) {
	for i := 64 + number; i >= 65; i-- {
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
	printInvertedLetterPattern(readNumberOfLetter())
}
