package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v3"
)

// simpleArgsLookup is a tiny helper to read flag-like args without relying
// on the library's parsing APIs (keeps the example robust across minor API changes).
func simpleArgsLookup(name string, short string, args []string) string {
	for i, a := range args {
		if a == "--"+name || a == short {
			if i+1 < len(args) {
				return args[i+1]
			}
		}
		if strings.HasPrefix(a, "--"+name+"=") {
			return strings.SplitN(a, "=", 2)[1]
		}
	}
	return ""
}

func main() {
	root := &cli.Command{
		Name:  "example-cli",
		Usage: "A small demo of urfave/cli v3",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Value: "World", Usage: "name to greet"},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			// fallback to simple arg parsing for this small example
			name := simpleArgsLookup("name", "-n", os.Args)
			if name == "" {
				name = "World"
			}
			fmt.Printf("Hello, %s!\n", name)
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "greet",
				Aliases: []string{"g"},
				Usage:   "greet someone",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					// positional args: use os.Args to find the command's positional argument
					// naive: last arg if not a flag
					target := ""
					for i := len(os.Args) - 1; i >= 0; i-- {
						a := os.Args[i]
						if strings.HasPrefix(a, "-") {
							continue
						}
						if a == "greet" || a == "g" || strings.HasSuffix(a, "main.go") {
							break
						}
						target = a
						break
					}
					if target == "" {
						target = "stranger"
					}
					fmt.Printf("Greetings, %s!\n", target)
					return nil
				},
			},
		},
	}

	if err := root.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
