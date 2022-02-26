package utils

import (
	"os"
	"time"

	"gopkg.in/ezzarghili/recaptcha-go.v4"
)

func CaptchaVerifyToken(token string, action string) (bool, error) {
	captcha, err := recaptcha.NewReCAPTCHA(os.Getenv("RECAPTCHA_SECRET_KEY"), recaptcha.V3, 10*time.Second) // for v3 API use https://g.co/recaptcha/v3 (apperently the same admin UI at the time of writing)
	if err != nil {
		return false, err
	}

	err = captcha.VerifyWithOptions(token, recaptcha.VerifyOption{Action: action, Threshold: 0.8, Hostname: "localhost"})
	if err != nil {
		return false, err
	}

	return true, nil
}
