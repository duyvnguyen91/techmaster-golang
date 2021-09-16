package main

import (
	"fmt"
	term "github.com/nsf/termbox-go"
	"os"
)

func reset() {
	term.Sync() // cosmestic purpose
}

func main() {
	err := term.Init()
	if err != nil {
		panic(err)
	}

	defer term.Close()

	fmt.Println("Enter any key to see their ASCII code or press ESC button to quit")
	fmt.Println("Press any key to see their ASCII code follow by Enter")

	//for i := 1; i <= 3; i++ {
	//	ev := term.PollEvent()
	//	if ev.Key == term.KeyF5 && i == 1 {
	//		fmt.Println("Running container")
	//	}
	//
	//	if ev.Key == term.KeyF5 && i == 2 {
	//		fmt.Println("Stop container")
	//	}
	//
	//	if ev.Key == term.KeyF5 && i == 3 {
	//		i = 0
	//		fmt.Println("All container")
	//	}
	//
	//	if ev.Key == term.KeyEsc || ev.Key == term.KeyCtrlC {
	//		fmt.Println("Exiting...")
	//		os.Exit(0)
	//	}
	//}
	i := 0
	for {
		ev := term.PollEvent()
		if ev.Key == term.KeyF5{
			switch i % 3 {
			case 0:
				fmt.Println("Running container")
			case 1:
				fmt.Println("Stop container")
			case 2:
				fmt.Println("All container")
			}
		} else if ev.Key == term.KeyEsc || ev.Key == term.KeyCtrlC{
			fmt.Println("Exiting...")
			os.Exit(0)
		}
		i ++
	}
}
