# ASCII-Art-Justify Project

Welcome to the **ASCII-Art-Justify** repository! This project builds upon the ASCII-Art suite by introducing advanced text alignment options, allowing dynamic formatting based on terminal size. It was collaboratively developed by me **Tala Amm**, **Noor Halabi**, and **Amro Khweis**.

---

## Overview

The **ASCII-Art-Justify** program generates ASCII art with customizable text alignment options. Users can specify alignment styles such as *left*, *right*, *center*, or *justify* to format their ASCII art dynamically within the terminal.

---

## Objectives

This project emphasizes:
- Advanced **data manipulation** techniques for alignment.
- Utilizing the **file system (fs)** API in Go for template handling.
- Dynamically **adapting text representation** to fit terminal dimensions.

---

## Features

1. **Text Alignment**:
   - Support for alignment types: **left**, **right**, **center**, and **justify**.
   - Text formatting dynamically adapts to terminal width.

2. **Dynamic Terminal Adaptation**:
   - Automatically adjusts output to fit the terminal's size.

3. **Error Handling**:
   - Returns a usage message for incorrect input formats.

4. **Compatibility with Other Projects**:
   - Accepts correctly formatted options and arguments if implemented alongside other ASCII-Art extensions.

---

## Usage

### Run the Program:
```bash
go run . --align=<type> [STRING] [BANNER]
```

### Examples:
1. **Center Alignment**:
   ```bash
   go run . --align=center "hello" standard
   ```

   Output (centered in terminal):
   ```
                       _                _    _
                      | |              | |  | |
                      | |__      ___   | |  | |    ___
                      |  _ \    / _ \  | |  | |   / _ \
                      | | | |  |  __/  | |  | |  | (_) |
                      |_| |_|   \___|  |_|  |_|   \___/
   ```

2. **Left Alignment**:
   ```bash
   go run . --align=left "Hello There" standard
   ```

3. **Right Alignment**:
   ```bash
   go run . --align=right "hello" shadow
   ```

   Output (aligned to the right):
   ```
                                                                                          _|                _| _|
                                                                                          _|_|_|     _|_|   _| _|   _|_|
                                                                                          _|    _| _|_|_|_| _| _| _|    _|
                                                                                          _|    _| _|       _| _| _|    _|
                                                                                          _|    _|   _|_|_| _| _|   _|_|
   ```

4. **Justify Alignment**:
   ```bash
   go run . --align=justify "how are you" shadow
   ```

   Output (justified across terminal width):
   ```
   _|_|_|     _|_|   _|      _|      _|                  _|_|_| _|  _|_|   _|_|
   _|    _| _|    _| _|      _|      _|                _|    _| _|_|     _|_|_|_|
   _|    _| _|    _|   _|  _|  _|  _|                  _|    _| _|       _|
   _|    _|   _|_|       _|      _|                      _|_|_| _|         _|_|_|
   ```

5. **Error Example**:
   ```bash
   go run . --align=invalid "hello" standard
   ```

   Output:
   ```
   Usage: go run . --align=<type> [STRING] [BANNER]
   ```

---

## Collaboration

This project was developed by:
- me: **Tala Amm**
- **Noor Halabi**
- **Amro Khweis**

---

We hope you enjoy working with **ASCII-Art-Justify**! 🚀
