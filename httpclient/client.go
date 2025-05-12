package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strconv"
	"time"

	"github.com/fatih/color"
)

func CreateRequest(method, rootURL, urlPath string, args []string) (*http.Request, error) {
	fullURL, err := joinURL(rootURL, urlPath)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	// Parse query parameters
	queryParams := parseArgsToQueryParams(args)

	// Add query parameters to URL
	if len(queryParams) > 0 {
		q := fullURL.Query()
		for key, value := range queryParams {
			q.Add(key, value)
		}
		fullURL.RawQuery = q.Encode()
	}

	var body io.Reader
	if method == "POST" || method == "PUT" || method == "PATCH" {
		jsonBody := parseArgsToJSON(args)
		bodyData, err := json.Marshal(jsonBody)
		if err != nil {
			return nil, fmt.Errorf("error marshalling JSON body: %w", err)
		}
		body = bytes.NewBuffer(bodyData)
	}

	req, err := http.NewRequest(method, fullURL.String(), body)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	headers := parseArgsToHeaders(args)

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return req, nil
}

func ExecuteRequest(req *http.Request, outputStatus bool) error {
	client := &http.Client{
		Timeout: 10 * time.Minute,
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if outputStatus {
		colorizeStatus(resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)

	contentType := resp.Header.Get("Content-Type")
	if contentType == "application/json" {
		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, body, "", "  "); err != nil {
			fmt.Println("Error formatting JSON response.")
		} else {
			fmt.Println(prettyJSON.String())
		}
	} else {
		fmt.Println(string(body))
	}
	return nil
}

func joinURL(root, urlPath string) (*url.URL, error) {
	baseURL, err := url.Parse(root)
	if err != nil {
		return nil, err
	}
	baseURL.Path = path.Join(baseURL.Path, urlPath)
	return baseURL, nil
}

func parseArgsToJSON(args []string) map[string]interface{} {
	jsonBody := make(map[string]interface{})
	jsonBodyRe := regexp.MustCompile(`^([a-zA-Z0-9_-]+)=(.*)$`)

	for _, arg := range args {
		if matches := jsonBodyRe.FindStringSubmatch(arg); matches != nil {
			key, value := matches[1], matches[2]
			jsonBody[key] = parseValue(value)
		}
	}
	return jsonBody
}

func parseArgsToQueryParams(args []string) map[string]string {
	queryParams := make(map[string]string)
	queryParamRe := regexp.MustCompile(`^([a-zA-Z0-9_-]+)\|(.*)$`)

	for _, arg := range args {
		if matches := queryParamRe.FindStringSubmatch(arg); matches != nil {
			key, value := matches[1], matches[2]
			queryParams[key] = value
		}
	}
	return queryParams
}

func parseArgsToHeaders(args []string) map[string]string {
	headers := make(map[string]string)
	headerRe := regexp.MustCompile(`^([a-zA-Z0-9_-]+):(.*)$`)

	for _, arg := range args {
		if matches := headerRe.FindStringSubmatch(arg); matches != nil {
			key, value := matches[1], matches[2]
			headers[key] = value
		}
	}
	return headers
}

func parseValue(value string) interface{} {
	if value == "true" || value == "false" {
		return value == "true"
	}
	if i, err := strconv.Atoi(value); err == nil {
		return i
	}
	if f, err := strconv.ParseFloat(value, 64); err == nil {
		return f
	}
	return value
}

func colorizeStatus(code int) {
	switch {
	case code >= 200 && code < 300:
		color.New(color.FgGreen).Printf("Response Status: %d\n", code)
	case code >= 400 && code < 500:
		color.New(color.FgYellow).Printf("Response Status: %d\n", code)
	case code >= 500:
		color.New(color.FgRed).Printf("Response Status: %d\n", code)
	default:
		fmt.Printf("Response Status: %d\n", code)
	}
}
