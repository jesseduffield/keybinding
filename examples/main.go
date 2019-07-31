package main

import (
	"fmt"

	"github.com/jesseduffield/keybinding"
)

// show will call the ParseAll function
func show(keyStr string) {
	keys, err := keybinding.ParseAll(keyStr)
	if err != nil {
		fmt.Println("Error parsing", keyStr, ":", err)
	} else {
		fmt.Println("Key: ", keyStr, "=", keys)
	}
}

// must will call the MustParseAll function
func must(keyStr string) {
	defer func() {
        if r := recover(); r != nil {
            fmt.Println("Error caught: ", r)
        }
	}()

	keys := keybinding.MustParseAll(keyStr)
	fmt.Println("Key: ", keyStr, "=", keys)
}

func main() {
	fmt.Println("The show calls:")
	show("ctrl+a")
	show("ctrl+b")
	show("ctrl+/, tab")
	show("ctrl+   alt +/")
	show("jibber+   jabber +/") // This will fail
	fmt.Println("The must calls:")
	must("ctrl+a")
	must("ctrl+b")
	must("ctrl+/, tab")
	must("ctrl+   alt +/")
	must("jibber+   jabber +/") // This will fail
}
