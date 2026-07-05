// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 49: atm_login_unlimited.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Loops & Validation | الحلقات والتحقق
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that asks the user to enter a PIN code. If the PIN is "1234",
// print the account balance. If wrong, keep asking indefinitely until the correct
// PIN is entered.
//
// AR:
// اكتب برنامجًا يطلب من المستخدم إدخال رقم PIN. إذا كان الرقم "1234"، اطبع رصيد
// الحساب. إذا كان خاطئًا، استمر في السؤال إلى ما لا نهاية حتى يتم إدخال الرقم الصحيح.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  1111
// Output: Wrong PIN
// Input:  1234
// Output: Your account balance is 7500
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • The correct PIN code is "1234" | رقم PIN الصحيح هو "1234"
// • Account balance is 7500 | رصيد الحساب هو 7500
// • Use an infinite loop for wrong attempts | استخدم حلقة لا نهائية للمحاولات الخاطئة
// • Print error message for wrong PIN | اطبع رسالة خطأ عند إدخال PIN خاطئ
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPinCode(pin string) string
// func login(input, pin string) bool
//
// ────────────────────────────────────────────────────────────────────────────
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ======================
//   UTILITY
// ======================

const (
	PIN            = "1234"
	AccountBalance = 7500.0
)

var reader = bufio.NewReader(os.Stdin)

func printError(prompt string) {
	fmt.Printf("❌ Error : %s\n", prompt)
}

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

func readPinCode(pin string) string {
	for {
		input, err := readString("Please enter your PIN")

		if err != nil {
			printError("Invalid input please enter again!")
			continue
		}

		if login(input, pin) {
			return input
		}

		printError("Wrong PIN")
	}
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func login(input, pin string) bool {
	return input == pin
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(accountBalance float64) {
	fmt.Printf("Your account balance is %.2f", accountBalance)
}

// ======================
//         MAIN
// ======================

func main() {
	readPinCode(PIN)
	printResult(AccountBalance)
}
