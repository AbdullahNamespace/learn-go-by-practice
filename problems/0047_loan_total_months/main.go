// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 47: loan_total_months.go
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
// Write a program that reads a loan amount and a monthly installment from the user,
// then calculates the total number of months required to pay off the loan.
//
// AR:
// اكتب برنامجًا يقرأ مبلغ القرض ومبلغ القسط الشهري من المستخدم، ثم يحسب إجمالي
// عدد الأشهر المطلوبة لسداد القرض.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  LoanAmount = 5000, MonthlyInstallment = 500
// Output: Total Months to pay = 10
//   Why:    5000 / 500 = 10 months
//
// Example 2:
// Input:  LoanAmount = 1000, MonthlyInstallment = 150
// Output: Total Months to pay = 6.666...
//   Why:    1000 / 150 = 6.666...
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Both inputs must be positive numbers (> 0) | يجب أن تكون كلا المدخلين أرقامًا موجبة
// • Formula: Total Months = LoanAmount / MonthlyInstallment | الصيغة: الأشهر = القرض / القسط
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPositiveNumber(prompt string) float64
// func totalMonths(loanAmount, monthlyInstallment float64) float64
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

func readLoanTotalMonths() (float64, float64) {
	loanAmount := readPositiveNumber("Please enter loan amount")
	monthlyInstallment := readPositiveNumber("Please enter monthly installment")

	return loanAmount, monthlyInstallment
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func totalMonths(loanAmount, monthlyInstallment float64) float64 {
	return loanAmount / monthlyInstallment
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(loanAmount, monthlyInstallment float64) {
	fmt.Printf("Total Months to pay = %.2f\n", totalMonths(loanAmount, monthlyInstallment))
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readLoanTotalMonths())
}

// ======================
//   UTILITY FUNCTIONS
// ======================

var reader = bufio.NewReader(os.Stdin)

func printError(prompt string) {
	fmt.Printf("❌ Error : %s\n", prompt)
}
