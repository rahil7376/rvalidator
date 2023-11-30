
# rvalidator - Enhanced Struct Validation in Go

rvalidator is a Go package that extends the functionality of the github.com/go-playground/validator/v10 package. It provides an enhanced way to perform struct validation with custom error messages. This package is especially useful for projects where clear, user-friendly validation feedback is crucial, such as web applications, APIs, and data processing tools.


# Features

* Custom Error Messages: Easily define custom error messages for each field in your struct.
* Reflect-Based Validation: Dynamically validates structs using reflection, offering flexibility for various applications.
* Clear Error Distinction: Differentiates between validation errors and other types of errors for better error handling.
* Simple and Intuitive API: Designed to be straightforward and easy to integrate into Go projects.


# Installation
```
go get -u github.com/rahil7376/rvalidator

```

# Usage

To use rvalidator, import it into your Go project and pass your structs to the Validate function. The function differentiates between validation errors and other types of errors for efficient error handling.

Here's a basic usage example:

```
package main

import (
    "fmt"
    "github.com/yourusername/rvalidator"
)

type User struct {
    Name  string `validate:"required" errormessage:"Name is required"`
    Email string `validate:"required,email" errormessage:"Valid email is required"`
    Age   int    `validate:"gte=18" errormessage:"Age must be at least 18"`
}

func main() {
    user := User{
        Name:  "",
        Email: "invalid-email",
        Age:   16,
    }

    errorsStrings, err := rvalidator.Validate(user)
    if err != nil {
        // Handle non-validation errors
        fmt.Println("Error:", err)
    } else if len(errorsStrings) > 0 {
        // Handle validation errors
        fmt.Println("Validation Errors:")
        for _, e := range errorsStrings {
            fmt.Println(e)
        }
    } else {
        // No errors, struct is valid
        fmt.Println("User is valid")
    }
}


```



