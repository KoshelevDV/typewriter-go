set -Eeoux pipefail

docker build -t typewriter:$(git describe --tags --abbrev=0) $1