package irpapi

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// query функция реализующая непосредственно сам HTTP запрос
func (api *IrpApi) query(ctx context.Context, reqpath string, query []byte, method string) ([]byte, int, error) {
	var (
		res *http.Response
		err error
	)

	ctx, cancel := context.WithTimeout(ctx, time.Duration(c.settings.timeout)*time.Second)
	defer cancel()

	url := fmt.Sprintf("http://%s:%d%s", api.settings.host, api.settings.port, reqpath)
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(query))
	if err != nil {
		return nil, 0, err
	}

	req.Header.Add("Authorization", "Bearer "+api.settings.apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err = client.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("attempts to establish a network connection - %w", err)
	}
	defer func(body io.ReadCloser, err error) {
		if errClose := body.Close(); errClose != nil {
			errors.Join(err, errClose)
		}
	}(res.Body, err)

	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusCreated {
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, res.StatusCode, err
		}

		return resBody, res.StatusCode, nil
	}

	var msg string
	if m, err := supportingfunctions.GetDetailedMessage(res.Body); err == nil {
		msg = m
	}

	return nil, res.StatusCode, fmt.Errorf("error request, status is '%s' %s", res.Status, msg)
}
