# pulumi-12709

Prototyping for https://github.com/pulumi/pulumi/issues/12709

The approach demonstrated here introduces a codegen option XReturnPlainResource for methods. A resource provider can
then use this option to build a method that returns an explicitly configured Provider Resource.

Usage is demonstrated in examples:

- [TypeScript example](./examples/ts-example/index.ts)
- [Go example](./examples/go-example/main.go)
- [Python example](./examples/py-example/__main__.py)

There are also some [slides](./slides/resource_methods.org) on the wider problem.

The implementaion of the method can be found in [configurer.go](./provider/configurer.go).
