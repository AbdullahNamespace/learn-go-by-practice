// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 99: custom_ceil_function.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Math & Custom Implementation | الرياضيات والتنفيذ المخصص
// **Difficulty:** Medium | متوسط
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that reads a floating-point number from the user and
// computes its ceiling value — the smallest integer greater than or equal
// to the number — using a custom-written function (without relying on Go's
// built-in ceiling function). The program should print the result of the
// custom function alongside the result of the standard library equivalent
// for comparison.
//
//
// العربية:
// اكتب برنامجًا يقرأ عددًا عشريًا من المستخدم ويحسب قيمته السقفية (Ceiling)
// — أصغر عدد صحيح أكبر من أو يساوي الرقم — باستخدام دالة مكتوبة يدويًا
// (دون الاعتماد على دالة ceiling المدمجة في غو). يجب أن يطبع البرنامج
// نتيجة الدالة المخصصة جنبًا إلى جنب مع نتيجة المكافئ في المكتبة القياسية
// للمقارنة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  4.3
// Output: My Ceil Result : 5
//         Go math.Ceil Result: 5
//   Why:    Positive number with a fraction rounds up
//
// Example 2 (Edge Case - Negative with Fraction):
// Input:  -4.7
// Output: My Ceil Result : -4
//         Go math.Ceil Result: -4
//   Why:    Ceiling of a negative fractional number rounds toward zero (truncation)
//
// Example 3 (Edge Case - Already a Whole Number):
// Input:  6.0
// Output: My Ceil Result : 6
//         Go math.Ceil Result: 6
//   Why:    No fractional part, so the number is returned unchanged
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Input can be any valid floating-point number (positive, negative, or zero). | يمكن أن يكون المدخل أي عدد عشري صالح.
// • Must NOT use math.Ceil inside the custom function itself. | يجب عدم استخدام math.Ceil داخل الدالة المخصصة نفسها.
// • Must correctly distinguish positive fractional numbers from negative fractional numbers. | يجب التمييز بشكل صحيح بين الأعداد الكسرية الموجبة والأعداد الكسرية السالبة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readNumber(prompt string) float64
// func getFractionPart(number float64) float64
// func myCeil(number float64) int
//
// ────────────────────────────────────────────────────────────────────────────
package main

import (
	"fmt"
	"math"
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

func readNumber(prompt string) float64 {
	for {
		input := readString(prompt)

		inputNumber, err := strconv.ParseFloat(input, 64)

		if err != nil {
			printError("Invalid input please enter a valid number!")
			continue
		}

		return inputNumber
	}
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func getFractionPart(number float64) float64 {
	return number - float64(int(number))
}

func myCeil(number float64) int {
	fg := getFractionPart(number)
	if math.Abs(fg) > 0 {
		if number > 0 {
			return int(number) + 1
		}
	}

	return int(number)
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(number float64) {
	fmt.Printf("Output: My Ceil Result : %d\n", myCeil(number))
	fmt.Printf("Go math.Ceil Result: %.0f\n", math.Ceil(number))
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readNumber("Please enter number"))
}
