// Code generated by go-autowire. DO NOT EDIT.

package autowire_gen

import (
	"github.com/Just-maple/go-autowire/example/dependencies"
	"github.com/Just-maple/go-autowire/example/dependencies/test_b"
	"github.com/Just-maple/go-autowire/example/dependencies/test_b/test"
	test2 "github.com/Just-maple/go-autowire/example/dependencies/test_b/test/test"
	test_b2 "github.com/Just-maple/go-autowire/example/dependencies/test_b/test/test_b"
	test3 "github.com/Just-maple/go-autowire/example/dependencies/test_b/test/test_b/test"
	"github.com/Just-maple/go-autowire/example/dependencies/test_c"
	"github.com/google/wire"
)

var StructSet = wire.NewSet(
	wire.Struct(new(dependencies.Test), "*"),

	wire.Struct(new(dependencies.Test2), "*"),
	wire.Bind(new(dependencies.TestInterface1), new(*dependencies.Test2)),

	dependencies.ConstTest3,

	dependencies.NewTest4,

	wire.Struct(new(test_b.Test), "*"),

	wire.Struct(new(test_b.Test2), "*"),

	wire.Struct(new(test.Test), "*"),

	wire.Struct(new(test.Test2), "*"),

	wire.Struct(new(test2.Test), "*"),

	wire.Struct(new(test2.Test2), "*"),

	wire.Struct(new(test_b2.Test), "*"),

	wire.Struct(new(test_b2.Test2), "*"),

	wire.Struct(new(test3.Test), "*"),

	wire.Struct(new(test3.Test2), "*"),

	wire.Struct(new(test_c.Test), "*"),

	wire.Struct(new(test_c.Test2), "*"),
)
