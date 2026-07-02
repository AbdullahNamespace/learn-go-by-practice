// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 9: sum_of_3_numbers.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Basic Math / Functions | الرياضيات الأساسية / الدوال
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that reads three integer numbers from the user. The program
// should calculate the sum of these three numbers and print the total result
// in the format: "The total sum of numbers is: {sum}".
//
// العربية:
// اكتب برنامجاً يقرأ ثلاثة أرقام صحيحة من المستخدم. يجب على البرنامج
// حساب مجموع هذه الأرقام الثلاثة وطباعة النتيجة النهائية بالصيغة:
// "The total sum of numbers is: {المجموع}".
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  10, 20, 30
// Output: "The total sum of numbers is: 60"
//
// Example 2:
// Input:  5, 5, 5
// Output: "The total sum of numbers is: 15"
//
// Example 3 (Edge Case - Negatives):
// Input:  -10, 10, 0
// Output: "The total sum of numbers is: 0"
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Inputs are integers. | المدخلات أرقام صحيحة.
// • Standard integer range applies. | نطاق الأعداد الصحيحة القياسي ينطبق.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readNumbers() (int, int, int)
// func sumOf3Numbers(num1, num2, num3 int) int
// func printResults(total int)
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
	fmt.Printf("%s", fmt.Sprintf("%s :", prompt))
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func readNumber(prompt string) int {
	for {
		input, err := readString(prompt)
		if err != nil {
			fmt.Println("Error: invalid input please enter again!")
			continue
		}
		inputNumber, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error: invalid input please enter a number!")
			continue
		}
		return inputNumber
	}
}

func readNumbers() (int, int, int) {
	input1 := readNumber("enter number 1")
	input2 := readNumber("enter number 2")
	input3 := readNumber("enter number 3")
	return input1, input2, input3
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func sumOf3Numbers(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResults(total int) {
	fmt.Printf("\nThe total sum of numbers is: %d\n", total)
}

// ======================
//         MAIN
// ======================

func main() {
	printResults(sumOf3Numbers(readNumbers()))
}
