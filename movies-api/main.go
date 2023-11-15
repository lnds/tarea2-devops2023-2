package main

import (
	"github.com/spf13/cobra"
	"os"
)

func main() {

 	bind  := os.Getenv("BIND_IP")

	port  := os.Getenv("BIND_PORT")

	var rootCmd = &cobra.Command{
		Use:   "api",
		Short: "run movie server",
		Long:  "run a RESTful API server of movies",
		Run: func(cmd *cobra.Command, args []string) {
			server(bind, port)
		},
	}
	rootCmd.Flags().StringVarP(&bind, "bind", "b", "", "bind address")
	rootCmd.MarkFlagRequired("bind")
	rootCmd.Flags().StringVarP(&port, "port", "p", port, "bind port")

	rootCmd.Execute()
}


