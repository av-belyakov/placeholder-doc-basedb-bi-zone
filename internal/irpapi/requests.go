package irpapi

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

// GetIncident запрос основной информации по инциденту
func (api *IrpApi) GetIncident(ctx context.Context, incidentId string) ([]byte, int, error) {
	return api.query(ctx, "issue/"+incidentId, nil, "GET")
}

// query функция реализующая непосредственно сам HTTP запрос
func (api *IrpApi) query(ctx context.Context, reqPath string, query []byte, method string) ([]byte, int, error) {
	var (
		url string = fmt.Sprintf("https://%s:%d/api/v0/irp/%s", api.settings.host, api.settings.port, reqPath)
		res *http.Response

		err error
	)

	ctx, cancel := context.WithTimeout(ctx, time.Duration(api.settings.connectionTimeout)*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(query))
	if err != nil {
		return nil, 0, err
	}

	req.Header.Add("Authorization", "Token "+api.settings.apiKey)
	req.Header.Set("Content-Type", "application/json")

	res, err = api.client.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("attempts to establish a network connection - %w", err)
	}
	defer supportingfunctions.CloseBodyHTTPResponse(res)

	if res.Body != nil {
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, res.StatusCode, err
		}

		return resBody, res.StatusCode, nil
	}

	return nil, res.StatusCode, nil
}
