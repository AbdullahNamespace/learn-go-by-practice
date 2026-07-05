// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 29: sum_even_numbers_1_to_n.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Loops & Math | الحلقات والرياضيات
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that reads a positive number N from the user and calculates
// the sum of all even numbers from 1 to N using three different loop styles.
//
// AR:
// اكتب برنامجاً يقرأ رقماً موجباً N من المستخدم ويحسب مجموع جميع الأرقام الزوجية
// من 1 إلى N باستخدام ثلاثة أساليب مختلفة من الحلقات.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  10
// Output: 30
//   Why:    Even numbers from 1 to 10 are: 2, 4, 6, 8, 10 → sum = 30
//
// Example 2:
// Input:  5
// Output: 6
//   Why:    Even numbers from 1 to 5 are: 2, 4 → sum = 6
//
// Example 3 (Edge Case):
// Input:  2
// Output: 2
//   Why:    Only one even number (2) in the range
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • N must be a positive integer (N > 0) | يجب أن يكون N عدداً صحيحاً موجباً
// • Use three different loop styles | استخدم ثلاثة أساليب مختلفة من الحلقات
// • Only sum even numbers | اجمع الأرقام الزوجية فقط
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPositiveNumber() int
// func isEven(number int) bool
// func sumEvenUsingClassicFor(n int) int
// func sumEvenUsingWhileStyle(n int) int
// func sumEvenUsingInfiniteFor(n int) int
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

func readNumber(prompt string) int {
	for {
		input, err := readString(prompt)

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

func isEven(number int) bool {
	return number%2 == 0
}

func sumEvenUsingClassicFor(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		if isEven(i) {
			sum += i
		}
	}
	return sum
}

func sumEvenUsingWhileStyle(n int) int {
	i := 1
	sum := 0

	for i <= n {
		if isEven(i) {
			sum += i
		}
		i++
	}

	return sum
}

func sumEvenUsingInfiniteFor(n int) int {
	i := 1
	sum := 0

	for {
		if i > n {
			break
		}
		if isEven(i) {
			sum += i
		}
		i++
	}

	return sum
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(n int) {

	fmt.Print("\n1) Classic For Loop : ")
	fmt.Printf("%d\n", sumEvenUsingClassicFor(n))

	fmt.Print("2) While-Style Loop  : ")
	fmt.Printf("%d\n", sumEvenUsingWhileStyle(n))

	fmt.Print("3) Infinite For Loop : ")
	fmt.Printf("%d\n", sumEvenUsingInfiniteFor(n))
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readPositiveNumber("Enter Number"))
}

// ======================
//   UTILITY FUNCTIONS
// ======================

func printError(prompt string) {
	fmt.Printf("❌ Error: %s\n", prompt)
}
