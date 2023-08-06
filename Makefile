UNIT_TEST_COVERAGE_MIN ?= 95

.PHONY: help
help:
	@echo "Usage: make <target>"
	@echo " "
	@echo "  where <target> is one of:"
	@echo "    - help                       - display this help message"
	@echo "    - clean                      - remove temporary files"
	@echo "    - test                       - run unit tests from inside a docker container"
	@echo "    - test/no-docker             - run unit tests (assumes local environment has installed all prequisites)"
	@echo " "

clean:
	go clean

.PHONY: test
test:
	docker build -f ./Dockerfile -t bgauth:latest . 
	@docker run \
		--rm \
		-t \
		-v $$(pwd):/workspace/bgauth \
		bgauth:latest "cd bgauth && make test/no-docker"

test/no-docker:
	echo "Unit tests not yet implemented"
