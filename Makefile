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

preview.ts::	bin/pulumi-resource-awsconf
	(cd examples/ts-example && bash preview.sh)

preview.py:: 	bin/pulumi-resource-awsconf ./examples/py-example/venv
	(cd examples/py-example && bash preview.sh)

./examples/py-example/venv:
	(cd ./examples/py-example && python3 -m venv venv && ./venv/bin/python -m pip install pulumi_aws)
	(cd ./examples/py-example && ./venv/bin/python -m pip install -e ../../sdk/python/bin)

up.ts::	bin/pulumi-resource-awsconf
	(cd examples/ts-example && bash up.sh)

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

gen.py::	gen.schema
	rm -rf ./sdk/python
	mkdir -p ./sdk/python
	bin/pulumi-gen-awsconf python ./sdk/python schema/schema.json 0.0.1

build.ts::	gen.ts
	cd sdk/nodejs/ && \
		printf "module fake_nodejs_module // Exclude this directory from Go tools\n\ngo 1.17\n" > go.mod && \
		yarn install && \
		yarn link @pulumi/pulumi && \
		yarn run tsc && \
		mkdir -p bin && \
		cp package.json yarn.lock ./bin/ && \
		sed -i.bak -e "s/\$${VERSION}/$(VERSION)/g" ./bin/package.json

build.py::	gen.py
	cd sdk/python/ && \
		printf "module fake_python_module // Exclude this directory from Go tools\n\ngo 1.17\n" > go.mod && \
		python3 setup.py clean --all 2>/dev/null && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		sed -i.bak -e 's/^VERSION = .*/VERSION = "$(PYPI_VERSION)"/g' -e 's/^PLUGIN_VERSION = .*/PLUGIN_VERSION = "$(VERSION)"/g' ./bin/setup.py && \
		rm ./bin/setup.py.bak && rm ./bin/go.mod && \
		cd ./bin && python3 setup.py build sdist

install.ts::	build.ts
	(cd ./sdk/nodejs/bin && yarn unlink || echo "unlink not needed")
	yarn link --cwd ./sdk/nodejs/bin

tidy::
	(cd provider && go mod tidy)
	(cd sdk && go mod tidy)
	(cd examples/go-example && go mod tidy)
