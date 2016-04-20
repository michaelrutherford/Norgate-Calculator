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
                fmt.Println (calc)
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
        var opcount int64 = 0
        opcount = countOperations (a)
        for opcount >= 0 {
                a, opcount, answer = squareroot (a, opcount, answer)
                a, opcount, answer = exponent (a, opcount, answer)
                a, opcount, answer = absolute (a, opcount, answer)
                a, opcount, answer = sine (a, opcount, answer)
                a, opcount, answer = cosine (a, opcount, answer)
                a, opcount, answer = tangent (a, opcount, answer)
                a, opcount, answer = factorial (a, opcount, answer)
                a, opcount, answer = multiply (a, opcount, answer)
                a, opcount, answer = divide (a, opcount, answer)
                a, opcount, answer = addition (a, opcount, answer)
                a, opcount, answer = subtraction (a, opcount, answer)
                a, opcount, answer = modulus (a, opcount, answer)
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

func countOperations (a []string) int64 {
        var count int64 = 0
        for t := 0; t < len (a); t++ {
                switch a[t] {
                case "*":
                        count++
                case "/":
                        count++
                case "+":
                        count++
                case "-":
                        count++
                case "%":
                        count++
                case "sqrt":
                        count++
                case "abs":
                        count++
                case "^":
                        count++
                case "sin":
                        count++
                case "cos":
                        count++
                case "tan":
                        count++
                case "!":
                        count++
                }
        }
        count--
        return count
}

func printEquation (a []string) {
        fmt.Println (strings.Trim (fmt.Sprint (a), "[]\n"))
}

func squareroot (a []string, b int64, c float64) ([]string, int64, float64) {
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
                        if one > MAX || one < MIN {
                                ecode = 3
                                err (ecode)
                        }
                        if one < 0 {
                                ecode = 5
                                err (ecode)
                        }
                        one = math.Sqrt (one)
                        a[i + fdist] = "."
                        a[i] = fmt.Sprintf ("%v", one);
                        c = one
                        b--
                        i = 0
                        printEquation (a)
                }
        }
        return a, b, c
}

func addition (a []string, b int64, c float64) ([]string, int64, float64) {
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
                        two = one + two
                        if two > MAX || two < MIN {
                                ecode = 3
                                err (ecode)
                        }
                        a[i] = "."
                        a[i + fdist] = fmt.Sprintf ("%v", two);
                        a[i - bdist] = "."
                        c = two
                        b--
                        i = 0
                        printEquation (a)
                }
        }
        return a, b, c
}

func subtraction (a []string, b int64, c float64) ([]string, int64, float64) {
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
                        two = one - two
                        if two > MAX || two < MIN {
                                ecode = 3
                                err (ecode)
                        }
                        a[i] = "."
                        a[i + fdist] = fmt.Sprintf ("%v", two);
                        a[i - 1] = "."
                        c = two
                        b--
                        i = 0
                        printEquation (a)
                }
        }
        return a, b, c
}

func multiply (a []string, b int64, c float64) ([]string, int64, float64) {
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
                        two = one * two
                        if two > MAX || two < MIN {
                                ecode = 3
                                err (ecode)
                        }
                        a[i] = "."
                        a[i + fdist] = fmt.Sprintf ("%v", two);
                        a[i - bdist] = "."
                        c = two
                        b--
                        i = 0
                        printEquation (a)
                }
        }
        return a, b, c
}

func divide (a []string, b int64, c float64) ([]string, int64, float64) {
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
                        if two == 0 {
                                ecode = 4
                                err (ecode)
                        }
                        if one == 0 {
                                two = 0
                                a[i] = "."
                                a[i + fdist] = fmt.Sprintf ("%v", two);
                                a[i - bdist] = "."
                                c = two
                                b--
                                i = 0
                                printEquation (a)
                                return a, b, c
                        }
                        two = one / two
                        if two > MAX || two < MIN {
                                ecode = 3
                                err (ecode)
                        }
                        a[i] = "."
                        a[i + fdist] = fmt.Sprintf ("%v", two);
                        a[i - bdist] = "."
                        c = two
                        b--
                        i = 0
                        printEquation (a)
                }
        }
        return a, b, c
}

func modulus (a []string, b int64, c float64) ([]string, int64, float64) {
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
                        two = math.Mod (one, two)
                        if two > MAX || two < MIN {
                                ecode = 3
                                err (ecode)
                        }
                        a[i] = "."
                        a[i + fdist] = fmt.Sprintf ("%v", two);
                        a[i - bdist] = "."
                        c = two
                        b--
                        i = 0
                        printEquation (a)
                }
        }
        return a, b, c
}

func factorial (a []string, b int64, c float64) ([]string, int64, float64) {
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
                        if one < 0 {
                                ecode = 2
                                err (ecode)
                        }
                        for track := one - 1; track > 0; track-- {
                                one = one * track
                        }
                        if one > MAX || one < MIN {
                                ecode = 3
                                err (ecode)
                        }
                        a[i - bdist] = "."
                        a[i] = fmt.Sprintf ("%v", one);
                        c = one
                        b--
                        i = 0
                        printEquation (a)
                }
        }
        return a, b, c
}

func sine (a []string, b int64, c float64) ([]string, int64, float64) {
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
                        one = math.Sin (one)
                        if one > MAX || one < MIN {
                                ecode = 3
                                err (ecode)
                        }
                        a[i + fdist] = "."
                        a[i] = fmt.Sprintf ("%v", one);
                        c = one
                        b--
                        i = 0
                        printEquation (a)
                }
        }
        return a, b, c
}

func cosine (a []string, b int64, c float64) ([]string, int64, float64) {
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
                        one = math.Cos (one)
                        if one > MAX || one < MIN {
                                ecode = 3
                                err (ecode)
                        }
                        a[i + fdist] = "."
                        a[i] = fmt.Sprintf ("%v", one);
                        c = one
                        b--
                        i = 0
                        printEquation (a)
                }
        }
        return a, b, c
}

func tangent (a []string, b int64, c float64) ([]string, int64, float64) {
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
                        one = math.Tan (one)
                        if one > MAX || one < MIN {
                                ecode = 3
                                err (ecode)
                        }
                        a[i + fdist] = "."
                        a[i] = fmt.Sprintf ("%v", one);
                        c = one
                        b--
                        i = 0
                        printEquation (a)
                }
        }
        return a, b, c
}

func absolute (a []string, b int64, c float64) ([]string, int64, float64) {
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
                        one = math.Abs (one)
                        if one > MAX || one < MIN {
                                ecode = 3
                                err (ecode)
                        }
                        a[i + fdist] = "."
                        a[i] = fmt.Sprintf ("%v", one);
                        c = one
                        b--
                        i = 0
                        printEquation (a)
                }
        }
        return a, b, c
}

func exponent (a []string, b int64, c float64) ([]string, int64, float64) {
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
                        two, errorTwo := strconv.ParseFloat (a[i + fdist], 64)
                        if errorTwo != nil {
                                ecode = 6
                                err (ecode)
                        }
                        two = math.Pow (one, two)
                        if two > MAX || two < MIN {
                                ecode = 3
                                err (ecode)
                        }
                        a[i - bdist] = "."
                        a[i] = fmt.Sprintf ("%v", two);
                        a[i + fdist] = "."
                        c = two
                        b--
                        i = 0
                        printEquation (a)
                }
        }
        return a, b, c
}
