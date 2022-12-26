# Simple album fetch CLI tool written in go. 

Base albums(from a test api) range between 1-100, requesting albums outside of range will throw an error. Included in the repo is a sample build of the project.

## Basic CLI Commands

### Get albums
#### To Run without pre-compilation

Run `go run main.go 8`

#### To Run with compilation(Default output path listed in readme)

Run `./album 4`

#### To run the compilation installed to global scope, or with a custom global scope name

##### Default

Run `album 56`


## Build Prerequisites
Have Golang in a working state. Installation instructions at: https://go.dev/doc/install

## Install dependencies
Install go local dependencies for project simply run `go get` command 

## To Build For local system
To build the program have go installed on your local system run 

`go build`

this will output a compiled program with the name `album` that can be run with a command like `./album <albumNumber>`

optional:  `-o` flag used to set directory/filename using standard path. i.e. `go build -o <customName>` will drop the complied program in the working directory with a custom name. You can also specify any custom path.

## To Build and install on the global scope using default name(goCli)

To install the cli for use anywhere on the local machine clone this repository and have Golang in a working state. Clone this repo and simply use the `go install` command. This will compile the program and install it to `$GOPATH/bin` directory. To figure out where your $GOPATH is set run the `go env GOPATH`. You can delete installed programs from there. You don't need to install within the $GOPATH but it is the default for go projects. If installed to another directory path you'll need to setup a custom alias in your `.bashrc` or `.zshrc`. 

## To Build and install install with a custom name within the global scope 
Run the command `go build -o $GOPATH/bin/<customName>` i.e. `foo` as the custom name. If this is used you can call the cli with `foo <albumNumber>` command from any working directory on the machine.

## To Compile for specific architectures(not local machine)
Find valid compilation target combinations at:
https://go.dev/doc/install/source#environment:~:text=%24GOBIN.-,%24GOOS%20and%20%24GOARCH,-The%20name%20of

### Specific architecture compile example: Linux 64bit

`GOOS=linux GOARCH=amd64 go build -o album`
