// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 110: matrix_sum_all_elements.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Matrices & Math | المصفوفات الثنائية والرياضيات
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that fills a 3x3 matrix with random numbers between 1 and
// 10, prints it, then computes and prints the sum of every element in the
// entire matrix (all rows and all columns combined).
//
// العربية:
// اكتب برنامجًا يملأ مصفوفة 3x3 بأرقام عشوائية بين 1 و10، يطبعها، ثم يحسب
// ويطبع مجموع كل عنصر في المصفوفة بأكملها (كل الصفوف والأعمدة مجتمعة).
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Matrix = [[1 2 3] [4 5 6] [7 8 9]]
// Output: Sum of Matrix1 is: 45
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Random values are in the range [1, 10] | القيم العشوائية في المدى [1، 10]
// • Must account for every cell exactly once (no double-counting, no
//   skipped cells) | يجب احتساب كل خلية مرة واحدة بالضبط (بدون احتساب مزدوج
//   أو تخطي خلايا)
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// const MatrixSize = 3
// func randomNumber(from, to int) int
// func newMatrix() [][]int
// func fillMatrixWithRandomNumbers(matrix [][]int)
// func matrixSum(matrix [][]int) int
// func printMatrix(matrix [][]int)
//
// ────────────────────────────────────────────────────────────────────────────
package main

import (
	"fmt"
	"math/rand/v2"
)

// ======================
//   UTILITY
// ======================

const MatrixSize = 3

func randomNumber(min, max int) int {
	if max < min {
		max, min = min, max
	}
	return rand.IntN(max-min+1) + min
}

// ======================
//     INPUT FUNCTIONS
// ======================

func newMatrix() [][]int {
	matrix := make([][]int, MatrixSize)
	for i := range matrix {
		matrix[i] = make([]int, MatrixSize)
	}
	return matrix
}

func fillMatrixWithRandomNumbers(matrix [][]int) {
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = randomNumber(1, 10)
		}
	}
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func matrixSum(matrix [][]int) int {
	sum := 0
	for i := range matrix {
		for j := range matrix[i] {
			sum += matrix[i][j]
		}
	}
	return sum
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printMatrix(matrix [][]int) {
	for _, valueI := range matrix {
		for _, valueJ := range valueI {
			fmt.Printf("%d\t", valueJ)
		}
		fmt.Println()
	}
}

// ======================
//         MAIN
// ======================

func main() {
	matrix := newMatrix()
	fillMatrixWithRandomNumbers(matrix)
	printMatrix(matrix)
	fmt.Printf("Sum of Matrix1 is: %d\n", matrixSum(matrix))
}
