bin/pulumi-resource-awsconf::
	mkdir -p bin
	(cd provider/cmd/pulumi-resource-awsconf && go build -o ../../../bin/pulumi-resource-awsconf)

bin/pulumi-gen-awsconf::
	mkdir -p bin
	(cd provider/cmd/pulumi-gen-awsconf && go build -o ../../../bin/pulumi-gen-awsconf)

preview.yaml::	bin/pulumi-resource-awsconf
	(cd examples/yaml-example && bash preview.sh)

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
