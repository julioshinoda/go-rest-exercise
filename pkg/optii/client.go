package optii

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"github.com/julioshinoda/go-rest-exercise/internal/model"
)

type Client struct {
	user   string
	secret string
}

// NewClient creates a Aarin client.
func NewClient(
	u string,
	p string,
) Client {
	return Client{
		user:   u,
		secret: p,
	}
}

func (c Client) GetToken() (string, error) {
	url := "https://test.optii.io/oauth/authorize"

	payload := strings.NewReader("client_secret=oGDXMy5jsD4qEEjfWr1AOW1fuzKKZvQ9y1RcmhbYPNiesfi0haNYO7e9sL6FtlGr&client_id=IXhXcAAfhY4WUIck8Sp1CgavF1zMVjVoihMoIvSuePQN6dpr&grant_type=client_credentials&scope=openapi")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var responseBody AuthResponse

	err := json.Unmarshal(body, &responseBody)
	if err != nil {
		return "",
			fmt.Errorf("failed to parse json response body: %s. body: %q", err, string(body))
	}
	return responseBody.AccessToken, nil
}

func (c Client) GetDepartment(ctx context.Context, p map[string]string) (model.Departments, error) {
	departments := model.Departments{}
	urlA, err := url.Parse("https://test.optii.io/api/v1/departments")
	if err != nil {
		log.Fatal(err)
	}
	values := urlA.Query()

	for i, v := range p {
		values.Set(i, v)
	}

	urlA.RawQuery = values.Encode()

	req, _ := http.NewRequest("GET", urlA.String(), nil)

	req.Header.Add("accept", "application/json")

	token, err := c.GetToken()
	if err != nil {
		return departments, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return departments, err
	}

	switch res.StatusCode {
	case http.StatusOK:
	case http.StatusBadRequest:
		slog.Info(string(body))
		return departments, errors.New("bad request")

	}

	var responseBody Response

	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return model.Departments{},
			fmt.Errorf("failed to parse json response body: %s. body: %q", err, string(body))
	}

	if len(responseBody.Items) > 0 {
		for _, val := range responseBody.Items {
			departments.List = append(departments.List, model.Department{
				ID:   val.ID,
				Name: val.Name,
			})
		}

	}

	return departments, nil
}

func (c Client) GetJobItem(ctx context.Context, p map[string]string) (model.Departments, error) {
	departments := model.Departments{}

	urlA, err := url.Parse("https://test.optii.io/api/v1/jobitems")
	if err != nil {
		log.Fatal(err)
	}
	values := urlA.Query()

	for i, v := range p {
		values.Set(i, v)
	}

	urlA.RawQuery = values.Encode()

	req, _ := http.NewRequest("GET", urlA.String(), nil)

	req.Header.Add("accept", "application/json")

	token, err := c.GetToken()
	if err != nil {
		return departments, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return departments, err
	}

	switch res.StatusCode {
	case http.StatusOK:
	case http.StatusBadRequest:
		slog.Info(string(body))
		return departments, errors.New("bad request")

	}

	var responseBody Response

	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return model.Departments{},
			fmt.Errorf("failed to parse json response body: %s. body: %q", err, string(body))
	}

	if len(responseBody.Items) > 0 {
		for _, val := range responseBody.Items {
			departments.List = append(departments.List, model.Department{
				ID:   val.ID,
				Name: val.Name,
			})
		}

	}

	return departments, nil
}

func (c Client) CreateJob(ctx context.Context, req model.CreateJobRequest) error {
	return nil
}
