package main

import (
  "fmt"
  "net"
  "io"
)

type InputData struct {
  input []byte
}

func (d *InputData) Write(p []byte) (i int, err error) {
  addendum := []byte(" has been added to inputdata\n")
  d.input = p
  d.input = append(d.input, addendum...)
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
  io.Copy(&data,c)
  s := string(data.input)
}
