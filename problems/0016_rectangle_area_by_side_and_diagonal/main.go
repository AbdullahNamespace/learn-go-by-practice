// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 16: rectangle_area_by_side_and_diagonal.go
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
// Write a program that calculates the area of a rectangle using one side (A)
// and its diagonal (D). The program should:
// 1. Read the side length (A) from the user
// 2. Read the diagonal length (D) from the user
// 3. Calculate the area using the formula: Area = A × √(D² - A²)
// 4. Display the computed area
//
// The formula is derived from the Pythagorean theorem where the other side
// can be calculated as: B = √(D² - A²)
//
// AR:
// اكتب برنامجًا يحسب مساحة مستطيل باستخدام ضلع واحد (A) والقطر (D).
// يجب على البرنامج:
// 1. قراءة طول الضلع (A) من المستخدم
// 2. قراءة طول القطر (D) من المستخدم
// 3. حساب المساحة باستخدام الصيغة: المساحة = A × √(D² - A²)
// 4. عرض المساحة المحسوبة
//
// الصيغة مشتقة من نظرية فيثاغورس حيث يمكن حساب الضلع الآخر كـ: B = √(D² - A²)
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Side A = 3.0, Diagonal D = 5.0
// Output: Rectangle Area = 12.0
//   Why:  Other side B = √(5² - 3²) = √(25 - 9) = √16 = 4.0
//         Area = 3.0 × 4.0 = 12.0
//
// Example 2:
// Input:  Side A = 5.0, Diagonal D = 13.0
// Output: Rectangle Area = 60.0
//   Why:  Other side B = √(13² - 5²) = √(169 - 25) = √144 = 12.0
//         Area = 5.0 × 12.0 = 60.0
//
// Example 3 (Edge Case - Invalid Diagonal):
// Input:  Side A = 10.0, Diagonal D = 5.0
// Output: Error: Diagonal must be greater than side
//   Why:  A diagonal cannot be shorter than or equal to the side of a rectangle
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Side A must be positive (A > 0) | يجب أن يكون الضلع A موجباً
// • Diagonal D must be greater than side A (D > A) | يجب أن يكون القطر D أكبر من الضلع A
// • Input should be valid floating-point numbers (float64) | يجب أن تكون القيم أرقاماً عشرية صحيحة
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readRectangleMeasurements() (float64, float64)
// func calculateRectangleAreaBySideAndDiagonal(sideA, diagonalD float64) float64
// func printRectangleArea(sideA, diagonalD float64)
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

func readPositiveNumber(prompt string) float64 {
	for {
		input := readNumber(prompt)
		if input <= 0 {
			fmt.Println("Error: invalid input please enter a positive number!")
			continue
		}
		return input
	}
}

func readRectangleMeasurements() (float64, float64) {
	side := readPositiveNumber("Side A")

	for {
		diagonal := readPositiveNumber("Diagonal D")
		if diagonal <= side {
			fmt.Println("❌ Error: Diagonal must be greater than side A!")
			continue
		}
		return side, diagonal
	}
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func calculateRectangleAreaBySideAndDiagonal(sideA float64, diagonalD float64) float64 {
	b := math.Sqrt(math.Pow(diagonalD, 2) - math.Pow(sideA, 2))
	return sideA * b
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printRectangleArea(sideA, diagonalD float64) {
	area := calculateRectangleAreaBySideAndDiagonal(sideA, diagonalD)
	fmt.Println("\n=== Rectangle Information ===")
	fmt.Printf("sideA:  %.2f\n", sideA)
	fmt.Printf("diagonalD: %.2f\n", diagonalD)
	fmt.Printf("Area:   %.2f\n", area)
}

// ======================
//         MAIN
// ======================

func main() {
	printRectangleArea(readRectangleMeasurements())
}
