package otp

import (
	"fmt"
	"github.com/jangirvipin/go-otp/api/utils"
	"github.com/jangirvipin/go-otp/lib"
	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"os"
)

type Client struct {
	client  *twilio.RestClient
	phone   string
	rClient *lib.RedisClient
}

func NewOtpClient() *Client {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	accountSID := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")
	fromPhone := os.Getenv("FROM_PHONE")

	// create a new Twilio client
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSID,
		Password: authToken,
	})

	r := lib.NewRedisClient()

	return &Client{
		client:  client,
		phone:   fromPhone,
		rClient: r,
	}
}

func (c *Client) SendOtp(phone string) (string, error) {
	// Generate OTP
	otp := utils.GenerateOTP()

	params := &openapi.CreateMessageParams{}
	params.SetTo(phone)
	params.SetFrom(c.phone)
	params.SetBody("Your OTP is: " + otp)

	_, err := c.client.Api.CreateMessage(params)
	if err != nil {
		return "", fmt.Errorf("Error sending OTP: %w", err)
	}
	fmt.Println("✅ OTP sent successfully!")

	err = c.rClient.SetOtp(phone, otp)
	if err != nil {
		return "", fmt.Errorf("Error saving OTP to Redis: %w", err)
	}

	return otp, nil
}

func (c *Client) VerifyOtp(inputOtp string, phone string) (error, bool) {
	result, err := c.rClient.GetOtp(inputOtp)
	if err != nil {
		fmt.Println("❌ Error verifying OTP:", err)
		return err, false
	}
	if result != phone {
		fmt.Println("❌ OTP does not match for phone:", phone)
		return fmt.Errorf("OTP does not match"), false
	}
	return nil, true
}
