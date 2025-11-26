# Concurrency in Go

This repository provides a hands-on approach to learning about concurrency in the Go programming language. It includes practical examples that demonstrate the use of Goroutines, WaitGroups, and Mutexes for managing concurrent operations.

## Project Structure

The repository is organized into the following main components:

-   `main.go`: The entry point of the application. It provides a simple command-line interface (CLI) to navigate through the different concurrency examples.
-   `waitgroups/`: This directory contains examples related to Goroutines and WaitGroups.
-   `mutex/`: This directory contains examples that illustrate the use of Mutexes to prevent race conditions.

## Getting Started

To run the examples in this repository, you need to have Go installed on your system.

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/rishisingh34/concurrency-with-go.git
    cd concurrency-with-go
    ```

2.  **Run the application:**

    ```bash
    go run main.go
    ```

    This will start the interactive CLI.

## Flow of the Program

The `main.go` file acts as a router to the different concurrency examples. When you run the program, you will be presented with a menu of options to choose from:

-   **Option 1: Goroutines and WaitGroups:** This option runs the code in the `waitgroups/` directory, demonstrating how to create and manage Goroutines, and how to use WaitGroups to wait for them to complete.
-   **Option 2: Mutex and Race Conditions:** This option executes the examples in the `mutex/` directory, which show how to use Mutexes to protect shared data and prevent race conditions.

### `waitgroups/`

This directory contains the following examples:

-   `simple_goroutine.go`: A basic example of how to create a Goroutine.
-   `goroutine_with_sleep.go`: An example that shows how a Goroutine can work in the background.
-   `wait_group.go`: Demonstrates the use of `sync.WaitGroup` to wait for a collection of Goroutines to finish executing.
-   `goroutines_usage.go`: A more complex example that combines the concepts of Goroutines and WaitGroups.

### `mutex/`

This directory contains the following examples:

-   `race_condition.go`: An example that intentionally creates a race condition to show its effects.
-   `mutex.go`: Shows how to use a `sync.Mutex` to prevent the race condition seen in the previous example.
-   `sync_mutex.go`: A more detailed example of using Mutexes to protect shared resources.
-   `complex_example.go`: A more complex scenario demonstrating the use of Mutexes in a more realistic application.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.
