# Modules and Packages

This repository is dedicated to learning about Go modules and packages. It is designed to mimic a project enviornemt to simulate expanding

## Overview

Before we get into creating packages lets take a look at what they really are and how we can benifit from them. The main use of a package, is to create repeatable code, that can be published and utilized by yourself or others for different projects.

To create a package, we "bundle" all of our code up and "label it" with a package name. We can then import that package from other projects as needed using that package name. 

A module is a collection of packages, or can be treated as a package itself, that will contain nested packages. Modules allow us to control the version of each package, use multiple versions of a package, and create dependancy lists, to name a few.


## Creating the Module

First we need to create a directory for a project to sit in. This will be the root folder of a new program we are writing. You can do that with the mkdir command if your using linux. For this example I will be using the folder name "hello".

```
mkdir hello && cd hello
```
Once created lets start with the script that starts the project/program. This is where the main function will live. I will use a file name "Script.go". To get started we will fill it with code similar to previous examples in other directories.

Inside the "Script.go" file add the following code
```go
package main

import "fmt"

func main() {
	fmt.Println("We're making way here!")
}
```

Once the code above is added to the file we will need to create the module. We can do this with the ``` go mod init ``` command. When initualizing we will want a module name that allows us to keep track of where the code can be found, because my code will be on github I will use a name such as "github.com/Syssos/hello".

```
go mod init github.com/Syssos/hello
```

Note the module name used above is different then the location for the local directory on the machine for the module. Modules do not need to be located in a specific path, if you do further research on this topic keep in mind modules were finalized in the 1.3 version

If the module initualizes correctly we should now be able to run the following command, to build the project.

```
go install .
```

If all the steps above were followed, we should see that it returns nothing at all. This is because go install will compile the project and create an executable binary with the name hello, and place it in the ``` $GOPATH/bin ``` folder. This will most likely be ``` $HOME/go/bin ``` for linux users.

If we see the hello file in our /bin folder we can run the binary and this should give us the output "We're makin way here!". This means the start of our "Proram" is working and we can expand it with a package.

For further reading on modules, the blog post I read for fact checking while writing this is a great place to find more information.

[written by Uday Hiwarale](https://medium.com/rungo/anatomy-of-modules-in-go-c8274d215c16)


## example/

The code within this directory is what composes the package that expands our code. While the readme in this directory will go over this content in-depth, I would like to use this space to talk about spreading code out over multiple files.

Let's say we are creating a package that has a lot of code invloved, and we want to break it up into 3 major categories. Loops, syscalls, requests, are the 3 categories we've decided on going with. We can create a 'name'.go file for each and within it place the corresponding code. As long as each page starts with ``` package project_name ``` the functions will be able to interact with eachother as if they are in the same file.

## Script

This file contains the starting point of our program. After adding the package some changes are required if we are to use it inside of our module. Lets look at the example to see those changes.

The first noticable change is the import statement, we can see a new import "github.com/Syssos/hello/example". With the package imported we can now call exportable functions from within it. The next concept I want to touch on is the init function.

### Init Function

Just like the main() function in go the init function has an important role. It will run when the package is first used in our code, becuase our init function is in the main pacakge, it will execute before the main function. This can be an important tool as packages build in complexity and certain things need to be established, like database connections, before the package can be used.

Keep in mind the init function is only called once, meaning if you store code in there to run everytime you call the package, it will not yeild the result you are after.

Inside of the main function we see that the only line is example.Example_func(). This uses the package we created to print a string to the screen.

Once all of the changes are made we can now build/install, and run our code in the same manor as above. From inside the "hello/" directory

``` 
go install .
```
the execute the "hello" binary (stored in $GOPATH/bin)

## Conclusion

Utilizing packages in Go is one way to benifit from some of Go's capabilities. Utilizing packages gives you an even greater advantage. This ability is one that gives Go an edge on other programming languages as it allows for sharing of code that has not been acheived before. Packages and modules are going to be a major part of where time is spent developing in Go, knowing how to use them will be critical. Spend time getting comfortable creating and using modules and packages. The next directory will almost directly expand on this topic so knowing them will not only be important for real life applications, but also to understand further examples.