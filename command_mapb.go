package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func commandMapb(c *config) error {
	if c.previous == "" {
		return errors.New("you're on the first page")
	}

	url := c.previous

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error while creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error while reading json body: %w", err)
	}

	var locationData LocationData
	if err := json.Unmarshal(data, &locationData); err != nil {
		return fmt.Errorf("error reading from unmarshaling: %w", err)
	}

	for _, location := range locationData.Results {
		fmt.Println(location.Name)
	}

	c.next = locationData.Next
	if locationData.Previous == nil {
		c.previous = ""
	} else {
		c.previous = *locationData.Previous
	}

	return nil
}
