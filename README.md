# GoFigure

`GoFigure` is a `Go` package that provides an easy way to format large numbers into more readable strings with abbreviations like `K`, `M`, `B`, `T`, `Q`, and `Qi`.

## Installation

Install GoFigure using the following command:

```bash
go get github.com/theHamdiz/gofigure
```


### Usage

> Here's how to use `GoFigure` in your `Go` programs:

```go
package main

import (
    "fmt"

    "github.com/theHamdiz/gofigure"
)

func main() {
    fmt.Println(gofigure.FormatFigure(123))  // Output: "123"
    fmt.Println(gofigure.FormatFigure(1234)) // Output: "1.2K"
    fmt.Println(gofigure.FormatFigure(1234567)) // Output: "1.2M"
    fmt.Println(gofigure.FormatFigure(1234567890)) // Output: "1.2B"
    fmt.Println(gofigure.FormatFigure(1234567890123)) // Output: "1.2T"
    fmt.Println(gofigure.FormatFigure(1234567890123456)) // Output: "1.2Q"
    fmt.Println(gofigure.FormatFigure(1234567890123456789)) // Output: "1.2Qi"
}

```
