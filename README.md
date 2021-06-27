# pisvalidator

Brazilian pis validator (PIS) - validation package in Golang.

It is an essential package to validate PIS number in your application.

## Information about PIS

- [Caixa - PIS](http://www.caixa.gov.br/beneficios-trabalhador/pis/Paginas/default.aspx)

## Installation

Use the `go tool` for do that:

```bash
go get github.com/patricksferraz/pisvalidator
```

## Usage

You need to import the _package_ `github.com/patricksferraz/pisvalidator` like that:

```go
import (
    pisvalidator "github.com/patricksferraz/pisvalidator"
)
```

## Example

```go
package main

import (
    "fmt"
    pisvalidator "github.com/patricksferraz/pisvalidator"
)

func main() {

    // This will validate the PIS and clean the formatting.
    valid := pisvalidator.ValidatePis("47711617275")

    // Verifies if it is a valid PIS
    if !valid {
        panic(fmt.Errorf("It isn't valid: %v", err))
    }
}
```

## License

[![License](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat)](LICENSE)
