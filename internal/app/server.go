package app

import (
	"fmt"
)

func (a *Application) Run() error {
	return a.Router.Run(fmt.Sprintf(":%d", a.Config.App.Port))
}
