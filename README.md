# Norgate-Calculator
A simple, cross-platform, CLI calculator written in Go.

# TODO
* Improve error handling.
* Create a makefile.
* Create more functions.
* Create applicaton icon.

# COMPILING
_Before running anything, ensure that Go is installed and your GOPATH is set._

In a terminal, navigate to the directory in your GOPATH where norgate.go resides and run the following command:

> go install norgate.go

If your GOPATH is not set, you can install Norgate Calculator using the Go build tool:

> go build -o ~/Go/bin/norgate

To run Norgate after compilation, enter the following command into the terminal:

> ./norgate

If you want to run Norgate without saving the compiled program, type the following command into the terminal:

> go run norgate.go

# RUNNING
_All input must be seperated by spaces. See the example below._

> Enter an equation.

> 2 + 2 * 2 =

> 2 + . . 4 =

> . . . . 6 =

> 6

# LINKS
[Go](https://golang.org/)

[GPL](http://www.gnu.org/licenses/gpl-3.0.html)

# LICENSE
Released under the Apache License v2.0 by Michael Rutherford (2015) -- see LICENSE
