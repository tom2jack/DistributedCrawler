package main

import (
	"fmt"
	"net/http"
	"log"
)

type StatisticsServer struct {
	port     string
}

func InitServer(port string) *StatisticsServer {
  s = &StatisticsServer{
		port:    port,
	}
  return s
}

func (s *StatisticsServer) Start() {
	go func () {
		log.Printf("STARTING REDISMQ SERVER ON PORT %s", s.port)
		err := http.ListenAndServe(":"+s.port, nil)
		if err != nil {
			log.Fatalf("REDISMQ SERVER SHUTTING DOWN [%s]\n\n", err.Error())
		}
	}()
}

// Can be run in 3 ways:
// 1) Sequential (e.g., go run wc.go master x.txt sequential)
// 2) Master (e.g., go run wc.go master x.txt localhost:7777)
// 3) Worker (e.g., go run wc.go worker localhost:7777 localhost:7778 &)
func main() {
  if len(os.Args) != 4 {
    fmt.Printf("%s: see usage comments in file\n", os.Args[0])
  } else if os.Args[1] == "master" {
    if os.Args[3] == "sequential" {
      redismq.RunSingle(5, 3, os.Args[2], Map, Reduce)
    } else {
      mr := redismq.Makeredismq(5, 3, os.Args[2], os.Args[3])
      // Wait until MR is done
      <- mr.DoneChannel
    }
  } else {
    redismq.RunWorker(os.Args[2], os.Args[3], Map, Reduce, 100)
  }
}