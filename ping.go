package houstn

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Houstn) Ping(metadata any) {
	username := fmt.Sprintf("%s+%s+%s", h.options.Organisation, h.options.Application, h.options.Environment)
	auth := fmt.Sprintf("%s:%s", username, h.options.Token)

	authorization := base64.StdEncoding.EncodeToString([]byte(auth))

	body, err := json.Marshal(metadata)

	if err != nil {
		fmt.Printf("Error parsing metadata: %s\n", err)
		return
	}

	request, err := http.NewRequest("POST", h.options.Url, bytes.NewBuffer(body))

	if err != nil {
		fmt.Printf("Error creating heartbeat request: %s\n", err)
		return
	}

	request.Header.Set("Authorization", fmt.Sprintf("Basic %s", authorization))
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
