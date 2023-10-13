package provider

import (
	"github.com/Lofanmi/go-autowire/example_iwanta"
)

// @autowire()
func ProvideIwantaDep() example_iwanta.Dep {
	return example_iwanta.Dep{}
}
