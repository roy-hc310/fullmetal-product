package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if _, err := exec.LookPath("kitex"); err != nil {
		fmt.Println("⚠️  kitex not found — installing now...")
		installCmd := exec.Command("go", "install", "github.com/cloudwego/kitex/tool/cmd/kitex@latest")
		installCmd.Stdout = os.Stdout
		installCmd.Stderr = os.Stderr

		if err := installCmd.Run(); err != nil {
			fmt.Printf("❌  Failed to install kitex: %v\n", err)
			os.Exit(1)
		}
	}

	if _, err := exec.LookPath("protoc-gen-go"); err != nil {
		fmt.Println("⚠️  protoc-gen-go not found — installing now...")
		installCmd := exec.Command("go", "install", "google.golang.org/protobuf/cmd/protoc-gen-go@latest")
		installCmd.Stdout = os.Stdout
		installCmd.Stderr = os.Stderr

		if err := installCmd.Run(); err != nil {
			fmt.Printf("❌  Failed to install protoc-gen-go: %v\n", err)
			os.Exit(1)
		}
	}

	cmd := exec.Command("kitex",
		"-module", "github.com/roy-hc310/fullmetal-product",
		"-gen-path", "./pkg/gen/kitex",
		"-I", "../fullmetal-shared/idl/fullmetal-product",
		"../fullmetal-shared/idl/fullmetal-product/product.proto")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	if err := cmd.Run(); err != nil {
		fmt.Printf("❌  kitex generation failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ kitex generation completed successfully.")
}
