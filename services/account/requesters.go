package account

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func doRequest(url string, method string, body string) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(url)
	req.Header.SetMethod(method)
	req.SetBodyString(body)

	fasthttp.Do(req, resp)

	bodyBytes := resp.Body()
	fmt.Println(string(bodyBytes))
	// User-Agent: fasthttp
	// Body:
}
