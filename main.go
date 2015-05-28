// 1337 project main.go
package main

import (
	"flag"
	"fmt"
	"github.com/iwankgb/1337/chuck"
	"github.com/iwankgb/1337/twilio"
)

var path = flag.String("path", "", "Path to Chuck Norris file")
var sid = flag.String("sid", "", "Twilio AccountSID")
var token = flag.String("token", "", "Twilio authentication token")
var to = flag.String("to", "", "Receipient phone numbet")
var from = flag.String("from", "1337", "Sender ID")

func main() {
	flag.Parse()
	generator := chuck.NewGenerator(*path)
	chuck := generator.Get()
	fmt.Println(chuck)
	client := twilio.NewClient(sid, token)
	client.Send(to, from, &chuck)
}
