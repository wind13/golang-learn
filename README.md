# Learn golang notes

Learning code for [Golang_Puzzlers](https://github.com/hyper0x/Golang_Puzzlers)

## Part1 usage:

### Check your `GOPATH` setting

```
echo $GOPATH
```

### Add this path to `GOPATH`

```
export GOPATH=$GOPATH:/your-home/path-to/golang-learn
```

### Install libs

```
go install part1/libs
```

This will create a folder like `./pkg/darwin_amd64/part1`

### Run main.go

```
go run src/part1/main.go -name "Simon" -n 49
```

### Build and run main

```
go build src/part1/main.go 
./main -name "Simon" -n 25
```