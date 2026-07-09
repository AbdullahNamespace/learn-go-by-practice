// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 70: random_character_generator.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Math & Random & ASCII | الرياضيات والأرقام العشوائية ورموز ASCII
// **Difficulty:** Medium | متوسط
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that generates and prints a random character from each of the
// following categories: a lowercase letter (a-z), an uppercase letter (A-Z),
// a special character (e.g., !, @, #), and a digit (0-9). Use an enumeration
// (or constants) to define the character types and a unified function to fetch
// a random character based on the requested type.
//
// العربية:
// اكتب برنامجًا يقوم بتوليد وطباعة حرف عشوائي من كل فئة من الفئات التالية:
// حرف صغير (a-z)، حرف كبير (A-Z)، حرف خاص (مثل: !، @، #)، ورقم (0-9).
// استخدم تعدادًا (أو ثوابت) لتعريف أنواع الأحرف ودالة موحدة لجلب حرف عشوائي
// بناءً على النوع المطلوب.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1 (Possible Output):
// Input:  (No user input)
// Output:
// g
// D
// #
// 5
//
// Example 2 (Possible Output on another run):
// Input:  (No user input)
// Output:
// m
// Q
// !
// 0
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • No user input is required for this specific problem. | لا تتطلب هذه المشكلة مدخلات من المستخدم.
// • Special characters are limited to ASCII range 33 to 47. | الأحرف الخاصة محدودة بنطاق ASCII من 33 إلى 47.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// type CharType int
// const (SmallLetter CharType = iota + 1; CapitalLetter; SpecialCharacter; Digit)
// func randomNumber(from int, to int) int
// func getRandomCharacter(charType CharType) rune
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

type CharType int

const (
	SmallLetter CharType = iota + 1
	CapitalLetter
	SpecialCharacter
	Digit
	ArabicLetter
)

// ======================
//   PROCESSING FUNCTIONS
// ======================

func randomNumber(from, to int) int {
	return rand.IntN(to-from+1) + from
}

func getRandomCharacter(charType CharType) rune {
	switch charType {
	case ArabicLetter:
		return rune(randomNumber(1569, 1610))
	case SmallLetter:
		return rune(randomNumber(97, 122)) // a - z
	case CapitalLetter:
		return rune(randomNumber(65, 90)) // A - Z
	case SpecialCharacter:
		return rune(randomNumber(33, 47)) // ! - /
	case Digit:
		return rune(randomNumber(48, 57)) // 0 - 9
	default:
		return 0
	}
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResult() {
	// fmt.Println(string(getRandomCharacter(ArabicLetter)))
	fmt.Println(string(getRandomCharacter(SmallLetter)))
	fmt.Println(string(getRandomCharacter(CapitalLetter)))
	fmt.Println(string(getRandomCharacter(SpecialCharacter)))
	fmt.Println(string(getRandomCharacter(Digit)))
}

// ======================
//         MAIN
// ======================

func main() {
	printResult()
}
