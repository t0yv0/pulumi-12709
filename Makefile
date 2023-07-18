bin/pulumi-resource-awsconf::
	mkdir -p bin
	(cd provider/cmd/pulumi-resource-awsconf && go build -o ../../../bin/pulumi-resource-awsconf)

bin/pulumi-gen-awsconf::
	mkdir -p bin
	(cd provider/cmd/pulumi-gen-awsconf && go build -o ../../../bin/pulumi-gen-awsconf)

preview.yaml::	bin/pulumi-resource-awsconf
	(cd examples/yaml-example && bash preview.sh)

preview.go::	bin/pulumi-resource-awsconf
	(cd examples/go-example && bash preview.sh)

up.go::	bin/pulumi-resource-awsconf
	(cd examples/go-example && bash up.sh)

gen.schema::	bin/pulumi-gen-awsconf
	rm -rf ./schema
	mkdir -p ./schema
	bin/pulumi-gen-awsconf schema ./schema

gen.go::	gen.schema
	rm -rf ./sdk/go
	mkdir -p ./sdk/go
	bin/pulumi-gen-awsconf go ./sdk/go schema/schema.json 0.0.1

gen.ts::	gen.schema
	rm -rf ./sdk/nodejs
	mkdir -p ./sdk/nodejs
	bin/pulumi-gen-awsconf typescript ./sdk/nodejs schema/schema.json 0.0.1

tidy::
	(cd provider && go mod tidy)
	(cd sdk && go mod tidy)
	(cd examples/go-example && go mod tidy)
