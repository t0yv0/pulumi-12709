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
	rm -rf ./sdk/go/awsconf
	mkdir -p ./sdk/go/awsconf
	bin/pulumi-gen-awsconf go ./sdk/go/awsconf schema/schema.json 0.0.1
