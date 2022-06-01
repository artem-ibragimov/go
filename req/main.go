package req

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Req struct {
	tr     *http.Transport
	client *http.Client
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
	// body, _ := io.ReadAll(response.Body)
	// fmt.Println(string(body))
	defer response.Body.Close()
	return goquery.NewDocumentFromReader(response.Body)
}

func getUA() string {
	uas := []string{
		"Mozilla/5.0 (Linux; Android 7.0; SM-G930V Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.125 Mobile Safari/537.36",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 6_1_4 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10B350 Safari/8536.25",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:57.0) Gecko/20100101 Firefox/57.0",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36",
		"Mozilla/5.0 (Linux; Android 5.1.1; Nexus 5 Build/LMY48B; wv) AppleWebKit/537.36 (KHTML, like Gecko)  Version/4.0 Chrome/43.0.2357.65 Mobile Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/85 Version/11.1.1 Safari/605.1.15",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/600.7.12 (KHTML, like Gecko) Version/8.0.7 Safari/600.7.12",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/603.1.23 (KHTML, like Gecko) Version/10.0 Mobile/14E5239e Safari/602.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X)  AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1",
	}
	return uas[rand.Intn(len(uas))]
}
