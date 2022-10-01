package proxy

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/cyborch/corp/internal/config"
	"github.com/smallnest/ringbuffer"
)

func getVirtualHost(
	req *http.Request,
	conf config.Configurations) *config.VirtualHostConfigurations {

	for _, vh := range conf.VirtualHosts {
		if req.Host == vh.Hostname {
			return &vh
		}
	}
	return nil
}

// set req Host, URL and Request URI to forward a request to the origin server
func setOriginUrl(req *http.Request, vh config.VirtualHostConfigurations) {
	uri, _ := url.Parse(vh.Origin)
	req.Host = uri.Host
	req.URL.Host = uri.Host
	req.URL.Scheme = uri.Scheme
	req.RequestURI = ""
}

func removeProxyHeaders(req *http.Request, vh config.VirtualHostConfigurations) {
	req.Header.Set("Referer", vh.Origin)
	req.Header.Del("X-Scheme")
	req.Header.Del("X-Forwarded-Host")
	req.Header.Del("X-Forwarded-Proto")
	req.Header.Del("X-Real-Ip")
	req.Header.Del("X-Forwarded-Port")
	req.Header.Del("X-Request-Id")
	req.Header.Del("X-Forwarded-For")
}

// Write body from origin request body where origin urls are replaced
// proxy hostname urls.
func writeWithReplacedOrigin(req *http.Request,
	res *http.Response,
	rw http.ResponseWriter,
	vh config.VirtualHostConfigurations) {

	if strings.HasPrefix(res.Header.Get("Content-Type"), "text/") {
		rw.Header().Del("Content-Length")
		rw.Header().Add("Transfer-Encoding", "chunked")
		rw.WriteHeader(res.StatusCode)

		length := len(vh.Origin)
		buf := make([]byte, length)
		count, _ := res.Body.Read(buf)

		for count == length {
			if string(buf) == vh.Origin {
				rw.Write([]byte(vh.Scheme))
				rw.Write([]byte("://"))
				rw.Write([]byte(vh.Hostname))
				count, _ = res.Body.Read(buf)
			} else {
				rw.Write(buf[:1])
				ring := ringbuffer.New(length)
				ring.Write(buf[1:length])
				next := make([]byte, 1)
				count, _ = res.Body.Read(next)
				if count > 0 {
					ring.Write(next)
					count, _ = ring.Read(buf)
				}
			}
		}
		rw.Write(buf)
	} else {
		rw.WriteHeader(res.StatusCode)
		io.Copy(rw, res.Body)
	}
	req.Body.Close()
}

func logRequest(req *http.Request, since time.Time, status int) {
	elapsed := time.Now().Sub(since).Seconds()
	fmt.Printf("[corp] [%s]: %s / %s %f %d %s\n",
		since,
		req.Method,
		req.Proto,
		float32(elapsed),
		status,
		req.URL.String())
}

func handlerFunc(conf config.Configurations) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		requestTime := time.Now()
		vh := getVirtualHost(req, conf)
		if vh == nil {
			rw.WriteHeader(http.StatusBadGateway)
			rw.Write([]byte(fmt.Sprintf("No origin found for %s", req.Host)))
			logRequest(req, requestTime, http.StatusBadGateway)
			return
		}

		setOriginUrl(req, *vh)
		removeProxyHeaders(req, *vh)

		// save the response from the origin server
		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
		originServerResponse, err := client.Do(req)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(rw, err)
			logRequest(req, requestTime, http.StatusInternalServerError)
			return
		}

		// copy response headers to the client
		for key, values := range originServerResponse.Header {
			skip := false
			for _, header := range vh.SkipHeaders {
				if strings.ToLower(key) == strings.ToLower(header) {
					// Skip any unwanted headers
					skip = true
					continue
				}
			}
			if skip {
				continue
			}

			if vh.EnableCors && strings.ToLower(key) == "access-control-allow-origin" {
				continue
			}

			if strings.ToLower(key) == "location" {
				location := strings.Replace(values[0], vh.Origin, fmt.Sprintf("%s://%s", vh.Scheme, vh.Hostname), 0)
				rw.Header().Add(key, location)
				continue
			}

			for _, value := range values {
				rw.Header().Add(key, value)
			}
		}
		if vh.EnableCors {
			rw.Header().Add("Access-Control-Allow-Origin", "*")
		}

		// return response body to the client
		writeWithReplacedOrigin(req, originServerResponse, rw, *vh)
		logRequest(req, requestTime, originServerResponse.StatusCode)
	})
}

func Listen() {
	conf, err := config.Config()
	if err == nil {
		address := fmt.Sprintf("0.0.0.0:%d", conf.Server.Port)
		fmt.Printf("[corp] [%s]: Listening on %s\n", time.Now(), address)
		http.ListenAndServe(address, handlerFunc(conf))
	}
}
