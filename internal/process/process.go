package process

import (
	"fmt"
	"os"
	"os/exec"
)

type Server struct {
	Cmd *exec.Cmd
}

func (s *Server) Start(command string) error {

	fmt.Println("Starting server:", command)

	cmd := exec.Command("cmd", "/C", command)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting server:", err)
		return err
	}

	s.Cmd = cmd

	go cmd.Wait()

	return nil
}

func (s *Server) Stop() {

	if s.Cmd == nil {
		return
	}

	fmt.Println("Stopping server")

	exec.Command("taskkill", "/T", "/F", "/PID", fmt.Sprint(s.Cmd.Process.Pid)).Run()

	s.Cmd = nil
}