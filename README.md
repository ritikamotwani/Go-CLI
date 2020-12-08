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

- go mod init HelloWorld
- go build
- go install HelloWorld


## Usage

![image](https://user-images.githubusercontent.com/22127980/101438705-fc5f7c00-3938-11eb-9bb5-55b129316a05.png)

Example - 
Run `HelloWorld create NodeJS` a nodeJS Hello World app gets created in the home directory.

