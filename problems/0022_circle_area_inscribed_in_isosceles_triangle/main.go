// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 22: circle_area_inscribed_in_isosceles_triangle.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Geometry / Mathematics | الهندسة / الرياضيات
// **Difficulty:** Medium | متوسط
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that calculates the area of a circle inscribed in an isosceles
// triangle given the equal side (A) and base (B).
// The program should:
// 1. Read the equal side length (A) and base (B) from the user
// 2. Calculate the area using the formula: Area = π × (B²/4) × ((2A-B)/(2A+B))
// 3. Display the computed area
//
// An isosceles triangle has two equal sides and one different side (base).
// The inscribed circle (incircle) touches all three sides of the triangle.
//
// AR:
// اكتب برنامجًا يحسب مساحة دائرة محاطة بمثلث متساوي الساقين
// معطى الضلع المتساوي (A) والقاعدة (B).
// يجب على البرنامج:
// 1. قراءة طول الضلع المتساوي (A) والقاعدة (B) من المستخدم
// 2. حساب المساحة باستخدام الصيغة: المساحة = π × (B²/4) × ((2A-B)/(2A+B))
// 3. عرض المساحة المحسوبة
//
// المثلث متساوي الساقين له ضلعان متساويان وضلع مختلف (القاعدة).
// الدائرة المحاطة (الدائرة الداخلية) تلامس جميع أضلاع المثلث الثلاثة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Side A = 5.0, Base B = 6.0
// Output: Circle Area ≈ 7.07
//   Why:  Area = π × (36/4) × ((10-6)/(10+6)) = π × 9 × (4/16) = π × 2.25 ≈ 7.07
//
// Example 2:
// Input:  Side A = 10.0, Base B = 8.0
// Output: Circle Area ≈ 21.54
//   Why:  Area = π × (64/4) × ((20-8)/(20+8)) = π × 16 × (12/28) ≈ 21.54
//
// Example 3 (Edge Case):
// Input:  Side A = 5.0, Base B = 5.0 (equilateral triangle)
// Output: Circle Area ≈ 6.54
//   Why:  Area = π × (25/4) × ((10-5)/(10+5)) = π × 6.25 × (5/15) ≈ 6.54
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Both A and B must be positive (A > 0, B > 0) | يجب أن يكون A و B موجبين
// • Must satisfy triangle inequality: 2A > B | يجب تحقق متباينة المثلث: 2A > B
// • Input should be valid floating-point numbers (float64) | يجب أن يكون الإدخال أرقاماً عشرية صحيحة
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readTriangleData() (float64, float64)
// func circleAreaInscribedInIsoscelesTriangle(sideA, baseB float64) (float64, error)
// func validateIsoscelesTriangle(sideA, baseB float64) bool
// func printResult(sideA, baseB float64)
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
			printError("Invalid input please enter a valid input!")
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

func readTriangleData() (float64, float64) {
	sideA := readPositiveNumber("Side A")
	baseB := readPositiveNumber("Base B")
	return sideA, baseB
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func circleAreaInscribedInIsoscelesTriangle(sideA, baseB float64) (float64, error) {
	if validateIsoscelesTriangle(sideA, baseB) {
		return math.Pi * (baseB * baseB / 4) * ((2*sideA - baseB) / (2*sideA + baseB)), nil
	}
	return 0, fmt.Errorf("satisfy triangle inequality you must enter '2A > B'")
}

func validateIsoscelesTriangle(sideA, baseB float64) bool {
	return 2*sideA > baseB
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult(sideA, baseB float64) {
	fmt.Println("\n=== Isosceles Triangle Information ===")
	area, err := circleAreaInscribedInIsoscelesTriangle(sideA, baseB)
	if err != nil {
		printError(err.Error())
		return
	}
	fmt.Printf("sideA:  %.2f\n", sideA)
	fmt.Printf("baseB:  %.2f\n", baseB)
	fmt.Printf("Area:   %.2f\n", area)
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
