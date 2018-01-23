// +build !unit

package acceptance

import (
	"os"
	"os/user"
	"testing"

	"github.com/DATA-DOG/godog"
)

func FeatureContext(mainSuit *godog.Suite) {
	context := &testContext{}

	mainSuit.BeforeScenario(context.resetContext)

	CommonFeatureContext(mainSuit, context)
	PersonCrudFeatureContext(mainSuit, context)
}

func TestMain(m *testing.M) {
	u, _ := user.Current()
	status := godog.RunWithOptions("godogs", func(mainSuit *godog.Suite) {
		FeatureContext(mainSuit)
	}, godog.Options{
		NoColors: u.Name == "jenkins",
		Format:   "pretty",
		Paths:    []string{"features"},
	})

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}
