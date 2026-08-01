// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 106: ordered_matrix.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Matrices | المصفوفات الثنائية
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that creates a 3x3 matrix and fills it with sequential
// numbers from 1 to 9 (row by row, left to right), then prints the matrix.
//
// العربية:
// اكتب برنامجًا يُنشئ مصفوفة 3x3 ويملأها بأرقام متسلسلة من 1 إلى 9 (صفًا
// صفًا، من اليسار لليمين)، ثم يطبع المصفوفة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Output: The following is a 3x3 ordered matrix:
//         1 2 3
//         4 5 6
//         7 8 9
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Numbers must run from 1 to MatrixSize*MatrixSize with no gaps or
//   repeats | يجب أن تمتد الأرقام من 1 إلى MatrixSize*MatrixSize بدون فجوات
//   أو تكرار
// • Filling order is row-major (complete row 1 before starting row 2) |
//   ترتيب التعبئة صفي (إكمال الصف 1 قبل بدء الصف 2)
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// const MatrixSize = 3
// func newMatrix() [][]int
// func fillMatrixWithOrderedNumbers(matrix [][]int)
// func printMatrix(matrix [][]int)
//
// ────────────────────────────────────────────────────────────────────────────
package main

import "fmt"

// ======================
//   UTILITY
// ======================

const MatrixSize = 3

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

// ======================
//   PROCESSING FUNCTIONS
// ======================

func fillMatrixWithOrderedNumbers(matrix [][]int) {
	counter := 1
	for i := range matrix {
		for j := range matrix[i] {
			//matrix[i][j] = i*len(matrix[i]) + j + 1
			matrix[i][j] = counter
			counter++
		}
	}
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
	fillMatrixWithOrderedNumbers(matrix)
	printMatrix(matrix)
}
