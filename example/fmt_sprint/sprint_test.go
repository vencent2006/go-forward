/**
 * @Author: vincent
 * @Description:
 * @File:  sprint_test
 * @Version: 1.0.0
 * @Date: 2021/10/28 14:38
 */

package fmt_sprint

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestSprint1(t *testing.T) {
	// Declaring some const variables
	const name, dept = "GeeksforGeeks", "CS"

	// Calling Sprint() function
	s := fmt.Sprint(name, " is a ", dept, " Portal.\n")

	// Calling WriteString() function to write the
	// contents of the string "s" to "os.Stdout"
	io.WriteString(os.Stdout, s)
}

func TestSprint2(t *testing.T) {
	// Declaring some const variables
	const num1, num2, num3 = 5, 10, 15

	// Calling Sprint() function
	s := fmt.Sprint(num1, num2, num3)

	// Calling WriteString() function to write the
	// contents of the string "s" to "os.Stdout"
	io.WriteString(os.Stdout, s)
}

func TestSprint3(t *testing.T) {
	// Declaring some const variables
	const str1, str2, str3 = "a", "b", "c"

	// Calling Sprint() function
	s := fmt.Sprint(str1, str2, str3)

	// Calling WriteString() function to write the
	// contents of the string "s" to "os.Stdout"
	io.WriteString(os.Stdout, s)
}
