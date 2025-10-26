package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if _, err := exec.LookPath("sqlc"); err != nil {
		fmt.Println("⚠️  sqlc not found — installing now...")
		installCmd := exec.Command("go", "install", "github.com/sqlc-dev/sqlc/cmd/sqlc@latest")
		installCmd.Stdout = os.Stdout
		installCmd.Stderr = os.Stderr

		if err := installCmd.Run(); err != nil {
			fmt.Printf("❌  Failed to install sqlc: %v\n", err)
			os.Exit(1)
		}
	}

	cmd := exec.Command("sqlc", "generate", "-f", "sqlc.yaml")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	if err := cmd.Run(); err != nil {
		fmt.Printf("❌  sqlc generation failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅  sqlc generation completed successfully.")
}
