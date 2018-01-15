# Mock of sd-cmd

You need to run two server. 
One is Dummy Screwdriver Store API. Second one is Screwdriver API.
You can see how lunch Screwdriver Store API from this directory. How to lunch Screwdriver API is written on api/README.md

This repository expect to run sd-cmd as below
```
$ sudo --preserve-env go run sd-cmd.go foo/binary@1.0.2
$ sudo --preserve-env go run sd-cmd.go exec foo/binary@1.0.2 -la

$ sudo --preserve-env go run sd-cmd.go exec foo/text@1.0.1
```

If you change api or store maybe you cannot use these commands.


# Dummy Screwdriver Store API

If you want to change the output of sd-cmd you can change fake.go or shell.sh.

## Usage

### Install
```
$ git clone https://github.com/yuichi10/mock_sd-cmd.git
$ mv mock_sd-cmd $GOPATH/src/github.com/yuichi10/mock_sd-cmd
$ cd $GOPATH/src/github.com/yuichi10/mock_sd-cmd
$ go build -o output/fake fake.go
```

### Environment variables
```
export SD_API_URL=http://localhost:3000/
export SD_STORE_URL=http://localhost:8080/
export SD_TOKEN=DcdjHW9u--dyFydVpVSaCaGSwj9VPlUFOExF8YQZWrY
SD_ARTIFACTS_DIR=/opt/
```
The log file will be created at `/opt/.sd/commands/...`
The command file is downloaded at `/opt/sd/commands/...`

### Execute
```
$ go run main.go
```

### Test
no test
