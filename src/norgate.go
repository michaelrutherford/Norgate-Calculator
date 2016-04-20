/*
Copyright 2015-2016 Michael Rutherford

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
        "fmt"
        "os"
        "strconv"
        "strings"
        "math"
        "bufio"
)

const MAX = 99999999
const MIN = -99999999

var ecode int = 0
var history [50]float64;
var hist int = 0

func main () {
        clear ()
        start:
        fmt.Print ("Enter an equation or operation.\n")
        input := ""
        inread := bufio.NewReader (os.Stdin)
        in, inerr := inread.ReadString ('\n')
        if inerr != nil {
                ecode = 6
                err (ecode)
        }
        input = in
        if strings.Contains (input, "history") {
                for hval := 0; hval < hist; hval++ {
                        fmt.Println (history[hval])
                }
        } else if strings.Contains (input, "exit") {
                os.Exit (0)
        } else {
                calc := solve (splice (input))
                if hist >= 50 {
                        clear ()
                }
                history[hist] = calc
                hist++
                fmt.Print (calc, "\n")
        }
        goto start
}

func splice (a string) []string {
        var eqcoll []string
        if strings.Contains (a, "=") == true {
                eqcoll = strings.Split (a, " ")
        } else {
                ecode = 2
                err (ecode)
        }
        return eqcoll
}

func solve (a []string) float64 {
        var answer float64 = 0
        var opcount float64 = 0
        for t := 0; t < len (a); t++ {
                switch a[t] {
                case "*":
                        opcount++
                case "/":
                        opcount++
                case "+":
                        opcount++
                case "-":
                        opcount++
                case "%":
                        opcount++
                case "sqrt":
                        opcount++
                case "abs":
                        opcount++
                case "^":
                        opcount++
                case "sin":
                        opcount++
                case "cos":
                        opcount++
                case "tan":
                        opcount++
                case "!":
                        opcount++
                }
        }
        opcount--
        for opcount >= 0 {
                for i := 0; i < len (a); i++ {
                        fdist := 1
                        if a[i] == "sqrt" {
                                for y := i; y < len (a); y++ {
                                        if a[i + fdist] == "." {
                                                fdist++
                                        }
                                }
                                one, errorOne := strconv.ParseFloat (a[i + fdist], 64)
                                if errorOne != nil {
                                        ecode = 6
                                        err (ecode)
                                }
                                one = squareroot (one)
                                a[i + fdist] = "."
                                a[i] = fmt.Sprintf ("%v", one);
                                answer = one
                                opcount--
                                fmt.Println (strings.Trim (fmt.Sprint (a), "[]"))
                                i = 0
                        }
                }
                for i := 0; i < len (a); i++ {
                        fdist := 1
                        bdist := 1
                        if a[i] == "^" {
                                for y := i; y < len (a); y++ {
                                        if a[i + fdist] == "." {
                                                fdist++
                                        }
                                }
                                for y := i; y >= 0; y-- {
                                        if a[i - bdist] == "." {
                                                bdist++
                                        }
                                }
                                one, errorOne := strconv.ParseFloat (a[i - bdist], 64)
                                if errorOne != nil {
                                        ecode = 6
                                        err (ecode)
                                }
                                one = factorial (one)
                                a[i - bdist] = "."
                                a[i] = fmt.Sprintf ("%v", one);
                                answer = one
                                opcount--
                                fmt.Println (strings.Trim (fmt.Sprint (a), "[]"))
                                i = 0
                        }
                }
                for i := 0; i < len (a); i++ {
                        fdist := 1
                        if a[i] == "abs" {
                                for y := i; y < len (a); y++ {
                                        if a[i + fdist] == "." {
                                                fdist++
                                        }
                                }
                                one, errorOne := strconv.ParseFloat (a[i + fdist], 64)
                                if errorOne != nil {
                                        ecode = 6
                                        err (ecode)
                                }
                                one = absolute (one)
                                a[i + fdist] = "."
                                a[i] = fmt.Sprintf ("%v", one);
                                answer = one
                                opcount--
                                fmt.Println (strings.Trim (fmt.Sprint (a), "[]"))
                                i = 0
                        }
                }
                for i := 0; i < len (a); i++ {
                        fdist := 1
                        if a[i] == "sin" {
                                for y := i; y < len (a); y++ {
                                        if a[i + fdist] == "." {
                                                fdist++
                                        }
                                }
                                one, errorOne := strconv.ParseFloat (a[i + fdist], 64)
                                if errorOne != nil {
                                        ecode = 6
                                        err (ecode)
                                }
                                one = sine (one)
                                a[i + fdist] = "."
                                a[i] = fmt.Sprintf ("%v", one);
                                answer = one
                                opcount--
                                fmt.Println (strings.Trim (fmt.Sprint (a), "[]"))
                                i = 0
                        }
                }
                for i := 0; i < len (a); i++ {
                        fdist := 1
                        if a[i] == "cos" {
                                for y := i; y < len (a); y++ {
                                        if a[i + fdist] == "." {
                                                fdist++
                                        }
                                }
                                one, errorOne := strconv.ParseFloat (a[i + fdist], 64)
                                if errorOne != nil {
                                        ecode = 6
                                        err (ecode)
                                }
                                one = cosine (one)
                                a[i + fdist] = "."
                                a[i] = fmt.Sprintf ("%v", one);
                                answer = one
                                opcount--
                                fmt.Println (strings.Trim (fmt.Sprint (a), "[]"))
                                i = 0
                        }
                }
                for i := 0; i < len (a); i++ {
                        fdist := 1
                        if a[i] == "tan" {
                                for y := i; y < len (a); y++ {
                                        if a[i + fdist] == "." {
                                                fdist++
                                        }
                                }
                                one, errorOne := strconv.ParseFloat (a[i + fdist], 64)
                                if errorOne != nil {
                                        ecode = 6
                                        err (ecode)
                                }
                                one = tangent (one)
                                a[i + fdist] = "."
                                a[i] = fmt.Sprintf ("%v", one);
                                answer = one
                                opcount--
                                fmt.Println (strings.Trim (fmt.Sprint (a), "[]"))
                                i = 0
                        }
                }
                for i := 0; i < len (a); i++ {
                        fdist := 1
                        bdist := 1
                        if a[i] == "!" {
                                for y := i; y < len (a); y++ {
                                        if a[i + fdist] == "." {
                                                fdist++
                                        }
                                }
                                for y := i; y >= 0; y-- {
                                        if a[i - bdist] == "." {
                                                bdist++
                                        }
                                }
                                one, errorOne := strconv.ParseFloat (a[i - bdist], 64)
                                if errorOne != nil {
                                        ecode = 6
                                        err (ecode)
                                }
                                one = factorial (one)
                                a[i - bdist] = "."
                                a[i] = fmt.Sprintf ("%v", one);
                                answer = one
                                opcount--
                                fmt.Println (strings.Trim (fmt.Sprint (a), "[]"))
                                i = 0
                        }
                }
                for i := 0; i < len (a); i++ {
                        fdist := 1
                        bdist := 1
                        if a[i] == "*" {
                                for y := i; y < len (a); y++ {
                                        if a[i + fdist] == "." {
                                                fdist++
                                        }
                                }
                                for y := i; y >= 0; y-- {
                                        if a[i - bdist] == "." {
                                                bdist++
                                        }
                                }
                                one, errorOne := strconv.ParseFloat (a[i - bdist], 64)
                                if errorOne != nil {
                                        ecode = 6
                                        err (ecode)
                                }
                                two, errorTwo := strconv.ParseFloat (a[i + fdist], 64)
                                if errorTwo != nil {
                                        ecode = 6
                                        err (ecode)
                                }
                                two = multiply (one, two)
                                a[i] = "."
                                a[i + fdist] = fmt.Sprintf ("%v", two);
                                a[i - bdist] = "."
                                answer = two
                                opcount--
                                fmt.Println (strings.Trim (fmt.Sprint (a), "[]"))
                                i = 0
                        }
                }
                for i := 0; i < len (a); i++ {
                        fdist := 1
                        bdist := 1
                        if a[i] == "/" {
                                for y := i; y < len (a); y++ {
                                        if a[i + fdist] == "." {
                                                fdist++
                                        }
                                }
                                for y := i; y >= 0; y-- {
                                        if a[i - bdist] == "." {
                                                bdist++
                                        }
                                }
                                one, errorOne := strconv.ParseFloat (a[i - bdist], 64)
                                if errorOne != nil {
                                        ecode = 6
                                        err (ecode)
                                }
                                two, errorTwo := strconv.ParseFloat (a[i + fdist], 64)
                                if errorTwo != nil {
                                        ecode = 6
                                        err (ecode)
                                }
                                two = divide (one, two)
                                a[i] = "."
                                a[i + fdist] = fmt.Sprintf ("%v", two);
                                a[i - bdist] = "."
                                answer = two
                                opcount--
                                fmt.Println (strings.Trim (fmt.Sprint (a), "[]"))
                                i = 0
                        }
                }
                for i := 0; i < len (a); i++ {
                        fdist := 1
                        bdist := 1
                        if a[i] == "+" {
                                for y := i; y < len (a); y++ {
                                        if a[i + fdist] == "." {
                                                fdist++
                                        }
                                }
                                for y := i; y >= 0; y-- {
                                        if a[i - bdist] == "." {
                                                bdist++
                                        }
                                }
                                one, errorOne := strconv.ParseFloat (a[i - bdist], 64)
                                if errorOne != nil {
                                        ecode = 6
                                        err (ecode)
                                }
                                two, errorTwo := strconv.ParseFloat (a[i + fdist], 64)
                                if errorTwo != nil {
                                        ecode = 6
                                        err (ecode)
                                }
                                two = add (one, two)
                                a[i] = "."
                                a[i + fdist] = fmt.Sprintf ("%v", two);
                                a[i - bdist] = "."
                                answer = two
                                opcount--
                                fmt.Println (strings.Trim (fmt.Sprint (a), "[]"))
                                i = 0
                        }
                }
                for i := 0; i < len (a); i++ {
                        fdist := 1
                        bdist := 1
                        if a[i] == "-" {
                                for y := i; y < len (a); y++ {
                                        if a[i + fdist] == "." {
                                                fdist++
                                        }
                                }
                                for y := i; y >= 0; y-- {
                                        if a[i - bdist] == "." {
                                                bdist++
                                        }
                                }
                                one, errorOne := strconv.ParseFloat (a[i - bdist], 64)
                                if errorOne != nil {
                                        ecode = 6
                                        err (ecode)
                                }
                                two, errorTwo := strconv.ParseFloat (a[i + fdist], 64)
                                if errorTwo != nil {
                                        ecode = 6
                                        err (ecode)
                                }
                                two = subtract (one, two)
                                a[i] = "."
                                a[i + fdist] = fmt.Sprintf ("%v", two);
                                a[i - 1] = "."
                                answer = two
                                opcount--
                                fmt.Println (strings.Trim (fmt.Sprint (a), "[]"))
                                i = 0
                        }
                }
                for i := 0; i < len (a); i++ {
                        fdist := 1
                        bdist := 1
                        if a[i] == "%" {
                                for y := i; y < len (a); y++ {
                                        if a[i + fdist] == "." {
                                                fdist++
                                        }
                                }
                                for y := i; y < len (a); y++ {
                                        if a[i - bdist] == "." {
                                                bdist++
                                        }
                                }
                                one, errorOne := strconv.ParseFloat (a[i - bdist], 64)
                                if errorOne != nil {
                                        ecode = 6
                                        err (ecode)
                                }
                                two, errorTwo := strconv.ParseFloat (a[i + fdist], 64)
                                if errorTwo != nil {
                                        ecode = 6
                                        err (ecode)
                                }
                                two = modulus (one, two)
                                a[i] = "."
                                a[i + fdist] = fmt.Sprintf ("%v", two);
                                a[i - bdist] = "."
                                answer = two
                                opcount--
                                fmt.Println (strings.Trim (fmt.Sprint (a), "[]"))
                                i = 0
                        }
                }
        }
        return answer
}

func err (e int) {
        switch (e) {
        case 1:
                fmt.Println ("Error 1: Overflow.")
                os.Exit (1)
        case 2:
                fmt.Println ("Error 2: Invalid operation.")
                os.Exit (2)
        case 3:
                fmt.Println ("Error 3: Post-operation overflow.")
                os.Exit (3)
        case 4:
                fmt.Println ("Error 4: Divide by zero.")
                os.Exit (4)
        case 5:
                fmt.Println ("Error 5: Imaginary number.")
                os.Exit (5)
        case 6:
                fmt.Println ("Error 6: Error parsing.")
                os.Exit (6)
        default:
                fmt.Println ("Error 0: Generic error.")
                os.Exit (0)
        }
}

func clear () {
        for hval := 0; hval < len (history); hval++ {
                history[hval] = 0.0
        }
        hist = 0
}

func add (a, b float64) float64 {
        if a + b > MAX || a + b < MIN {
                ecode = 3
                err (ecode)
        }
        return a + b
}

func subtract (a, b float64) float64 {
        if a - b > MAX || a - b < MIN {
                ecode = 3
                err (ecode)
        }
        return a - b
}

func multiply (a, b float64) float64 {
        if a * b > MAX || a * b < MIN {
                ecode = 3
                err (ecode)
        }
        return a * b
}

func divide (a, b float64) float64 {
        if b == 0 {
                ecode = 4
                err (ecode)
        } else if a == 0 {
                return 0
        }
        if a / b > MAX || a / b < MIN {
                ecode = 3
                err (ecode)
        }
        return a / b
}

func modulus (a, b float64) float64 {
        if math.Mod (a, b) > MAX || math.Mod (a, b) < MIN {
                ecode = 3
                err (ecode)
        }
        return math.Mod (a, b)
}

func factorial (a float64) float64 {
        if a < 0 {
                ecode = 2
                err (ecode)
        }
        for track := a - 1; track > 0; track-- {
                a = a * track
        }
        if a > MAX || a < MIN {
                ecode = 3
                err (ecode)
        }
        return a
}

func squareroot (a float64) float64 {
        if math.Sqrt (a) > MAX || math.Sqrt (a) < MIN {
                ecode = 3
                err (ecode)
        }
        if a < 0 {
                ecode = 5
                err (ecode)
        }
        return math.Sqrt (a)
}

func sine (a float64) float64 {
        if math.Sin (a) > MAX || math.Sin (a) < MIN {
                ecode = 3
                err (ecode)
        }
        return math.Sin (a)
}

func cosine (a float64) float64 {
        if math.Cos (a) > MAX || math.Cos (a) < MIN {
                ecode = 3
                err (ecode)
        }
        return math.Cos (a)
}

func tangent (a float64) float64 {
        if math.Tan (a) > MAX || math.Tan (a) < MIN {
                ecode = 3
                err (ecode)
        }
        return math.Tan (a)
}

func absolute (a float64) float64 {
        if math.Abs (a) > MAX || math.Abs (a) < MIN {
                ecode = 3
                err (ecode)
        }
        return math.Abs (a)
}

func exponent (a, b float64) float64 {
        if math.Pow (a, b) > MAX || math.Pow (a, b) < MIN {
                ecode = 3
                err (ecode)
        }
        return math.Pow (a, b)
}
