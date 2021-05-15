package gemini

import (
	"crypto/hmac"
	"crypto/sha512"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type (

	// ClientOptions Used to create a new client.
	// NOTE: Client ID on Zabos config page IS NOT the secret.
	ClientOptions struct {
		PublicKey   string
		Secret      string
		Environment Environment
		HTTPClient  *http.Client
	}

	// Client Holds information required to interact with the Zabo services.
	Client struct {
		publicKey   string
		secret      string
		environment Environment
		httpClient  *http.Client
	}

	BaseRequest struct {
		Nonce   string `json:"nonce"`
		Request string `json:"request"`
	}
)

func (c *Client) createNonce() string {
	return strconv.Itoa(int(time.Now().UnixNano() / int64(time.Millisecond)))
}

// NewClient Initalizes a new client with http.Client.
func NewClient(options ClientOptions) (client *Client, err error) {
	if !options.Environment.IsValid() {
		fmt.Println("WARNING: Invalid Environment specified! Please use gemini.Sandbox or gemini.Live!")
	}

	if options.HTTPClient == nil {
		options.HTTPClient = &http.Client{}
	}

	client = &Client{
		secret:      options.Secret,
		publicKey:   options.PublicKey,
		environment: options.Environment,
		httpClient:  options.HTTPClient,
	}

	return
}

// Call Executes a basic HTTP call to the endpoint.
func (c *Client) Call(method, endpoint string, body []byte, v interface{}) error {
	req, err := c.newRequest(method, endpoint, body)
	if err != nil {
		return err
	}

	return c.do(req, v)
}

// newRequest is used by Call to generate a http.Request with appropriate headers.
func (c *Client) newRequest(method, endpoint string, body []byte) (*http.Request, error) {
	if !strings.HasPrefix(endpoint, "/") {
		endpoint = "/" + endpoint
	}

	req, err := http.NewRequest(method, string(c.environment)+endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Cache-Control", "no-cache")

	// Gemini API docs state that POST requests are usually identified as private requests.
	if method == "POST" {
		c.createAuthorizationHeaders(body, req)
	}
	return req, nil
}

// do is used by Call to execute an http.Request and parse its response .
// Also handles parsing of the gemini error format.
func (c *Client) do(req *http.Request, v interface{}) error {
	res, err := c.httpClient.Do(req)

	if err != nil {
		return err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	// Successful response
	if res.StatusCode == 200 || res.StatusCode == 201 {
		return json.NewDecoder(res.Body).Decode(v)
	}

	// Attempt to unmarshal into Gemini error format
	var geminiError Error
	if err = json.NewDecoder(res.Body).Decode(&geminiError); err != nil {
		return err
	}
	geminiError.StatusCode = res.StatusCode

	return geminiError
}

func (c *Client) createAuthorizationHeaders(body []byte, req *http.Request) {

	bs64Payload := b64.URLEncoding.EncodeToString(body)

	h := hmac.New(sha512.New384, []byte(c.secret))
	h.Write([]byte(bs64Payload))
	signature := hex.EncodeToString(h.Sum(nil))

	req.Header.Add("Content-Type", "text/plain")
	req.Header.Add("Content-Length", "0")
	req.Header.Add("X-GEMINI-APIKEY", c.publicKey)
	req.Header.Add("X-GEMINI-SIGNATURE", signature)
	req.Header.Add("X-GEMINI-PAYLOAD", bs64Payload)
}
