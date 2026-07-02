// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 27: print_range_n_to_1.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Loops & Iteration | الحلقات والتكرار
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that reads a positive number N from the user and prints all
// numbers from N down to 1 using three different loop styles in Go:
// 1. Classic C-style for loop
// 2. While-style for loop
// 3. Infinite for loop with break
//
// AR:
// اكتب برنامجاً يقرأ رقماً موجباً N من المستخدم ويطبع جميع الأرقام من N إلى 1
// (تنازلياً) باستخدام ثلاثة أساليب مختلفة من الحلقات في Go:
// 1. حلقة for الكلاسيكية
// 2. حلقة for بنمط while
// 3. حلقة for اللانهائية مع break
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  5
// Output: 5 4 3 2 1 (printed three times using different loops)
//   Why:    Program prints all numbers from 5 down to 1
//
// Example 2:
// Input:  3
// Output: 3 2 1 (printed three times using different loops)
//   Why:    Program prints all numbers from 3 down to 1
//
// Example 3 (Edge Case):
// Input:  1
// Output: 1 (printed three times using different loops)
//   Why:    When N=1, only one number is printed
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • N must be a positive integer (N > 0) | يجب أن يكون N عدداً صحيحاً موجباً
// • Use three different loop styles | استخدم ثلاثة أساليب مختلفة من الحلقات
// • Print numbers in descending order | اطبع الأرقام بترتيب تنازلي
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPositiveNumber() int
// func printRangeUsingClassicFor(n int)
// func printRangeUsingWhileStyle(n int)
// func printRangeUsingInfiniteFor(n int)
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
			printError("invalid input please enter a positive number!")
			continue
		}

		return input
	}
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printRangeUsingClassicFor(n int) {
	for i := n; i >= 1; i-- {
		fmt.Printf("%d \t", i)
	}
	fmt.Println()
}

func printRangeUsingWhileStyle(n int) {
	i := n

	for i >= 1 {
		fmt.Printf("%d \t", i)
		i--
	}
	fmt.Println()
}

func printRangeUsingInfiniteFor(n int) {
	i := n

	for {
		if i <= 0 {
			break
		}
		fmt.Printf("%d \t", i)
		i--
	}
	fmt.Println()
}

func printResult(n int) {
	fmt.Print("\n1) Classic For Loop:   ")
	printRangeUsingClassicFor(n)

	fmt.Print("2) While-Style Loop:   ")
	printRangeUsingWhileStyle(n)

	fmt.Print("3) Infinite For Loop:  ")
	printRangeUsingInfiniteFor(n)
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readPositiveNumber("Please enter number to print from it to 1"))
}

// ======================
//   UTILITY FUNCTIONS
// ======================

func printError(prompt string) {
	fmt.Printf("❌ Error: %s\n", prompt)
}
