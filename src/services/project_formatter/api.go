package project_formatter

import (
	"context"
	"os/exec"
	"time"
)

func All(path string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return exec.CommandContext(ctx, "go", "fmt", ".../").Run()
}
