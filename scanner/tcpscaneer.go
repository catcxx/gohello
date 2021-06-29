package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func isOpen(host string, port int, timeout time.Duration) bool {
	time.Sleep(time.Millisecond * 1)
	//fmt.Println("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err == nil {
		_ = conn.Close()
		return true
	}

	return false
}

func main() {
	openPorts := make(map[string][]int)
	//ips := []string{"10.15.1.48", "127.0.0.1","10.15.1.32","172.20.152.206"}
	ips := []string{ "127.0.0.1","172.20.152.206"}


	timeout := time.Millisecond * 1000
	for i:= range ips {
		fmt.Println(i)
		wg := &sync.WaitGroup{}
		mutex := &sync.Mutex{}
		for port := 1; port < 6500; port++ {
			wg.Add(1)
			go func(p int) {
				opened := isOpen(ips[i], p, timeout)
				if opened {
					mutex.Lock()
					openPorts[ips[i]] = append(openPorts[ips[i]], p)
					mutex.Unlock()
				}
				wg.Done()
			}(port)
		}
		wg.Wait()
	}



	//fmt.Printf("opened ports: %s %v\n", openPorts)
	for i := range openPorts {
		fmt.Println(i, "openports:", openPorts[i])
	}

}