# Golang TypeWriter

This app allows you to type content of the files to terminal with the fixed time for every line

## Example. Raw

>You shold specify the input text file as the `first` parameter and typing delay in _milliseconds_ as `second` parameter

```bash
go run cmd/main.go text.txt 100
```

## Example. Docker

> First, you should build docker image using `scripts/build.sh` script from the repo root directory

```bash
./scripts/build.sh .
```

> After that you can pass parameters to docker container. Note that you should specify volume with your file and map it to container. In example above file located in current working directory

```bash
docker run --rm "$PWD:$PWD" -w "$PWD" typewriter:<tag> <filename> <delay>
```
