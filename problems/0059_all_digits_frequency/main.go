// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 59: all_digits_frequency.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Math & Loops | الرياضيات والحلقات
// **Difficulty:** Medium | متوسط
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that reads a positive number from the user and prints the frequency
// of each digit (0-9) that appears in the number. Only digits with a frequency
// greater than 0 should be displayed.
//
// AR:
// اكتب برنامجًا يقرأ عددًا موجبًا من المستخدم ويطبع تكرار كل رقم (0-9) يظهر في
// العدد. يجب عرض الأرقام التي لها تكرار أكبر من 0 فقط.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  1223321
// Output: Digit 1 Frequency is 2 Time(s).
//         Digit 2 Frequency is 3 Time(s).
//         Digit 3 Frequency is 2 Time(s).
//
// Example 2:
// Input:  5005
// Output: Digit 0 Frequency is 2 Time(s).
//         Digit 5 Frequency is 2 Time(s).
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Number must be a positive integer (> 0) | الرقم يجب أن يكون عددًا صحيحًا موجبًا
// • Loop through all digits 0-9 and check frequency | المرور عبر جميع الأرقام 0-9 وفحص التكرار
// • Only print digits that appear at least once | اطبع فقط الأرقام التي تظهر مرة واحدة على الأقل
// • Reuse countDigitFrequency logic from Problem 58 | أعد استخدام منطق countDigitFrequency من المسألة 58
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPositiveNumber(message string) int
// func countDigitFrequency(digitToCheck int, number int) int
// func printAllDigitsFrequency(number int)
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
	fmt.Printf("❌ Error : %s\n", prompt)
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

func countDigitFrequency(digitToCheck int, number int) int {
	count := 0

	for number > 0 {
		if number%10 == digitToCheck {
			count++
		}
		number /= 10
	}

	return count
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printAllDigitsFrequency(number int) {
	for i := 0; i <= 9; i++ {
		countNumber := countDigitFrequency(i, number)

		if countNumber > 0 {
			fmt.Printf("Digit %d Frequency is %d Time(s).\n", i, countNumber)
		}
	}
}

// ======================
//         MAIN
// ======================

func main() {
	printAllDigitsFrequency(readPositiveNumber("Please enter number"))
}
