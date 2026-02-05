# Advent of Code 2025 — Runner

This program runs solutions for Advent of Code 2025 days by reading input from text files and dynamically calling the corresponding methods.

---

## Features

- Run a specific day (`-day N`, default: `1`)
- Use example inputs with `-E`
- Load input files from a custom directory (`-dir ./inputs`)
- Display both parts of the solution

---

## Expected Structure

```
./inputs/
├── day01.txt
├── day02.txt
├── ...
├── example_day01.txt
├── example_day02.txt
└── ...
```

Files must be named `dayXX.txt` or `example_dayXX.txt`.

---

## Usage

```bash
go run main.go -day 1
go run main.go -day 1 -E
go run main.go -day 1 -dir ./myinputs
```
