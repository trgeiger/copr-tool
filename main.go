/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/spf13/afero"
	"github.com/trgeiger/copr-tool/cmd"
)

func main() {
	fs := afero.NewOsFs()
	cmd.Execute(fs)
}
