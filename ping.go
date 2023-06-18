package houstn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Houstn) Ping(metadata any) {
	path := fmt.Sprintf("%s/%s/%s", h.options.Project, h.options.Application, h.options.Environment)

	body, err := json.Marshal(metadata)

	if err != nil {
		fmt.Printf("Error parsing metadata: %s\n", err)
		return
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", h.options.Url, path), bytes.NewBuffer(body))

	if err != nil {
		fmt.Printf("Error creating heartbeat request: %s\n", err)
		return
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", h.options.ApiKey))
	request.Header.Set("X-Houstn-Deployment", h.options.Deployment)
	request.Header.Set("X-Houstn-Environment", h.options.Environment)
	request.Header.Set("X-Houstn-Application", h.options.Application)
	request.Header.Set("Content-Type", "application/json")

	_, err = h.client.Do(request)

	if err != nil {
		fmt.Printf("Error sending heartbeat: %s\n", err)
		return
	}
}
