package ssh

import (
  "fmt"
)

type SSHConfig struct {
  Username string
  Publickey string
}

func Apply(config SSHConfig) error {
  fmt.Print("Applying some SSH changes for username: ")
  fmt.Println(config.Username)

  return nil
}
