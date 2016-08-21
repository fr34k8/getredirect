package getredirect

import (
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
)

// Question struct
type Question struct {
	JobURL     string    `json:"JobURL"`
	JobStatus  string    `json:"JobStatus"`
	JobMessage string    `json:"JobMessage"`
	JobTime    time.Time `json:"JobTime"`
}

// Answer struct
type Answer struct {
	HTTPStatus     int      `json:"HTTPStatus"`
	HTTPStatusText string   `json:"HTTPStatusText"`
	HTTPRedirect   string   `json:"HTTPRedirect,omitempty"`
	Location       string   `json:"Location,omitempty"`
	HTTPProto      string   `json:"HTTPProto"`
	IPAddress      []string `json:"IPAddress,omitempty"`
	Connection     string   `json:"Connection,omitempty"`
}

// ReturnMessage for retunring
type ReturnMessage struct {
	Question Question `json:"Question"`
	Answer   Answer   `json:"Answer"`
}

// From function
func From(testurl string) *ReturnMessage {
	h := new(ReturnMessage)
	h.Question.JobTime = time.Now()
	h.Question.JobURL = testurl

	url, err := url.Parse(testurl)
	if err != nil {
		panic(err)
	}

	// to get the Port number, split the Host
	hostport := strings.Split(url.Host, ":")
	hostname := hostport[0]
	// port := hostport[1]

	if govalidator.IsURL(testurl) == false {
		log.Println("Failed: Not a valid URL.")
		h.Question.JobStatus = "Failed"
		h.Question.JobMessage = "Not a valid URL"
		return h
	}

	addrs, err := net.LookupHost(hostname)
	if err != nil {
		log.Println("DNS lookup failed: .", err)
		h.Question.JobStatus = "Failed"
		h.Question.JobMessage = "DNS lookup failed"
		return h
	}

	// Get HTTP status.
	req, err := http.NewRequest("GET", testurl, nil)
	req.Header.Add("User-Agent", "AboutSecurity.nl Checker")

	if err != nil {
		log.Println("Failed: Can not connect.", err)
		h.Question.JobStatus = "Failed"
		h.Question.JobMessage = err.Error()
		return h
	}

	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Println("Failed connect.", err)
		h.Question.JobStatus = "Failed"
		h.Question.JobMessage = err.Error()
		return h
	}

	h.Answer.HTTPStatus = resp.StatusCode
	h.Answer.HTTPStatusText = resp.Status
	h.Answer.HTTPProto = resp.Proto
	h.Answer.IPAddress = addrs

	// TODO: Do not always get a Location
	h.Answer.Location = resp.Header.Get("Location")

	timeout := time.Duration(2 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp2, err := client.Get(testurl)
	if err != nil {
		h.Question.JobStatus = "Failed"
		h.Question.JobMessage = err.Error()
		return h
	}
	h.Answer.HTTPRedirect = resp2.Request.URL.String()

	h.Answer.Connection = resp.Header.Get("Connection")

	h.Question.JobStatus = "OK"
	h.Question.JobMessage = "Job done!"

	return h
}
