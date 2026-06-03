// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 7: half_number_calculator.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Basic Math / Type Casting | الرياضيات الأساسية / تحويل أنواع البيانات
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that asks the user to enter an integer number. Calculate half of that
// number and display the result. Since dividing an integer by 2 might result in a fraction
// (e.g., 5 / 2 = 2.5), ensure the calculation handles floating-point numbers correctly
// by casting the number to float64 before division.
//
// العربية:
// اكتب برنامجاً يطلب من المستخدم إدخال رقم صحيح. حساب نصف هذا الرقم وعرض النتيجة.
// نظراً لأن قسمة عدد صحيح على 2 قد تنتج كسراً (مثلاً: 5 / 2 = 2.5)، تأكد من أن
// الحساب يعالج الأرقام العشرية بشكل صحيح عن طريق تحويل الرقم إلى float64 قبل القسمة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  10
// Output: Half of 10 is 5
//
// Example 2:
// Input:  5
// Output: Half of 5 is 2.5
//
// Example 3 (Edge Case):
// Input:  0
// Output: Half of 0 is 0
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Input is a valid integer. | المدخل هو عدد صحيح صالح.
// • Output must handle decimal points. | المخرجات يجب أن تعالج الكسور العشرية.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readNumber(prompt string) int
// func calculateHalfNumber(num int) float64
// func printResults(num float64)
//
// ────────────────────────────────────────────────────────────────────────────
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// ======================
//     INPUT FUNCTIONS
// ======================

func readString(prompt string) (string, error) {
	fmt.Printf("%s", fmt.Sprintf("%s  :", prompt))
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func readNumber(prompt string) int {
	for {
		input, err := readString(prompt)
		if err != nil {
			fmt.Println("Invalid input, Please enter a valid input!")
			continue
		}
		inputNumber, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input, Please enter a valid number!")
			continue
		}
		return inputNumber
	}
}

// ======================
//	PROCESSING FUNCTIONS
// ======================

func calculateHalfNumber(num int) float64 {
	return float64(num) / 2.0
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResults(num int) {
	fmt.Printf("Half of %d is %f\n", num, calculateHalfNumber(num))
}

// ======================
//         MAIN
// ======================

func main() {
	num := readNumber("Please Enter a Number")
	printResults(num)
}
