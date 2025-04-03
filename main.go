package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

const (
	VERSION byte = 0x01
)

func GetLedStatus(s []byte) string {
	str := string(s)
	if str == "1" {
		return "ON"
	} else if str == "0" {
		return "OFF"
	}
	log.Fatalf(ByteParserError, str, s)
	return ""
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	p := tea.NewProgram(initModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
