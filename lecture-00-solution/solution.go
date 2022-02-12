package solution

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func GetMessage() string {
	hello := emoji.Sprint("Hello :world_map:")
	fmt.Println(hello)
}
