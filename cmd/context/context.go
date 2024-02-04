package context

import (
	"github.com/gabrieljuliao/godo/cmd/configuration"
	"github.com/gabrieljuliao/godo/cmd/utils"
)

type Context struct {
	Configuration configuration.Configuration
}

// ApplicationContext is the global context
var ApplicationContext Context

func NewContext() *Context {
	ApplicationContext = Context{
		Configuration: configuration.NewConfiguration(),
	}
	return &ApplicationContext
}

func ListConfiguration() {
	utils.PrettyPrintYaml(ApplicationContext.Configuration)
}
