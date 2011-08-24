package Trailhead

import (
	"app"
	"smtp"
	"log"
)

func init(){
    app.Start()
    app.Get("/index",index)
    app.Post("/request",request)
    app.Get("/request",index)
}

func index(ctx app.Context){
    ctx.Show("html/index")
}

func request(ctx app.Context){
	auth := smtp.PlainAuth(
		"",
		"leighpots@trailheadpottery.com",
		"weiner27",
		"smtp.gmail.com",
	)
	//headers := "MIME-Version: 1.0\r\nContent-Type: text/html\r\n"
	msg := "<html><body><h1>This is the email body</h1></body></html>"
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"aninneman@curtmfg.com",
		[]string{"ninnemana@gmail.com"},
		[]byte(msg),
	)
	if err != nil {
		log.Fatal(err)
	}
    ctx.Redirect("/")
}