package sse

import (
	"bufio"
	"encoding/json"
)

func sendEvent(w *bufio.Writer, event string, data interface{}) error {
	dataBytes, _ := json.Marshal(data)

	_, err := w.WriteString("event:" + event + "\ndata:" + string(dataBytes) + "\n\n")
	if err != nil {
		return err // client disconnected
	}

	if err = w.Flush(); err != nil {
		return err // client lost
	}

	return nil
}
