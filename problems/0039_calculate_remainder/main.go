// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 39: calculate_remainder.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Basic Math | الرياضيات الأساسية
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that reads the total bill and the total cash paid, then calculates
// and prints the remainder (change).
//
// AR:
// اكتب برنامجًا يقرأ إجمالي الفاتورة وإجمالي النقد المدفوع، ثم يحسب ويطبع
// المتبقي (الباقي).
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  TotalBill = 50, TotalCashPaid = 100
// Output: Remainder = 50
//   Why:    100 - 50 = 50
//
// Example 2:
// Input:  TotalBill = 15.5, TotalCashPaid = 20
// Output: Remainder = 4.5
//   Why:    20 - 15.5 = 4.5
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Both inputs must be positive numbers (> 0) | كلا المدخلين يجب أن يكونا أرقاماً موجبة
// • Formula: Remainder = CashPaid - TotalBill | الصيغة: الباقي = المدفوع - الفاتورة
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPositiveNumber(message string) float64
// func calculateRemainder(totalBill, totalCashPaid float64) float64
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

// ======================
//     INPUT FUNCTIONS
// ======================

func readString(prompt string) (string, error) {
	fmt.Printf("%s : ", prompt)

	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

func readNumber(prompt string) float64 {
	for {
		input, err := readString(prompt)

		if err != nil {
			printError("Invalid input please enter again!")
			continue
		}

		inputNumber, err := strconv.ParseFloat(input, 64)

		if err != nil {
			printError("Invalid input please enter a valid number!")
			continue
		}

		return inputNumber
	}
}

func readPositiveNumber(prompt string) float64 {
	for {
		input := readNumber(prompt)

		if input <= 0 {
			printError("Invalid input please enter a positive number!")
			continue
		}

		return input
	}
}

func readBillData() (float64, float64) {
	totalBill := readPositiveNumber("Please enter total bill")
	totalCashPaid := readPositiveNumber("Please enter total cash paid")

	return totalBill, totalCashPaid
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func calculateRemainder(totalBill, totalCashPaid float64) float64 {
	return totalCashPaid - totalBill
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(totalBill, totalCashPaid float64) {
	fmt.Printf("Remainder = %f\n", calculateRemainder(totalBill, totalCashPaid))
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readBillData())
}

// ======================
//   UTILITY FUNCTIONS
// ======================

var reader = bufio.NewReader(os.Stdin)

func printError(prompt string) {
	fmt.Printf("❌ Error : %s\n", prompt)
}
