# ASCII Art Color

A command-line tool written in Go that renders text as ASCII art with optional color highlighting for a specified substring.

---

## Description

ASCII Art Color takes a text input and displays it in ASCII art format using a standard font file (`standard.txt`). It allows you to highlight a specific substring within the text using ANSI terminal colors.

---

## Requirements

- [Go](https://golang.org/) 1.18 or higher
- A terminal that supports ANSI color codes
- `standard.txt` font file in the same directory as the program

---

## Installation

1. Clone the repository:
```bash
git clone http://your-gitea-server/groupleader/ascii-art-color.git
cd ascii-art-color
```

2. Ensure `standard.txt` is present in the project directory.

---

## Usage

```bash
go run . --color=COLOR SUBSTRING TEXT
```

### Arguments

| Argument    | Description                                      | Required |
|-------------|--------------------------------------------------|----------|
| `--color`   | The color to apply to the substring              | Yes      |
| `SUBSTRING` | The part of the text you want to colorize        | Yes      |
| `TEXT`      | The full text to render in ASCII art             | Yes      |

---

## Supported Colors

| Color    | Code          |
|----------|---------------|
| `red`    | ANSI Red      |
| `green`  | ANSI Green    |
| `blue`   | ANSI Blue     |
| `yellow` | ANSI Yellow   |

---

## Examples

### Colorize a substring in red
```bash
go run . --color=red "ello" "Hello World"
```
This renders **Hello World** in ASCII art, with **ello** highlighted in red.

### Colorize a word in blue
```bash
go run . --color=blue "World" "Hello World"
```

### Multi-line text using `\n`
```bash
go run . --color=green "Hi" "Hi\nThere"
```
This renders two lines of ASCII art, with **Hi** highlighted in green.

---

## How It Works

1. The program reads the `--color`, `SUBSTRING`, and `TEXT` arguments from the command line.
2. It looks up the ANSI escape code for the requested color.
3. It reads the `standard.txt` font file which contains ASCII art representations of each character.
4. It splits the text by `\n` to handle multi-line input.
5. For each line, it finds all occurrences of the substring and records their positions.
6. It prints each character row by row in ASCII art, applying the chosen color to characters that fall within the substring positions.
7. The ANSI reset code is applied after each colored character to restore normal terminal color.

---

## Project Structure

```
ascii-art-color/
├── main.go          # Main program logic
├── standard.txt     # ASCII art font file
└── README.md        # Project documentation
```

---

## Limitations

- Only printable ASCII characters (32–126) are supported.
- Only four colors are currently supported: red, green, blue, yellow.
- The `standard.txt` font file must be present in the same directory.

---

## Authors

- **Charles Locko** (clocko)
- *(Omitogun Oluwatobi)*
- *(Adejumo Segun- Group Leader)*

---

## License

This project is open source and available for educational purposes.