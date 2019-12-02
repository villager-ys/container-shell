package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/1179325921/kubeutil/cmd/controllers/app"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	command := app.NewKubeCommand()

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
