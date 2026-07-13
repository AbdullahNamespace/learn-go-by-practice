// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 100: custom_sqrt_function.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Math & Custom Implementation | الرياضيات والتنفيذ المخصص
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that reads a floating-point number from the user and
// computes its square root using a custom-written function that leverages
// the general power function instead of a dedicated square-root function.
// The program should print the result of the custom function alongside the
// result of the standard library's dedicated square-root function, for
// comparison.
//
// العربية:
// اكتب برنامجًا يقرأ عددًا عشريًا من المستخدم ويحسب جذره التربيعي باستخدام
// دالة مكتوبة يدويًا تستفيد من دالة الأس العامة بدلاً من دالة جذر تربيعي
// مخصصة. يجب أن يطبع البرنامج نتيجة الدالة المخصصة جنبًا إلى جنب مع نتيجة
// دالة الجذر التربيعي المخصصة في المكتبة القياسية، للمقارنة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  25
// Output: My Sqrt Result : 5
//         Go math.Sqrt Result: 5
//   Why:    5^2 = 25, so sqrt(25) = 5
//
// Example 2:
// Input:  2
// Output: My Sqrt Result : 1.414214
//         Go math.Sqrt Result: 1.414214
//   Why:    x^0.5 and math.Sqrt agree for non-perfect squares too
//
// Example 3 (Edge Case - Negative Input):
// Input:  -9
// Output: My Sqrt Result : NaN
//         Go math.Sqrt Result: NaN
//   Why:    The square root of a negative number is undefined in real numbers
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Input can be any valid floating-point number (including negative, which yields NaN). | يمكن أن يكون المدخل أي عدد عشري صالح (بما في ذلك السالب، والذي يُنتج NaN).
// • Must use math.Pow(number, 0.5) instead of math.Sqrt inside the custom function. | يجب استخدام math.Pow(number, 0.5) بدلاً من math.Sqrt داخل الدالة المخصصة.
// • No manual square-root algorithm (e.g., Newton's method) is required for this exercise. | لا يُطلب خوارزمية جذر تربيعي يدوية (مثل طريقة نيوتن) لهذا التمرين.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readNumber(prompt string) float64
// func mySqrt(number float64) float64
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

func mySqrt(number float64) float64 {
	return math.Pow(number, 0.5)
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(number float64) {
	fmt.Printf("Output: My Sqrt Result : %f\n", mySqrt(number))
	fmt.Printf("Go math.Sqrt Result: %f\n", math.Sqrt(number))
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readNumber("Please enter number"))
}
