package manager

import "net/http"

func IsServerHealthy(serverURL string) bool {
	response, err := http.Get(serverURL + "/health")
	if err != nil {
		return false
	}
	defer response.Body.Close()

	return response.StatusCode == http.StatusOK
}
