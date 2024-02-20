package main

import (
	"bufio"
	"fmt"
	"os"

	htgotts "github.com/hegedustibor/htgo-tts"
)

func main() {
	fmt.Println("Type something to say:")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Type text")
	text, _ := reader.ReadString('\n')

	fmt.Println("Name of the file to save the speech to:")
	var filename string
	fmt.Scanln(&filename)
	fmt.Println("You typed:", filename)
	var language string
	fmt.Println("Type language")
	fmt.Scanln(&language)
	speech := htgotts.Speech{Folder: "audio", Language: language}

	// Zapisywanie przemówienia pod konkretą nazwą pliku
	speech.CreateSpeechFile(text, filename)
	fmt.Println("Speech has been spoken and saved to:", filename)
	fmt.Println("Press enter to exit")
	fmt.Scanln()
}
