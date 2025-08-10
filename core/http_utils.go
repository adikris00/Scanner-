package core

import (
	"crypto/tls"
	"io"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"

	utls "github.com/refraction-networking/utls"
)

var acceptLangs = []string{
	"en-US,en;q=0.9",
	"en-GB,en;q=0.8",
	"id-ID,id;q=0.9,en-US;q=0.8,en;q=0.7",
	"fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7",
	"de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
}

var referers = []string{
	"https://www.google.com/",
	"https://www.bing.com/",
	"https://duckduckgo.com/",
	"https://search.yahoo.com/",
	"https://yandex.com/",
}

func NewHTTPClient() *http.Client {
	dialTLS := func(network, addr string) (net.Conn, error) {
		tcpConn, err := net.DialTimeout(network, addr, 10*time.Second)
		if err != nil {
			return nil, err
		}
		serverName := strings.Split(addr, ":")[0]
		config := &tls.Config{ServerName: serverName}

		utlsConfig := &utls.Config{
			ServerName:         config.ServerName,
			InsecureSkipVerify: config.InsecureSkipVerify,
			MinVersion:         config.MinVersion,
			MaxVersion:         config.MaxVersion,
			CipherSuites:       config.CipherSuites,
		}

		helloIDs := []utls.ClientHelloID{
			utls.HelloChrome_Auto,
			utls.HelloFirefox_Auto,
			utls.HelloIOS_Auto,
			utls.HelloAndroid_11_OkHttp,
			utls.HelloRandomized,
			utls.HelloRandomizedALPN,
		}
		uTLSConn := utls.UClient(tcpConn, utlsConfig, helloIDs[rand.Intn(len(helloIDs))])
		if err := uTLSConn.Handshake(); err != nil {
			return nil, err
		}
		return uTLSConn, nil
	}

	return &http.Client{
		Transport: &http.Transport{
			DialTLS: dialTLS,
		},
		Timeout: 15 * time.Second,
	}
}

func SendRequest(client *http.Client, url, path string) (string, error) {
	time.Sleep(time.Duration(rand.Intn(1000)+500) * time.Millisecond)

	req, err := http.NewRequest("GET", url+path, nil)
	if err != nil {
		return "", err
	}

	for key, value := range Headers {
		if strings.ToLower(key) != "user-agent" {
			req.Header.Set(key, value)
		}
	}

	if len(UserAgents) > 0 {
		req.Header.Set("User-Agent", UserAgents[rand.Intn(len(UserAgents))])
	}

	req.Header.Set("Accept-Language", acceptLangs[rand.Intn(len(acceptLangs))])

	req.Header.Set("Referer", referers[rand.Intn(len(referers))])

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}
