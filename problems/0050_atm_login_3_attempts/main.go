// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 50: atm_login_3_attempts.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Loops & Validation | الحلقات والتحقق
// **Difficulty:** Medium | متوسط
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that asks the user to enter a PIN code with a maximum of 3
// attempts. If the user enters the correct PIN, print the balance. If they fail
// 3 times, block the card and print a blocked message.
//
// AR:
// اكتب برنامجًا يطلب من المستخدم إدخال رقم PIN بحد أقصى 3 محاولات. إذا أدخل
// المستخدم الرقم الصحيح، اطبع الرصيد. إذا فشل 3 مرات، احظر البطاقة واطبع رسالة الحظر.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1 (Successful Login):
// Input:  0000
// Output: Wrong PIN, you have 2 more tries
// Input:  1234
// Output: Your account balance is 7500
//
// Example 2 (Blocked Card):
// Input:  0000
// Output: Wrong PIN, you have 2 more tries
// Input:  1111
// Output: Wrong PIN, you have 1 more tries
// Input:  2222
// Output: Your card is blocked. Call the bank for help.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Maximum of 3 attempts allowed | الحد الأقصى للمحاولات المسموح بها هو 3
// • The correct PIN code is "1234" | رقم PIN الصحيح هو "1234"
// • Account balance is 7500 | رصيد الحساب هو 7500
// • Print remaining attempts on failure | اطبع المحاولات المتبقية عند الفشل
// • Print blocked message after 3 failures | اطبع رسالة الحظر بعد 3 إخفاقات
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPinCode() string
// func loginWith3Attempts() bool
//
// ────────────────────────────────────────────────────────────────────────────
package main

import (
	"fmt"
)

// ======================
//   UTILITY
// ======================

const (
	MaximumOfAttempts = 3
	PIN               = "1234"
	AccountBalance    = 7500.0
)

func printError(prompt string) {
	fmt.Printf("❌ Error : %s\n", prompt)
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

func readPinCode() string {
	return readString("Please enter your PIN")
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func loginWith3Attempts() bool {
	for i := 1; i <= MaximumOfAttempts; i++ {
		input := readPinCode()

		if input == PIN {
			return true
		}

		fmt.Printf("Wrong PIN, you have %d more tries\n", MaximumOfAttempts-i)
	}

	return false
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(loginStatus bool) {
	if loginStatus {
		fmt.Printf("\nYour account balance is %.2f\n", AccountBalance)
	} else {
		fmt.Printf("\nYour card is blocked. Call the bank for help.\n")
	}
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(loginWith3Attempts())
}
