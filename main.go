package main

import "github.com/zserge/webview"

func main() {
	// Open wikipedia in a 800x600 resizable window
	webview.Open("Todoist",
		"https://todoist.com", 800, 600, true)
}
