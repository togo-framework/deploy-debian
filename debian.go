// Package debian is a VPS deploy driver for togo deploy targeting Debian hosts
// (the "deipanos"/Debian family). Same SSH + systemd flow as deploy-ubuntu, with
// apt for prerequisites. Select with DEPLOY_PROVIDER=debian.
package debian

import (
	"github.com/togo-framework/deploy"
	ubuntu "github.com/togo-framework/deploy-ubuntu"
	"github.com/togo-framework/togo"
)

func init() {
	deploy.RegisterDriver("debian", func(k *togo.Kernel) (deploy.Deployer, error) {
		return ubuntu.New("debian", "apt-get update -y && apt-get install -y ca-certificates rsync"), nil
	})
}
