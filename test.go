package main
import	(
	"fmt"
	"time"
)

func main() {

	fmt.Println("Hello World")
	<-time.After(5*time.Second)
	fmt.Println("Hello Again for GO")
}


