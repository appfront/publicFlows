// curl -X POST --data-urlencode  https://hooks.slack.com/services/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

package main

import (
	"log"
	"os"

	"github.com/parnurzeal/gorequest"
)

func postToSlack(url string, message string, channel string) {

	if url == "" {
		log.Fatal("envs")
	}

	request := gorequest.New()
	payload := `{"text":"` + message + `","channel":"#` + channel + `","username":"appfront", "icon_url": "https://s3-us-west-2.amazonaws.com/slack-files2/avatars/2016-04-13/34421041538_30998f3e7bb695f2ee02_96.png"}`

	resp, _, errs := request.Post(url).
		Set("Content-Type", "application/x-www-form-urlencoded").
		Send("payload=" + payload).
		End()

	if errs != nil {

		log.Fatal(errs)
	}

	if resp.Status != "200 OK" {
		println(resp.Status)
		log.Fatal("error")
	}

}

func main() {
	url := os.Getenv("HOOK")
	message := os.Getenv("MESSAGE")
	channel := os.Getenv("CHANNEL")

	postToSlack(url, message, channel)
	//postToSlack("https://hooks.slack.com/services/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", "test", "testone")
}
