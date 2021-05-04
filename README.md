# Learning Go
### The tasks I used to learn the basics
Go is an extremely powerful langauge to have in your toolbelt as a developer on any level. It has a bunch of packages built in that can be utilized and its goal is to create lightweight executables without the hastle of other lower level langanges. It deals with memory management (which coming from C is a very nice touch), and will hold you responsible for imports or variables that are declared and never used.

I could turn this readme into a list with all the reasons Go is a good langauge to learn, but for the sake of time and to get back to learning I'll just move along. In the following directories are Go scripts, grouped together by category, I used to learn the langauge. Alongside anther readme describing the point of each script.

 - [Overview](#learning-go)
 - [Download](#download)
 - [Installation](#installation)
 	* [Linux](#linux)
 	* [Windows](#windows)
 	* [Mac](#mac)
 - [Getting Started](#getting-started)
 	* [Enviornment](#enviornment-setup)
 	* [Code Execution](#code-execution)
 		- [Running](#running)
 		- [Compiling](#compiling)
 - [Directories](#directories)
 	* [0x00_Hello_World](#0x00_hello_world)
 	* [0x01_Variables](#0x01_variables)

## Download

https://golang.org/dl/

The above link will bring you to the downloads page for Go, in the section below I will highlight some common OS installations such as Linux, Windows and Mac. Make sure the version you download is compatible with your OS.

## Installation

### Linux
**Please note, this is for common distributions of linux such as Ubuntu, using the linux precompiled installer**

For more installation on installing Go from the source, please refer to their documentation
https://golang.org/doc/install/source

1. Download and extract the archive file from the Go website, to "/usr/local". The file downloaded should look similar to ``` go[version #].linux-amd64.tar.gz ```

	**Important:** Files stored in the location /usr/local/go will need to be removed if they exist to ensure the installation has no errors. To ensure the location is cleared and no files are at that location we will remove that file and then extract the tar.gz file to that location.

	\* Note that I am using the Go version 1.16.3 tar.gz at the time of creating this, so the file name "go1.16.3.linux-amd64.tar.gz" may need to be replaced if you copy and paste this snippet
	
	``` sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.16.3.linux-amd64.tar.gz ```

2. Now that the Go files are in the local directory we need to add the PATH environment variable. To do this open either $HOME/.profile (\~/.prfile) or /etc/profile in a text editor of your choice and append the following line to the end of the file.
	
	```	export PATH=$PATH:/usr/local/go/bin	```

	or you can run the command below to echo and append the line to the end of the $HOME/.profile file.
	
	```	echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.profile ```

	Due to the method of adding go to the path, the changes will not take affect until a reboot, to have them take affect immediatly run the export command from the command line.

	**When trying to do this export with zsh I found that putting the export at the end of .zshrc file had the same result as this is called when zsh is started**

3. Ensure Go is installed and refrenced correctly in the PATH by getting the go version with the command below.
	
	``` go version ```

	This should return something along the lines of "go version go1.16.3 linux/amd64"

### Windows
The windows installation process is fairly simple. In order to get started make sure you have the .msi file, available from the link under the "Download" section.

1. Open the MSI file, this will open up a generic windows installation windows. The default location of install will be in the C:\\Windows\\Program Files (x86) or C:\\Windows\\Program Files directory. You can change that if needed. Please note that any open command prompts will need to be closed and re-opened for enviornment variables to update.

2. To ensure the installation went smoothly, go into a Command Prompt Window and run the following command

``` go version ```

If go is installed correctly the correct version should be printed in the cmd window

### Mac
The Mac install of go does essentially the same thing as the linux install, however it is automated for the most part.

1. The Mac download file from the go website should be of type .pkg, Once it is download open the package file and follow the prompts that appear, just like in linux the install should end up in the /usr/local/go directory and the installer should ad the enviornment path to the same location.

2. To ensure the install worked open up a command prompt and enter the following command

``` go version ```

This should return the current version of go installed.

## Getting Started
After we have Go installed we need to have a couple things initialized before we can get into actually writing code. Below I will go over how I initialized the enviornment for this repository, note that there will be a section dedicated to understanding this further when we get into working with external packages. This will cover setting up a local module in order to run the code we write.

### Enviornment Setup
The first thing we need to do is create a location in which we will have the project source code. Lets look at the current repository for example. The Root is "Learing_Go", and our first directory is "0x00_Hello_World".

Now cd to the parent directory of "Learning Go", this is where we will initialize the Go module. We can do this by running the following command

```
go mod init Learning_Go
```

This command is doing some very important things as far setting up an enviornment. It is creating a **Module** for our project. After running the command you will notice a "go.mod" file was created, and placed in the current directory. This file is super important to Go as it is what manages the dependencies the **Module** you are creating requires. For now however the generic go.mod file is fine as all of our code will be within one file, and using only built in functions.

The scripts can be saved in sub-directories of "Learning_Go" for now. As we learn more about Go, and how the packages/imports work, changes will be made to reflect the best practices for module and package management. For now lets just focus on learning the syntax of Go, and get familiar with how it works.

### Code Execution
Once we have the enviornment setup we can start focusing on running the code we write. Lets take a look at running the "Hello_World.go" script, inside the "0x00_Hello_World" directory.

Go is an interpreted and compilied langauge, this means we can either run the go script through an interpreter written in Go, or we can compile the code into a binary and execut the file.

Lets start with running a script.

#### Running
To run a simple Go script we will make use of the ``` go run [Path/to/file]``` command. This command will tell Go to start an interpreter and execute the script passed in. This is super useful when it comes to testing changes, or finding bugs within a script.

For more information on this command check out the link below

https://golang.org/cmd/go/#hdr-Compile_and_run_Go_program

#### Compiling
Compiling the script will create an executable file. On windows this will generate an EXE and on linux and mac it will generate an ELF file. To compile the script we will use a command close to running however the second argument goes from run to build.

```
go build [path/to/file]
```

This will create and place a new executable with the same file name as the script in the current directory.

## Directories

### 0x00_Hello_World
Now that we have an enviornment setup to write, run, and compile code we can start getting our feet wet. In this directory we will go over some simple print functions

### 0x01_Variables
This Directory is all about varibales and datypes in Go. There are a few that are different from C that I would like to highlight, as well as cover the basics of how go handles common datatypes.