# GoScout

GoScout is a powerful and efficient file search tool written in Go. It allows users to search for specific words across all files in a given directory, regardless of file type.

## Features

- Scans all files in a directory
- Supports word search within files
- Outputs only lines containing the searched word
- Fast and optimized for performance

## Installation

### Prerequisites

- Go (1.18 or later)
- Git

### Clone the repository

```sh
git clone https://github.com/dracuxan/GoScout.git && cd GoScout
```

### Install dependencies

```sh
make deps
```

### Build the project

```sh
make build
```

## Usage

Run the application with the directory path and search term:

```sh
./bin/fs -p <directory_path> -q <search_query>
```

### Run tests

```sh
make test
```

### Check code coverage

```sh
make cover
```

## Contributing

Contributions are welcome! Follow these steps:

1. Fork the repository
2. Create a new branch (`feature-name`)
3. Commit changes and push
4. Create a pull request

## License

Check [LICENSE](LICENSE) for more information.

---

Feel free to submit issues or feature requests via GitHub Issues!
