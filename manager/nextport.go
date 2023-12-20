package manager

import (
	"strconv"
	"sync"
)

var (
	port         = ":6666"
	serverPorts  = []int{6667, 6668, 6669}
	currentIndex = 0
	mu           sync.Mutex
	requestCount int
)

func GetNextPort() string {
	mu.Lock()
	defer mu.Unlock()

	currentIndex++
	if currentIndex >= len(serverPorts) {
		currentIndex = 0
	}

	return strconv.Itoa(serverPorts[currentIndex])
}
