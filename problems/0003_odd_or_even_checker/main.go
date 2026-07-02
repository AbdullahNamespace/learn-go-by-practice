// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 3: odd_or_even_checker.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Control Flow, Constants | التحكم في التدفق، الثوابت
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that determines whether a user-entered number is odd or even
// using constants to represent types.
//
// العربية:
// اكتب برنامجًا يحدد ما إذا كان الرقم المدخل من المستخدم فرديًا أو زوجيًا
// باستخدام ثوابت لتمثيل الأنواع.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  8
// Output: Number is Even.
//
// Example 2:
// Input:  5
// Output: Number is Odd.
//
// Example 3 (Negative Number):
// Input:  -4
// Output: Number is Even.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Input is a standard integer (int). | المدخل عدد صحيح قياسي (int).
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// const (Odd, Even)
// func readNumber(prompt string) int
// func checkNumberType(num int) string
// func printNumberType(numberType string)
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

const (
	Odd  = "Odd"
	Even = "Even"
)

var reader = bufio.NewReader(os.Stdin)

// ======================
//     INPUT FUNCTIONS
// ======================

func readString(prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func readNumber() int {
	for {
		inputString, err := readString("\nPlease Enter a number: ")

		if err != nil {
			fmt.Println("Please Enter a valid number!")
			continue
		}

		inputNumber, err := strconv.Atoi(inputString)

		if err != nil {
			fmt.Println("Please Enter a valid number!")
			continue
		}
		return inputNumber
	}
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func checkNumberType(num int) string {
	if num%2 == 0 {
		return Even
	}
	return Odd
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printNumberType(numberType string) {
	fmt.Printf("\nNumber is %s\n", numberType)
}

// ======================
//         MAIN
// ======================

func main() {
	printNumberType(checkNumberType(readNumber()))
}
