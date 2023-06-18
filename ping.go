package houstn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Houstn) Ping(metadata any) {
	options := GetOptions(h.options)

	path := fmt.Sprintf("%s/%s/%s", options.Project, options.Application, options.Environment)

	body, err := json.Marshal(metadata)

	if err != nil {
		fmt.Printf("Error parsing metadata: %s\n", err)
		return
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", options.Url, path), bytes.NewBuffer(body))

	if err != nil {
		fmt.Printf("Error creating heartbeat request: %s\n", err)
		return
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", options.ApiKey))
	request.Header.Set("X-Houstn-Project", options.Project)
	request.Header.Set("X-Houstn-Environment", options.Environment)
	request.Header.Set("X-Houstn-Application", options.Application)
	request.Header.Set("Content-Type", "application/json")

	_, err = h.client.Do(request)

	if err != nil {
		fmt.Printf("Error sending heartbeat: %s\n", err)
		return
	}
}
