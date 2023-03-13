package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type TagApiResponse struct {
	Guild          string
	MultipleGuilds map[string]string
}

type TagApiError struct {
	InvalidTag string
}

func (e *TagApiError) Error() string {
	return "No guild found with tag `" + e.InvalidTag + "`"
}

func newInvalidTagError(tag string) error {
	return &TagApiError{InvalidTag: tag}
}

func getGuildFromTag(tag string) (TagApiResponse, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("http://avicia.ga/api/tag/?tag=%s", tag), nil)
	if err != nil {
		return TagApiResponse{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return TagApiResponse{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return TagApiResponse{}, err
	}

	var body string = string(bodyBytes)

	// check if we have multiple options
	var res map[string]string
	if err := json.Unmarshal(bodyBytes, &res); err == nil {
		return TagApiResponse{MultipleGuilds: res}, nil
	}

	// if we don't have multiple options, we have a single guild or null
	if body == "\"null\"" {
		return TagApiResponse{}, errors.New(fmt.Sprintf("Invalid tag: %s", tag))
	} else {
		gu := strings.Trim(body, "\"")
		return TagApiResponse{Guild: gu}, nil
	}
}

func main() {
	if len(os.Args) < 2 {
		panic("No tag provided")
	}
	var query = os.Args[1]

	res, err := getGuildFromTag(query)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)

}
