package ssh

import (
  "fmt"
)

type SSHConfig struct {
  sshUsername string
  sshPublickey []byte
}

func Apply(config []byte) error {
  fmt.Println("Applying some SSH changes")
  return nil
}
