// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 96: custom_abs_function.go
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
// computes its absolute value using a custom-written function (without
// relying on any built-in absolute-value function). The program should then
// print the result of the custom function side by side with the result of
// Go's standard library equivalent, to let the user visually confirm both
// produce the same output.
//
// العربية:
// اكتب برنامجًا يقرأ عددًا عشريًا من المستخدم ويحسب قيمته المطلقة باستخدام
// دالة مكتوبة يدويًا (دون الاعتماد على أي دالة مدمجة للقيمة المطلقة). يجب
// أن يطبع البرنامج بعد ذلك نتيجة الدالة المخصصة جنبًا إلى جنب مع نتيجة
// المكافئ في المكتبة القياسية لغو، للسماح للمستخدم بالتحقق بصريًا من أن
// كلاهما يُنتج نفس المخرجات.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  -7.5
// Output: My abs Result : 7.5
//         Go math.Abs Result: 7.5
//
// Example 2:
// Input:  12.3
// Output: My abs Result : 12.3
//         Go math.Abs Result: 12.3
//   Why:    Positive numbers are returned unchanged
//
// Example 3 (Edge Case - Zero):
// Input:  0
// Output: My abs Result : 0
//         Go math.Abs Result: 0
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Input can be any valid floating-point number (positive, negative, or zero). | يمكن أن يكون المدخل أي عدد عشري صالح (موجب، سالب، أو صفر).
// • Must NOT use math.Abs (or any built-in) inside the custom function itself. | يجب عدم استخدام math.Abs (أو أي دالة مدمجة) داخل الدالة المخصصة نفسها.
// • Use float64 for consistency with Go's math package. | استخدم float64 للاتساق مع حزمة math في غو.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readNumber(prompt string) float64
// func myAbs(number float64) float64
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

func myAbs(number float64) float64 {
	if number < 0 {
		return -number
	}
	return number
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(number float64) {
	fmt.Printf("My abs Result : %.2f\n", myAbs(number))
	fmt.Printf("Go math.Abs Result: %.2f\n", math.Abs(number))
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readNumber("Please enter number"))
}
