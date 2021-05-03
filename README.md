# Learning Go
### The tasks I used to learn the basics
Go is an extremely powerful langauge to have in your toolbelt as a developer on any level. It has a bunch of packages built in that can be utilized and its goal is to create lightweight executables without the hastle of other lower level langanges. It deals with memory management (which coming from C is a very nice touch), and will hold you responsible for imports or variables that are declared and never used.

I could turn this readme into a list with all the reasons Go is a good langauge to learn, but for the sake of time and to get back to learning I'll just move along. In the following directories are Go scripts, grouped together by category, I used to learn the langauge. Alongside anther readme describing the point of each script.

 - [Download](#download)
 - [Installation](#installation)
 	* [Linux](#linux)
 	* [Windows](#windows)
 	* [Mac](#mac)
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

1. Extract the information within the downloaded archive file downloaded from the Go website, to /usr/local this should look like ``` go[version #].linux-amd64.tar.gz ```

	**Important:** Files stored in the location /usr/local/go will need to be removed or overwritten if they exist to ensure the installation has no errors. To ensure the location is cleared and no files are there we will remove that file and then extract the tar.gz file

	\* Note that in this case I am using the Go version 1.16.3 tar.gz at the time of creating this, so the file name may need to be replaced if you copy and paste this snippet
	
	``` sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.16.3.linux-amd64.tar.gz ```

2. Now that the Go files are in the local directory we need to add the PATH environment variable. To do this open either $HOME/.profile (\~/.prfile) or /etc/profile in a text editor of your choice and append the following line to the end of the file.
	
	```	export PATH=$PATH:/usr/local/go/bin	```

	or you can run the command below to echo and append the line.
	
	```	echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.profile ```

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

## Directories

### 0x00_Hello_World
This Directory is dedicated to learning the basics of Go. I Learned how to run and build the code as well as the basic syntax and file structure. I also looked into printing strings to the standard output as this will be used to show everything we are doing as we learn and progress.

### 0x01_Variables
This Directory is all about varibales and datypes in Go. There are a few that are different from C that I would like to highlight, as well as cover the basics of how go handles common datatypes.