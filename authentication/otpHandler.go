package authentication

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/subosito/twilio"
)

var (
	AccountSid = "ACdc8df276bc645fa623463aa214593f4c"
	AuthToken  = "e81860b62714c717095afba59ac31e3a"
)

func GenerateOTP() string {
	c := twilio.NewTwilio(AccountSid, AuthToken)
	otpGen := strconv.Itoa(genRandomNumber())

	log.Println(otpGen)
	params := twilio.SMSParams{
		StatusCallback: "Your OTP for blah.com is : " + otpGen + "\nThank you for using our service.\n",
	}
	response, err := c.SendSMS("+12017405668", "+918050078481", otpGen, params)
	if err != nil {
		log.Fatal(response, err)
	}
	return otpGen
}

func genRandomNumber() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return r1.Intn(1000)
}
