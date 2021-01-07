package main

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func main() {

	uid := uuid.New()
	no := strings.Replace(uid.String(), "-", "", -1)
	fmt.Println(uid.String())
	fmt.Println(no)
	fmt.Scanln()
}
