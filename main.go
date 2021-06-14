package main

import (
	"bufio"
	"context"
	"encoding/json"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	in, err := reader.ReadString('\n')
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return
	}
	var event Event
	if err := json.Unmarshal([]byte(in), &event); err != nil {
		os.Stderr.WriteString(err.Error())
		return
	}
	if _, err := run(context.Background(), event); err != nil {
		os.Stderr.WriteString(err.Error())
		return
	}
}
