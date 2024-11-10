package main

import (
	"github.com/fatih/color"
)

func main() {
	deadLinks := Links{}
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&deadLinks)
	color.Green("Active links are as follows\n")
	for _, activeLinks := range ActiveLinks {
		color.Green("%v\n", activeLinks)
	}
	color.Red("Dead links are as follows\n")
	for _, activeLinks := range ActiveLinks {
		color.Red("%v\n", activeLinks)
	}
}
