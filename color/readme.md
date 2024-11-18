# ASCII-Art-Color Project

Welcome to the **ASCII-Art-Color** repository! This project builds upon the original ASCII-Art project by adding color manipulation functionality to enhance the visual representation of ASCII art. It was collaboratively developed by **Tala Amm**, **Noor Halabi**, **Amro Khweis**, and **[Your Name]**.

## Overview

The **ASCII-Art-Color** program takes a string as input and outputs a colorful graphical representation using ASCII art. It introduces the ability to selectively apply colors to parts of the string using a simple flag-based system.

## Objectives

This project focuses on:
- Advanced **terminal display** manipulation.
- Working with **color code systems** (e.g., RGB, HSL, ANSI).
- Enhancing **data manipulation** skills.
- Strengthening proficiency with Go's **file system API**.

## Features

1. **Colorful ASCII Art**:
   - Use the `--color=<color> <substring>` flag to specify the color and substring to apply it to.
   - Supports multiple color systems, allowing for flexible and vibrant outputs.

2. **Substring Highlighting**:
   - Apply colors to specific parts of the string.
   - If no substring is specified, the entire string is colored.

3. **Backward Compatibility**:
   - Compatible with other optional ASCII-Art projects.
   - Can run with a single `[STRING]` argument without the color flag.

4. **Usage Message**:
   - Returns a detailed usage message if the format is incorrect.

## Usage

### Run the Program:
```bash
go run . [OPTION] [STRING]
```

### Examples:
1. **Simple ASCII Art**:
   ```bash
   go run . "Hello, World!"
   ```

2. **Colored ASCII Art**:
   ```bash
   go run . --color=red Hello "Hello, World!"
   ```

   Output: The substring `Hello` in `Hello, World!` is colored red.

3. **Full String Coloring**:
   ```bash
   go run . --color=blue "Hello, World!"
   ```

   Output: The entire string `Hello, World!` is colored blue.

4. **Incorrect Format**:
   ```bash
   go run . --color red
   ```

   Output: 
   ```
   Usage: go run . [OPTION] [STRING]
   EX: go run . --color=<color> <substring to be colored> "something"
   ```

## Collaboration

This project was developed by:
- me: **Tala Amm**
- **Noor Halabi**
- **Amro Khweis**

---

Thank you for checking out the **ASCII-Art-Color** project! Contributions, suggestions, and improvements are always welcome. Let us know your thoughts. 🌈
