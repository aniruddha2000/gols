package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func listFileDirectory(path string, all, moreThanOnePath bool) {
	if moreThanOnePath {
		fmt.Println(path)
	}
	list, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalf("Failed to open directory: %s", err)
	}
	for _, name := range list {
		if name.IsDir() {
			dirColor := color.New(color.FgCyan, color.Bold)
			dirColor.Printf("%s  ", name.Name())
		} else {
			fmt.Printf("%s  ", name.Name())
		}
	}
	fmt.Println()
}

func printPermissions(path string, moreThanOnePath bool) {
	if moreThanOnePath {
		fmt.Println(path)
	}
	list, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalf("Failed to open directory: %s", err)
	}
	for _, name := range list {
		if name.IsDir() {
			dirColor := color.New(color.FgCyan, color.Bold)
			dirColor.Printf("%s   %04o   %s\n", name.Mode(), name.Mode().Perm(), name.Name())
		} else {
			fmt.Printf("%s   %04o   %s\n", name.Mode(), name.Mode().Perm(), name.Name())
		}
	}
	fmt.Println()
}

func main() {
	app := &cli.App{
		Name:  "gols",
		Usage: "A fun ls command using go",
		Action: func(c *cli.Context) error {
			if c.Args().Len() > 1 {
				for i := 0; i < c.Args().Len(); i++ {
					path := c.Args().Get(i)
					listFileDirectory(path, false, true)
				}
			} else if c.Args().Len() == 1 {
				path := c.Args().Get(0)
				listFileDirectory(path, false, false)
			} else {
				path, err := os.Getwd()
				if err != nil {
					log.Fatalf("Failed to get the directory name: %s", err)
				}
				listFileDirectory(path, false, false)
			}
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "Show all files",
				Action: func(c *cli.Context) error {
					if c.Args().Len() > 1 {
						for i := 0; i < c.Args().Len(); i++ {
							path := c.Args().Get(i)
							listFileDirectory(path, true, true)
						}
					} else if c.Args().Len() == 1 {
						path := c.Args().Get(0)
						listFileDirectory(path, true, false)
					} else {
						path, err := os.Getwd()
						if err != nil {
							log.Fatalf("Failed to get the directory name: %s", err)
						}
						listFileDirectory(path, true, false)
					}
					return nil
				},
			},
			{
				Name:    "long",
				Aliases: []string{"l"},
				Action: func(c *cli.Context) error {
					if c.Args().Len() > 1 {
						for i := 0; i < c.Args().Len(); i++ {
							path := c.Args().Get(i)
							printPermissions(path, true)
						}
					} else if c.Args().Len() == 1 {
						path := c.Args().Get(0)
						printPermissions(path, false)
					} else {
						path, err := os.Getwd()
						if err != nil {
							log.Fatalf("Failed to get the directory name: %s", err)
						}
						printPermissions(path, false)
					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
