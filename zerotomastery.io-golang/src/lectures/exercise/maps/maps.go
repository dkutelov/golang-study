//--Summary:
//  Write a program to display server status.
//
//--Requirements:
//* Create a function to print server status, including:
//  - Number of servers
//  - Number of servers for each status (Online, Offline, Maintenance, Retired)
//* Store the existing slice of servers in a map
//* Default all of the servers to `Online`
//* Perform the following status changes and display server info:
//  - display server info
//  - change `darkstar` to `Retired`
//  - change `aiur` to `Offline`
//  - display server info
//  - change all servers to `Maintenance`
//  - display server info

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func printServerStatus(serverList map[string]int) {
	serversCount := len(serverList)

	serverStatus := map[string]int {
			"Online "     : 0,
			"Offline"     : 0,
			"Maintenance" : 0,
			"Retired"     : 0,
			}
	
	for _, status := range serverList {

		serverStatus[] += 1
	}


	fmt.Println("We have", serversCount, "servers.")
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serverList := make(map[string]int)

	for _, server := range servers {
		serverList[server] = Online
	}

	fmt.Println(serverList)
}
