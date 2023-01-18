package main

import (

	"github.com/aeone1/rotech-post-comment/initializers"
	"github.com/aeone1/rotech-post-comment/pkg/cmd"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	cmd.RunServer()
}
