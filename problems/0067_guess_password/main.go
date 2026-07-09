// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 67: guess_password.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Nested Loops & Strings | الحلقات المتداخلة والنصوص
// **Difficulty:** Medium | متوسط
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that asks the user to enter a 3-letter uppercase password.
// The program should then attempt to guess the password using a brute-force approach,
// iterating through all combinations from "AAA" to "ZZZ". It must print each trial
// number and the guessed word. When the password is found, it prints a success message
// including the number of trials it took.
//
// العربية:
// اكتب برنامجًا يطلب من المستخدم إدخال كلمة مرور مكونة من 3 أحرف كبيرة.
// يجب أن يحاول البرنامج تخمين كلمة المرور باستخدام أسلوب القوة الغاشمة (Brute-force)،
// بالمرور على جميع التركيبات من "AAA" إلى "ZZZ". يجب أن يطبع رقم كل محاولة والكلمة
// المخمّنة. عند العثور على كلمة المرور، يطبع رسالة نجاح تتضمن عدد المحاولات التي استغرقها.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  ABC
// Output:
// Trial [1] : AAA
// Trial [2] : AAB
// ...
// Trial [28] : ABC
//
// Password is ABC
// Found after 28 Trial(s)
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Password must be exactly 3 uppercase English letters. | كلمة المرور يجب أن تكون 3 أحرف إنجليزية كبيرة بالضبط.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPassword() string
// func guessPassword(originalPassword string) bool
//
// ────────────────────────────────────────────────────────────────────────────
package main

import (
	"fmt"
	"unicode"
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

func readPassword() string {
	for {
		input := readString("Please enter password 3 digit capital only")

		if len(input) != 3 {
			printError("Invalid input you must enter 3 digit only!")
			continue
		}

		if !isAllLetterIsCapital(input) {
			printError("Invalid input you must enter a capital letter only!")
			continue
		}

		return input
	}
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func isAllLetterIsCapital(input string) bool {
	for _, i := range input {
		if !unicode.IsUpper(i) {
			return false
		}
	}

	return true
}

func findPasswordIn3Digit(password string) (bool, int) {
	counter := 0
	for i := 65; i <= 90; i++ {
		for j := 65; j <= 90; j++ {
			for k := 65; k <= 90; k++ {
				counter++
				word := fmt.Sprintf("%c%c%c", i, j, k)
				fmt.Printf("Trial [%d] : %s\n", counter, word)
				if word == password {
					return true, counter
				}
			}
		}
	}
	return false, 0
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(password string) {
	found, counter := findPasswordIn3Digit(password)
	if found {
		fmt.Printf("Password is %s\nFound after %d Trial(s)\n", password, counter)
	}
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readPassword())
}
