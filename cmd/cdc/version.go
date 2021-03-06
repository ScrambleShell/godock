package main

import "github.com/sqp/godock/libs/cdglobal"

var cmdVersion = &Command{
	Run:       runVersion,
	UsageLine: "version",
	Short:     "print cdc version",
	Long:      `Version prints the cdc version.`,
}

func runVersion(cmd *Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
	}

	println(cdglobal.AppVersion)
}
