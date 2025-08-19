DOCKERCMD        ?= docker
DOCKERCOMPOSECMD ?= ${DOCKERCMD} compose
GOCMD            ?= go
GOLANGCICMD      ?= golangci-lint

.PHONY: start test generate tidy \
        docker/check-engine docker/setup docker/stop docker/clean \
        docker/start docker/test docker/generate docker/tidy \
        docker/lint docker/dev-shell docker/light

# start: starts the game locally
start:
	${GOCMD} run main.go

#test: tests the game locally
test:
	${GOCMD} test -shuffle=on -cover ./...

#generate: regenerate the mocks
generate:
	-find . -name "*_mock.go" -delete
	${GOCMD} generate ./...

#tidy: synchronizes go dependencies
tidy:
	${GOCMD} mod tidy

# -------------------------------------------------------------------------------

#docker/check-engine: checks if docker is running before calling other commands
docker/check-engine:
	@if ! ${DOCKERCMD} info >/dev/null 2>&1; then \
		echo "ERROR: Docker engine is not running. Please start Docker manually."; \
		exit 1; \
	fi

# START HERE
# docker/setup: sets initial dependencies and creates the game container
docker/setup: docker/check-engine
	${DOCKERCOMPOSECMD} build
	$(MAKE) docker/tidy
	${DOCKERCOMPOSECMD} up --no-start development
	$(MAKE) docker/test

# docker/start: start the development (game) container already created in docker/setup
docker/start: docker/check-engine
	${DOCKERCMD} start -ai rock-paper-scissors

# docker/test: runs tests in docker
docker/test: docker/check-engine
	${DOCKERCOMPOSECMD} run --rm test

# docker/lint: runs the linter in a container
docker/lint: docker/check-engine
	${DOCKERCOMPOSECMD} run --rm development ${GOLANGCICMD} run

# docker/stop: stop containers in docker
docker/stop: docker/check-engine
	${DOCKERCOMPOSECMD} down

# docker/clean: clean containers and images in docker
docker/clean: docker/check-engine
	${DOCKERCOMPOSECMD} down --volumes --remove-orphans --rmi local

# docker/generate: regenerates the mocks
docker/generate: docker/check-engine
	${DOCKERCOMPOSECMD} run --rm development make generate

# docker/tidy synchronizes dependencies
docker/tidy: docker/check-engine
	${DOCKERCOMPOSECMD} run --rm development ${GOCMD} mod tidy

#docker/dev-shell enables the shell for development purposes
docker/dev-shell: docker/check-engine
	${DOCKERCOMPOSECMD} run --rm development sh

# docker/light: is the light version of the game without development tools
docker/light:
	${DOCKERCOMPOSECMD} run --rm -it light