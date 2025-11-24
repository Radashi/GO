package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func clearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Slice(slc []int) string {
	var res strings.Builder
	for i := 0; i < len(slc); i++ {
		y := fmt.Sprintf("[%d x %d = %d]\n", slc[i], i, slc[i]*i)
		res.WriteString(y)
	}

	/*x := []int{1, 2, 3, 4}
	y := make([]int, 2)
	copy(y, x)
	fmt.Println(y)
	*/
	return res.String()
}

func Maps() {
	users := make(map[string][]string)
	users["users"] = []string{"admin", "guest"}

	for k, v := range users {
		for i := range len(users) {
			fmt.Printf("user %d:, %v %v\n", i+1, k, v)
		}
	}
}

func main() {
	/*slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(Slice(slc)); i++ {
		time.Sleep(100 * time.Millisecond)
		clearTerminal()
		fmt.Println(Slice(slc)[:i])
	}*/

	Maps()
}
