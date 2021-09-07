package main

import "fmt"

type messgeType int

const (
	INFO messgeType = 0 + iota
	WARRNING
	ERROR
)

const (
	infoColor     = "\033[1;34m%s\033[0m"
	warrningColor = "\033[1;33m%s\033[0m"
	errorColor    = "\033[1;31m%s\033[0m"
)

func main() {
	showMessage(ERROR, "testMessage")
}

func showMessage(messge_type messgeType, message string) {
	switch messge_type {
	case INFO:
		printMessage := fmt.Sprintf("\nInformation: \n %s\n", message)
		fmt.Println(infoColor, printMessage)
	case WARRNING:
		printMessage := fmt.Sprintf("\nWarrning: \n %s\n", message)
		fmt.Println(infoColor, printMessage)
	case ERROR:
		printMessage := fmt.Sprintf("\nError: \n %s\n", message)
		fmt.Println(infoColor, printMessage)
	}
}
