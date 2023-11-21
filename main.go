package main

import "time"

func main() {

	// Sleep for 20 seconds for the proxy to run
	duration := 20 * time.Second
	time.Sleep(duration)

	USER_AGENT_TOR := "Mozilla/5.0 (Windows NT 10.0; rv:68.0) Gecko/20100101 Firefox/68.0"

	url := "http://4usoivrpy52lmc4mgn2h34cmfiltslesthr56yttv2pxudd3dapqciyd.onion/"

	crawler(url, USER_AGENT_TOR)

	// If we want to use threads
	// go crawler(...)

}
