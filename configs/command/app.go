package command

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/wsaefulloh/go-solid-principle/routers"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start api server",
	RunE:  serve,
}

func serve(cmd *cobra.Command, args []string) error {
	mainRutes := routers.New()

	fmt.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", mainRutes)
	if err != nil {
		return err
	}

	return nil
}
