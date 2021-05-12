package main

import (
  "encoding/json"
  "fmt"
  "net"
  "os"
  "sync"

  "github.com/zant/json-tcp/common"
)

func main() {
  var wg sync.WaitGroup
  message := common.Message{
    Channel: "Random",
    Body:    "Heyyyy",
  }

  if len(os.Args) != 2 {
    fmt.Println("Usage :", os.Args[0], "host:port")
    os.Exit(1)
  }
  service := os.Args[1]

  conn, err := net.Dial("tcp", service)
  common.CheckError(err)

  go sendAndPrint(message, conn, &wg)

  fmt.Printf("Immediatly\n")

  wg.Wait()
  os.Exit(0)
}

func sendAndPrint(m common.Message, conn net.Conn, wg *sync.WaitGroup) {
  wg.Add(1)
  encoder := json.NewEncoder(conn)
  decoder := json.NewDecoder(conn)

  encoder.Encode(m)
  var newMessage common.Message
  decoder.Decode(&newMessage)
  fmt.Println(newMessage.String())
  defer wg.Done()
}
