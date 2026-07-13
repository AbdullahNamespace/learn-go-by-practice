// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 98: custom_floor_function.go
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
// computes its floor value — the greatest integer less than or equal to the
// number — using a custom-written function (without relying on Go's
// built-in floor function). The program should print the result of the
// custom function alongside the result of the standard library equivalent
// for comparison.
//
// العربية:
// اكتب برنامجًا يقرأ عددًا عشريًا من المستخدم ويحسب قيمته الأرضية (Floor)
// — أكبر عدد صحيح أصغر من أو يساوي الرقم — باستخدام دالة مكتوبة يدويًا
// (دون الاعتماد على دالة floor المدمجة في غو). يجب أن يطبع البرنامج نتيجة
// الدالة المخصصة جنبًا إلى جنب مع نتيجة المكافئ في المكتبة القياسية
// للمقارنة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  4.7
// Output: My Floor Result : 4
//         Go math.Floor Result: 4
//   Why:    Positive numbers: floor equals the truncated integer part
//
// Example 2 (Edge Case - Negative with Fraction):
// Input:  -4.3
// Output: My Floor Result : -5
//         Go math.Floor Result: -5
//   Why:    Truncation alone gives -4, but floor rounds toward negative infinity, giving -5
//
// Example 3 (Edge Case - Negative Whole Number):
// Input:  -4.0
// Output: My Floor Result : -4
//         Go math.Floor Result: -4
//   Why:    No fractional part, so no adjustment is needed
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Input can be any valid floating-point number (positive, negative, or zero). | يمكن أن يكون المدخل أي عدد عشري صالح.
// • Must NOT use math.Floor inside the custom function itself. | يجب عدم استخدام math.Floor داخل الدالة المخصصة نفسها.
// • Must correctly distinguish negative fractional numbers from negative whole numbers. | يجب التمييز بشكل صحيح بين الأعداد الكسرية السالبة والأعداد الصحيحة السالبة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readNumber(prompt string) float64
// func myFloor(number float64) int
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

func myFloor(number float64) int {
	truncated := int(number)

	// نطرح 1 فقط إذا كان الرقم سالباً وليس عدداً صحيحاً
	if number < 0 && number != float64(truncated) {
		return truncated - 1
	}

	return truncated
}

// ======================
//	OUTPUT FUNCTIONS
// ======================

func printResult(number float64) {
	fmt.Printf("Output: My Floor Result : %d\n", myFloor(number))
	fmt.Printf("Go math.Floor Result: %.0f\n", math.Floor(number))
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readNumber("Please enter number"))
}
