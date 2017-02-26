package main

import (
  "fmt"
  "net"
  "io"
  "log"
  "encoding/json"
)

type Config struct {
  configType string
  configApply bool
}

type InputData struct {
  input []byte
}

func (d *InputData) Write(p []byte) (i int, err error) {
  d.input = p
  // TODO: check for errors!
  return 0, nil
}

func main() {
  ln, err := net.Listen("tcp", ":8080")
  if err != nil {
  	fmt.Printf("An error!")
  }
  for {
  	conn, err := ln.Accept()
  	if err != nil {
  		fmt.Printf("An error!")
  	}
  	go handleConnection(conn)
  }
}

func handleConnection(c net.Conn) {
  defer c.Close()

  var data InputData
  var config Config

  io.Copy(&data,c)
  //s := string(data.input)

  // Do something with the input data:
  if err := json.Unmarshal(data.input, &config); err != nil {
		log.Fatal(err)
	}
  fmt.Sprintf(config.configType)
}
