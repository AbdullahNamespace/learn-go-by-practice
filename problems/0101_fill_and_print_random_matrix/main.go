// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 101: fill_and_print_random_matrix.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Matrices & Random | المصفوفات الثنائية والعشوائية
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that creates a 3x3 matrix, fills every cell with a random
// number between 1 and 100, then prints the matrix in a neatly aligned grid.
//
// In Go, a fixed-size matrix like this is represented as a slice of slices
// ([][]int). A [][]int must be explicitly allocated row by row (each inner
// slice created separately), so it's good practice to centralize that
// allocation in a single constructor function instead of repeating the
// allocation logic everywhere a matrix is needed. Since a slice already knows
// its own dimensions via len(matrix) and len(matrix[0]), there is no need to
// pass separate Rows/Cols parameters to every function that operates on it.
//
// العربية:
// اكتب برنامجًا يُنشئ مصفوفة 3x3، يملأ كل خلية فيها برقم عشوائي بين 1 و100،
// ثم يطبع المصفوفة بشكل مُصفّف ومُحاذى بدقة.
//
// في لغة غو، أفضل تمثيل لمصفوفة ثابتة الحجم كهذه هو شريحة من الشرائح
// ([][]int). يجب تخصيص [][]int بشكل صريح صفًا صفًا (كل شريحة داخلية تُنشأ على
// حدة)، لذا من الأفضل تجميع منطق التخصيص في دالة إنشاء (constructor) واحدة
// بدلاً من تكراره في كل مكان تُحتاج فيه مصفوفة. وبما أن الشريحة تعرف أبعادها
// بنفسها عبر len(matrix) و len(matrix[0])، فلا حاجة لتمرير معاملي Rows/Cols
// منفصلين لكل دالة تعمل عليها.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Output: The following is a 3x3 random matrix:
//          45  12  89
//           3  67  91
//          28  56   7
//
// (Note: Actual values vary due to randomness)
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • The matrix size is fixed at 3x3, defined via a named constant, not a
//   magic number scattered through the code | حجم المصفوفة ثابت 3x3، ويُعرَّف
//   عبر ثابت مُسمّى وليس رقمًا سحريًا متناثرًا في الكود
// • Random values are in the range [1, 100] | القيم العشوائية في المدى [1، 100]
// • Printed columns must be aligned regardless of digit count (1, 2, or 3
//   digits) | يجب أن تكون الأعمدة المطبوعة محاذاة بغض النظر عن عدد الأرقام
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

// ======================
//   PROCESSING FUNCTIONS
// ======================

func fillMatrixWithRandomNumbers(matrix [][]int) {
	for rowIndex, value := range matrix {
		for columnIndex := range value {
			matrix[rowIndex][columnIndex] = randomNumber(1, 100)
		}
	}
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, column := range row {
			fmt.Printf("%d\t", column)
		}
		fmt.Println()
	}
}

// ======================
//         MAIN
// ======================

func main() {
	m := newMatrix()
	fillMatrixWithRandomNumbers(m)
	printMatrix(m)
}
