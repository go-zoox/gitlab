package request

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/go-zoox/fetch"
)

var defaulService = ""
var defaultToken = ""
var defaultVersion = "v4"

type Config struct {
	Action   string
	Resource string
	//
	Headers map[string]string
	// Query   map[string]string
	// Params  map[string]string
	// Body    map[string]interface{}
	//
	Service string
	Token   string
	Version string
}

type Payload struct {
	Headers map[string]string
	Query   map[string]string
	Params  map[string]string
	Body    map[string]interface{}
}

func Request(cfg *Config, payload *Payload) (*fetch.Response, error) {
	service := cfg.Service
	token := cfg.Token
	version := cfg.Version

	if service == "" {
		service = os.Getenv("GITLAB_SERVICE")
	}

	if token == "" {
		token = os.Getenv("GITLAB_TOKEN")
	}

	if service == "" {
		service = defaulService
	}
	if token == "" {
		token = defaultToken
	}
	if version == "" {
		version = defaultVersion
	}

	if service == "" {
		return nil, fmt.Errorf("gitlab api service is required")
	}

	if token == "" {
		return nil, fmt.Errorf("gitlab api token is required")
	}

	if version == "" {
		return nil, fmt.Errorf("gitlab api version is required")
	}

	method := cfg.Action
	url := fmt.Sprintf("%s/api/%s/%s", service, version, cfg.Resource)

	client := fetch.New()
	client.SetMethod(method)
	client.SetURL(url)
	client.SetHeader("PRIVATE-TOKEN", token)

	if cfg.Headers != nil {
		for k, v := range cfg.Headers {
			client.SetHeader(k, v)
		}
	}

	// if cfg.Query != nil {
	// 	for k, v := range cfg.Query {
	// 		client.SetQuery(k, v)
	// 	}
	// }

	// if cfg.Params != nil {
	// 	for k, v := range cfg.Params {
	// 		client.SetParam(k, v)
	// 	}
	// }

	// if cfg.Body != nil {
	// 	b, err := json.Marshal(cfg.Body)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	client.SetBody(b)
	// }

	if payload.Headers != nil {
		for k, v := range payload.Headers {
			client.SetHeader(k, v)
		}
	}

	if payload.Query != nil {
		for k, v := range payload.Query {
			client.SetQuery(k, v)
		}
	}

	if payload.Params != nil {
		for k, v := range payload.Params {
			client.SetParam(k, v)
		}
	}

	if payload.Body != nil {
		client.SetBody(payload.Body)
	}

	key := fmt.Sprintf("%s %s", method, url)

	if DEBUG {
		cfg, _ := client.Config()
		cfgBytes, _ := json.MarshalIndent(cfg, "", "  ")
		fmt.Printf("[%s][REQUEST][START]\n", key)
		fmt.Println(string(cfgBytes))
		fmt.Printf("[%s][REQUEST][END]\n", key)
		fmt.Println("")
	}

	response, err := client.Send()
	if err != nil {
		return nil, err
	}

	if DEBUG {
		fmt.Printf("[%s][RESPONSE][START]\n", key)
		fmt.Println("status:", response.Status)
		fmt.Println("headers:", response.Headers)
		fmt.Println("body:", response.String())
		fmt.Printf("[%s][RESPONSE][END]\n", key)
		fmt.Println("")
	}

	if !response.Ok() {
		if DEBUG {
			fmt.Println("response error:", response.Get("Response.Error").String())
		}

		Code := response.Get("code").String()
		Message := response.Get("message").String()
		if Code == "" {
			Code = strconv.Itoa(response.Status)
		}
		if Message == "" {
			Message = "no detail message"
		}

		return nil, &ResponseError{Code, Message}
	}

	return response, nil
}

func SetService(service string) {
	defaulService = service
}

func SetToken(token string) {
	defaultToken = token
}

func SetVersion(version string) {
	defaultVersion = version
}
