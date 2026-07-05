// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 34: commission_percentage.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Math & Conditional Logic | الرياضيات والمنطق الشرطي
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that reads the total sales amount and calculates the commission
// based on specific percentage brackets, then displays the total commission amount.
//
// AR:
// اكتب برنامجًا يقرأ إجمالي المبيعات ويحسب العمولة بناءً على فئات نسبية محددة،
// ثم يعرض إجمالي مبلغ العمولة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  600000
// Output: Commission Percentage = 2%
//         Total Commission = 12000
//   Why:    600000 falls in the 500k-999k bracket (2%)
//
// Example 2:
// Input:  40000
// Output: Commission Percentage = 0%
//         Total Commission = 0
//   Why:    Less than 50k has 0% commission
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Total sales must be a positive float64 | إجمالي المبيعات يجب أن يكون رقم عشري موجب
// • Brackets: >=1M(1%), >=500K(2%), >=100K(3%), >=50K(5%), <50K(0%) | الفئات كما هو محدد
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPositiveNumber(message string) float64
// func getCommissionPercentage(totalSales float64) float64
// func calculateTotalCommission(totalSales float64) float64
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
			printError("Invalid input please enter positive number!")
			continue
		}

		return input
	}
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func getCommissionPercentage(totalSales float64) float64 {
	switch {
	case totalSales < 50000:
		return 0
	case totalSales >= 1000000:
		return 1
	case totalSales >= 500000:
		return 2
	case totalSales >= 100000:
		return 3
	default:
		return 5
	}
}
func calculateTotalCommission(totalSales float64) float64 {
	return totalSales * getCommissionPercentage(totalSales) / 100
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(totalSales float64) {
	fmt.Println(strings.Repeat("=", 30))
	fmt.Printf("Commission Percentage = %.0f%%\n", getCommissionPercentage(totalSales))
	fmt.Printf("Total Commission = %.2f\n", calculateTotalCommission(totalSales))

	fmt.Println(strings.Repeat("=", 30))
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readPositiveNumber("Please enter total sales"))
}

// ======================
//   UTILITY FUNCTIONS
// ======================

func printError(prompt string) {
	fmt.Printf("❌ Error : %s\n", prompt)
}
