# Password Generator (CLI)

[![Go](https://github.com/KnYaZ-95/PasswordGenerator/actions/workflows/go.yml/badge.svg)](https://github.com/KnYaZ-95/PasswordGenerator/actions/workflows/go.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENCE)

A simple CLI password generator written in Go with configurable length and special character support

## 📦 Features
- Generate passwords of custom length
- Optional special character inclusion
- Interactive CLI interface

## 🚀 Installation & Usage

### 1. Clone the repository
```bash
git https://github.com/KnYaZ-95/PasswordGenerator.git
cd PasswordGenerator
```

### 2. Run without building
```bash
go run main.go
```

### 3. Build executable
```bash
go build -o password-generator . # Linux
go build -o password-generator.exe  # Windows
```
## 🛠 How to Use

After launching, the program will prompt you to:
1.  Enter password length
2.  Choose whether to include special characters (y/n)
3.  Receive your password with option to generate another

## 🧪 Testing

Run tests with:
```bash
go test -v ./passwordGenerator
```
## 📝 License
This project is licensed under the MIT License - see the [LICENCE](./LICENCE) file for details