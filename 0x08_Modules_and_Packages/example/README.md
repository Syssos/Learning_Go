# 00_Package

## Overview

When dealing with a package inside of a module, it is known as a "nested package". To keep things organized packages will usually be contained inside of a folder in the root directory of the module.

The path of the package will then become relative to the path of the module, ie we can use "module_name/nested_package_name" to import the package. In this case that would look like

```go
import "github.com/Syssos/hello/example"
``` 

where "github.com/Syssos/hello" is the module name, also what we used for the go mod init command, and "example" is the name of the package.

## pretty_print

This file contains the entirety of our package. If we were to want to add more files to help manage the readability of the package thats expanding, we can add those files in this directory. Just ensure those files start with the line package 'package_name'.

This script outside of the changeing of the package name is nothing new to us. Right? 

Well kind of, In this example the name of the function we created is a capital letter. Why is that important? 

This tells the compiler that the function is exportable, and can be called from outside of the package. Meaning if we were to name the function ``` example_func() ```, instead of ``` Example_func() ```, we would not be able to call this function from the Script.go file in the root of the module folder.

This is important to know as it applys to all functions we see in the Go language. For example look at the built in print function we use in this package, ``` fmt.Println() ```. The function starts with a capital letter.

## Conclusion

While this package is not nessesary at all to complete the task, I hope it outlines how a nested package would be used if needed. If the module is placed on github, and just this package is wanted in another project, you can use the path to this package to skip grabbing everything unrelated to this package. There are some amazing onine resources that go over useing packages in go, if you need to do further research keep in mind that modules were added to Go after 1.3, this changed how both packages and modules are used so if your following guides for version of go older then 1.3 keep in mind modules may not have been accessable at the time of creating that article.