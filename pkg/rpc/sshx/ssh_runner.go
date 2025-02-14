package sshx

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"strings"

	"golang.org/x/crypto/ssh"
)

// SshRunner provides methods to execute commands and transfer files over an SSH connection.
// It requires an established SSH client connection to function.
// The connection must remain valid throughout the lifetime of the SshRunner.
type SshRunner struct {
	sshClient *ssh.Client
}

// NewSshRunner creates a new SshRuner instance with the provided SSH client.
//
// Parameters:
//   - user: SSH username
//   - host: SSH host
//   - opts: Optional SSH client options
//
// Returns an error if the SSH client creation fails.
func NewSshRunner(user, host string, opts ...SSHClientOption) (*SshRunner, error) {
	sshClient, err := MakeSSHClient(user, host, opts...)
	if err != nil {
		return nil, err
	}
	return &SshRunner{
		sshClient: sshClient,
	}, nil
}

// NewSshRunnerWithClient creates a new SshRunner using an existing SSH client.
// This is useful when you want to reuse an existing connection or need
// custom SSH client configuration.
//
// The provided SSH client must:
// - Be already connected and authenticated
// - Remain valid for the lifetime of the SshRunner
// - Not be closed while the SshRunner is in use
//
// Example:
//
//	client, _ := ssh.Dial("tcp", "host:22", &ssh.ClientConfig{...})
//	runner := NewSshRunnerWithClient(client)
func NewSshRunnerWithClient(sshClient *ssh.Client) *SshRunner {
	return &SshRunner{
		sshClient: sshClient,
	}
}

// Run executes a command over SSH and returns its combined output (stdout + stderr) as a string.
// The session is automatically closed after the command completes.
//
// Example:
//
//	output, err := runner.Run("ls -la")
//	// output contains both stdout and stderr from the 'ls -la' command
func (s *SshRunner) Run(ctx context.Context, cmd string) (string, error) {
	session, err := s.sshClient.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	stderr := bytes.Buffer{}
	session.Stderr = &stderr

	output, err := session.CombinedOutput(cmd)
	if stderr.Len() > 0 {
		log.Warnf(ctx, "run with stderr: %s", stderr.String())
	}
	return string(output), err
}

// RunLog executes a command over SSH and writes its output to the provided writers.
// stdout and stderr are written to their respective writers.
// The session is automatically closed after the command completes.
//
// Example:
//
//	var stdout, stderr bytes.Buffer
//	err := runner.RunLog("ls -la", &stdout, &stderr)
//	// stdout contains standard output
//	// stderr contains error output
func (s *SshRunner) RunLog(ctx context.Context, cmd string, stdOut, stdErr io.Writer) error {
	session, err := s.sshClient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	session.Stdout = stdOut
	session.Stderr = stdErr

	return session.Run(cmd)
}

var (
	// ErrUpdateScriptFailed is returned when a script upload operation
	// completes but the verification step fails.
	// This usually indicates that the remote server did not receive
	// the complete file or had permission issues.
	ErrUpdateScriptFailed = errors.New("update script failed")
)

// success is the expected string response from the remote server
// indicating a successful file operation.
const success = "success"

// UpdateScript uploads a script to the remote server, makes it executable,
// and verifies the upload was successful.
//
// Parameters:
//   - ctx: Context for logging
//   - scriptPath: Remote path where the script should be saved
//   - scriptData: Content of the script
//
// The function will return ErrUpdateScriptFailed if the upload verification fails.
// The uploaded script will be made executable (chmod +x) automatically.
//
// Example:
//
//	scriptData := []byte("#!/bin/bash\necho 'Hello World'")
//	err := runner.UpdateScript(ctx, "/tmp/hello.sh", scriptData)
func (s *SshRunner) UpdateScript(ctx context.Context, scriptPath string, scriptData []byte) error {
	session, err := s.sshClient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	stderr := bytes.Buffer{}
	session.Stdin = bytes.NewBuffer(scriptData)
	session.Stderr = &stderr

	cmd := fmt.Sprintf("cat > %s && chmod +x %s && echo %s", scriptPath, scriptPath, success)
	output, err := session.Output(cmd)
	if stderr.Len() > 0 {
		log.Warnf(ctx, "update script with stderr: %s", stderr.String())
	}
	if err != nil {
		log.Errorf(ctx, "update script failed: %v", err)
		return err
	}

	if !strings.HasPrefix(string(output), success) {
		log.Infof(ctx, "ssh update script not success, output: %s", string(output))
		return ErrUpdateScriptFailed
	}

	return nil
}

// UploadFile uploads a file to the remote server and verifies the upload was successful.
//
// Parameters:
//   - ctx: Context for logging
//   - filePath: Remote path where the file should be saved
//   - fileData: Content of the file
//
// Returns ErrUpdateScriptFailed if the upload verification fails.
//
// Example:
//
//	fileData := []byte("Hello World")
//	err := runner.UploadFile(ctx, "/tmp/hello.txt", fileData)
func (s *SshRunner) UploadFile(ctx context.Context, filePath string, fileData []byte) error {
	session, err := s.sshClient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	stderr := bytes.Buffer{}
	session.Stdin = bytes.NewBuffer(fileData)
	session.Stderr = &stderr

	output, err := session.Output(fmt.Sprintf("cat > %s && echo %s", filePath, success))
	if stderr.Len() > 0 {
		log.Warnf(ctx, "update script with stderr: %s", stderr.String())
	}
	if err != nil {
		log.Errorf(ctx, "upload file failed: %v", err)
		return err
	}

	if !strings.HasPrefix(string(output), success) {
		log.Infof(ctx, "ssh upload file not success, output: %s", string(output))
		return ErrUpdateScriptFailed
	}

	return nil
}

// DownloadFile retrieves a file from the remote server.
//
// Parameters:
//   - ctx: Context for logging
//   - filePath: Remote path of the file to download
//
// Returns the file contents as bytes. If the file cannot be read or doesn't exist,
// returns an empty byte array and an error.
//
// Example:
//
//	data, err := runner.DownloadFile(ctx, "/etc/hosts")
//	// data contains the contents of /etc/hosts if successful
func (s *SshRunner) DownloadFile(ctx context.Context, filePath string) ([]byte, error) {
	session, err := s.sshClient.NewSession()
	if err != nil {
		return []byte(""), err
	}
	defer session.Close()

	stderr := bytes.Buffer{}
	session.Stderr = &stderr

	output, err := session.Output(fmt.Sprintf("cat %s", filePath))
	if stderr.Len() > 0 {
		log.Warnf(ctx, "download file with stderr: %s", stderr.String())
	}
	if err != nil {
		log.Errorf(ctx, "download file failed: %v", err)
		return []byte(""), err
	}

	return output, nil
}
