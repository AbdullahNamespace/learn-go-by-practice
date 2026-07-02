// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 19: circle_area_by_diameter.go
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
// Write a program that calculates the area of a circle using its diameter (D).
// The program should:
// 1. Read the diameter (D) from the user
// 2. Calculate the area using the formula: Area = (π × D²) / 4
// 3. Display the computed area
//
// This formula is derived from the standard circle area formula (π × R²) by
// substituting R = D/2, which gives: Area = π × (D/2)² = (π × D²) / 4
//
// AR:
// اكتب برنامجًا يحسب مساحة دائرة باستخدام القطر (D).
// يجب على البرنامج:
// 1. قراءة القطر (D) من المستخدم
// 2. حساب المساحة باستخدام الصيغة: المساحة = (π × D²) / 4
// 3. عرض المساحة المحسوبة
//
// هذه الصيغة مشتقة من صيغة مساحة الدائرة القياسية (π × R²) بالتعويض
// R = D/2، مما يعطي: المساحة = π × (D/2)² = (π × D²) / 4
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Diameter D = 10.0
// Output: Circle Area ≈ 78.54
//   Why:  Area = (π × 10²) / 4 = (3.14159 × 100) / 4 = 78.53975
//         (Same as radius = 5.0)
//
// Example 2:
// Input:  Diameter D = 20.0
// Output: Circle Area ≈ 314.16
//   Why:  Area = (π × 20²) / 4 = (3.14159 × 400) / 4 = 314.159
//         (Same as radius = 10.0)
//
// Example 3 (Edge Case):
// Input:  Diameter D = 2.0
// Output: Circle Area ≈ 3.14
//   Why:  Area = (π × 2²) / 4 = (3.14159 × 4) / 4 = π ≈ 3.14159
//         (Unit circle with radius = 1.0)
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Diameter D must be positive (D > 0) | يجب أن يكون القطر D موجباً
// • Input should be a valid floating-point number (float64) | يجب أن يكون الإدخال رقماً عشرياً صحيحاً
// • Formula: Area = (π × D²) / 4 | الصيغة: المساحة = (π × D²) / 4
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readCircleDiameter() float64
// func calculateCircleAreaByDiameter(diameter float64) float64
// func printCircleArea(diameter float64)
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
			printError("Invalid input please enter a number greater than 0!")
			continue
		}

		return input
	}
}

func readCircleDiameter() float64 {
	return readPositiveNumber("Diameter D")
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func calculateCircleAreaByDiameter(diameter float64) float64 {
	return (math.Pi * math.Pow(diameter, 2)) / 4
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printCircleArea(diameter float64) {
	area := calculateCircleAreaByDiameter(diameter)
	fmt.Println("\n=== Circle Information ===")
	fmt.Printf("diameter:  %.2f\n", diameter)
	fmt.Printf("Area:   %.2f\n", area)
}

// ======================
//         MAIN
// ======================

func main() {
	printCircleArea(readCircleDiameter())
}

// ======================
//   UTILITY FUNCTIONS
// ======================

func printError(prompt string) {
	fmt.Printf("❌ Error: %s\n", prompt)
}
