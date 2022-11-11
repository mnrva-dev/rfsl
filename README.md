# Really F*cking Simple Logger

I wrote this package because I wanted an easy logger that would write out to files and move on to
new files at a certain point so that I could safely store past logs. There are probably alternative
solutions to this problem, and maybe they are better than writing my own package, but why not, right?

## Basic Usage

First, call rfsl.Config()

```go
func main() {
    rfsl.Config("/var/logs/mylogs.txt", "", 500000000, rfsl.LOG_LVL_INFO) // 500MB log files
    ...
```

Then you can use the logger with the included functions

```go
...
rfsl.Trace("This is a trace log!")
...
rfsl.Warningf("Something unexpected happened! Variable 'x' is %d.", x)
...
rfsl.Panic("Something very bad happened!")
...
rfsl.Fatal("A critical error! The program will shut down.")
...
```

mylogs.txt:

    2022-11-11 00:15:18.3473 TRACE	↦ This is a trace!
    2022-11-11 00:15:18.3473 WARN	↦ Something unexpected happened! Variable 'x' is 10.
    2022-11-11 00:15:18.3473 PANIC	↦ Something very bad happened!
    2022-11-11 00:15:18.3473 FATAL	↦ A critical error! The program will shut down.
