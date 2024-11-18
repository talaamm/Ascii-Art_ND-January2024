# ASCII-Art-FS Project

Welcome to the **ASCII-Art-FS** repository! This project extends the ASCII-Art project by introducing support for different graphical templates (banners) specified as an additional argument. It was collaboratively developed by me **Tala Amm**, and **Noor Halabi**.

---

## Overview

The **ASCII-Art-FS** program generates ASCII art using customizable templates (banners). Users can specify a string and a banner type, and the program outputs the corresponding ASCII representation.

---

## Objectives

This project emphasizes:
- Working with the **file system (fs)** API in Go.
- Enhanced **data manipulation** through file parsing.
- Understanding and utilizing **custom templates** for ASCII art.

---

## Features

1. **Template (Banner) Support**:
   - Choose from predefined templates such as `standard`, `shadow`, and `thinkertoy`.
   - Easily extensible to support additional templates.

2. **Flexible Usage**:
   - Specify a string and template to customize output.
   - Supports other correctly formatted options and arguments if implemented alongside optional ASCII-Art projects.

3. **Error Handling**:
   - Returns a usage message for incorrect input formats.

---

## Usage

### Run the Program:
```bash
go run . [STRING] [BANNER]
```

### Examples:
1. **Default ASCII Art**:
   ```bash
   go run . "hello" standard
   ```

2. **Custom Template**:
   ```bash
   go run . "Hello There!" shadow
   ```

   Output: 
   ```
   _|    _|          _| _|                _|_|_|_|_| _|                                  _|
   _|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _|
   _|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _|
   _|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          _
   _|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _|  
   ```

3. **Another Template**:
   ```bash
   go run . "Hello There!" thinkertoy
   ```

4. **Error Example**:
   ```bash
   go run . "Hello" unknown
   ```

   Output:
   ```
   Usage: go run . [STRING] [BANNER]
   ```

---

## Collaboration

This project was developed by:
- me: **Tala Amm**
- **Noor Halabi**

---

We hope you enjoy working with **ASCII-Art-FS**! Suggestions, issues, and contributions are welcome. 🎨
