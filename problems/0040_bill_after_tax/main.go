// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 40: bill_after_tax.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Math & Percentages | الرياضيات والنسب المئوية
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that reads a total bill and calculates the final bill after
// applying a 10% service charge and a 16% tax.
//
// AR:
// اكتب برنامجًا يقرأ إجمالي الفاتورة ويحسب الفاتورة النهائية بعد إضافة
// رسوم خدمة بنسبة 10٪ وضريبة بنسبة 16٪.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  100
// Output: Total Bill = 100
//         Total Bill After Service and Tax = 127.6
//   Why:    100 * 1.1 (service) = 110 -> 110 * 1.16 (tax) = 127.6
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Bill must be a positive number | الفاتورة يجب أن تكون رقماً موجباً
// • Apply service charge first, then tax on the new amount | تطبيق رسوم الخدمة أولاً، ثم الضريبة على المبلغ الجديد
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPositiveNumber(message string) float64
// func totalBillAfterServiceAndTax(totalBill float64) float64
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

// ======================
//   PROCESSING FUNCTIONS
// ======================

func totalBillAfterServiceAndTax(totalBill float64) float64 {
	return totalBill * Service * Tax
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

// Output: Total Bill = 100
//
//	Total Bill After Service and Tax = 127.6
func printResult(totalBill float64) {
	fmt.Printf("\nTotal Bill = %.2f", totalBill)
	fmt.Printf("\nTotal Bill After Service and Tax = %.2f\n", totalBillAfterServiceAndTax(totalBill))
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readPositiveNumber("Please enter total bill"))
}

// ======================
//   UTILITY FUNCTIONS
// ======================

const (
	Service = 1.1
	Tax     = 1.16
)

var reader = bufio.NewReader(os.Stdin)

func printError(prompt string) {
	fmt.Printf("❌ Error : %s\n", prompt)
}
