package press_enter_to_exit

import (
	"bufio"
	"fmt"
	"os"
)

func WaitForEnter() {
	fmt.Println("Press enter to exit...")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}