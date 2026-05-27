// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 2: read_and_print_name.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** String Manipulation | معالجة النصوص
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that asks the user for their name and prints it back to them.
// The program must be capable of reading the full name, including spaces between words.
//
// العربية:
// اكتب برنامجًا يطلب من المستخدم اسمه ويقوم بطباعته له.
// يجب أن يكون البرنامج قادرًا على قراءة الاسم الكامل، بما في ذلك المسافات بين الكلمات.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Alice
// Output: Your Name is: Alice
//
// Example 2:
// Input:  John Doe
// Output: Your Name is: John Doe
//
// Example 3 (Arabic Name):
// Input:  أحمد محمد
// Output: Your Name is: أحمد محمد
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • The input must handle spaces (e.g., First Name + Last Name). | يجب أن تتعامل المدخلات مع المسافات (مثل الاسم الأول واسم العائلة).
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readName() string
// func printName(name string)
//
// ────────────────────────────────────────────────────────────────────────────

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// ======================
//     INPUT FUNCTIONS
// ======================

func readString(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func readName() string {
	return readString("\nPlease Enter Your Name: ")
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printName(name string) {
	fmt.Printf("Your Name is %s\n", name)
}

// ======================
//         MAIN
// ======================

func main() {
	printName(readName())
}
