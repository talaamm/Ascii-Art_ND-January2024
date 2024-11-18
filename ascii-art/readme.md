# ASCII-Art Project

Welcome to the **ASCII-Art** repository! This project was collaboratively developed by me **Tala Amm**, and **Noor Halabi**, and successfully passed audit review. 🎉

## Overview

**ASCII-Art** is a Go-based program that takes a string as input and outputs a graphical representation of the string using ASCII art. The program supports letters, numbers, spaces, special characters, and newline characters (`\n`) to create visually striking text art.

## Objectives

This project focuses on:
- Working with Go's **file system API**.
- Manipulating and structuring data effectively.
- Learning and adhering to Go best practices.

## Features

1. **Multiple Banner Styles**:
   - `shadow`
   - `standard`
   - `thinkertoy`
2. Handles all input types, including:
   - Letters (uppercase and lowercase)
   - Numbers
   - Special characters
   - Newline characters

3. **Scalable Output**:
   - Creates dynamic ASCII art for any input string.

4. **Robust Testing**:
   - Includes unit tests to ensure accuracy and reliability.

## Usage

1. Clone the repository:
   ```bash
   git clone https://adam-jerusalem.nd.edu/git/taamm/ascii-art.git
   cd ascii-art
   ```

2. Run the program:
   ```bash
   go run . "<your_input_string>"
   ```

3. Example:
   ```bash
   go run . "Hello, World!\n"
   ```

## Examples

### Input:
```
"Hello There"
```

### Output:
```
 _    _          _   _               _______   _                           
| |  | |        | | | |             |__   __| | |                          
| |__| |   ___  | | | |   ___          | |    | |__     ___   _ __    ___  
|  __  |  / _ \ | | | |  / _ \         | |    |  _ \   / _ \ | '__|  / _ \ 
| |  | | |  __/ | | | | | (_) |        | |    | | | | |  __/ | |    |  __/ 
|_|  |_|  \___| |_| |_|  \___/         |_|    |_| |_|  \___| |_|     \___| 
```

## File Format for Banners

Each character in the banner has:
- A height of 8 lines.
- Characters separated by a newline (`\n`).

Example for `' '`, `'!'`, and `'"'`:
```
......
......
......
......
......
......
......
......

._..
|.|.
|.|.
|.|.
|_|.
(_).
....
....

._._..
(.|.).
.V.V..
......
......
......
......
......
```

## Collaboration

This project was a joint effort by:
- me: **Tala Amm**
- **Noor Halabi**


Feel free to explore, use, and improve the code. Contributions are welcome! 😊

--- 

Thank you for checking out our ASCII-Art project! Let us know if you have any questions or suggestions. 💡
