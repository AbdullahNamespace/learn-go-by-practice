// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 48: loan_monthly_installment.go
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
// Write a program that reads a loan amount and the number of months for repayment,
// then calculates the required monthly installment.
//
// AR:
// اكتب برنامجًا يقرأ مبلغ القرض وعدد الأشهر للسداد، ثم يحسب القسط الشهري المطلوب.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  LoanAmount = 5000, HowManyMonths = 10
// Output: Monthly Installment = 500
//   Why:    5000 / 10 = 500
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Both inputs must be positive numbers (> 0) | يجب أن تكون كلا المدخلين أرقامًا موجبة
// • Formula: MonthlyInstallment = LoanAmount / HowManyMonths | الصيغة: القسط = القرض / عدد الأشهر
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPositiveNumber(message string) float64
// func monthlyInstallment(loanAmount, howManyMonths float64) float64
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

func readLoanMonthlyInstallment() (float64, float64) {
	loanAmount := readPositiveNumber("Please enter loan amount")
	howManyMonths := readPositiveNumber("Please enter how many months")

	return loanAmount, howManyMonths
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func monthlyInstallment(loanAmount, howManyMonths float64) float64 {
	return loanAmount / howManyMonths
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(loanAmount, howManyMonths float64) {
	fmt.Printf("Monthly Installment = %.2f\n", monthlyInstallment(loanAmount, howManyMonths))
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readLoanMonthlyInstallment())
}

// ======================
//   UTILITY FUNCTIONS
// ======================

var reader = bufio.NewReader(os.Stdin)

func printError(prompt string) {
	fmt.Printf("❌ Error : %s\n", prompt)
}
