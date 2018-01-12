# Mock of sd-cmd

You need to run two server. 
One is Dummy Screwdriver Store API. Second one is Screwdriver API.
You can see how lunch Screwdriver Store API from this directory. How to lunch Screwdriver API is written on api/README.md


# Dummy Screwdriver Store API

If you want to change the output of sd-cmd you can change fake.go or shell.sh.

## Usage

### Install
```
$ go get github.com/yuichi10/mock_sd-cmd.go
$ cd $GOPATH/src/github.com/yuichi10/mock_sd-cmd.go
$ go build -o output/fake fake.go
```

### Execute
```
$ go run main.go
```

### Test
no test