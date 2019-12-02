package app

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	cmdutil "github.com/1179325921/container-shell/cmd"
	"github.com/1179325921/container-shell/controllers/demo"
	"github.com/1179325921/container-shell/initialize"
	"github.com/1179325921/container-shell/kube"
	"github.com/1179325921/container-shell/options"
	"github.com/1179325921/container-shell/utils"
)

// NewKubeCommand creates a *cobra.Command object with default parameters
func NewKubeCommand() *cobra.Command {
	opt, err := options.NewkubeOptions()
	if err != nil {
		log.Fatalf("unable to initialize command options: %v", err)
	}
	var flags *pflag.FlagSet

	cmd := &cobra.Command{
		Use:  "container-shell",
		Long: `kube-util is utils for kubernetes.`,
		Run: func(cmd *cobra.Command, args []string) {
			if opt.Version {
				printVersion()
			}
			var stopCh = make(chan struct{})
			go run(stopCh)
			cmdutil.Wait(func() { fmt.Println("exiting.") }, stopCh)
		},
	}
	flags = cmd.Flags()
	flags.BoolVarP(&opt.Version, "version", "v", false, "Print version information and quit")
	// flags.BoolVar(&opt.Version, "version", false, "Print version information and quit")

	return cmd
}

func printVersion() {
	fmt.Printf("container-shell version: %s\n", initialize.Version)
	os.Exit(0)
}

func printHelp() {
	fmt.Printf("container-shell help \n")
	os.Exit(0)
}

func run(stopCh <-chan struct{}) {
	kubeConfig, _ := utils.ReadFile("./config")
	kubeC, _ := kube.NewKubeOutClusterClient(kubeConfig)
	sharedInformerFactory, _ := kube.NewSharedInformerFactory(kubeC)
	podInformer := sharedInformerFactory.Core().V1().Pods()
	demoController := demo.NewDemoController(podInformer)
	go sharedInformerFactory.Start(stopCh)
	demoController.Run(5, stopCh)
	fmt.Println("exit")
}
