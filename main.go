package main

import (
	"fmt"

	"github.com/zzayne/Zayne.Blog/config"
)

func main() {

	fmt.Println(config.DBConfig.Port)
}
