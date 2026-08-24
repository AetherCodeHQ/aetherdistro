# 🌌 aetherdistro

![Go](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=flat-square&logo=go&logoColor=white)
![Version](https://img.shields.io/badge/Version-v2.1.0-00ADD8?style=flat-square)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)
![PRs](https://img.shields.io/badge/PRs-Welcome-brightgreen?style=flat-square)

> AetherCore tool by [AetherCodeHQ](https://github.com/AetherCodeHQ)

`aethercore` `security` `cli` `golang` `exec`

---

## What is aetherdistro?

**aetherdistro** is part of the AetherCore security ecosystem — a zero-dependency, pure Go security utility.

## Features

- ✅ `printUsage()` — Printusage
- ✅ `crossBuild()` — Crossbuild
- ✅ `createRelease()` — Createrelease
- ✅ `updateHomebrew()` — Updatehomebrew
- ✅ `updateScoop()` — Updatescoop
- ✅ `showInfo()` — Showinfo
- ✅ Formatted output
- ✅ External command execution
- 🚀 **Zero dependencies** — only Go standard library
- 📦 **Single binary** — compile and run anywhere
- 🔄 **Offline capable** — no internet required

## Installation

```bash
# Clone
git clone https://github.com/AetherCodeHQ/aetherdistro.git
cd aetherdistro

# Build
go build -o aetherdistro .

# Run
./aetherdistro Usage:
```

### Or directly with `go run`:
```bash
go run main.go Usage:
```

## Usage

```bash
# Basic usage
./aetherdistro Usage:

# With flags
./aetherdistro Usage: value Usage:
```

### Example Output

```
$ ./aetherdistro Usage:
AetherDistro - Cross-platform binary distribution
Usage:
  aetherdistro build       - Cross-compile for all platforms
```

## Project Structure

```
aetherdistro/
  main.go          # Entry point (108 lines)
  go.mod            # Go module definition
  go.sum            # Dependency checksums
  README.md         # This file
  LICENSE           # MIT License
  CHANGELOG.md      # Version history
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License - see [LICENSE](LICENSE) for details.

---

Built with ❤️ by [AetherCodeHQ](https://github.com/AetherCodeHQ)
