# Norgate-Calculator
A simple, cross-platform, CLI calculator written in Go.

# TODO
* Improve error handling.
* Create more functions.
* Allow non-spaced input.
* Store input as well as output.

# COMPILING
_Before running anything, ensure that Go is installed and your GOPATH is set._

In a terminal, navigate to the directory in your GOPATH where norgate.go resides and run the following command:

> go install norgate.go

If your GOPATH is not set, you can install Norgate Calculator using the Go build tool:

> go build -o ~/Go/bin/norgate

# RUNNING
To run Norgate after compilation, enter the following command into the terminal:

> ./norgate

If you want to run Norgate without saving the compiled program, type the following command into the terminal:

> go run norgate.go

_All input must be separated by spaces. See the example below._

> Enter an equation.

> 2 + 2 * 2 =

> 2 + . . 4 =

> . . . . 6 =

> 6

# LINKS
[Apache](http://www.apache.org/licenses/LICENSE-2.0)

[Go](https://golang.org/)

# LICENSE
Released under version 3 of the GPL by Michael Rutherford (2015-2017) -- see LICENSE
