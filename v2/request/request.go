package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/poroping/forti-sdk-go/v2/config"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func CreateUpdate(c config.Config, r *models.CmdbRequest) (*models.CmdbResponse, error) {
	req, err := newRequest(c, r)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPCon.Do(req)
	if err != nil {
		// handle 404s (search for resource on delete at least)
		return nil, err
	}

	log.Printf("[DEBUG] Status code: %d", res.StatusCode)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &models.CmdbResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("[ERROR] Error reading response body during CREATE/UPDATE %s", err)
		return nil, err
	}

	// if response != nil {
	// 	if response.HTTPStatus == int64(500) && r.Mkey != nil {
	// 		if r.Params.AllowAppend != nil {
	// 			b := *r.Params.AllowAppend
	// 			if b {
	// 				response2, err := Read(c, r)
	// 				if err != nil {
	// 					return nil, err
	// 				}
	// 				return response2, nil
	// 			}
	// 		}
	// 	}
	// }

	err = fortiErrorCheck(body, response)
	if err != nil {
		return nil, err
	}

	return response, err
}

func Read(c config.Config, r *models.CmdbRequest) (*models.CmdbResponse, error) {
	req, err := newRequest(c, r)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPCon.Do(req)
	if err != nil {
		// handle 404s (search for resource on delete at least)
		return nil, err
	}

	log.Printf("[DEBUG] Status code: %d", res.StatusCode)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &models.CmdbResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		log.Printf("[ERROR] Error reading response body during READ")
		return nil, err
	}

	err = fortiErrorCheck(body, response)
	if err != nil {
		return nil, err
	}

	return response, err
}

func Delete(c config.Config, r *models.CmdbRequest) (err error) {
	req, err := newRequest(c, r)
	if err != nil {
		return err
	}

	res, err := c.HTTPCon.Do(req)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Status code: %d", res.StatusCode)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	response := models.CmdbResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("[ERROR] Error reading response body during DELETE")
		return err
	}

	err = fortiErrorCheck(body, &response)
	if err != nil {
		if res.StatusCode == 404 {
			log.Printf("[WARN] Resource not found")
			return nil
		}
		return err
	}

	return nil
}

func fortiErrorCheck(body []byte, res *models.CmdbResponse) (err error) {
	if res.Status == "success" {
		return nil
	}

	switch hs := res.HTTPStatus; hs {
	case int64(400):
		err = fmt.Errorf("bad Request - Request cannot be processed by the API (%d)", res.HTTPStatus)
	case int64(401):
		err = fmt.Errorf("not Authorized - Request without successful login session (%d)", res.HTTPStatus)
	case int64(403):
		err = fmt.Errorf("forbidden - Request is missing CSRF token or administrator is missing access profile permissions (%d)", res.HTTPStatus)
	case int64(404):
		err = fmt.Errorf("resource Not Found - Unable to find the specified resource (%d)", res.HTTPStatus)
	case int64(405):
		err = fmt.Errorf("method Not Allowed - Specified HTTP method is not allowed for this resource (%d)", res.HTTPStatus)
	case int64(413):
		err = fmt.Errorf("request Entity Too Large - Request cannot be processed due to large entity (%d)", res.HTTPStatus)
	case int64(424):
		err = fmt.Errorf("failed Dependency - Fail dependency can be duplicate resource, missing required parameter, missing required attribute, invalid attribute value (%d)", res.HTTPStatus)
	case int64(429):
		err = fmt.Errorf("access temporarily blocked - Maximum failed authentications reached. The offended source is temporarily blocked for certain amount of time (%d)", res.HTTPStatus)
	case int64(500):
		return parseError500(body)
	default:
		err = fmt.Errorf("unknown Error (%d)", res.HTTPStatus)
	}

	return
}

func parseError500(body []byte) error {
	fortiError := &models.CmdbError500{}
	err := json.Unmarshal(body, &fortiError)
	if err != nil {
		log.Printf("[ERROR] Failed to parse error 500 response")
		return err
	}
	errorCode := parseErrorCode(fortiError.Error)
	log.Printf("[ERROR] Error code: %d context: %s", fortiError.Error, errorCode)
	log.Printf("[DEBUG] cli_error: %s", fortiError.CLIError)
	err = fmt.Errorf("error code: %d context: %s cli_error: %s (%d)", fortiError.Error, errorCode, fortiError.CLIError, fortiError.HTTPStatus)
	return err
}

func parseErrorCode(errorCode int64) string {
	switch e := errorCode; e {
	case int64(-8):
		return "Invalid IP"
	case int64(-9):
		return "Invalid netmask"
	case int64(-15):
		return "Duplicated entry"
	case int64(-651):
		return "Missing required attribute or attribute incorrect"
	default:
		return "No additional context available"
	}
}

// -8 invalid IP
// -9 invalid mask /33
// -15 duplicate entry
// -651 "enables" // missing required attribute // attribute incorrect
//

func newRequest(c config.Config, r *models.CmdbRequest) (*http.Request, error) {
	headers := buildHeaders(c, r)
	body := bytes.NewBuffer(r.Payload)
	log.Printf("[DEBUG] BODY: %s", body.String())
	// set URL w/ url queries
	url := buildURL(c, r)
	log.Printf("[DEBUG] %s %s", r.HTTPMethod, url)

	req, err := http.NewRequest(r.HTTPMethod, "", body)
	if err != nil {
		return nil, err
	}
	req.Header = *headers
	req.URL = url

	return req, nil
}

func buildHeaders(c config.Config, r *models.CmdbRequest) *http.Header {
	headers := http.Header{}
	bearer := "Bearer " + c.Auth.Token
	headers.Set("Authorization", bearer)
	headers.Set("Content-Type", "application/json")
	headers.Set("User-Agent", c.UserAgent)
	// for k, v := range *h {
	// 	for _, b := range v {
	// 		headers.Add(k, b)
	// 	}
	// }

	return &headers
}

func buildURL(c config.Config, r *models.CmdbRequest) *url.URL {
	u := url.URL{}
	u.Scheme = "https"
	u.Host = c.FwTarget
	u.Path = r.Path
	q := marshalParams(&r.Params)
	if r.Params.Vdom != "" {
		q.Set("vdom", r.Params.Vdom)
	} else {
		q.Set("vdom", c.Auth.Vdom)
	}
	if r.NoVdom {
		q.Del("vdom")
	}

	u.RawQuery = q.Encode()

	return &u
}

func marshalParams(params *models.CmdbRequestParams) url.Values {
	urlQuery := url.Values{}
	if params.AllowAppend != nil {
		v := strconv.FormatBool(*params.AllowAppend)
		urlQuery.Set("allow_append", v)
	}
	if params.PlainTextPassword != nil {
		v := strconv.FormatBool(*params.PlainTextPassword)
		urlQuery.Set("plain-text-password", v)
	}
	if params.Action != "" {
		urlQuery.Add("action", params.Action)
	}
	for _, v := range params.Filter {
		urlQuery.Add("filter", v)
	}
	for _, v := range params.Format {
		urlQuery.Add("format", v)
	}
	return urlQuery
}
