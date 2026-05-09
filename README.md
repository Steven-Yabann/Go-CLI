# gtodo CLI

A lightweight, efficient Command Line Interface (CLI) todo application built in Go. This tool allows you to manage tasks directly from your terminal, storing them locally in a JSON file in your home directory.

Inspiration is from: `[text](https://github.com/heybran/gtodo)`

## Features

* **Persistent Storage**: Tasks are saved to `~/todos.json`.
* **Subcommand Architecture**: Uses Go's `flag.FlagSet` for isolated commands like `add`, `update`, and `delete`.
* **Automated ID Management**: Automatically tracks and assigns unique IDs to your tasks.
* **Safety First**: Built-in initialization checks to ensure your data storage is ready.

## Installation

Ensure you have Go installed on your system. To install `gtodo` globally:

```bash
# Clone the repository
git clone [https://github.com/Steven-Yabann/gtodo.git](https://github.com/Steven-Yabann/gtodo.git)
cd gtodo

# Update dependencies and install
go mod tidy
go install .