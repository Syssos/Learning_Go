# Basic Example [</>](https://github.com/Syssos/Learning_Go/blob/main/0x11_Concurrency/00_Basic_Example/example.go)

This example was designed to outline 2 things. First Goroutines can be super simple to impliment in your code. Theres hardly any setup, and with good variable and function names the code easy to read. Second Goroutines are meant to help schedule events that will take time to complete, as to not waste time in the end result.

Synchronously, without the use of the go keyword, if it wasn't for the phrase "Starting" you may not have known for up to eight seconds if the code was actually running. This is becuase immediatly after printing the intital phrase the program would sleep for eight seconds. Then print "From Go Routine", then print "Not in Routine", indefinetly (or until told otherwise) every second.

If this example is ran as is, you will notice that, from the start, as each second passes the phrase "Not in Routine" is printed. However after 8 seconds have pased you will notice the phrase "From Go Routine" is printed.

Example output:

```
Starting
Not in Routine
Not in Routine
Not in Routine
Not in Routine
Not in Routine
Not in Routine
Not in Routine
From Go Routine
Not in Routine
Not in Routine

```

## Use Case's

While this particular example may not have real world implications, it deminstrates how easily a goroutine can be utilized by a developer. In the next example, [01_Advanced_Example](https://github.com/Syssos/Learning_Go/tree/main/0x11_Concurrency/01_Advanced_Example), you will see a use case for Go routines that speed up the process of making multiple HTTP requests.