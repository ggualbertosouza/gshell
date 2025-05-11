package prompt

import (
	"fmt"
	"os"
)

func Prompt() {
	pwd, _ := os.Getwd()
	fmt.Printf("🚀 %s > ", Print(pwd, Yellow))
}
