// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 60: print_digits_ordered.go
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
// in their original order (left to right), one digit per line. This is achieved by
// first reversing the number, then printing the digits of the reversed number
// (which extracts them from right to left, effectively giving the original order).
//
// AR:
// اكتب برنامجًا يقرأ عددًا موجبًا من المستخدم ويطبع أرقامه بترتيبها الأصلي (من
// اليسار إلى اليمين)، رقمًا واحدًا في كل سطر. يتم ذلك عن طريق عكس العدد أولاً،
// ثم طباعة أرقام العدد المعكوس (مما يستخرجها من اليمين إلى اليسار، مما يعطي
// الترتيب الأصلي فعليًا).
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  1234
// Output: 1
//         2
//         3
//         4
//   Why:    Reverse(1234) = 4321, then print digits of 4321: 1, 2, 3, 4
//
// Example 2:
// Input:  506
// Output: 5
//         0
//         6
//   Why:    Reverse(506) = 605, then print digits of 605: 5, 0, 6
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Number must be a positive integer (> 0) | الرقم يجب أن يكون عددًا صحيحًا موجبًا
// • Must use reverseNumber then printDigits approach | يجب استخدام نهج reverseNumber ثم printDigits
// • Reuse functions from Problems 55 and 57 | أعد استخدام الدوال من المسألتين 55 و 57
// • Print each digit on a new line | اطبع كل رقم في سطر جديد
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPositiveNumber(message string) int
// func reverseNumber(number int) int
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
			printError("Invalid input please enter a valid number")
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

func reverseNumber(number int) int {
	reversed := 0

	for number > 0 {
		reversed = reversed*10 + number%10
		number /= 10
	}

	return reversed
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printDigits(number int) {
	for number > 0 {
		fmt.Println(number % 10)
		number /= 10
	}
}

// ======================
//         MAIN
// ======================

func main() {
	printDigits(reverseNumber(readPositiveNumber("Please enter number")))
}
