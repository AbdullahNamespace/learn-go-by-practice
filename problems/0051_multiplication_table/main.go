// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 51: multiplication_table.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Loops & Tables | الحلقات والجداول
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// EN:
// Write a program that prints a multiplication table from 1 to 10. The table should
// include a formatted header with column labels (1-10), a separator line, and each
// row should display the row number followed by the products of that row with
// columns 1 through 10.
//
// AR:
// اكتب برنامجًا يطبع جدول ضرب من 1 إلى 10. يجب أن يتضمن الجدول رأسًا منسقًا
// مع تسميات الأعمدة (1-10)، وخط فاصل، وكل صف يعرض رقم الصف متبوعًا بمنتجات
// ذلك الصف مع الأعمدة من 1 إلى 10.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  (No input required)
// Output:
//           Multiplication Table From 1 to 10
//
//           1   2   3   4   5   6   7   8   9   10
//   ___________________________________________________
//    1 |   1   2   3   4   5   6   7   8   9   10
//    2 |   2   4   6   8  10  12  14  16  18   20
//    ...
//   10 |  10  20  30  40  50  60  70  80  90  100
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Table is fixed from 1 to 10 | الجدول ثابت من 1 إلى 10
// • Use nested loops for rows and columns | استخدم حلقات متداخلة للصفوف والأعمدة
// • Format spacing adjusts for single-digit vs double-digit row numbers | تنسيق المسافات يتكيف مع أرقام الصفوف ذات رقم واحد أو رقمين
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func printTableHeader()
// func columnSeparator(row int) string
// func printMultiplicationTable()
//
// ────────────────────────────────────────────────────────────────────────────
package main

import (
	"fmt"
	"strings"
)

// ======================
//     OUTPUT FUNCTIONS
// ======================

func columnSeparator(row int) string {
	if row >= 10 {
		return fmt.Sprintf("%d  |\t", row)
	}
	return fmt.Sprintf("%d   |\t", row)
}

func printTableHeader() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("\t%d", i)
	}
	fmt.Printf("\n%s\n", strings.Repeat("_", 90))
}

func printTableBody() {
	for i := 1; i <= 10; i++ {

		fmt.Printf("%s", columnSeparator(i))

		for j := 1; j <= 10; j++ {
			fmt.Printf("%d\t", i*j)
		}

		fmt.Println()
	}
}

func printMultiplicationTable() {
	printTableHeader()
	printTableBody()
}

// ======================
//         MAIN
// ======================

func main() {
	printMultiplicationTable()
}
