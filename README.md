# Go-CLI
Command Line Utility to create Hello World apps.


<img src="https://ordina-jworks.github.io/img/make-your-own-cli-with-golang-and-cobra/banner.jpg" />

## Installation

       
### Go 
Instructions: https://golang.org/doc/install

### Setting up environment variables

In your `.zshrc` or `.bashrc` file, append your system’s $PATH so that the command can be invoked from anywhere.

`export PATH=${PATH}:$HOME/go/bin`

Restart your terminal once to source them into your machine environment or just `source ~/.zshrc`.


### Cobra
To install cobra, run: 

`go get -u github.com/spf13/cobra/cobra`

You can run `cobra help` or just `cobra` to get more familiar with it.


[Reference: https://dzone.com/articles/how-to-create-a-cli-in-go-in-few-minutes]





## Steps to install the project:

After cloning the repo, go to the project's directory (HelloWorld) and run:

One Time Commands:

- go mod init HelloWorld
- go build
- go install HelloWorld

After editing files to run the command again reflecting new changes, run:
- go install HelloWorld to build the new version.


<b> This project also uses to create binary files of the hello world files in different languages: https://github.com/go-bindata/go-bindata </b>

After adding a new file in cmd/data for a new language
To install that, run:

1. go get -u github.com/go-bindata/go-bindata/...
2. go-bindata data/

Then install the module again.



## Usage

![image](https://user-images.githubusercontent.com/22127980/101438705-fc5f7c00-3938-11eb-9bb5-55b129316a05.png)

Example - 
Run `HelloWorld create NodeJS` a nodeJS Hello World app gets created in the home directory.

