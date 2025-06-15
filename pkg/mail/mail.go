package mail

import (
	"fmt"
	"net/smtp"
)

func SendRecoveryEmail(to string, newPassword string) error {
	from := "rescuefriendsuio@gmail.com"
	password := "vtfv twff qntj gwwp"

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	subject := "Subject: 🔐 Password Recovery - Shaggy Mission\n"
	
	body := fmt.Sprintf(`
<html>
<body style="font-family: Arial, sans-serif; background-color: #f4f7fa; color: #333; margin: 0; padding: 20px;">
  <div style="max-width: 600px; margin: auto; background-color: #ffffff; padding: 30px; border-radius: 12px; box-shadow: 0 4px 12px rgba(0,0,0,0.1);">
    
    <!-- Header -->
    <div style="text-align: center; margin-bottom: 30px;">
      <h1 style="color: #2c5aa0; margin: 0; font-size: 28px;">🐾 Shaggy Mission</h1>
      <p style="color: #666; margin: 5px 0 0 0; font-size: 14px;">Pet Rescue & Adoption Platform</p>
    </div>

    <!-- Greeting -->
    <div style="margin-bottom: 25px;">
      <h2 style="color: #333; font-size: 22px; margin: 0 0 10px 0;">Hello!</h2>
      <p style="color: #555; font-size: 16px; line-height: 1.5; margin: 0;">
        We received a request to reset your password for your Shaggy Mission account.
      </p>
    </div>

    <!-- Password Section -->
    <div style="background-color: #f8f9fa; padding: 25px; border-radius: 8px; margin: 25px 0; border-left: 4px solid #2c5aa0;">
      <h3 style="color: #2c5aa0; margin: 0 0 15px 0; font-size: 18px;">🔑 Your Temporary Password</h3>
      <div style="background-color: #ffffff; padding: 15px; border-radius: 6px; border: 2px dashed #2c5aa0; text-align: center;">
        <code style="font-size: 20px; font-weight: bold; color: #e74c3c; letter-spacing: 2px;">%s</code>
      </div>
    </div>

    <!-- Instructions -->
    <div style="margin: 25px 0;">
      <h3 style="color: #333; font-size: 18px; margin: 0 0 15px 0;">📝 Next Steps</h3>
      <ol style="color: #555; font-size: 16px; line-height: 1.6; margin: 0; padding-left: 20px;">
        <li style="margin-bottom: 8px;">Log in to your account using this temporary password</li>
        <li style="margin-bottom: 8px;">Go to your account settings</li>
        <li style="margin-bottom: 8px;"><strong>Change your password immediately</strong> for security</li>
      </ol>
    </div>

    <!-- Security Notice -->
    <div style="background-color: #fff3cd; padding: 20px; border-radius: 8px; border-left: 4px solid #ffc107; margin: 25px 0;">
      <h4 style="color: #856404; margin: 0 0 10px 0; font-size: 16px;">⚠️ Security Notice</h4>
      <p style="color: #856404; font-size: 14px; margin: 0; line-height: 1.5;">
        This temporary password will expire in 24 hours. If you didn't request this password reset, 
        please contact our support team immediately.
      </p>
    </div>

    <!-- Footer -->
    <div style="margin-top: 30px; padding-top: 20px; border-top: 1px solid #e9ecef; text-align: center;">
      <p style="color: #666; font-size: 14px; margin: 0 0 10px 0;">
        Thank you for using Shaggy Mission to help rescue animals! 🐶🐱
      </p>
      <p style="color: #999; font-size: 12px; margin: 0;">
        — The Shaggy Mission Team
      </p>
    </div>

  </div>
</body>
</html>
`, newPassword)

	message := []byte(subject + 
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n" +
		body)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	return err
}