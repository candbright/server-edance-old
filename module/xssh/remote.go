package xssh

import (
	"bytes"
	"edance"
	"fmt"
	"github.com/candbright/gin-util/xlog"
	"golang.org/x/crypto/ssh"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type Remote struct {
	IsLocal   bool
	SshClient *ssh.Client
	Config    *SessionConfig
	Mutex     *sync.Mutex
}

func NewRemote(config *SessionConfig) (*Remote, error) {
	remote := &Remote{
		Mutex:  &sync.Mutex{},
		Config: config,
	}
	if config.host() == "127.0.0.1" || config.host() == config.localIp() {
		remote.IsLocal = true
		return remote, nil
	}
	sshKeyPath := "/root/.ssh/id_rsa"
	var auth []ssh.AuthMethod
	if config.password() != "" {
		auth = []ssh.AuthMethod{ssh.Password(config.password())}
	} else {
		auth = []ssh.AuthMethod{publicKeyAuth(sshKeyPath)}
	}
	sshConfig := &ssh.ClientConfig{
		Timeout:         time.Second * 10,
		User:            config.user(),
		Auth:            auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	addr := fmt.Sprintf("%s:%d", config.host(), config.port())
	sshClient, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		return nil, xlog.Wrap(err, addr)
	}
	remote.SshClient = sshClient
	return remote, nil
}

func publicKeyAuth(keyPath string) ssh.AuthMethod {
	key, err := os.ReadFile(keyPath)
	if err != nil {
		xlog.Error("ssh key file read failed, error:%s", err)
		return nil
	}
	singer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		xlog.Error("parse ssh private key failed, error:%s", err)
		return nil
	}
	return ssh.PublicKeys(singer)
}

type SessionConfig struct {
	Host     string
	User     string
	Password string
	Port     int
	LocalIp  string
}

func (cfg *SessionConfig) host() string {
	if cfg.Host == "" {
		cfg.Host = "127.0.0.1"
	}
	return cfg.Host
}

func (cfg *SessionConfig) localIp() string {
	if cfg.LocalIp == "" {
		cfg.LocalIp = "127.0.0.1"
	}
	return cfg.LocalIp
}

func (cfg *SessionConfig) user() string {
	if cfg.User == "" {
		cfg.User = "root"
	}
	return cfg.User
}
func (cfg *SessionConfig) password() string {
	return cfg.Password
}
func (cfg *SessionConfig) port() int {
	if cfg.Port == 0 {
		cfg.Port = 22
	}
	return cfg.Port
}

type FileInfo struct {
	Name         string
	AbsolutePath string
}

func (f *FileInfo) IsDir() bool {
	return strings.HasSuffix(f.Name, "/")
}

func (s *Remote) Run(name string, arg ...string) error {
	cmd := name
	if arg != nil {
		cmd += " " + strings.Join(arg, " ")
	}
	if s.IsLocal {
		return xlog.Wrap(exec.Command(name, arg...).Run(), cmd)
	}
	if s.SshClient == nil {
		return xlog.Wrap(edance.ErrNilSshClient)
	}
	session, err := s.SshClient.NewSession()
	if err != nil {
		return xlog.Wrap(err)
	}
	defer session.Close()
	var stderr bytes.Buffer
	session.Stderr = &stderr
	err = session.Run(cmd)
	if err != nil {
		return xlog.Wrap(err, cmd)
	}
	return nil
}
