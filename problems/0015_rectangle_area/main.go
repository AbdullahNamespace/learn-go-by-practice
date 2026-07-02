// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 15: rectangle_area.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Basic Math / Geometry | الرياضيات الأساسية / الهندسة
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that calculates the area of a rectangle. The program should
// read the width (A) and length (B) of a rectangle from the user as floating-point
// numbers, calculate the area using the formula Area = Width × Length, and display
// the computed area. This demonstrates basic geometric calculations with floating-point
// arithmetic and proper handling of decimal values.
//
// AR:
// اكتب برنامجًا يحسب مساحة المستطيل. يجب على البرنامج قراءة العرض (A) والطول (B)
// للمستطيل من المستخدم كأرقام عشرية، وحساب المساحة باستخدام الصيغة المساحة = العرض × الطول،
// وعرض المساحة المحسوبة. هذا يوضح الحسابات الهندسية الأساسية مع العمليات الحسابية
// العشرية والتعامل الصحيح مع القيم العشرية.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Width = 5.0, Length = 10.0
// Output: Rectangle Area = 50
//   Why:    5.0 × 10.0 = 50.0
//
// Example 2:
// Input:  Width = 3.14, Length = 2.5
// Output: Rectangle Area = 7.85
//   Why:    3.14 × 2.5 = 7.85
//
// Example 3 (Edge Case):
// Input:  Width = 6.0, Length = 6.0
// Output: Rectangle Area = 36
//   Why:    Square is special rectangle, 6 × 6 = 36
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Dimensions should be positive floating-point numbers (float64) | الأبعاد يجب أن تكون أرقام عشرية موجبة
// • Use floating-point arithmetic (float64) | استخدام الحساب العشري
// • Formula: Area = Width × Length | الصيغة: المساحة = العرض × الطول
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readDimensions() (float64, float64)
// func calculateRectangleArea(width, length float64) float64
// func printResult(width, length float64)
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

func readNumber(prompt string) float64 {
	for {
		input, err := readString(prompt)
		if err != nil {
			fmt.Println("Error: invalid input please enter again!")
			continue
		}
		inputNumber, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Error: invalid input please enter a valid number!")
			continue
		}
		return inputNumber
	}
}

func readPositiveFloat(prompt string) float64 {
	for {
		input := readNumber(fmt.Sprintf("%s a positive float number", prompt))
		if input < 0 {
			fmt.Println("Error: invalid positive float please enter a positive number!")
			continue
		}
		return input
	}
}

func readDimensions() (float64, float64) {
	width := readPositiveFloat("Enter width")
	length := readPositiveFloat("Enter length")
	return width, length
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func calculateRectangleArea(width, length float64) float64 {
	return width * length
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(width, length float64) {
	area := calculateRectangleArea(width, length)
	fmt.Println("\n=== Rectangle Information ===")
	fmt.Printf("Width:  %.2f\n", width)
	fmt.Printf("Length: %.2f\n", length)
	fmt.Printf("Area:   %.2f\n", area)
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readDimensions())
}
