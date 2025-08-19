# Rock, Paper, Scissors

This is a command-line golang implementation of the classic game "Rock, Paper, Scissors", with the player facing off against
a computer-controlled opponent. The computer uses [a more sophisticated strategy](https://arstechnica.com/science/2014/05/win-at-rock-paper-scissors-by-knowing-thy-opponent/) to bet the opponents.

## Quick start:
1. Run `make start` if you want to play locally OR
2. Run `make docker/setup` and `make docker/start` to play in a docker container

## Game rules:
- The game starts by asking the player to enter their name.
- Entering 0 opens the exit menu, which requires confirmation ("Y") to quit.
- A game continues until one player reaches the winning score or chooses to exit.
- The winning score defaults to 3 but can be changed before the game starts.
- A round is a single throw from both the player and the computer.
- Rounds automatically continue until the game ends.
- After a game ends, the player can choose to start a new game or exit.

## Makefile
| Command                 | Description                                                                             |
| ----------------------- |-----------------------------------------------------------------------------------------|
| `make start`            | Starts the game locally using Go on your machine.                                       |
| `make test`             | Runs all tests locally.                                                                 |
| `make generate`         | Regenerates Go mocks by deleting old `_mock.go` files and running `go generate`.        |
| `make tidy`             | Synchronizes Go dependencies using `go mod tidy`.                                       |
| `make docker/setup`     | Builds all Docker images, syncs dependencies and starts the development container       |
| `make docker/start`     | Starts the development container already created via `docker/setup` and attaches to it. |
| `make docker/test`      | Runs all tests inside a Docker container.                                               |
| `make docker/lint`      | Runs the Go linter inside a Docker container.                                           |
| `make docker/stop`      | Stops all running containers.                                                           |
| `make docker/clean`     | Stops and removes all containers, volumes, orphaned resources, and local images.        |
| `make docker/generate`  | Regenerates Go mocks inside a Docker container.                                         |
| `make docker/tidy`      | Synchronizes Go dependencies inside a Docker container.                                 |
| `make docker/dev-shell` | Opens a shell (`sh`) inside the development container for manual commands.              |
| `make docker/light`     | Runs the light (production-like) version of the game without development tools.         |
