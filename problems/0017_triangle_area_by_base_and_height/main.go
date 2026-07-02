// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 17: triangle_area_by_base_and_height.go
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
// Write a program that calculates the area of a triangle using its base (A)
// and height (H). The program should:
// 1. Read the base length (A) from the user
// 2. Read the height (H) from the user
// 3. Calculate the area using the formula: Area = (1/2) × Base × Height
// 4. Display the computed area
//
// This is the most common formula for calculating triangle area when the
// base and perpendicular height are known.
//
// AR:
// اكتب برنامجًا يحسب مساحة مثلث باستخدام القاعدة (A) والارتفاع (H).
// يجب على البرنامج:
// 1. قراءة طول القاعدة (A) من المستخدم
// 2. قراءة الارتفاع (H) من المستخدم
// 3. حساب المساحة باستخدام الصيغة: المساحة = (1/2) × القاعدة × الارتفاع
// 4. عرض المساحة المحسوبة
//
// هذه هي الصيغة الأكثر شيوعًا لحساب مساحة المثلث عندما تكون القاعدة والارتفاع العمودي معروفين.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Base A = 10.0, Height H = 5.0
// Output: Triangle Area = 25.0
//   Why:  Area = (1/2) × 10.0 × 5.0 = 5.0 × 5.0 = 25.0
//
// Example 2:
// Input:  Base A = 6.0, Height H = 8.0
// Output: Triangle Area = 24.0
//   Why:  Area = (1/2) × 6.0 × 8.0 = 3.0 × 8.0 = 24.0
//
// Example 3 (Edge Case):
// Input:  Base A = 1.0, Height H = 1.0
// Output: Triangle Area = 0.5
//   Why:  This is the smallest meaningful triangle
//         Area = (1/2) × 1.0 × 1.0 = 0.5
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Base A must be positive (A > 0) | يجب أن تكون القاعدة A موجبة
// • Height H must be positive (H > 0) | يجب أن يكون الارتفاع H موجباً
// • Both values should be valid floating-point numbers (float64) | يجب أن تكون القيم أرقاماً عشرية صحيحة
// • The result is always half the product of base and height | النتيجة دائماً نصف حاصل ضرب القاعدة والارتفاع
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readTriangleMeasurements() (float64, float64)
// func calculateTriangleArea(baseA, heightH float64) float64
// func printTriangleArea(baseA, heightH float64)
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
			printError("Invalid input please enter a valid input again!")
			continue
		}

		inputNumber, err := strconv.ParseFloat(input, 64)
		if err != nil {
			printError("Invalid input please enter a number!")
			continue
		}

		return inputNumber
	}
}

func readPositiveNumber(prompt string) float64 {
	for {
		input := readNumber(prompt)

		if input <= 0 {
			printError("Invalid input Please enter number greater than 0")
			continue
		}

		return input
	}
}

func readTriangleMeasurements() (float64, float64) {
	baseA := readPositiveNumber("base A")
	heightH := readPositiveNumber("height H")

	return baseA, heightH
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func calculateTriangleArea(baseA, heightH float64) float64 {
	return 0.5 * baseA * heightH
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printTriangleArea(baseA, heightH float64) {
	area := calculateTriangleArea(baseA, heightH)
	fmt.Println("\n=== Triangle Information ===")
	fmt.Printf("baseA:  %.2f\n", baseA)
	fmt.Printf("heightH: %.2f\n", heightH)
	fmt.Printf("Area:   %.2f\n", area)
}

// ======================
//         MAIN
// ======================

func main() {
	printTriangleArea(readTriangleMeasurements())
}

// ======================
//   UTILITY FUNCTIONS
// ======================

func printError(prompt string) {
	fmt.Printf("❌ Error: %s\n", prompt)
}
