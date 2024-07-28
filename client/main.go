package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	reqBody, _ := json.Marshal(map[string]string{"key1": "val1", "key2": "val2"})

	resp, _ := http.Post(":8091", "application/json", bytes.NewReader(reqBody))
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Printf("resp: %s", respBody)
}
