// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 18: circle_area_by_radius.go
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
// Write a program that calculates the area of a circle using its radius (R).
// The program should:
// 1. Read the radius (R) from the user
// 2. Calculate the area using the formula: Area = π × R²
// 3. Display the computed area
//
// Use the standard library math.Pi for precise calculations.
// This is the fundamental formula for calculating the area of a circle.
//
// AR:
// اكتب برنامجًا يحسب مساحة دائرة باستخدام نصف القطر (R).
// يجب على البرنامج:
// 1. قراءة نصف القطر (R) من المستخدم
// 2. حساب المساحة باستخدام الصيغة: المساحة = π × R²
// 3. عرض المساحة المحسوبة
//
// استخدم الثابت math.Pi من المكتبة القياسية لحسابات دقيقة.
// هذه هي الصيغة الأساسية لحساب مساحة الدائرة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Radius R = 5.0
// Output: Circle Area ≈ 78.54
//   Why:  Area = π × 5² = 3.14159265 × 25 = 78.539816
//
// Example 2:
// Input:  Radius R = 10.0
// Output: Circle Area ≈ 314.16
//   Why:  Area = π × 10² = 3.14159265 × 100 = 314.159265
//
// Example 3 (Edge Case):
// Input:  Radius R = 1.0
// Output: Circle Area ≈ 3.14
//   Why:  Area = π × 1² = π ≈ 3.141593 (unit circle)
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Radius R must be positive (R > 0) | يجب أن يكون نصف القطر R موجباً
// • Use math.Pi from the standard library for precision | استخدم math.Pi من المكتبة القياسية للدقة
// • Input should be a valid floating-point number (float64) | يجب أن يكون الإدخال رقماً عشرياً صحيحاً
// • Result should maintain floating-point precision | يجب الحفاظ على دقة الأرقام العشرية
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readCircleRadius() float64
// func calculateCircleAreaByRadius(radius float64) float64
// func printCircleArea(radius float64)
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
			printError("Invalid input please enter a valid input!")
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
			printError("Invalid input please enter a number greater than 0!")
			continue
		}

		return input
	}
}

func readCircleRadius() float64 {
	return readPositiveNumber("Radius R")
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func calculateCircleAreaByRadius(radius float64) float64 {
	return math.Pi * math.Pow(radius, 2)
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printCircleArea(radius float64) {
	area := calculateCircleAreaByRadius(radius)
	fmt.Println("\n=== Circle Information ===")
	fmt.Printf("radius:  %.2f\n", radius)
	fmt.Printf("Area:   %.2f\n", area)
}

// ======================
//         MAIN
// ======================

func main() {
	printCircleArea(readCircleRadius())
}

// ======================
//   UTILITY FUNCTIONS
// ======================

func printError(prompt string) {
	fmt.Printf("❌ Error: %s\n", prompt)
}
