// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 23: circumcircle_area.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Geometry | الهندسة
// **Difficulty:** Medium | متوسط
// **Source:** ProgrammingAdvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program to calculate the area of the circumscribed circle (circumcircle)
// around a triangle. The program reads the lengths of the three sides (a, b, c)
// and computes the circle area using the following derived formula:
//
// 1. Calculate semi-perimeter (P): P = (a + b + c) / 2
// 2. Calculate Triangle Area (S) using Heron's Formula: S = sqrt(P(P-a)(P-b)(P-c))
// 3. Calculate Circumradius (R): R = (a * b * c) / (4 * S)
// 4. Calculate Circle Area: Area = PI * R^2
//
// AR:
// اكتب برنامجاً لحساب مساحة الدائرة المحيطة (Circumcircle) بمثلث.
// يقرأ البرنامج أطوال الأضلاع الثلاثة (a, b, c) ويحسب مساحة الدائرة
// باستخدام المعادلات المشتقة التالية:
//
// 1. حساب نصف المحيط (P): P = (a + b + c) / 2
// 2. حساب مساحة المثلث (S) باستخدام صيغة هيرون: S = جذر(P(P-a)(P-b)(P-c))
// 3. حساب نصف قطر الدائرة المحيطة (R): R = (a * b * c) / (4 * S)
// 4. حساب مساحة الدائرة: المساحة = ط * نق^2
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  3 4 5
// Output: 19.6349...
//   Why:    P=6, S=6, R=30/24=2.5, Area ≈ 19.63
//
// Example 2:
// Input:  6 6 6
// Output: 37.6991...
//   Why:    P=9, S≈15.588, R≈3.464, Area ≈ 37.69
//
// Example 3:
// Input:  13 14 15
// Output: 207.388...
//   Why:    P=21, S=84, R=8.125, Area = pi * (8.125)^2 ≈ 207.39
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • 1 <= a, b, c <= 10^3 | الأضلاع أرقام موجبة
// • Valid Triangle (a+b>c, a+c>b, b+c>a) | المدخلات تشكل مثلثاً صحيحاً
// • Precision | النتيجة بدقة عشرية (float64)
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readTriangleData() (float64, float64, float64)
// func circumcircleArea(a, b, c float64) (float64, error)
// func validateTriangle(a, b, c float64) bool
// func printResult(a, b, c float64)
//
// ────────────────────────────────────────────────────────────────────────────
package main

import (
	"bufio"
	"fmt"
	"math"
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
			printError("Invalid input please enter again!")
			continue
		}

		inputNumber, err := strconv.ParseFloat(input, 64)

		if err != nil {
			printError("Invalid input please enter a valid number!")
			continue
		}

		return inputNumber
	}
}

func readPositiveNumber(prompt string) float64 {
	for {
		input := readNumber(prompt)

		if input <= 0 {
			printError("invalid input please enter a positive number!")
			continue
		}

		return input
	}
}

func readTriangleData() (float64, float64, float64) {
	a := readPositiveNumber("A")
	b := readPositiveNumber("B")
	c := readPositiveNumber("C")

	return a, b, c
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func circumcircleArea(a, b, c float64) (float64, error) {
	if !validateTriangle(a, b, c) {
		return 0, fmt.Errorf("Invalid Triangle Area 'Valid Triangle (a+b>c, a+c>b, b+c>a)'")
	}

	// 1. Calculate semi-perimeter (P): P = (a + b + c) / 2
	// 2. Calculate Triangle Area (S) using Heron's Formula: S = sqrt(P(P-a)(P-b)(P-c))
	// 3. Calculate Circumradius (R): R = (a * b * c) / (4 * S)
	// 4. Calculate Circle Area: Area = PI * R^2
	P := (a + b + c) / 2
	S := math.Sqrt(P * (P - a) * (P - b) * (P - c))
	R := (a * b * c) / (4 * S)
	area := math.Pi * (R * R)

	return area, nil
}

func validateTriangle(a, b, c float64) bool {
	if a+b < c {
		return false
	}
	if a+c < b {
		return false
	}
	if c+b < a {
		return false
	}
	return true
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(a, b, c float64) {
	fmt.Println("\n=== Circum Circle Information ===")
	area, err := circumcircleArea(a, b, c)

	if err != nil {
		printError(err.Error())
		return
	}

	fmt.Printf("a   : %.2f\n", a)
	fmt.Printf("b   : %.2f\n", b)
	fmt.Printf("c   : %.2f\n", c)
	fmt.Printf("Area: %.2f\n", area)
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readTriangleData())
}

// ======================
//   UTILITY FUNCTIONS
// ======================

func printError(prompt string) {
	fmt.Printf("❌ Error: %s\n", prompt)
}
