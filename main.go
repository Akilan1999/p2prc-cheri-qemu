package main

import (
    "github.com/Akilan1999/p2p-rendering-computation/p2p/frp"
    "time"
    "fmt"
)

// Trigger the FRP implementation
func main() {
    natssh, err := EscapeNATSSH()
    if err != nil {
        return
    }

    fmt.Println("ssh root@64.227.168.102 -p" + natssh)

    for {

    }
}

func EscapeNATSSH() (SSHport string, err error) {
    // Get free port from P2PRC server node
    serverPort, err := frp.GetFRPServerPort("http://64.227.168.102:8088")

    if err != nil {
        return
    }

    time.Sleep(1 * time.Second)

    //port for the barrierKVM server
    SSHport, err = frp.StartFRPClientForServer("64.227.168.102", serverPort, "10005")
    if err != nil {
        return
    }

    return
}
