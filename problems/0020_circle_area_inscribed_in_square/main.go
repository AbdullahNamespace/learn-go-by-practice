// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 20: circle_area_inscribed_in_square.go
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
// Write a program that calculates the area of a circle inscribed inside a
// square. The circle touches all four sides of the square. The program should:
// 1. Read the side length (A) of the square from the user
// 2. Calculate the circle's area using the formula: Area = (π × A²) / 4
// 3. Display the computed area
//
// When a circle is inscribed in a square, the diameter of the circle equals
// the side length of the square (D = A). Therefore, the circle's radius is
// R = A/2, and using Area = π × R², we get: Area = π × (A/2)² = (π × A²) / 4
//
// AR:
// اكتب برنامجًا يحسب مساحة دائرة محاطة بمربع. الدائرة تلامس جميع أضلاع المربع
// الأربعة. يجب على البرنامج:
// 1. قراءة طول ضلع المربع (A) من المستخدم
// 2. حساب مساحة الدائرة باستخدام الصيغة: المساحة = (π × A²) / 4
// 3. عرض المساحة المحسوبة
//
// عندما تكون الدائرة محاطة بمربع، فإن قطر الدائرة يساوي طول ضلع المربع (D = A).
// لذلك، نصف قطر الدائرة هو R = A/2، وباستخدام المساحة = π × R²، نحصل على:
// المساحة = π × (A/2)² = (π × A²) / 4
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Square side A = 10.0
// Output: Circle Area ≈ 78.54
//   Why:  The inscribed circle has diameter = 10.0 (radius = 5.0)
//         Area = (π × 10²) / 4 = (3.14159 × 100) / 4 = 78.53975
//
// Example 2:
// Input:  Square side A = 20.0
// Output: Circle Area ≈ 314.16
//   Why:  The inscribed circle has diameter = 20.0 (radius = 10.0)
//         Area = (π × 20²) / 4 = (3.14159 × 400) / 4 = 314.159
//
// Example 3 (Edge Case):
// Input:  Square side A = 2.0
// Output: Circle Area ≈ 3.14
//   Why:  The inscribed circle has diameter = 2.0 (radius = 1.0)
//         Area = (π × 2²) / 4 = π ≈ 3.14159 (unit circle)
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Square side A must be positive (A > 0) | يجب أن يكون ضلع المربع A موجباً
// • The circle's diameter equals the square's side (D = A) | قطر الدائرة يساوي ضلع المربع
// • Input should be a valid floating-point number (float64) | يجب أن يكون الإدخال رقماً عشرياً صحيحاً
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readSquareSide() float64
// func calculateInscribedCircleArea(squareSide float64) float64
// func printCircleArea(squareSide float64)
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
			printError("Invalid input please enter number greater than 0!")
			continue
		}

		return input
	}
}

func readSquareSide() float64 {
	return readPositiveNumber("Square side A")
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func calculateInscribedCircleArea(squareSide float64) float64 {
	return (math.Pi * math.Pow(squareSide, 2)) / float64(4)
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printCircleArea(squareSide float64) {
	area := calculateInscribedCircleArea(squareSide)
	fmt.Println("\n=== Inscribed Circle Information ===")
	fmt.Printf("squareSide:  %.2f\n", squareSide)
	fmt.Printf("Area:   %.2f\n", area)
}

// ======================
//         MAIN
// ======================

func main() {
	printCircleArea(readSquareSide())
}

// ======================
//   UTILITY FUNCTIONS
// ======================

func printError(prompt string) {
	fmt.Printf("❌ Error: %s\n", prompt)
}
