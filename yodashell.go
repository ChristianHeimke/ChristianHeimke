package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("WELCOME TO Yoda-Shell")
	fmt.Println("-")

	for {
		fmt.Print("YodaShell -->  ")

		// Tastatureingabe lesen

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Ausführung der Eingabe.
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

}

//// ErrNoPath wird zurückgegeben, wenn 'cd' ohne ein zweites Argument aufgerufen wurde.
var ErrNoPath = errors.New("Pfad erforderlich")

func execInput(input string) error {

	// Das Zeilenumbruchszeichen entfernen.
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	// Yerleşik komutları kontrol edin.
	switch args[0] {
	case "cd":

		if len(args) < 2 {
			return ErrNoPath
		}

		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}


	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
