// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 35: piggy_bank.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Structs & Math | الهياكل والرياضيات
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that reads the quantity of different coins and dollar bills in a
// piggy bank, calculates the total value in pennies, and then converts it to dollars.
//
// AR:
// اكتب برنامجًا يقرأ كمية العملات المعدنية المختلفة والأوراق النقدية في حصالة،
// ويحسب إجمالي القيمة بالبنسات، ثم يحولها إلى دولارات.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Pennies=10, Nickels=5, Dimes=3, Quarters=2, Dollars=1
// Output: Total Pennies = 165
//         Total Dollars = $1.65
//   Why:   (10*1) + (5*5) + (3*10) + (2*25) + (1*100) = 165
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Use a struct to store the piggy bank content | استخدم هيكلاً (struct) لتخزين محتويات الحصالة
// • Coin values: Penny=1, Nickel=5, Dime=10, Quarter=25, Dollar=100 | قيم العملات كما هو محددة
// • Inputs should be non-negative integers | المدخلات يجب أن تكون أعداد صحيحة غير سالبة
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// type PiggyBankContent struct { Pennies, Nickels, Dimes, Quarters, Dollars int }
// func readPiggyBankContent() PiggyBankContent
// func calculateTotalPennies(content PiggyBankContent) int
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

type PiggyBankContent struct {
	Pennies, Nickels, Dimes, Quarters, Dollars int
}

var reader = bufio.NewReader(os.Stdin)

const (
	NumberOfPenniesInNickel  = 5
	NumberOfPenniesInDime    = 10
	NumberOfPenniesInQuarter = 25
	NumberOfPenniesInDollar  = 100
)

// ======================
//     INPUT FUNCTIONS
// ======================

func ReadString(prompt string) (string, error) {
	fmt.Printf("%s : ", prompt)

	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

func ReadNumber(prompt string) int {
	for {
		input, err := ReadString(prompt)

		if err != nil {
			printError("Invalid input please enter again!")
			continue
		}

		inputNumber, err := strconv.Atoi(input)

		if err != nil {
			printError("Invalid input please enter a valid number!")
			continue
		}

		return inputNumber
	}
}

func readPositiveNumber(prompt string) int {
	for {
		input := ReadNumber(prompt)

		if input <= 0 {
			printError("Invalid input please enter a positive number!")
			continue
		}

		return input
	}
}

func (content *PiggyBankContent) readPiggyBankContent() {
	content.Pennies = readPositiveNumber("Please enter 'Pennies'")
	content.Nickels = readPositiveNumber("Please enter 'Nickel'")
	content.Dimes = readPositiveNumber("Please enter 'Dimes'")
	content.Quarters = readPositiveNumber("Please enter 'Quarters'")
	content.Dollars = readPositiveNumber("Please enter 'Dollars'")
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func calculateTotalPennies(content PiggyBankContent) int {
	numberOfPennies := 0

	numberOfPennies += content.Pennies                             // 10 = 10
	numberOfPennies += content.Nickels * NumberOfPenniesInNickel   // 10 + 5 *5 =35
	numberOfPennies += content.Dimes * NumberOfPenniesInDime       // 35 + 3 * 10 = 65
	numberOfPennies += content.Quarters * NumberOfPenniesInQuarter // 65 + 2 * 25 = 115
	numberOfPennies += content.Dollars * NumberOfPenniesInDollar   // 115 + 1

	return numberOfPennies
}

func (content *PiggyBankContent) PennToDollar() float64 {
	return float64(calculateTotalPennies(*content)) / NumberOfPenniesInDollar
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func (content *PiggyBankContent) printResult() {
	fmt.Printf("\nTotal Pennies = %d", calculateTotalPennies(*content))
	fmt.Printf("\nTotal Dollars = %.2f\n", content.PennToDollar())
}

// ======================
//         MAIN
// ======================

func main() {
	var p PiggyBankContent
	p.readPiggyBankContent()
	p.printResult()
}

// ======================
//   UTILITY FUNCTIONS
// ======================

func printError(prompt string) {
	fmt.Printf("❌ Error : %s\n", prompt)
}
