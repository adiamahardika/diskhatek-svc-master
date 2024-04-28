package command

import (
	"log"

	"github.com/spf13/cobra"

	application "svc-master/app"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)
}

var cmdRoot = &cobra.Command{
	Use:   "svc-master",
	Short: "The svc-master is a service to handle all master domain.",
	Long:  `The svc-master is a service to handle the requirements of the svc-master`,
	Run: func(cmd *cobra.Command, args []string) {
		app := application.New()
		err := app.Init()
		if err != nil {
			log.Fatalf("Error in initializing the application: %+v", err)
			return
		}

		err = app.Run()
		if err != nil {
			log.Fatalf("Error in running the application: %+v", err)
			return
		}
	},
}

func Execute() {
	if err := cmdRoot.Execute(); err != nil {
		log.Fatalf("Error in executing the root command: %+v", err)
	}
}
