// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 84: find_number_position_and_order_in_array.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Arrays & Searching | المصفوفات والبحث
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Write a program that fills an array with random numbers between 1 and 100,
// prints the array, then asks the user for a number to search for. The program
// searches the array and, if the number is found, prints its zero-based position
// as well as its "order" (position + 1, i.e. the 1-based rank). If the number is
// not found, it prints an appropriate message.
//
// العربية:
// اكتب برنامجًا يملأ مصفوفة بأرقام عشوائية بين 1 و 100، يطبع المصفوفة، ثم يسأل
// المستخدم عن رقم للبحث عنه. يبحث البرنامج في المصفوفة، وإذا وُجد الرقم، يطبع
// موقعه (يبدأ من الصفر) بالإضافة إلى "ترتيبه" (الموقع + 1، أي الترتيب من رقم 1).
// إذا لم يُوجد الرقم، يطبع رسالة مناسبة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  Array = [45 12 89 3 67], Number to search = 89
// Output: The number found at position: 2
//         The number found its order  : 3
//
// Example 2:
// Input:  Array = [10 20 30], Number to search = 99
// Output: The number is not found :-(
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
// • Array length must be a positive integer (> 0) | طول المصفوفة يجب أن يكون موجبًا
// • Return -1 as position when the number is not found | أرجع -1 كموقع عندما لا يُوجد الرقم
// • The search stops at the first matching occurrence | يتوقف البحث عند أول تطابق
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func fillArrayWithRandomNumbers(length int) []int
// func randomNumber(min, max int) int
// func findNumberPositionInArray(number int, arr []int) int
// func printSearchResult(position int)
// func printArray(arr []int)
//
// ────────────────────────────────────────────────────────────────────────────
package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

// ======================
//   UTILITY
// ======================

func printError(prompt string) {
	fmt.Printf("X Error :  %s\n", prompt)
}

// ======================
//     INPUT FUNCTIONS
// ======================

func readString(prompt string) string {
	fmt.Printf("%s : ", prompt)

	var input string

	fmt.Scan(&input)

	return input
}

func readNumber(prompt string) int {
	for {
		input := readString(prompt)

		inputNumber, err := strconv.Atoi(input)

		if err != nil {
			printError("Invalid input please enter a valid number!")
			continue
		}

		return inputNumber
	}
}

func readPositiveNumber(prompt string) int {
	for {
		input := readNumber(prompt)

		if input <= 0 {
			printError("Invalid input please enter a positive number!")
			continue
		}

		return input
	}
}

func fillArrayWithRandomNumbers(length int) []int {
	arr := make([]int, length)
	for key := range length {
		arr[key] = randomNumber(1, 100)
	}
	return arr
}

// ======================
//   PROCESSING FUNCTIONS
// ======================

func randomNumber(min, max int) int {
	return rand.IntN(max-min+1) + min
}

func findNumberPositionInArray(number int, arr []int) int {
	for key := range len(arr) {
		if arr[key] == number {
			return key
		}
	}

	return -1
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func printSearchResult(position int, number int) {
	fmt.Printf("Number you are looking for is : %d\n", number)
	if position != -1 {
		fmt.Printf("The number found at positions: %d\n", position)
		fmt.Printf("The number found its order: %d\n", position+1)

	} else {
		fmt.Println("The number is not found :-(")
	}
}

func printArray(arr []int) {
	fmt.Printf("Array = %v\n", arr)
}

// ======================
//         MAIN
// ======================

func main() {
	lengthOfArray := readPositiveNumber("Please enter numbers of element")

	arr := fillArrayWithRandomNumbers(lengthOfArray)

	printArray(arr)

	numberSearch := readPositiveNumber("Please enter a number to search for")

	index := findNumberPositionInArray(numberSearch, arr)

	printSearchResult(index, numberSearch)
}
