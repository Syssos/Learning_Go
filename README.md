# Learning Go

Go is an extremely powerful langauge to have in your toolbelt as a developer on any level. Go has amazing built in packages and an amazing method of utilizing external packages. Go's goal is to create lightweight executables without the hastles, or drawbacks other lower level langanges provide. It deals with garbage collection (which coming from C is a very nice touch), and will hold you responsible for imports or variables that are declared and never used. This helps limit the amount of resources required from the machine during code execution. Those are just some of the benefits Go has to offer, this repository will cover the basics of Go and a majority of the benifits Go offers developers.

## Table of Contents

 - [Overview](#learning-go)
 - [Table of Contents](#table-of-contents)
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
 	* [0x02_If_Else](#0x02_if_else)
 	* [0x03_Functions](#0x03_functions)
 	* [0x04_Loops_and_Arrays](#0x04_loops_and_arrays)
 	* [0x05_Slices](#0x05_slices)
 	* [0x06_Pointers](#0x06_pointers)
 	* [0x07_Structs](#0x07_structs)
 	* [0x08_Modules_and_Packages](#0x08_modules_and_packages)
 	* [0x09_Unit_Tests](#0x09_unit_tests)
 - [Conclusion](#conclusion)

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

## Directories

### 0x00_Hello_World
Hello world is all about the basics. In order for Go code to run there are a couple things required. This Directory covers what they are, and why they are needed. Along side learning the requirements for a script this directory will cover printing to the screen, or standard output. As you will see in the directories to come, this will be a very important concept. Even if you have some programming experiance, I recommend at least crusing through the code to see how Go does this.

### 0x01_Variables
This Directory is all about varibales and datypes in Go. While most of these variables are the same in concept as other "low level" programming languages its vital to understand how Go stores data in order to write extremely efficent code. This Directory is dedicated to breaking down variables and datatypes, as well as type conversions so you can "have the right type for the job".

### 0x02_If_Else
This Directory will get you started using if statement. It covers everything from a basic if else, up to a switch case. Which as you will see is essentually a fancy way of doing a bunch of if else statements. Why would you want a bunch of if else statements? Easy, there extremly useful. Not only will they be useful in different types of projects, but they can be used to check for errors, as you will see in 0x03_Functions. 

### 0x03_Functions
In this Directory I will cover what a function is and how we can utilize them in our code. We will also see how functions can be imported from another package, and even how to use a package from github. Functions are workhorse behind most of the projects you will work on, as they allow you to breakup code and make it easier to read.

## 0x04_Loops_and_Arrays
Loops and Arrays is dedicated to for loops and how to use them, and what arrays are. These two go side by side as often times you will be using a for loop do to "something" for each item in an array. This Directory covers all the good stuff you need to know for directorys to come that deal with slices, or arrays, heck even strings can be thought of in the same manor as an array...

## 0x05_Slices
Slices in Go are very similar to arrays but are far more practical to use. This directory is setup to cover why we as developers can benifit from using slices, and how to go about doing so. From creating, to adding, to removing values from a slices this should cover all the basic we need to know. Slices are extremly usefull so be sure to check out what they are and how they can be used.

## 0x06_Pointers
How pointers are used are not much different from programming langauge to programming language. The concept of a pointer is straight forward so there is not much that could change if there were to be differences. If your new to programming spending some time learning how pointers work will benifit you as it is fairly universal information when it comes to working with other languages. If you are a programmer that is just here to learn a new langauge stopping by and reading the material as a refresher it could never hurt.

## 0x07_Structs
If you are used to any object oreinted programming langauges then this is where home will feel. While structs do not make Go an object oriented langauge they offer some features that are closely related to this method of programming. Spending time understand structs and methods can greatly increase your ability to make more complex code. While using them wont be nessisary to complete a task, having them there and knowing how to use them can make a world of difference.

## 0x08_Modules_and_Packages
This directory will cover one of the most unique aspects of Go in my opinion, and that is modules and packages. Here we will see how to create shareable code, as well as learn the backbones we need to be able to create much larger and complex programs. While most online material will have you start here, I have it towards the end because the idea of them really only apply once you have a more complex project, and most of whats covered above will fit within one function.

## 0x09_Unit_Tests
Once we are at the point of writing and creating modules and packages, we will want a way to test all of our code. Thats when the contents of this repository will come into play. Writing unit tests are the best way of adding/changing the code you are writing and ensuring that whatever you are doing does not damage what you are working on. This can be important if you plan on sharing. An update to the code that breaks some functions can leave the people using your code out to dry.

## Conclusion

Because the repo is not finished I do not want to add a conclusion yet, Once the repo is finished the conclusion will be updated.


[Back to Top](#learning-go)