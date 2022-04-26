/*
Copyright © 2022 rfaychen justin928501@gmail.com

*/
package main

import (
	"bookworm/commands"
	"os"
)

func main() {
	commands.Execute(os.Args[1:])
}
