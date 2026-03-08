package builder

import (
	"fmt"
	"os/exec"
)

func Build(command string) error {

	fmt.Println("Running build command:", command)

	cmd := exec.Command("powershell", "-Command", command)

	output, err := cmd.CombinedOutput()

	fmt.Println(string(output))

	if err != nil {
		fmt.Println("Build failed")
		return err
	}

	fmt.Println("Build successful")

	return nil
}