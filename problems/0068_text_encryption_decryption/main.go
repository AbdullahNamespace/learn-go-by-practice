// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 68: text_encryption_decryption.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Strings & ASCII | النصوص ورموز ASCII
// **Difficulty:** Medium | متوسط
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that reads a line of text from the user and encrypts it using
// a simple Caesar cipher by shifting each character's ASCII value by a specific
// encryption key. Then, it should decrypt the encrypted text back to the original
// by subtracting the same key. Finally, print the text before encryption, after
// encryption, and after decryption.
//
// العربية:
// اكتب برنامجًا يقرأ سطرًا نصيًا من المستخدم ويقوم بتشفيره باستخدام شفرة قيصر
// البسيطة عن طريق إزاحة قيمة كل حرف في جدول ASCII بمفتاح تشفير محدد.
// ثم يقوم بفك تشفير النص المشفر وإعادته إلى أصله بطرح نفس المفتاح.
// أخيرًا، اطبع النص قبل التشفير، وبعد التشفير، وبعد فك التشفير.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  "ABC" (Encryption Key = 2)
// Output:
// Text Before Encryption : ABC
// Text After Encryption  : CDE
// Text After Decryption  : ABC
//
// Example 2:
// Input:  "Hello" (Encryption Key = 2)
// Output:
// Text Before Encryption : Hello
// Text After Encryption  : Jgnnq
// Text After Decryption  : Hello
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • The input text can contain spaces (use appropriate reading methods). | النص المدخل يمكن أن يحتوي على مسافات (استخدم طرق قراءة مناسبة).
// • Encryption key is a constant integer (e.g., 2). | مفتاح التشفير هو عدد صحيح ثابت (مثلاً: 2).
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readString(prompt string) string
// func encryptText(text string, encryptionKey int) string
// func decryptText(text string, encryptionKey int) string
// func printResults(original, encrypted, decrypted string)
//
// ────────────────────────────────────────────────────────────────────────────
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ======================
//   UTILITY
// ======================

const KEY = 20441

func printError(prompt string) {
	fmt.Printf("X Error : %s\n", prompt)
}

var reader = bufio.NewReader(os.Stdin)

// ======================
//     INPUT FUNCTIONS
// ======================

func readString(prompt string) string {
	fmt.Printf("%s : ", prompt)

	for {
		input, err := reader.ReadString('\n')

		if err != nil {
			printError("Invalid input please enter again!")
			continue
		}

		return strings.TrimSpace(input)
	}
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func encryptText(text string, encryptionKey int) string {
	encrypt := []rune(text)

	for i := 0; i < len(encrypt); i++ {
		encrypt[i] += rune(encryptionKey)
	}

	return string(encrypt)
}

func decryptText(text string, encryptionKey int) string {
	decrypt := []rune(text)

	for i := 0; i < len(decrypt); i++ {
		decrypt[i] -= rune(encryptionKey)
	}

	return string(decrypt)
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printResults(original, encrypted, decrypted string) {
	fmt.Printf("Text Before Encryption : %s\n", original)
	fmt.Printf("Text After Encryption  : %s\n", encrypted)
	fmt.Printf("Text After Decryption  : %s\n", decrypted)
}

// ======================
//         MAIN
// ======================

func main() {
	original := readString("Please enter text")
	encrypted := encryptText(original, KEY)
	decrypted := decryptText(encrypted, KEY)
	printResults(original, encrypted, decrypted)
}
