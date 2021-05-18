# Packages

This repository is dedicated to learning all about what a package is, and how to start creating one. Hopefully by now, you have been on your way to creating some functions that help you either print stuff, do mathimatics, maybe even a mini question based game. Lets say you want to share that so others can use it in there code. How do you do that? Well you start by creating a package.

## Overview

Before we get into creating packages lets take a look at what they really are and how we can benifit from them. The main use of a package, is to create repeatable code, that can be published and utilized by yourself or others for different projects.

To create a package, we "bundle" all of our code up and "label it" with a package name. We can then import that package from other projects as needed using that package name. I have a wonderful blog post going into much greater detail about all of the steps discussed in this directory, you can find it here -> [Under construction, sorry](#overview)

## Project Enviornment

So theres one more thing we need to touch on before we continue, thats the project enviornment. A package is intented to be used in a project, its not the project itself. That being said we need to do some preparing for this directories examples.

First we need to create a directory for the project to sit in. You can do that with the mkdir command if your using linux. For this example I will be using hello.

```
mkdir hello && cd hello
```

Once created lets start with the script that starts the main project/program. This is where the main function will live. I will use a file name "Start.go". Inside this script will sit the following code
```go
package main

import "fmt"

func main() {
	fmt.Println("We're making way here!")
}
```

Once the script is added to the folder we will need to create a module. We can do this with the ``` go mod init ``` command. When initualizing we will want a path that allows for the code to be found, because the my code will be on github I will use a path such as "github.com/Syssos/'name'"

```
go mod init github.com/Syssos/hello
```

Note the path used above is different then the location local on the machine for the package, while creating and testing this should not effect the ability to build and run locally, however if this code were intented to be shared on github it would allow others to be able to utilize the package. This is because the Go compiler can follow the link to github and retreive the package.

If everything above went correctly we should now be able to run the following command

```
go install .
```
If we see errors please ensure to handle them, however if all the steps above are followed, and the output is correct, we should see that it returns nothing at all. This is because go install will compile the project and create an executable binary with the name hello, and place it in the ``` $GOPATH/bin ``` folder. This will most likely be ``` $HOME/go/bin ``` for linux users.

If we see the hello file in our /bin folder we can run the binary and this should give us the output "We're makin way here!"

If everything above is working for you, we can now start working on creating a package.

## 00_Packages
<!-- Cover creating package -->

## 01_Multiple_files
<!-- Cover how multiple files work within a package -->

## 02_github
<!-- Quick cover on creating repo, then how to add it in a way others can impliment within seconds -->

## Conclusion