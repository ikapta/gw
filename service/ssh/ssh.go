package ssh

import (
	"bytes"
	"time"

	"golang.org/x/crypto/ssh"
)

type SSHConf struct {
  NAME string
  HOST string
  PASS string
}

func NewClient(conf SSHConf) (*ssh.Client, error) {
  var (
    auth         []ssh.AuthMethod
    addr         string
    clientConfig *ssh.ClientConfig
    client       *ssh.Client
    err          error
  )

  auth = make([]ssh.AuthMethod, 0)
  auth = append(auth, ssh.Password(conf.PASS))

  clientConfig = &ssh.ClientConfig{
    User: conf.NAME,
    Auth: auth,
    Timeout: 10 * time.Second,
    HostKeyCallback: ssh.InsecureIgnoreHostKey(),
  }

  addr = conf.HOST

  if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
    return nil, err
  }

  return client, nil
}

func SessionRun(client *ssh.Client, command string) (string, error) {
  var (
    session *ssh.Session
    buf     bytes.Buffer
    err     error
  )

  if session, err = client.NewSession(); err != nil {
    return "", err
  }

  session.Stdout = &buf
  defer session.Close()

  if err = session.Run(command); err != nil {
    return "", err
  }

  return buf.String(), nil
}

func Once(cfg SSHConf, command string) (string, error) {
  var (
    client *ssh.Client
    err    error
  )

  if client, err = NewClient(cfg); err != nil {
    return "", err
  }
  defer client.Close()

  return SessionRun(client, command)
}