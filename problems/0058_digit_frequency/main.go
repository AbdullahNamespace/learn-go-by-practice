// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 58: digit_frequency.go
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
// Write a program that reads a positive number and a single digit from the user, then
// counts how many times that digit appears in the number and prints the frequency.
//
// AR:
// اكتب برنامجًا يقرأ عددًا موجبًا ورقمًا واحدًا من المستخدم، ثم يحسب عدد مرات
// ظهور ذلك الرقم في العدد ويطبع التكرار.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Number = 1223321, Digit = 2
// Output: Digit 2 Frequency is 3 Time(s).
//   Why:    2 appears at positions: 2nd, 3rd, 6th digit
//
// Example 2:
// Input:  Number = 55555, Digit = 5
// Output: Digit 5 Frequency is 5 Time(s).
//
// Example 3:
// Input:  Number = 12345, Digit = 9
// Output: Digit 9 Frequency is 0 Time(s).
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Main number must be a positive integer (> 0) | العدد الرئيسي يجب أن يكون عددًا صحيحًا موجبًا
// • Digit to check should be a single digit (0-9) | الرقم المطلوب فحصه يجب أن يكون رقمًا واحدًا (0-9)
// • Extract digits using modulo and division | استخراج الأرقام باستخدام باقي القسمة والقسمة
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPositiveNumber(message string) int
// func countDigitFrequency(digitToCheck int, number int) int
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

func readSingleDigit(prompt string) int {
	for {
		input := readNumber(prompt)

		if input > 9 || input < 0 {
			printError("Invalid input please enter a number in 0-9")
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

func printResult(number int, digitToCheck int) {
	fmt.Printf("Digit %d Frequency is %d Time(s).\n", digitToCheck, countDigitFrequency(digitToCheck, number))
}

// ======================
//         MAIN
// ======================

func main() {
	number := readPositiveNumber("Please enter a number")
	digit := readSingleDigit("Please enter a digit to check (0-9)")
	printResult(number, digit)
}
