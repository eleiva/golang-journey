package main

import (
	"bufio"
	"examples"
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"os"
	"reflect"
	"strings"
)

func main() {

	fmt.Println("[ASJ] A Shell Journey (by: eleiva@gmail.com)")
	fmt.Println("+------------------------------------------+")
	fmt.Println("| Enter `help` if you dont know what to do |")
	fmt.Println("+------------------------------------------+")

	CmdMainLoop()
}

func CmdMainLoop() {

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		substrings 	:= strings.Split(text, " ")

		switch substrings[0] {
		case "arrays":
			var array * examples.Array
			prompt(array)
			break
		case "reflection":
			var reflection * examples.Reflection
			prompt(reflection)
			break
		case "time":
			var time * examples.Time
			prompt(time)
			break
		case "strings":
			var string * examples.String
			prompt(string)
			break
		case "purpose":
			color.New(color.FgBlue).Println("The only reason of doing this project, is for helping me remember things on this great Go Journey!")
			break
		case "help":
			help()
			break
		case "exit", "quit":
			os.Exit(1)
		}
	}
}

func prompt(currentStruct interface{}) {

	options := []string{}

	fooType := reflect.TypeOf(currentStruct)

	for i := 0; i < fooType.NumMethod(); i++ {
		method := fooType.Method(i)
		options = append(options, method.Name)
	}

	prompt := promptui.Select{
		Label: "Select An Option",
		Items: options,
	}

	_, selected, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	output := reflect.ValueOf(currentStruct).MethodByName(selected).Call([]reflect.Value{})[0]

	red := color.New(color.FgRed).Add(color.Bold)

	red.Println(output)

}

func help() {
	green := color.New(color.FgGreen)
	fmt.Println("\nThe commands are: \n")
	green.Println("\t arrays \t A collection of examples initializing and playing with arrays")
	green.Println("\t reflection \t Cheatsheet for some reflective use cases")
	green.Println("\t strings \t Some examples manipulating strings")
	green.Println("\t purpose \t Why im doing this?")
	green.Println("\t help	\t This help \n")
}
