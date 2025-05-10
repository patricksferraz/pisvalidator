# PIS Validator

[![Go Report Card](https://goreportcard.com/badge/github.com/patricksferraz/pisvalidator)](https://goreportcard.com/report/github.com/patricksferraz/pisvalidator)
[![GoDoc](https://godoc.org/github.com/patricksferraz/pisvalidator?status.svg)](https://godoc.org/github.com/patricksferraz/pisvalidator)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.16+-blue.svg)](go.mod)
[![Test Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen.svg)](pis_test.go)

A lightweight, high-performance Go package for validating Brazilian PIS (Programa de Integração Social) numbers. This package is designed to be simple to use, efficient, and reliable for validating PIS numbers in your Go applications.

## 📋 About PIS

The PIS (Programa de Integração Social) is a Brazilian social integration program number used for various social benefits and tax purposes. It's a crucial identifier in Brazilian financial and social systems, similar to a social security number.

For more information about PIS, visit:
- [Caixa - PIS](http://www.caixa.gov.br/beneficios-trabalhador/pis/Paginas/default.aspx)

## ✨ Features

- 🔍 **Accurate Validation**: Implements the official PIS validation algorithm
- 🧹 **Format Cleaning**: Removes formatting characters (dots, dashes, slashes)
- ⚡ **High Performance**: Optimized for speed and memory efficiency
- 🧪 **100% Test Coverage**: Thoroughly tested with comprehensive test cases
- 📦 **Zero Dependencies**: No external dependencies required
- 🛡️ **Type Safety**: Written in pure Go with strong type checking

## 🚀 Installation

```bash
go get github.com/patricksferraz/pisvalidator
```

## 💻 Usage

### Basic Example

```go
package main

import (
    "fmt"
    pisvalidator "github.com/patricksferraz/pisvalidator"
)

func main() {
    // Validate a PIS number
    valid := pisvalidator.ValidatePis("47711617275")
    if !valid {
        fmt.Println("Invalid PIS number")
        return
    }
    fmt.Println("Valid PIS number")

    // Clean a formatted PIS number
    cleaned := pisvalidator.Clean("477.116.172-75")
    fmt.Println(cleaned) // Output: 47711617275
}
```

## 📚 API Reference

### Functions

#### `ValidatePis(pis string) bool`
Validates a PIS number according to the official Brazilian validation algorithm.
- Returns `true` if the PIS number is valid
- Returns `false` if the PIS number is invalid or malformed
- Handles both formatted and unformatted PIS numbers

#### `Clean(s string) string`
Removes formatting characters from a PIS number.
- Removes dots (.), dashes (-), and slashes (/)
- Returns a clean string containing only digits
- Useful for standardizing input before validation

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## ⭐ Show your support

Give a ⭐️ if this project helped you!

## 📫 Contact

Patrick Sferraz - [GitHub Profile](https://github.com/patricksferraz)

Project Link: [https://github.com/patricksferraz/pisvalidator](https://github.com/patricksferraz/pisvalidator)
