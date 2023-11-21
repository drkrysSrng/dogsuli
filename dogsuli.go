package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

// If we do not use a container
/*
const (

	TOR1 = "socks5://127.0.0.1:9051"
	TOR2 = "socks5://127.0.0.1:9052"
	TOR3 = "socks5://127.0.0.1:9053"

)
*/
const (
	TOR1 = "socks5://tor1:9050"
	TOR2 = "socks5://tor2:9050"
	TOR3 = "socks5://tor3:9050"
)

/*
 *	Main crawler function
 *
 */
func crawler(urlSearch string, user_agent string) {

	u, _ := url.Parse(urlSearch)

	// Instantiate default collector
	c := colly.NewCollector(
		colly.UserAgent(user_agent),
		colly.AllowedDomains(u.Host),
	)

	// Do not revisit
	c.AllowURLRevisit = false

	// Disable keepAlive (slower but more secure)
	c.WithTransport(&http.Transport{
		DisableKeepAlives: true,
	})

	c.SetRequestTimeout(120 * time.Second)

	rp, err := proxy.RoundRobinProxySwitcher(TOR1, TOR2, TOR3)
	if err != nil {
		fmt.Print(err)
	}

	c.SetProxyFunc(rp)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		absoluteUrl := e.Request.AbsoluteURL(link)

		fmt.Println("Link a[href] found:", e.Text, link, absoluteUrl)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {

		fmt.Println("Visiting...", r.URL.String())
	})

	// extract status code
	c.OnResponse(func(r *colly.Response) {

		fmt.Println("Response received: ", r.StatusCode)

	})

	c.OnError(func(r *colly.Response, err error) {

		fmt.Println("URL not reachable:", r.Request.URL, "Status", r.StatusCode, "Error", err)

	})

	c.OnScraped(func(r *colly.Response) {

		fmt.Println("*************FINISHED", r.Request.URL.String())

	})

	// Start scraping
	c.Visit(urlSearch)

	//Wait until all threads are visited
	c.Wait()

}
