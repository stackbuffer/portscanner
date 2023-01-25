package main

import (
	"os"
	"fmt"
	"net"
	"sync"
	"strconv"
)


type PortMap struct {
	ports map[int]string
	mutex sync.Mutex
}


func scanPort(hostname string, port int, openPorts *PortMap){

	defer wg.Done()

	_, err := net.Dial("tcp", hostname + ":" + strconv.Itoa(port))

	if err != nil {
		//fmt.Printf("Port %d is closed.\n", port)
	} else {
		openPorts.mutex.Lock()
		openPorts.ports[port] = "open"
		openPorts.mutex.Unlock()
		//fmt.Printf("Port %d is open.\n", port)
	}
}


var wg sync.WaitGroup

func main(){

	args := os.Args

	hostname := args[1]
	portRange, _ := strconv.Atoi(args[2])

	if portRange > 65536 {
		portRange = 65536
	}

	fmt.Printf("scanning for open ports on %v from range 0 to %d\n", hostname, portRange)



	openPorts := PortMap{ports:make(map[int]string)}

	for i:=0; i<portRange; i++ {
		wg.Add(1)
		go scanPort(hostname, i, &openPorts)
	}

	wg.Wait()

	for key := range openPorts.ports {
		fmt.Printf("port : %d is %v\n", key, openPorts.ports[key])
	}
}