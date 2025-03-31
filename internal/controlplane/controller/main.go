package controller

//need to parse flags
//configs for gw and ingress
//start prom collector
//start prom http server
//start controller

import (
	"fmt"

	ctrl "sigs.k8s.io/controller-runtime"
)

func Start() error {

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{})
	if err != nil {
		return fmt.Errorf("failed to construct manager: %w", err)
	}

	return mgr.Start(ctrl.SetupSignalHandler())
}
