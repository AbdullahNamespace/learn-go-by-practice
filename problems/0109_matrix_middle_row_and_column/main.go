// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 109: matrix_middle_row_and_column.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Matrices & Arrays | المصفوفات الثنائية والأحادية
// **Difficulty:** Medium | متوسط
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that fills a 3x3 matrix with random numbers between 1 and
// 10, prints it, then extracts and prints the middle row and the middle
// column of the matrix as two separate 1D slices.
//
// العربية:
// اكتب برنامجًا يملأ مصفوفة 3x3 بأرقام عشوائية بين 1 و10، يطبعها، ثم
// يستخرج ويطبع الصف الأوسط والعمود الأوسط من المصفوفة كشريحتين أحاديتي
// البُعد منفصلتين.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Matrix = [[1 2 3] [4 5 6] [7 8 9]]
// Output: Middle Row of Matrix is: [4 5 6]
//         Middle Col of Matrix is: [2 5 8]
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Random values are in the range [1, 10] | القيم العشوائية في المدى [1، 10]
// • MatrixSize must be odd for a single, well-defined middle to exist |
//   يجب أن يكون MatrixSize فرديًا حتى يوجد وسط واحد مُعرَّف جيدًا
// • Both returned slices must have length equal to MatrixSize | يجب أن
//   يتساوى طول كلتا الشريحتين المُرجَعتين مع MatrixSize
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
// func middleRow(matrix [][]int) []int
// func middleColumn(matrix [][]int) []int
// func printMatrix(matrix [][]int)
// func printArray(arr []int)
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

func middleRow(matrix [][]int) []int {
	return matrix[len(matrix)/2]
}

func middleColumn(matrix [][]int) []int {
	column := make([]int, 0)
	mid := MatrixSize / 2
	for i := range matrix {
		column = append(column, matrix[i][mid])
	}
	return column
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

func printArray(prompt string, arr []int) {
	fmt.Printf("%s : %v\n", prompt, arr)
}

// ======================
//         MAIN
// ======================

func main() {
	matrix := newMatrix()
	fillMatrixWithRandomNumbers(matrix)
	printMatrix(matrix)
	printArray("Middle Col of Matrix is", middleColumn(matrix))
	printArray("Middle Row of Matrix is", middleRow(matrix))
}
