package client

import (
	"encoding/json"
	"fmt"
	neturl "net/url"
	"strconv"

	"github.com/beeper/hrobot-go/models"
)

func (c *Client) ServerGetList() ([]models.Server, error) {
	url := c.baseURL + "/server"
	bytes, err := c.doGetRequest(url)
	if err != nil {
		return nil, err
	}

	var servers []models.ServerResponse
	err = json.Unmarshal(bytes, &servers)
	if err != nil {
		return nil, err
	}

	var data []models.Server
	for _, server := range servers {
		data = append(data, server.Server)
	}

	return data, nil
}

func (c *Client) ServerGet(ip string) (*models.Server, error) {
	url := fmt.Sprintf(c.baseURL+"/server/%s", ip)
	bytes, err := c.doGetRequest(url)
	if err != nil {
		return nil, err
	}

	var serverResp models.ServerResponse
	err = json.Unmarshal(bytes, &serverResp)
	if err != nil {
		return nil, err
	}

	return &serverResp.Server, nil
}

func (c *Client) ServerGetCancellation(ip string) (*models.Cancellation, error) {
	url := fmt.Sprintf(c.baseURL+"/server/%s/cancellation", ip)

	bytes, err := c.doGetRequest(url)
	if err != nil {
		return nil, err
	}

	var cancelResp models.CancellationResponse
	err = json.Unmarshal(bytes, &cancelResp)
	if err != nil {
		return nil, err
	}
	return &cancelResp.Cancellation, nil
}

func (c *Client) ServerSetCancellation(ip string, input *models.CancellationSetInput) (*models.Cancellation, error) {
	url := fmt.Sprintf(c.baseURL+"/server/%s/cancellation", ip)

	formData := neturl.Values{}
	if input.CancellationReason != "" {
		formData.Set("cancellation_reason", input.CancellationReason)
	}
	formData.Set("cancellation_date", input.CancellationDate)
	formData.Set("reserve_location", strconv.FormatBool(input.ReserveLocation))

	bytes, err := c.doPostFormRequest(url, formData)
	if err != nil {
		return nil, err
	}

	var cancelResp models.CancellationResponse
	err = json.Unmarshal(bytes, &cancelResp)
	if err != nil {
		return nil, err
	}

	return &cancelResp.Cancellation, nil
}

func (c *Client) ServerSetName(ip string, input *models.ServerSetNameInput) (*models.Server, error) {
	url := fmt.Sprintf(c.baseURL+"/server/%s", ip)

	formData := neturl.Values{}
	formData.Set("server_name", input.Name)

	bytes, err := c.doPostFormRequest(url, formData)
	if err != nil {
		return nil, err
	}

	var serverResp models.ServerResponse
	err = json.Unmarshal(bytes, &serverResp)
	if err != nil {
		return nil, err
	}

	return &serverResp.Server, nil
}

func (c *Client) ServerReverse(ip string) (*models.Cancellation, error) {
	url := fmt.Sprintf(c.baseURL+"/server/%s/reversal", ip)

	bytes, err := c.doPostFormRequest(url, nil)
	if err != nil {
		return nil, err
	}

	var cancelResp models.CancellationResponse
	err = json.Unmarshal(bytes, &cancelResp)
	if err != nil {
		return nil, err
	}

	return &cancelResp.Cancellation, nil
}
