# Learning Go

Go is an extremely powerful langauge to have in your toolbelt as a developer on any level. Go's goal is to create lightweight executables without the hastles, or drawbacks other programming langanges provide. It deals with garbage collection(which coming from C is a very nice touch). Aside from garabge collection go was designed to utilize modern computer cpus, allowing us to write code that can be executed and ran faster.

## Table of Contents

 - [Overview](#learning-go)
 - [Table of Contents](#table-of-contents)
 - [Repo Goal](#repository-goal)
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
 - [Topics](#topics)
 - [Conclusion](#conclusion)
 - [Extra Sources](#helpful-sources)

## Repository Goal

The goal of this repository is to give any viewers the information they need to start writing go code. I will also personally use it as a refrence if I ever need a refresher on syntax or general usage. Many of the topics covered are also cover in the [Effective Go](https://golang.org/doc/effective_go) documentation. A list of the topics covered can be found under the [Topics](#topics) section.

## Download

https://golang.org/dl/

The above link will bring you to the downloads page for Go, in the section below I will highlight some common OS installations such as Linux, Windows and Mac. Make sure the version you download is compatible with your OS.

## Installation

### Linux
**Please note, this is for common distributions of linux such as Ubuntu, using the linux precompiled installer**

For more installation on installing Go from the source, please refer to their documentation
https://golang.org/doc/install/source

1. Extract the information within the downloaded archive file downloaded from the Go website, to /usr/local this should look like ``` go[version #].linux-amd64.tar.gz ```

	**Important:** Files stored in the location /usr/local/go will need to be removed or overwritten if they exist to ensure the installation has no errors. To ensure the location is cleared and no files are there we will remove that file and then extract the tar.gz file

	\* Note that in this case I am using the Go version 1.16.3 tar.gz at the time of creating this, so the file name may need to be replaced if you copy and paste this snippet
	
	``` sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.16.3.linux-amd64.tar.gz ```

2. Now that the Go files are in the local directory we need to add the PATH environment variable. To do this open either $HOME/.profile (\~/.prfile) or /etc/profile in a text editor of your choice and append the following line to the end of the file.
	
	```	export PATH=$PATH:/usr/local/go/bin	```
	
	Due to the method of adding go to the path, the changes will not take affect until a reboot, to have them take affect immediatly run the export command from the command line.

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

The end goal for this repository is to build your skills up to be able to write testable Go packages, and modules, that can be shared on a version control site such as github. To start off however creating a full blown package is not needed. Examples shown in directories 0x00-0x07 were created and run using the ``` go run ``` command, more on that below.

### Enviornment Setup

Modules in Go are usually stored in one folder, to keep everything it will relay on in one place. If you are following along with each directory below is how you would do such without creating modules. For this example I will be using the 00_Hello_World directory.

Lets say you are following along and make your way into the 00_Hello_World directory, and see our first file containing code. How do we run it?

We would want to start by creating a location on our local machine for the file to run. If you are on linux you can do that with the mkdir command, I will use a folder name "project". Once we have that folder created, navigate inside of it.

Our new location should be
```
project/
```
Once inside of this directory we are inside of our "enviornemt", when you get to the modules section you will see how directories can be turned into modules. For now this will do just fine. In this directory we can have as many of the example files as we want. Just be sure to specifiy the file name when running or compiling the scripts as seen below.

### Code Execution
Once we have the enviornment setup we can start focusing on running the code we write. Lets take a look at running the "00_Hello_World.go" script, inside the "0x00_Hello_World" directory.

Go is an interpreted and compilied langauge, this means we can either run the go script through an interpreter written in Go, or we can compile the code into a binary and execute the file.

Lets start with running a script.

#### Running
To run a simple Go script, like many in this repository, we will make use of the ``` go run [Path/to/file]``` command. This command will tell Go to start an interpreter and execute the script passed in. This is super useful when it comes to testing changes, or finding bugs within a script.

For more information on this command check out the link below

https://golang.org/cmd/go/#hdr-Compile_and_run_Go_program

#### Compiling
Compiling the script will create an executable file. On windows this will generate an EXE and on linux and mac it will generate an ELF file. To compile the script we will use a command close to running however the second argument goes from run to build.

```
go build [path/to/file]
```

This will create and place a new executable with the same file name as the script in the current directory.

## Topics

Below is a list of topics covered in this repository. If you are using this repo as a means to learn Go its recommended you start at the top and work your way down(0x00-0x10). This list is here because it will be an easy way to navigate to a specific topic without browsing for them.

- [Printing](https://github.com/Syssos/Learning_Go/tree/main/0x00_Hello_World#hello_world)
- [Comments](https://github.com/Syssos/Learning_Go/tree/main/0x03_Functions#comments)
- [Variables](https://github.com/Syssos/Learning_Go/tree/main/0x01_Variables#variables)
- [Control Structures](#topics)
	* [if else](https://github.com/Syssos/Learning_Go/tree/main/0x02_If_Else#if-else)
	* [for](https://github.com/Syssos/Learning_Go/tree/main/0x04_Loops_and_Arrays#00_for_loop)
	* [switch](https://github.com/Syssos/Learning_Go/tree/main/0x02_If_Else#02_switch)
- [Functions](https://github.com/Syssos/Learning_Go/tree/main/0x03_Functions#00_functions)
- [Return](https://github.com/Syssos/Learning_Go/tree/main/0x03_Functions#02_return)
- [Datatypes](#topics)
	* [string](https://github.com/Syssos/Learning_Go/tree/main/0x01_Variables#00_strings)
	* [int](https://github.com/Syssos/Learning_Go/tree/main/0x01_Variables#01_ints)
	* [float](https://github.com/Syssos/Learning_Go/tree/main/0x01_Variables#02_floats)
	* [Boolean](https://github.com/Syssos/Learning_Go/tree/main/0x01_Variables#03_booleans)
	* [array](https://github.com/Syssos/Learning_Go/tree/main/0x04_Loops_and_Arrays#02_arrays)
	* [slice](https://github.com/Syssos/Learning_Go/tree/main/0x05_Slices#slices)
	* [struct](https://github.com/Syssos/Learning_Go/tree/main/0x07_Structs_and_Maps#structs)
	* [map](https://github.com/Syssos/Learning_Go/tree/main/0x07_Structs_and_Maps#00_maps)
	* [pointers](https://github.com/Syssos/Learning_Go/tree/main/0x06_Pointers#pointers)
	* [blank identifier](https://github.com/Syssos/Learning_Go/tree/main/0x01_Variables#blank-identifiers)
- [Composite Literals](https://github.com/Syssos/Learning_Go/tree/main/0x01_Variables#07_Composite)
- [Init Function](https://github.com/Syssos/Learning_Go/tree/main/0x09_Modules_and_Packages#init-function)
- [Methods](https://github.com/Syssos/Learning_Go/tree/main/0x08_Methods_and_Interfaces#00_methods)
- [Interfaces](https://github.com/Syssos/Learning_Go/tree/main/0x08_Methods_and_Interfaces#01_interfaces)
- [Errors](https://github.com/Syssos/Learning_Go/tree/main/0x03_Functions#04_go_errors)

## Conclusion

Go is a langauge that can be vital to giving you the edge over your competitors. The syntax is relativly forgiving and straight forward. While the package management can take a little more time to get used to then say javascript or python that can utilize npm or pip, it can be more powerful in thewhen it comes to the share-ability of code. Theres a boat load of reasons Go is amazing, but I think the best way for you to find out is to get to writing some code.

## Helpful Sources


[Back to Top](#learning-go)