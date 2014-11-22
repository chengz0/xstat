// main.go
package main

import (
	"github.com/chengz0/xstat/collector"
	"log"
)

func main() {
	// disk, _ := collector.GetCurrentDisk()
	// log.Println(disk.Total / 1024 / 1024)

	log.Println("=========")
	a, _ := collector.LastLogin()
	log.Println(len(a))
	log.Println(a[0])
	log.Println(a[0].Username, " ", a[0].Latest)
}
