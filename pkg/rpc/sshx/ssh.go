package sshx

import (
	"fmt"
	"os"
	"time"

	"github.com/RyoJerryYu/go-utilx/pkg/utils/loggerx"
	"golang.org/x/crypto/ssh"
)

var log loggerx.Loggerf = loggerx.NoopLoggerf{}

// RegisterLogger sets the logger for the sshx package.
// The logger must implement loggerx.Loggerf interface.
func RegisterLogger(l loggerx.Loggerf) {
	log = l
}

// sshClientConfig holds the configuration for SSH client connection.
// All fields are required to establish a connection.
type sshClientConfig struct {
	port int
	auth []ssh.AuthMethod
}

// SSHClientOption is a function type that modifies sshClientConfig.
// It returns an error if the modification fails.
type SSHClientOption func(*sshClientConfig) error

// WithPort sets the SSH port for the connection.
// Default port is 22 if not specified.
//
// Example:
//
//	opts := WithPort(2222)  // Connect to port 2222 instead of default 22
func WithPort(port int) SSHClientOption {
	return func(c *sshClientConfig) error {
		c.port = port
		return nil
	}
}

// WithPrivateKeyBytes configures SSH authentication using a private key.
// The private key bytes should be in PEM format.
//
// Example:
//
//	keyBytes, _ := os.ReadFile("~/.ssh/id_rsa")
//	opts := WithPrivateKeyBytes(keyBytes)
func WithPrivateKeyBytes(privateKeyBytes []byte) SSHClientOption {
	return func(c *sshClientConfig) error {
		signer, err := ssh.ParsePrivateKey(privateKeyBytes)
		if err != nil {
			return err
		}
		c.auth = []ssh.AuthMethod{ssh.PublicKeys(signer)}
		return nil
	}
}

// WithDefaultAuth configures SSH authentication using the default private key
// from the user's home directory. It first tries to use ~/.ssh/id_rsa,
// and falls back to ~/.ssh/id_ed25519 if id_rsa doesn't exist.
//
// Returns an error if:
// - Unable to get user's home directory
// - Neither id_rsa nor id_ed25519 exists
// - Private key file cannot be read
// - Private key format is invalid
func WithDefaultAuth() SSHClientOption {
	return func(c *sshClientConfig) error {
		dir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		privateKeyFilePath := fmt.Sprintf("%s/.ssh/id_rsa", dir)
		if _, err := os.Stat(privateKeyFilePath); os.IsNotExist(err) {
			privateKeyFilePath = fmt.Sprintf("%s/.ssh/id_ed25519", dir)
		}
		privateKeyBytes, err := os.ReadFile(privateKeyFilePath)
		if err != nil {
			return err
		}
		return WithPrivateKeyBytes(privateKeyBytes)(c)
	}
}

// MakeSSHClient creates a new SSH client with the specified user, host and options.
// If no options are provided, it uses default port 22 and no authentication method.
//
// Example:
//
//	client, err := MakeSSHClient("user", "example.com", WithDefaultAuth(), WithPort(2222))
//
// Note: At least one authentication method must be provided through options,
// otherwise the connection will fail.
func MakeSSHClient(user, host string, opts ...SSHClientOption) (*ssh.Client, error) {
	cfg := &sshClientConfig{
		port: 22,
		auth: []ssh.AuthMethod{},
	}
	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return nil, err
		}
	}
	return makeSSHClientWithCustom(user, host, cfg)
}

// makeSSHClientWithCustom creates an SSH client with the given configuration.
// This is an internal function used by MakeSSHClient.
//
// The connection will:
// - Use InsecureIgnoreHostKey for host verification (not safe for production)
// - Timeout after 1 minute if connection cannot be established
// - Use TCP protocol for connection
//
// Example internal usage:
//
//	cfg := &sshClientConfig{
//	    port: 22,
//	    auth: []ssh.AuthMethod{ssh.Password("password")},
//	}
//	client, err := makeSSHClientWithCustom("user", "host.example.com", cfg)
func makeSSHClientWithCustom(user, host string, cfg *sshClientConfig) (*ssh.Client, error) {
	sshConfig := &ssh.ClientConfig{
		User:            user,
		Auth:            cfg.auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Minute,
	}
	sshcon, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, cfg.port), sshConfig)
	if err != nil {
		return nil, err
	}
	return sshcon, nil
}
