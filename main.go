package main

import (
	"strconv"

	"XJUZI/cmd"
	"github.com/samber/lo"
)

// set by LDFLAGS at compile time
var (
	gitCommit     string
	gitVersion    string
	embedFrontend string
)

func main() {
	cmd.Version = gitVersion
	cmd.Commit = gitCommit
	cmd.EmbedFrontend = lo.Must(strconv.ParseBool(embedFrontend))

	cmd.Execute()
}
