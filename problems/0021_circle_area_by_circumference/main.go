// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 21: circle_area_by_circumference.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Geometry / Mathematics | الهندسة / الرياضيات
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that calculates the area of a circle using its circumference (L).
// The program should:
// 1. Read the circumference (L) from the user
// 2. Calculate the area using the formula: Area = L² / (4π)
// 3. Display the computed area
//
// This formula is derived from the relationship between circumference (L = 2πr)
// and area (A = πr²). By eliminating r, we get A = L² / (4π).
//
// AR:
// اكتب برنامجًا يحسب مساحة دائرة باستخدام المحيط (L).
// يجب على البرنامج:
// 1. قراءة المحيط (L) من المستخدم
// 2. حساب المساحة باستخدام الصيغة: المساحة = L² / (4π)
// 3. عرض المساحة المحسوبة
//
// هذه الصيغة مشتقة من العلاقة بين المحيط (L = 2πr)
// والمساحة (A = πr²). بحذف r، نحصل على A = L² / (4π).
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Circumference L = 31.416
// Output: Circle Area ≈ 78.5398
//   Why:  Area = 31.416² / (4 × π) ≈ 986.917 / 12.56637 ≈ 78.5398
//
// Example 2:
// Input:  Circumference L = 62.83
// Output: Circle Area ≈ 314.16
//   Why:  Area = 62.83² / (4π) = 3947.6089 / 12.5664 ≈ 314.16
//
// Example 3 (Edge Case):
// Input:  Circumference L = 6.283185
// Output: Circle Area ≈ 3.141593
//   Why:  Area = (2π)² / (4π) = 4π² / 4π = π ≈ 3.141592653589793
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Circumference L must be positive (L > 0) | يجب أن يكون المحيط L موجباً
// • Input should be a valid floating-point number (float64) | يجب أن يكون الإدخال رقماً عشرياً صحيحاً
// • Result should maintain floating-point precision | يجب الحفاظ على دقة الأرقام العشرية
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readCircumference() float64
// func circleAreaByCircumference(circumference float64) float64
// func printResult(circumference float64)
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
			printError("Invalid input please enter a positive number!")
			continue
		}

		return input
	}
}

func readCircumference() float64 {
	return readPositiveNumber("Circumference L")
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func circleAreaByCircumference(circumference float64) float64 {
	return circumference * circumference / (4 * math.Pi)
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(circumference float64) {
	area := circleAreaByCircumference(circumference)
	fmt.Println("\n=== Circle Information ===")
	fmt.Printf("circumference:  %.2f\n", circumference)
	fmt.Printf("Area:   %.2f\n", area)
}

// ======================
//         MAIN
// ======================

func main() {
	printResult(readCircumference())
}

// ======================
//   UTILITY FUNCTIONS
// ======================

func printError(prompt string) {
	fmt.Printf("❌ Error: %s\n", prompt)
}
