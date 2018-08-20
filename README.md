# Learn golang notes

Learning code for [Golang_Puzzlers](https://github.com/hyper0x/Golang_Puzzlers)

## Part1 usage:

### Check your `GOPATH` setting

```
echo $GOPATH
```

### Add this path to GOPATH

```
export GOPATH=$GOPATH:/your-home/path-to/part1
```

### Install libs

```
go install libs
```

This will create a folder ./part1/pkg

### Run main.go

```
go run part1/src/main.go -name "Simon" -n 9
```

### Build and run main

```
go build part1/src/main.go 
./main -name "Simon" -n 9
```