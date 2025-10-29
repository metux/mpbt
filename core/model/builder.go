// SPDX-License-Identifier: AGPL-3.0-or-later
package model

import (
	"github.com/metux/go-magicdict/api"
)

type IBuilder interface {
	Init(p *Package, cf api.Entry)
	RunPrepare() error // eg. autogen.sh
	RunConfigure() error
	RunBuild() error
	RunInstall() error
}
