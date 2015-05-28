package twilio

import (
	//	"bytes"
	//	"encoding/json"
	"fmt"
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
	fmt.Println(*to)
	fmt.Println(*from)
	fmt.Println(*message)
	query := "To=" + *to + "&From=" + *from + "&Body=" + *message
	form, _ := url.ParseQuery(query)
	request := new(http.Request)
	//	request.SetBasicAuth(*c.sid, *c.token)
	request.URL.RawQuery = form.Encode()
	response, err := c.client.Do(request)
	fmt.Println(response)
	fmt.Println(err)
	//	jsonObj := &requestBody{From: *from, To: *to, Body: *message}
	//	jsonBody, _ := json.Marshal(jsonObj)
	//	fmt.Println(jsonBody)
	//	values := new(url.Values)
	//	values.Add("To", *to)
	//	values.Add("From", *from)
	//	values.Add("Body", *message)
	//	fmt.Println(values)
	//	c.client.PostForm()
	//	c.client.PostForm("https://api.twilio.com/2010-04-01/Accounts/"+*c.sid+"/Messages", values)
	//	request, _ := http.NewRequest("POST", "https://api.twilio.com/2010-04-01/Accounts/"+*c.sid+"/Messages", bytes.NewBuffer(jsonBody))
	//	request.Header.Add("Content-Type", "application/json")
	//	request.SetBasicAuth(*c.sid, *c.token)
	//	fmt.Println(request)
	//	result, err := c.client.Do(request)
	//	fmt.Println(result)
	//	fmt.Println(err)
	//	body := bytes.NewBuffer(make([]byte, 1000))
	//	enc := json.NewEncoder(body)
	//	json := &requestBody{From: *from, To: *to, Body: *message}
	//	fmt.Println(enc.Encode(json))
	//	fmt.Println(json)

	//	json := ""
	//	dupa := ""
	//	request, _ := http.NewRequest("POST", "https://api.twilio.com/2010-04-01/"+*c.sid+"/Messages", bytes.NewBuffer(dupa))
}
