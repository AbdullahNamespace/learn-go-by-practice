// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 31: power_of_2_3_4.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Mathematics | الرياضيات
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that reads a number from the user and calculates its
// square (power of 2), cube (power of 3), and fourth power (power of 4).
// Display all three results in a single line separated by spaces.
//
// AR:
// اكتب برنامجًا يقرأ رقمًا من المستخدم ويحسب مربعه (القوة 2)، ومكعبه (القوة 3)،
// والقوة الرابعة له. اعرض النتائج الثلاث في سطر واحد مفصولة بمسافات.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  2
// Output: 4 8 16
//   Why:    2² = 4, 2³ = 8, 2⁴ = 16
//
// Example 2:
// Input:  3
// Output: 9 27 81
//   Why:    3² = 9, 3³ = 27, 3⁴ = 81
//
// Example 3 (Edge Case):
// Input:  0
// Output: 0 0 0
//   Why:    Any power of 0 is 0
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Input must be an integer | يجب أن يكون المدخل عددًا صحيحًا
// • Number can be positive, negative, or zero | يمكن أن يكون الرقم موجبًا أو سالبًا أو صفرًا
// • Consider integer overflow for large numbers (use int64) | ضع في اعتبارك تجاوز الأعداد الصحيحة للأرقام الكبيرة (استخدم int64)
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readNumber() int
// func powerOf2_3_4(number int) (int64, int64, int64)
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
//   UTILITY
// ======================

func printError(prompt string) {
	fmt.Printf("❌ Error: %s\n", prompt)
}

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
			printError("invalid input please enter a valid number!")
			continue
		}
		return inputNumber
	}
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func powerOf2_3_4(number int) (int64, int64, int64) {
	var powerOf2 int64 = int64(number) * int64(number)
	var powerOf3 int64 = int64(number) * int64(number) * int64(number)
	var powerOf4 int64 = int64(number) * int64(number) * int64(number) * int64(number)

	return powerOf2, powerOf3, powerOf4
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(number int) {
	p2, p3, p4 := powerOf2_3_4(number)
	fmt.Printf("%s\n", strings.Repeat("=", 30))
	fmt.Printf("power number %d of 2, 3 and 4\n", number)
	fmt.Printf("%s\n", strings.Repeat("=", 30))
	fmt.Printf("The power number %d of 2 is : %d\n", number, p2)
	fmt.Printf("The power number %d of 3 is : %d\n", number, p3)
	fmt.Printf("The power number %d of 4 is : %d\n", number, p4)
	fmt.Printf("%s\n", strings.Repeat("=", 30))
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readNumber("Please enter number"))
}
