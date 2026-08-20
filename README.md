# AetherDistro

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-00ADD8?style=for-the-badge)
![CI](https://img.shields.io/badge/CI-GitHub%20Actions-2088FF?style=for-the-badge&logo=githubactions&logoColor=white)
![CodeQL](https://img.shields.io/badge/CodeQL-Security-00ADD8?style=for-the-badge)
![Version](https://img.shields.io/badge/Version-v1.0.0-00ADD8?style=for-the-badge)

> Automated cross-platform binary distribution with Homebrew, Scoop, and APT support

`distribution` `homebrew` `scoop` `cross-compilation` `release` `devops` `ci-cd` `golang`

---

## What is it?

**AetherDistro** is Tag a release and watch the magic: cross-compile for every OS/arch, auto-update Homebrew taps, Scoop buckets, and APT repos. One command to distribute to the world.

## Why should you care?

- **Fast** - Compiled Go binary, zero overhead
- **Secure** - CodeQL analysis + Dependabot
- **Offline-first** - Works without internet
- **Lightweight** - Single binary deployment
- **Developer-friendly** - Clean CLI with docs

---

## Features

- Automatic cross-compilation (Windows/macOS/Linux)
- Homebrew Tap auto-update
- Scoop Bucket manifest generation
- APT/YUM repository support
- Checksum generation (SHA256)
- Release notes automation
- GitHub Release asset upload
- Docker multi-arch images
- Formula version bumping
- Changelog generation

---

## Quick Start

### Prerequisites
- Go 1.21 or higher

### Install from source
```bash
git clone https://github.com/AetherCodeHQ/aetherdistro.git
cd aetherdistro
go build -o aetherdistro .
```

### Run
```bash
./aetherdistro --help
```

---

## Usage

./aetherdistro release --version v1.2.0 --binary mytool  OR  ./aetherdistro publish --homebrew --scoop --apt

---

## CLI Flags

| Flag | Description | Default |
|------|-------------|---------|
| `--path` | Target directory | `.` |
| `--format` | Output format (json, text) | `text` |
| `--output` | Output filename | `stdout` |
| `--verbose` | Enable verbose output | `false` |

---

## Development

```bash
git clone https://github.com/AetherCodeHQ/aetherdistro.git
cd aetherdistro
go build -o aetherdistro .
go test ./...
golangci-lint run
```

---

## Contributing

Contributions welcome! See [CONTRIBUTING.md](CONTRIBUTING.md).

## Security

Report to: aethercode.core@gmail.com | See [SECURITY.md](SECURITY.md)

## License

MIT License - see [LICENSE](LICENSE)

---

<p align="center">
  Built with love by <a href="https://github.com/AetherCodeHQ">AetherCode</a> | <a href="https://github.com/AetherCode-Core">AetherCode Core</a>
</p>


---

## What's New in v1.1.0

- Professional documentation with badges
- CI/CD pipeline with GitHub Actions
- Security analysis with CodeQL
- Dependency management with Dependabot
- Code quality with GolangCI-Lint
- Issue and PR templates
- Contributing guidelines

## Categories

| Category | Description |
|----------|-------------|
| DevOps & Infrastructure | This project is part of the AetherCode ecosystem |

## Related Projects

Part of [AetherCode Core](https://github.com/AetherCode-Core) ecosystem.
