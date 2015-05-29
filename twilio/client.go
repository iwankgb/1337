// twilio sends text messages using twilio and is ridicolously simple...
package twilio

import (
	"net/http"
	"net/url"
)

type client struct {
	client *http.Client
	sid    *string
	token  *string
}

type requestBody struct {
	From string
	To   string
	Body string
}

func NewClient(sid *string, token *string) *client {
	httpClient := new(http.Client)
	return &client{client: httpClient, sid: sid, token: token}
}

func (c *client) Send(to *string, from *string, message *string) {
	http.PostForm("https://"+*c.sid+":"+*c.token+"@api.twilio.com/2010-04-01/Accounts/"+*c.sid+"/Messages.json", url.Values{"To": {*to}, "From": {*from}, "Body": {*message}})
}
