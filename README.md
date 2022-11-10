# Really F*cking Simple Logger

I wrote this package because I wanted an easy logger that would write out to files and move on to
new files at a certain point so that I could safely store past logs. There are probably alternative
solutions to this problem, and maybe they are better than writing my open package, but why not, right?

## Basic Usage

First, call rfsl.Config()

    func main() {
        rfsl.Config("/var/logs/mylogs.txt", "", 500000000) // 500MB log files
        ...

Then you can use the logger with the included functions

    ...
    rfsl.Trace("This is a trace log!")
    rfsl.Warningf("Something bad happend! Variable 'x' is %d.", x)
    rfsl.Panic("Oh no!")
    ...

mylogs.txt:

    2006-01-02 15:04:05.0000    This is a trace log!
    2006-01-02 15:04:05.0000 WARNING:   Something bad happened! Variable 'x' is 10.
    2006-01-02 15:04:05.0000 **PANIC**	Oh no!