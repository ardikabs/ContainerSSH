package steps

import (
	"fmt"

	"github.com/containerssh/configuration"
	"github.com/containerssh/kuberun"
)

func (scenario *Scenario) ConfigureKubernetes(username string) error {
	if scenario.ConfigServer == nil {
		return fmt.Errorf("config server is not running")
	}

	appConfig := &configuration.AppConfig{}
	appConfig.Backend = "kuberun"

	if err := kuberun.SetConfigFromKubeConfig(&appConfig.KubeRun); err != nil {
		return err
	}

	scenario.ConfigServer.SetUserConfig(username, appConfig)
	return nil
}
