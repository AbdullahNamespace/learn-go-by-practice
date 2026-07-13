// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 97: custom_round_function.go
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
// rounds it to the nearest integer using a custom-written function (without
// relying on Go's built-in rounding function). The program should print the
// result of the custom function alongside the result of the standard library
// equivalent for comparison.
//
// العربية:
// اكتب برنامجًا يقرأ عددًا عشريًا من المستخدم ويقرّبه إلى أقرب عدد صحيح
// باستخدام دالة مكتوبة يدويًا (دون الاعتماد على دالة التقريب المدمجة في
// غو). يجب أن يطبع البرنامج نتيجة الدالة المخصصة جنبًا إلى جنب مع نتيجة
// المكافئ في المكتبة القياسية للمقارنة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  4.6
// Output: My Round Result : 5
//         Go math.Round Result: 5
//   Why:    Fractional part (0.6) >= 0.5, rounds up
//
// Example 2:
// Input:  4.3
// Output: My Round Result : 4
//         Go math.Round Result: 4
//   Why:    Fractional part (0.3) < 0.5, rounds toward zero
//
// Example 3 (Edge Case - Negative Number):
// Input:  -4.6
// Output: My Round Result : -5
//         Go math.Round Result: -5
//   Why:    Rounds away from zero in the negative direction
//
// Example 4 (Edge Case - Exactly 0.5):
// Input:  2.5
// Output: My Round Result : 3
//         Go math.Round Result: 3
//   Why:    Half-values round away from zero (up, in this case)
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Input can be any valid floating-point number (positive, negative, or zero). | يمكن أن يكون المدخل أي عدد عشري صالح.
// • Must NOT use math.Round inside the custom function itself. | يجب عدم استخدام math.Round داخل الدالة المخصصة نفسها.
// • Rounding must follow "round half away from zero" semantics, matching math.Round. | يجب أن يتبع التقريب دلالات "التقريب النصفي بعيدًا عن الصفر"، مطابقًا لـ math.Round.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readNumber(prompt string) float64
// func getFractionPart(number float64) float64
// func myRound(number float64) int
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

func myRound(number float64) int {
	if math.Abs(getFractionPart(number)) >= 0.5 {
		if number > 0 {
			return int(number) + 1
		}

		return int(number) - 1
	}

	return int(number)
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(number float64) {
	fmt.Printf("My Round Result : %d\n", myRound(number))
	fmt.Printf("Go math.Round Result: %.0f\n", math.Round(number))
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readNumber("Please enter number"))
}
