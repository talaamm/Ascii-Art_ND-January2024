# ASCII-Art-Output Project

Welcome to the **ASCII-Art-Output** repository! This project introduces the ability to write ASCII art output directly into a file, enabling easy storage and sharing of the generated artwork. This was developed by **Noor Halabi**, and me **Tala Amm**.

---

## Overview

The **ASCII-Art-Output** program extends the ASCII-Art suite by allowing users to save the generated ASCII art to a file specified via the `--output` flag.

---

## Objectives

This project emphasizes:
- Mastery of the **file system (fs)** API in Go.
- Advanced **data manipulation** techniques for formatted output.
- File handling and error management.

---

## Features

1. **Output to File**:
   - Write ASCII art output to a file specified by the `--output=<fileName.txt>` flag.

2. **Error Handling**:
   - Returns a usage message for incorrectly formatted input.

3. **Compatibility with Other Projects**:
   - Works with additional correctly formatted options and arguments from other ASCII-Art extensions.

---

## Usage

### Run the Program:
```bash
go run . --output=<fileName.txt> [STRING] [BANNER]
```

### Examples:
1. **Write Output to File**:
   ```bash
   go run . --output=banner.txt "hello" standard
   cat banner.txt
   ```

   Contents of `banner.txt`:
   ```
    _              _   _          
   | |            | | | |         
   | |__     ___  | | | |   ___   
   |  _ \   / _ \ | | | |  / _ \  
   | | | | |  __/ | | | | | (_) | 
   |_| |_|  \___| |_| |_|  \___/  
                                  
                                  
   ```

2. **Output with Shadow Banner**:
   ```bash
   go run . --output=shadow.txt "Hello There!" shadow
   cat shadow.txt
   ```

   Contents of `shadow.txt`:
   ```
   _|    _|          _| _|                _|_|_|_|_| _|                                  _|
   _|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _|
   _|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _|
   _|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          _
   _|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _|
   ```

3. **Error Example**:
   ```bash
   go run . --output fileName "hello" standard
   ```

   Output:
   ```
   Usage: go run . --output=<fileName.txt> [STRING] [BANNER]
   ```

---

## Collaboration

This project was developed by:
- **Noor Halabi**
- me: **Tala Amm**

---

Enjoy working with **ASCII-Art-Output**! 🚀
