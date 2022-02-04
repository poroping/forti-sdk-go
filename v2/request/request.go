package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/poroping/forti-sdk-go/v2/config"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func CreateUpdate(c *config.Config, r *models.CmdbRequest) (*models.CmdbResponse, error) {
	req, err := newRequest(*c, r)
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

func Read(c *config.Config, r *models.CmdbRequest) (*models.CmdbResponse, error) {
	req, err := newRequest(*c, r)
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

func ReadString(c *config.Config, r *models.CmdbRequest) (*string, error) {
	req, err := newRequest(*c, r)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPCon.Do(req)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] Status code: %d", res.StatusCode)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if body == nil {
		log.Printf("[ERROR] Error empty string response during READ")
		return nil, err
	}

	response := &models.CmdbResponse{}
	err = json.Unmarshal(body, response)
	if err == nil {
		log.Printf("[WARN] Expected string, got CmdbResponse")
		// skipping error parsing cause non existant cert response missing httpmethod
		// {
		// 	"path":"system",
		// 	"name":"certificate",
		// 	"action":"download",
		// 	"serial":"FGT60FTK20028507",
		// 	"version":"v7.0.2",
		// 	"build":234,
		// 	"status":"error",
		// 	"http_status":404
		//   }
		// err = fortiErrorCheck(body, response)
		// if err != nil {
		// 	return nil, err
		// }
		return nil, nil
	}

	responseString := string(body)

	return &responseString, nil
}

func Delete(c *config.Config, r *models.CmdbRequest) (err error) {
	req, err := newRequest(*c, r)
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
		return err
	}

	return nil
}

func Generic(c *config.Config, r *models.CmdbRequest) (*string, error) {
	req, err := newRequest(*c, r)
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

	err = fortiErrorCheck(body, response)
	if err != nil {
		return nil, err
	}

	resp := string(body)

	return &resp, err
}

// Currently only checks for errors, don't trust false, needs work.
func Exists(c *config.Config, r *models.CmdbRequest) bool {
	_, err := Read(c, r)
	return err == nil
}

// Currently only checks for errors, don't trust false, needs work.
func ExistsWithResponse(c *config.Config, r *models.CmdbRequest) (*models.CmdbResponse, bool) {
	resp, err := Read(c, r)
	if err != nil {
		return resp, false
	}
	return resp, true
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
		// err = fmt.Errorf("resource Not Found - Unable to find the specified resource (%d)", res.HTTPStatus)
		err = parseError404(body)
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

func parseError404(body []byte) error {
	fortiError := &models.CmdbError404{}
	err := json.Unmarshal(body, &fortiError)
	if err != nil {
		log.Printf("[ERROR] Failed to parse error 404 response")
		return err
	}
	if *fortiError.HTTPStatus == int64(404) && (*fortiError.HTTPMethod == "GET" || *fortiError.HTTPMethod == "DELETE") {
		// FortiOS returns 404 for non-existant resource. If the results unmarshal and status is 404 then resource non existant.
		// Generally return an empty resource struct/nil upstream
		log.Print("[INFO] 404 response seen with GET or DELETE but unmarshalled successfully, not returning error as most likely a non-existant resource")
		return nil
	}

	return nil
}

func parseError500(body []byte) error {
	fortiError := &models.CmdbError500{}
	err := json.Unmarshal(body, &fortiError)
	if err != nil {
		log.Printf("[ERROR] Failed to parse error 500 response")
		return err
	}
	errorCode := parseErrorCode(fortiError.Error, fortiError.CLIError)
	log.Printf("[ERROR] Error code: %d context: %s", fortiError.Error, errorCode)
	log.Printf("[DEBUG] cli_error: %s", fortiError.CLIError)
	err = fmt.Errorf("error code: %d context: %s cli_error: %s (%d)", fortiError.Error, errorCode, fortiError.CLIError, fortiError.HTTPStatus)
	return err
}

func parseErrorCode(errorCode int64, errorString string) string {
	switch e := errorCode; e {
	case int64(-3):
		// parse errorString and return troubled datasource
		datasource := ""
		find := *regexp.MustCompile(`.*value parse error before '(.*)'.*`)
		match := find.FindAllStringSubmatch(errorString, -1)
		if len(match) == 1 {
			datasource = match[0][1]
		}
		return fmt.Sprintf("Referenced datasource %q does not exist\n", datasource)
	case int64(-5):
		return "Mkey already exists\n"
	case int64(-8):
		return "Invalid IP\n"
	case int64(-9):
		return "Invalid netmask\n"
	case int64(-15):
		return "Duplicated entry\n"
	case int64(-23):
		// potentially look at listing references?
		return "Unable to remove, target is referenced elsewhere\n"
	case int64(-651):
		return "Missing required attribute or attribute incorrect\n"
	default:
		return "No additional context available\n"
	}
}

// -3 missing datasource/reference
// -5 mkey already exists
// -8 invalid IP
// -9 invalid mask /33
// -15 duplicate entry
// -23 cannot delete, target referenced
// -651 "enables" // missing required attribute // attribute incorrect
//

func newRequest(c config.Config, r *models.CmdbRequest) (*http.Request, error) {
	headers := buildHeaders(c, r)
	body := bytes.NewBuffer(r.Payload)
	if body != nil {
		log.Printf("[DEBUG] BODY: %s", body.String())
	}
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
	// this shitmix is needed cause the fortios needs to query escape the mkey even though it's part of the path
	path, _ := url.Parse(r.Path)
	u.Path = path.Path
	u.RawPath = r.Path
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
	if params.Action != "" {
		urlQuery.Add("action", params.Action)
	}
	if params.AllowAppend != nil {
		v := strconv.FormatBool(*params.AllowAppend)
		urlQuery.Set("allow_append", v)
	}
	if params.After != "" {
		urlQuery.Add("after", params.After)
	}
	if params.Before != "" {
		urlQuery.Add("before", params.Before)
	}
	if params.Datasource != nil {
		v := strconv.FormatBool(*params.Datasource)
		urlQuery.Set("datasource", v)
	}
	if params.ExcludeDefaultValues != nil {
		v := strconv.FormatBool(*params.ExcludeDefaultValues)
		urlQuery.Set("exclude-default-values", v)
	}
	if params.Meta != nil {
		v := strconv.FormatBool(*params.Meta)
		urlQuery.Set("meta", v)
	}
	if params.Mkey != "" {
		urlQuery.Add("mkey", params.Mkey)
	}
	if params.Scope != "" {
		urlQuery.Add("scope", params.Scope)
	}
	if params.Type != "" {
		urlQuery.Add("type", params.Type)
	}
	if params.PlainTextPassword != nil {
		v := strconv.FormatBool(*params.PlainTextPassword)
		urlQuery.Set("plain-text-password", v)
	}
	if params.WithMeta != nil {
		v := strconv.FormatBool(*params.WithMeta)
		urlQuery.Set("with_meta", v)
	}
	if params.Filter != nil {
		v := params.Filter
		urlQuery.Add("filter", strings.Join(*v, ","))
	}
	if params.Format != nil {
		v := params.Format
		urlQuery.Add("format", strings.Join(*v, "|"))
	}
	if params.Sort != nil {
		v := params.Sort
		urlQuery.Add("sort", strings.Join(*v, ","))
	}
	return urlQuery
}
