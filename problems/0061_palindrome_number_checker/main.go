// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 61: palindrome_number_checker.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Math & Logic | الرياضيات والمنطق
// **Difficulty:** Medium | متوسط
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that reads a positive number from the user and checks whether
// it is a palindrome number. A palindrome number reads the same backward as forward
// (e.g., 121, 1331, 12321). The program should reverse the digits of the number
// and compare it with the original number to determine the result.
//
// العربية:
// اكتب برنامجًا يقرأ عددًا موجبًا من المستخدم ويتحقق مما إذا كان عددًا متناظرًا (Palindrome).
// العدد المتناظر هو الذي يُقرأ بنفس الطريقة من اليمين واليسار (مثل: 121، 1331، 12321).
// يجب أن يقوم البرنامج بعكس أرقام العدد ومقارنته بالرقم الأصلي لتحديد النتيجة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  121
// Output: Yes, it is a Palindrome number.
//
// Example 2:
// Input:  123
// Output: No, it is NOT a Palindrome number.
//
// Example 3 (Single Digit):
// Input:  5
// Output: Yes, it is a Palindrome number.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Input must be a positive integer (> 0). | المدخل يجب أن يكون عددًا صحيحًا موجبًا (> 0).
// • Single digit numbers are considered palindromes. | الأرقام المفردة تُعتبر أعدادًا متناظرة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPositiveNumber(prompt string) int
// func reverseNumber(number int) int
// func isPalindromeNumber(number int) bool
// func printResult(number int)
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
			printError("Invalid input please enter a valid!")
			continue
		}

		return inputNumber
	}
}

func readPositiveNumber(prompt string) int {
	for {
		input := readNumber(prompt)

		if input <= 0 {
			printError("Invalid input please enter a positive number")
			continue
		}

		return input
	}
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func reverseNumber(number int) int {
	reversed := 0

	for number > 0 {
		reversed = reversed*10 + number%10
		number /= 10
	}

	return reversed
}

func isPalindromeNumber(number int) bool {
	return reverseNumber(number) == number
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(number int) {
	if isPalindromeNumber(number) {
		fmt.Printf("Yes, it is a Palindrome number.\n")
	} else {
		fmt.Printf("No, it is NOT a Palindrome number.\n")
	}
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readPositiveNumber("Please enter a number"))
}
