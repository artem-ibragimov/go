package req

import (
	"encoding/base64"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Req struct {
	tr     *http.Transport
	client *http.Client
}

type IReq interface {
	Get(url string) (*goquery.Document, error)
	GetImg(url string) (string, error)
}

func New() IReq {
	request := new(Req)
	request.Init()
	return request
}

func (r *Req) Init() {
	r.tr = &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    60 * time.Second,
		DisableCompression: true,
	}
	r.client = &http.Client{Transport: r.tr, Timeout: 120 * time.Second}
}

func (r *Req) Get(url string) (*goquery.Document, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		time.Sleep(time.Minute * 2)
		req, err = http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
	}

	req.Header.Set("User-Agent", getUA())

	response, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return goquery.NewDocumentFromReader(response.Body)
}

func getUA() string {
	uas := []string{
		"http",
		"crawler",
		"spider",
		"bot",
		"search",
		// "Mozilla/5.0 (iPhone; CPU iPhone OS 6_1_4 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10B350 Safari/8536.25",
		// "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:57.0) Gecko/20100101 Firefox/57.0",
		// "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36",
		// "Mozilla/5.0 (Linux; Android 5.1.1; Nexus 5 Build/LMY48B; wv) AppleWebKit/537.36 (KHTML, like Gecko)  Version/4.0 Chrome/43.0.2357.65 Mobile Safari/537.36",
		// "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/85 Version/11.1.1 Safari/605.1.15",
		// "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/600.7.12 (KHTML, like Gecko) Version/8.0.7 Safari/600.7.12",
		// "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/603.1.23 (KHTML, like Gecko) Version/10.0 Mobile/14E5239e Safari/602.1",
		// "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X)  AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1",
	}
	return uas[rand.Intn(len(uas))]
}

func (r *Req) GetImg(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		time.Sleep(time.Minute * 2)
		req, err = http.NewRequest("GET", url, nil)
		if err != nil {
			time.Sleep(time.Minute * 5)
			req, err = http.NewRequest("GET", url, nil)
			if err != nil {
				return "", err
			}
		}
	}

	req.Header.Set("User-Agent", getUA())

	response, err := r.client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding = "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding = "data:image/png;base64,"
	}

	// Append the base64 encoded output
	return base64Encoding + toBase64(bytes), nil
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
