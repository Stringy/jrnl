package commands

import (
	"fmt"
)

type HelpCommand struct {
	name    string
	desc    string
	usage   string
	longuse string
}

var cmdHelp = &HelpCommand{
	name:    "help",
	desc:    "shows help information about all commands",
	usage:   "journal help [name]",
	longuse: ``,
}

func (h *HelpCommand) Do(args []string) error {

	if len(args) == 0 {
		fmt.Println("Journal - journaling management software")
		fmt.Println("\nCommands:")
		for key, val := range Commands {
			fmt.Printf("  %s: %s\n", key, val.Desc())
		}
		fmt.Println("For specific help, use \"journal help <cmd>\"")
	} else {
		if _, ok := Commands[args[0]]; ok {
			fmt.Println(Commands[args[0]].LongUse())
		}
	}

	return nil
}

func (h *HelpCommand) Name() string    { return h.name }
func (h *HelpCommand) Desc() string    { return h.desc }
func (h *HelpCommand) Usage() string   { return h.usage }
func (h *HelpCommand) LongUse() string { return h.longuse }
