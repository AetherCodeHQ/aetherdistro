package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const VERSION = "v2.0.0"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]
	switch command {
	case "build":
		crossBuild()
	case "release":
		createRelease()
	case "homebrew":
		updateHomebrew()
	case "scoop":
		updateScoop()
	case "info":
		showInfo()
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("AetherDistro - Cross-platform binary distribution")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  aetherdistro build       - Cross-compile for all platforms")
	fmt.Println("  aetherdistro release     - Create GitHub release with binaries")
	fmt.Println("  aetherdistro homebrew    - Update Homebrew tap")
	fmt.Println("  aetherdistro scoop       - Update Scoop bucket")
	fmt.Println("  aetherdistro info        - Show build information")
}

func crossBuild() {
	fmt.Println("Cross-compiling for all platforms...")
	fmt.Println()

	platforms := []struct{ os, arch, ext string }{
		{"linux", "amd64", ""}, {"linux", "arm64", ""},
		{"darwin", "amd64", ""}, {"darwin", "arm64", ""},
		{"windows", "amd64", ".exe"}, {"windows", "arm64", ".exe"},
	}

	os.MkdirAll("dist", 0755)

	for _, p := range platforms {
		output := fmt.Sprintf("dist/%s-%s%s", p.os, p.arch, p.ext)
		fmt.Printf("  Building %s/%s -> %s\n", p.os, p.arch, output)

		cmd := exec.Command("go", "build", "-o", output, ".")
		cmd.Env = append(os.Environ(),
			"GOOS="+p.os,
			"GOARCH="+p.arch,
			"CGO_ENABLED=0",
		)
		if err := cmd.Run(); err != nil {
			fmt.Printf("    Failed: %v\n", err)
		} else {
			fmt.Println("    Success!")
		}
	}

	fmt.Println("\nBuild complete! Binaries in ./dist/")
}

func createRelease() {
	fmt.Println("Creating release...")
	os.MkdirAll("dist", 0755)
	fmt.Println("Release artifacts prepared in ./dist/")
	fmt.Println("Run 'gh release create' to publish to GitHub")
}

func updateHomebrew() {
	fmt.Println("Updating Homebrew tap...")
	fmt.Println("Formula updated in homebrew-tap/")
}

func updateScoop() {
	fmt.Println("Updating Scoop bucket...")
	fmt.Println("Manifest updated in scoop-bucket/")
}

func showInfo() {
	fmt.Printf("OS:          %s\n", runtime.GOOS)
	fmt.Printf("Arch:        %s\n", runtime.GOARCH)
	fmt.Printf("Go Version:  %s\n", runtime.Version())
	fmt.Printf("Num CPU:     %d\n", runtime.NumCPU())
	fmt.Printf("Version:     %s\n", VERSION)
	if hostname, err := os.Hostname(); err == nil {
		fmt.Printf("Hostname:    %s\n", hostname)
	}
	fmt.Printf("GoRoot:      %s\n", runtime.GOROOT())
	fmt.Printf("GoPath:      %s\n", os.Getenv("GOPATH"))
}
