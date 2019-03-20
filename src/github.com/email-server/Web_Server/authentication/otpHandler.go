package authentication

import (
	"log"	
	"time"
	"strconv"
	"math/rand"
	"github.com/subosito/twilio"
)

var (
	AccountSid = "ACdc8df276bc645fa623463aa214593f4c"
	AuthToken  = "e81860b62714c717095afba59ac31e3a"
)

func GenerateOTP() string {
	c := twilio.NewClient(AccountSid, AuthToken, nil)
	otpGen := strconv.Itoa(genRandomNumber())

	log.Println(otpGen)
	params := twilio.MessageParams{
		Body: "Your OTP for blah.com is : " + otpGen + "\nThank you for using our service.\n",
	}
	s, response, err := c.Messages.Send("+12017405668", "+918050078481", params)
	if err != nil {
		log.Fatal(s, response, err)
	}
	return otpGen
}

func genRandomNumber() int{
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	
	return r1.Intn(1000)
}

