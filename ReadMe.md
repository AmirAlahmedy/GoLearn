# Go Examples

This repository contains various examples of Go programming concepts and features. Each directory contains a Go program demonstrating a specific concept.

## Directory Structure

- `arrays/`: Examples of working with arrays in Go.
- `channels/`: Examples of using channels for communication between goroutines.
- `closures/`: Examples of using closures in Go.
- `constants/`: Examples of defining and using constants.
- `custom-errors/`: Examples of creating and handling custom errors.
- `enums/`: Examples of using enumerations.
- `errors/`: Examples of error handling.
- `for/`: Examples of using for loops.
- `functions/`: Examples of defining and using functions.
- `generics/`: Examples of using generics.
- `goroutines/`: Examples of using goroutines for concurrent programming.
- `hello-world/`: A simple "Hello, World!" program.
- `if-else/`: Examples of using if-else statements.
- `interfaces/`: Examples of defining and using interfaces.
- `maps/`: Examples of working with maps.
- `methods/`: Examples of defining and using methods.
- `mutiple-return-values/`: Examples of functions returning multiple values.
- `pointers/`: Examples of using pointers.
- `range-over-built-in-types/`: Examples of using range over built-in types.
- `range-over-iterators/`: Examples of using range over custom iterators.
- `recursion/`: Examples of using recursion.
- `slices/`: Examples of working with slices.
- `strings-runes/`: Examples of working with strings and runes.
- `struct-embedding/`: Examples of struct embedding.
- `structs/`: Examples of defining and using structs.
- `switch/`: Examples of using switch statements.
- `values/`: Examples of working with basic values.
- `variables/`: Examples of defining and using variables.
- `variadic-functions/`: Examples of using variadic functions.
- `watermill/`: An example project using the Watermill library for Azure Service Bus integration.

## Watermill Adapter for Azure Service Bus

The `watermill/` directory contains an example project demonstrating how to use the Watermill library to integrate with Azure Service Bus.

### Makefile Commands

- **Build the application:**

    ```bash
    make or make all
    ```
- **Compile the Go application into a binary in the build/ directory:**
    ```bash
    make build
    ```

- **Remove build artifacts and clean the Go test cache:**

    ```bash
    make clean
    ```

## Requirements

- Go 1.17 or later
- Make (optional for building the example project)

## Running the Examples

To run the examples, navigate to the respective directory and execute the Go program. 

```bash
cd arrays
go run arrays.go
```

## License
This project is licensed under the MIT License.