/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import "github.com/motiong-io/motiong-cli/cmd"
import _ "github.com/motiong-io/motiong-cli/cmd/bar"
import _ "github.com/motiong-io/motiong-cli/cmd/foo"

func main() {
	cmd.Execute()
}
