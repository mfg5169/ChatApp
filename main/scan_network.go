package main

import (
	"fmt"
	"os/exec"
	"sync"
)

func pingIP(ip string, wg *sync.WaitGroup, resultChan chan<- string) {
	defer wg.Done()
	cmd := exec.Command("ping", "-c", "1", "-W", "1", ip)
	err := cmd.Run()

	if err == nil {
		resultChan <- ip
	}
}

func main() {

	// Change this to your network's subnet, 
	subnet := "10.0.0."

	// Use a wait group to wait for all pinging goroutines to finish
	var wg sync.WaitGroup

	// Create a channel to receive results
	resultChan := make(chan string, 100)

	// Iterate over all IP addresses in the subnet 
	for i := 1; i <= 254; i++ {
		ip := fmt.Sprintf("%s%d", subnet, i)

		wg.Add(1)
		go pingIP(ip, &wg, resultChan) // Ping each IP 
	}

	// Wait for all pings to complete
	go func() {
		wg.Wait()
		close(resultChan) // Close the channel after all pings are done
	}()

	// Collect and display the results
	fmt.Println("Active IP addresses on the network:")
	for ip := range resultChan {
		fmt.Println(ip)
	}
}
