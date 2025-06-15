# 📧 Email Recovery Service - Shaggy Mission

<div align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go" />
  <img src="https://img.shields.io/badge/Gin-00ADD8?style=for-the-badge&logo=gin&logoColor=white" alt="Gin" />
  <img src="https://img.shields.io/badge/GORM-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="GORM" />
  <img src="https://img.shields.io/badge/Gmail_SMTP-EA4335?style=for-the-badge&logo=gmail&logoColor=white" alt="Gmail SMTP" />
  <img src="https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL" />
</div>

<div align="center">
  <h3>🔐 Secure Password Recovery Microservice for Pet Rescue Platform</h3>
  <p><em>Part of the Shaggy Mission ecosystem - Helping heroes regain access to save more lives! 🐾</em></p>
</div>

---

## 🌟 Overview

The **Email Recovery Service** is a secure password recovery microservice in the Shaggy Mission platform that handles user password resets through email delivery. This service ensures that volunteers, adopters, veterinarians, and administrators can quickly regain access to their accounts when passwords are forgotten, maintaining uninterrupted rescue mission coordination.

## 🎯 What This Service Does

- **Secure Password Recovery**: Generates temporary passwords and delivers them via email
- **Automated Email Delivery**: Sends professionally formatted HTML recovery emails
- **Random Password Generation**: Creates cryptographically secure temporary passwords
- **Database Integration**: Updates user passwords with bcrypt hashing
- **24-Hour Expiration**: Implements time-limited temporary access for security
- **SMTP Integration**: Utilizes Gmail SMTP for reliable email delivery
- **User Validation**: Verifies user existence before processing recovery requests

## 🛠️ Tech Stack

- **Runtime**: Go with Gin web framework
- **Database**: PostgreSQL with GORM ORM
- **Email Service**: Gmail SMTP with authentication
- **Security**: bcrypt for password hashing
- **Cryptography**: crypto/rand for secure password generation
- **Encoding**: base64 URL encoding for password generation
- **HTTP Framework**: Gin for RESTful API endpoints

## 📡 API Endpoints

### Password Recovery Endpoint
**`POST /password/recover`**
- Initiates password recovery process for registered users
- Generates secure temporary password
- Sends recovery email with HTML formatting
- Updates user password in database with bcrypt hashing

```json
{
  "email": "user@example.com"
}
```

**Response Success (200 OK):**
```json
{
  "message": "Password reset successfully. Please check your email."
}
```

**Response Error (404 Not Found):**
```json
{
  "message": "User not found or could not update password"
}
```

**Response Error (400 Bad Request):**
```json
{
  "message": "Invalid email format"
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "message": "Failed to send recovery email"
}
```

## 🔧 Core Functionality

### Password Recovery Process
The service handles secure password recovery by validating the provided email format, checking user existence in the database, generating a cryptographically secure random password, hashing the new password with bcrypt, updating the user record in the database, sending a professionally formatted HTML email with the temporary password, and providing clear next steps for the user.

### Email Template Features
The recovery email includes a branded Shaggy Mission header with pet rescue theming, clear temporary password display, step-by-step instructions for account access, security notice about 24-hour expiration, and professional HTML formatting for better user experience.

### Security Implementation
The service implements robust security through bcrypt password hashing with cost factor 10, cryptographically secure random password generation using crypto/rand, base64 URL encoding for safe password characters, database queries with email validation, and temporary password expiration notices.

## 🌐 Service Integration

This microservice serves as the password recovery component for the entire Shaggy Mission platform, working independently to handle forgotten password scenarios. It integrates with the main user authentication system and provides seamless account recovery for all platform users.

## 🔒 Security Features

- **Bcrypt Hashing**: All passwords encrypted with industry-standard hashing
- **Secure Random Generation**: Cryptographically secure password creation
- **Email Validation**: Validates email format before processing
- **User Verification**: Confirms user existence before password reset
- **Temporary Access**: 24-hour expiration for enhanced security
- **SMTP Authentication**: Secure email delivery through Gmail SMTP
- **Error Handling**: Comprehensive error management for security scenarios

## 📧 Email Configuration

The service uses Gmail SMTP for email delivery with the following configuration:
- **SMTP Host**: smtp.gmail.com
- **SMTP Port**: 587
- **Authentication**: App-specific password for enhanced security
- **Content Type**: HTML with UTF-8 encoding
- **Email Template**: Professional branded recovery email

## 🗃️ Database Operations

The service performs secure database operations by finding users by email address with unscoped queries, generating secure temporary passwords, applying bcrypt hashing to new passwords, updating user records with error handling, and maintaining data integrity throughout the recovery process.

## 🚀 Getting Started

### Prerequisites
- Go 1.19 or higher
- PostgreSQL database
- Gmail account with app-specific password
- GORM database connection configured

### Environment Setup
```go
// Configure your database connection in config package
db := config.ConnectDB()

// Update Gmail credentials in mail package
from := "your-email@gmail.com"
password := "your-app-specific-password"
```

### Running the Service
```bash
go run main.go
```

The service will start on port 4000 and be ready to handle password recovery requests.

---

<div align="center">
  <p><strong>Built with ❤️ for street dogs and cats everywhere 🐕🐱</strong></p>
  <p><em>Every password recovery keeps our heroes connected to rescue missions!</em></p>
</div>