/*
 * GuruMUD - A lightweight, fast MUD client in Go
**/

package main

import (
	"fmt"
	"net"

	"tmpfs/alias"
	"tmpfs/trigger"
)

const (
	version = 1.00
)

var (
	conn net.Conn
	isConnected bool

	aliases []alias
	triggers []trigger
)

func main() {
	fmt.Printf(">")
	for {
		fmt.Println(input())
	}
}

// Show header and version
func header() {
	fmt.Println("-------- GuruMUD --------")
	fmt.Printf("Version: %v\n", version)
	fmt.Println("-------------------------")
}

// Read one line of input from stdin. Blocking.
func input() string {
	var bfr string
	fmt.Scanf("%s\n", &bfr)

	return bfr
}

// Parse a string and execute it accordingly
func parse(s string) string {
	if len(s)>1 {
		if s[0]=="#" {
			kwarg:=strings.Split(standardize(s), " ")
			switch strings.ToLower(kwarg[0][1:len(s)-1]) {
				case "alias":
					cmd_alias(kwarg)
				case "commands":
					cmd_commands()
				case "help":
					cmd_help(kwarg)
				case "trigger":
					cmd_trigger(kwarg)
			}
		}
	}else{if isConnected{send(s)}}
}

// Standardize a string to use uniform whitespace
func standardize(s string) {
	for i:=len(s);i>0;i-- {
		s=strings.Replace(strings.Repeat(" ",i)," ")
	}

	return s
}

/// CMD FUNCTIONS ///

func cmd_commands() {
	fmt.Println("=============== COMMANDS ===============")
	fmt.Println("#alias - add/remove/show aliases")
	fmt.Println("#commands - shows this list")
	fmt.Println("#connect - connect to a MUD")
	fmt.Println("#end/#exit - exit")
	fmt.Println("#help - show help on a command")
	fmt.Println("#trigger - add/remove/show triggers")
	fmt.Println("========================================")
}

func cmd_help(kwarg []string) {
	if len(kwarg)==0 {
		fmt.Println("==== HELP ====")
		fmt.Println("Command: help <topic>"
		fmt.Println("--")
		fmt.Println("Displays help on a specific topic.")
	} else {
		fmt.Println("==== HELP: "+kwarg[1]+" ====")
		switch strings.ToLower(kwarg[1]) {
			case "alias":
				fmt.Println("Command: #alias list - list aliases")
				fmt.Println("Command: #alias add {aliasName} {command(s)} - add an alias")
				fmt.Println("Command: #alias remove <aliasName> - remove an alias")
				fmt.Println("--")
				fmt.Println("Aliases are things that you may type, which instead of being sent,")
				fmt.Println("will be replaced by the alias' command(s).")
				fmt.Println("--")
				fmt.Println("Example:")
				fmt.Println("'#alias add killAndLoot kill mob; loot corpse': when you type 'killAndLoot' the")
				fmt.Println("client shall send 'kill mob' and then 'loot corpse' instead of 'killAndLoot'.")
			case "commands":
				fmt.Println("Command: #commands - list commands")
				fmt.Println("--\nShows a list of available commands.")
			case "connect":
				fmt.Println("Command: #connect <host> <port> - connect to a MUD")
				fmt.Println("--\nConnects to a MUD.")
			case "end" || "exit":
				fmt.Println("Command: #end - exit")
				fmt.Println("Command: #exit - exit")
				fmt.Println("--\nExits GuruMUD client.")
			case "help":
				fmt.Println("Command: #help <topic> - display help on a topic")
				fmt.Println("--\nHelp displays a help article on a selected topic.")
			case "trigger":
				fmt.Println("Command: #trigger list - list triggers")
				fmt.Println("Command: #trigger <triggerName> 'trigger' 'response' - add a trigger")
				fmt.Println("Command: #trigger remove <triggerName> - remove a trigger")
				fmt.Println("--")
				fmt.Println("Triggers are things that when the MUD sends something that contains them,")
				fmt.Println("something is automatically done. For example, if you make a trigger by")
				fmt.Println("\"#trigger autoPat 'a dog wags its tail' 'pat dog'\", then whenever the MUD")
				fmt.Println("sends something that contains 'a dog wags its tail', then you will automatically")
				fmt.Println("pat the dog.")
		}
	}
}

func cmd_trigger(kwarg []string) {
	// TODO
}

func cmd_
