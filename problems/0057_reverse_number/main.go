// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 57: reverse_number.go
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
// Write a program that reads a positive number from the user and reverses its digits.
// For example, if the input is 123, the output should be 321. The reversed number
// is constructed by extracting digits and building a new number.
//
// AR:
// اكتب برنامجًا يقرأ عددًا موجبًا من المستخدم ويعكس أرقامه. على سبيل المثال،
// إذا كان الإدخال 123، يجب أن يكون المخرج 321. يتم بناء العدد المعكوس عن طريق
// استخراج الأرقام وبناء عدد جديد.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  123
// Output: Reverse is: 321
//   Why:    Extract 3, then 2, then 1 and build 321
//
// Example 2:
// Input:  506
// Output: Reverse is: 605
//   Why:    Extract 6, then 0, then 5 and build 605
//
// Example 3:
// Input:  1000
// Output: Reverse is: 1
//   Why:    Leading zeros in the reverse are dropped naturally
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Number must be a positive integer (> 0) | الرقم يجب أن يكون عددًا صحيحًا موجبًا
// • Build the reversed number mathematically, not as a string | بناء العدد المعكوس رياضيًا وليس كسلسلة نصية
// • Use formula: reversed = reversed * 10 + remainder | استخدم الصيغة: المعكوس = المعكوس * 10 + الباقي
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPositiveNumber(message string) int
// func reverseNumber(number int) int
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

func printResult(number int) {
	fmt.Printf("Reverse is: %d\n", reverseNumber(number))
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readPositiveNumber("Please enter number"))
}
