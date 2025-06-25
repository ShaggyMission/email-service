package handlers

import (
    "bytes"
    "encoding/json"
    "errors"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
)

func mockUpdateUserPasswordByEmailSuccess(email string) (string, error) {
    return "mockedPassword123", nil
}

func mockUpdateUserPasswordByEmailFail(email string) (string, error) {
    return "", errors.New("user not found")
}

func mockSendRecoveryEmailSuccess(to string, newPassword string) error {
    return nil
}

func mockSendRecoveryEmailFail(to string, newPassword string) error {
    return errors.New("failed to send email")
}

func RecoverPasswordHandlerTest(
    updatePasswordFunc func(string) (string, error),
    sendEmailFunc func(string, string) error,
) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req struct {
            Email string `json:"email" binding:"required,email"`
        }

        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email format"})
            return
        }

        newPassword, err := updatePasswordFunc(req.Email)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"message": "User not found or could not update password"})
            return
        }

        if err := sendEmailFunc(req.Email, newPassword); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to send recovery email"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully. Please check your email."})
    }
}

func TestRecoverPasswordHandler_InvalidEmail(t *testing.T) {
    gin.SetMode(gin.TestMode)
    router := gin.Default()

    router.POST("/password/recover", RecoverPasswordHandlerTest(mockUpdateUserPasswordByEmailSuccess, mockSendRecoveryEmailSuccess))

    payload := map[string]string{"email": "not-an-email"}
    jsonPayload, _ := json.Marshal(payload)

    req, _ := http.NewRequest("POST", "/password/recover", bytes.NewBuffer(jsonPayload))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    if w.Code != http.StatusBadRequest {
        t.Errorf("Expected status 400, got %d", w.Code)
    }
}

func TestRecoverPasswordHandler_UserNotFound(t *testing.T) {
    gin.SetMode(gin.TestMode)
    router := gin.Default()

    router.POST("/password/recover", RecoverPasswordHandlerTest(mockUpdateUserPasswordByEmailFail, mockSendRecoveryEmailSuccess))

    payload := map[string]string{"email": "user@example.com"}
    jsonPayload, _ := json.Marshal(payload)

    req, _ := http.NewRequest("POST", "/password/recover", bytes.NewBuffer(jsonPayload))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    if w.Code != http.StatusNotFound {
        t.Errorf("Expected status 404, got %d", w.Code)
    }
}

func TestRecoverPasswordHandler_EmailSendFail(t *testing.T) {
    gin.SetMode(gin.TestMode)
    router := gin.Default()

    router.POST("/password/recover", RecoverPasswordHandlerTest(mockUpdateUserPasswordByEmailSuccess, mockSendRecoveryEmailFail))

    payload := map[string]string{"email": "user@example.com"}
    jsonPayload, _ := json.Marshal(payload)

    req, _ := http.NewRequest("POST", "/password/recover", bytes.NewBuffer(jsonPayload))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    if w.Code != http.StatusInternalServerError {
        t.Errorf("Expected status 500, got %d", w.Code)
    }
}

func TestRecoverPasswordHandler_Success(t *testing.T) {
    gin.SetMode(gin.TestMode)
    router := gin.Default()

    router.POST("/password/recover", RecoverPasswordHandlerTest(mockUpdateUserPasswordByEmailSuccess, mockSendRecoveryEmailSuccess))

    payload := map[string]string{"email": "user@example.com"}
    jsonPayload, _ := json.Marshal(payload)

    req, _ := http.NewRequest("POST", "/password/recover", bytes.NewBuffer(jsonPayload))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Errorf("Expected status 200, got %d", w.Code)
    }
}
