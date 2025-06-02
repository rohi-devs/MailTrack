package Authentication

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"os"
	"rohidevs.engineer/mailTrack/Model"
	"rohidevs.engineer/mailTrack/Utlis"
	"time"
)

type Auth struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func Login(c echo.Context) error {
	_ = godotenv.Load()
	jwtSecret := os.Getenv("JWT_SECRET")
	email := c.FormValue("email")
	password := c.FormValue("password")
	var user Model.User
	db, _ := c.Get("db").(*gorm.DB)
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return c.JSON(401, map[string]string{"message": "Invalid email"})
	}
	if err := Utlis.CompareHashAndPassword(user.Password, password); err != nil {
		return c.JSON(401, map[string]string{"message": "Invalid password"})
	}
	if email == "" || password == "" {
		return c.JSON(400, map[string]string{"message": "Username and password are required"})
	}
	claims := &Auth{
		Email: user.Email,
		Role:  "user",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	tokenUnsigned := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenUnsigned.SignedString([]byte(jwtSecret))
	if err != nil {
		return c.JSON(500, map[string]string{"message": "Error signing token"})
	}
	return c.JSON(200, map[string]string{"token": token})
}

func Register(c echo.Context) error {
	db, _ := c.Get("db").(*gorm.DB)

	username := c.FormValue("username")
	password := c.FormValue("password")
	email := c.FormValue("email")
	if username == "" || password == "" || email == "" {
		return c.JSON(400, map[string]string{"message": "Username and password are required"})
	}
	passwordHash, _ := Utlis.Hash(password)
	user := Model.User{
		Name:     username,
		Password: passwordHash,
		Email:    email,
	}
	if err := db.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to register user"})
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"message": "user registered successfully",
		"user_id": user.ID,
	})
}
