// ════════════════════════════════════════════════════════════════════════════
// ## 🔷 Problem 6: full_name_constructor.go
// ════════════════════════════════════════════════════════════════════════════
// **Category:** Basic Programming / Structs | البرمجة الأساسية / الهياكل
// **Difficulty:** Easy | سهل
// **Source:** programmingadvices.com
// ════════════════════════════════════════════════════════════════════════════

// ────────────────────────────────────────────────────────────────────────────
// 📝 DESCRIPTION | الوصف
// ────────────────────────────────────────────────────────────────────────────
//
// English:
// Create a program to manage personal information. Define a struct `PersonInfo` to hold
// first and last names. Read these names from the user and combine them into a
// full name string. The program should use a function `GetFullName` that accepts
// a boolean (`reversed`) to determine the order (First + Last if false, Last + First if true).
// Finally, print the result.
//
// العربية:
// إنشاء برنامج لإدارة المعلومات الشخصية. تعريف هيكل `PersonInfo` لتخزين الاسم الأول واسم العائلة.
// قراءة هذه المعلومات من المستخدم ودمج الأسماء في سلسلة نصية واحدة.
// يجب أن يستخدم البرنامج دالة `GetFullName` تقبل قيمة منطقية لتحديد ترتيب الأسماء
// (الأول ثم العائلة إذا كانت false، والعائلة ثم الأول إذا كانت true). أخيراً طباعة النتيجة.
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 💡 EXAMPLES | الأمثلة
// ────────────────────────────────────────────────────────────────────────────
//
// Example 1:
// Input:  First Name: "Ali", Last Name: "Mohamed", Reversed: true
// Output: "Mohamed Ali"
//
// Example 2:
// Input:  First Name: "Sara", Last Name: "Smith", Reversed: false
// Output: "Sara Smith"
//
// Example 3 (Edge Case):
// Input:  First Name: "A", Last Name: "B", Reversed: true
// Output: "B A"
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// ⚠️ CONSTRAINTS | القيود
// ────────────────────────────────────────────────────────────────────────────
//
//
// ────────────────────────────────────────────────────────────────────────────

// ────────────────────────────────────────────────────────────────────────────
// 🔧 FUNCTION SIGNATURES | توقيعات الدوال
// ────────────────────────────────────────────────────────────────────────────
//
// func readPersonInfo() PersonInfo
// func getFullName(info PersonInfo, reversed bool) string
// func printFullName(fullName string)
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

type PersonInfo struct {
	FirstName string
	LastName  string
}

func readString(prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func readValidName(prompt string) string {
	for {
		input, err := readString(prompt)

		if err != nil {
			fmt.Println("Error: invalid input please enter again!")
			continue
		}
		if input == "" {
			fmt.Println("Error: invalid name please enter again not empty!")
			continue
		}
		return input
	}
}

func readBoolean(prompt string) bool {
	for {
		input, err := readString(fmt.Sprintf("%s (1 for Yes, 0 for No): ", prompt))

		if err != nil {
			fmt.Println("Error: invalid input please enter again!")
			continue
		}
		input = strings.ToUpper(input)
		switch input {
		case "Y", "YES", "T", "TRUE", "1":
			{
				return true
			}
		case "N", "NO", "F", "FALSE", "0":
			{
				return false
			}
		default:
			{
				fmt.Println("Please enter 'y' for yes or 'n' for no.")
			}
		}
	}
}

func NewPersonInfo() *PersonInfo {
	return &PersonInfo{
		FirstName: readValidName("Please enter your first name: "),
		LastName:  readValidName("Please enter your last name: "),
	}
}

// ======================
//	PROCESSING FUNCTIONS
// ======================

func getFullName(firstName string, lastName string, reverse bool) string {
	if reverse {
		return fmt.Sprint(lastName, " ", firstName)
	}
	return fmt.Sprint(firstName, " ", lastName)
}

// ======================
//     OUTPUT FUNCTIONS
// ======================

func (p *PersonInfo) printFullName() {
	reverse := readBoolean("\nDo you want to reverse name: ")
	fmt.Println(getFullName(p.FirstName, p.LastName, reverse))
}

// ======================
//         MAIN
// ======================

func main() {
	person := NewPersonInfo()
	person.printFullName()
}
