package cli

import (
	"github.com/spf13/cobra"
	cErrors "github.com/pjgg/go-errors/errors"
	"github.com/pjgg/go-rest-skel/configuration"
	"github.com/pjgg/go-rest-skel/handler"
)

var personCmd = &cobra.Command{
	Use: "person",
	Run: personHandler,
}

func init() {
	ServicesCmd.AddCommand(personCmd)
}

func personHandler(cmd *cobra.Command, args []string) {
	errorHandler := cErrors.NewErrorHandler(configuration.ConfigurationManagerInstance.Enviroment, configuration.ConfigurationManagerInstance.SentryDsn, configuration.ConfigurationManagerInstance.Version, true)
	r := handler.SetupPersonRouting(handler.NewPersonHandler(errorHandler))
	r.Run(configuration.ConfigurationManagerInstance.ServerAddress)
}
