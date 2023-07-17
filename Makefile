bin/pulumi-resource-awsconf::
	mkdir -p bin
	(cd provider/cmd/pulumi-resource-awsconf && go build -o ../../../bin/pulumi-resource-awsconf)

preview.yaml::	bin/pulumi-resource-awsconf
	(cd examples/hello && bash preview.sh)
