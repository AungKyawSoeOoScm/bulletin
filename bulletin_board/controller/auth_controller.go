package controller

import (
	"fmt"
	"gin_test/bulletin_board/data/request"
	"gin_test/bulletin_board/helper"
	service "gin_test/bulletin_board/service/auth"
	"path/filepath"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	AuthService service.Authservice
}

func NewAuthController(service service.Authservice) *AuthController {
	return &AuthController{
		AuthService: service,
	}
}

// Register Controller
func (controller *AuthController) Register(ctx *gin.Context) {
	username := ctx.PostForm("username")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	phone := ctx.PostForm("phone")
	address := ctx.PostForm("address")
	dob := ctx.PostForm("dob")
	utype := ctx.PostForm("type")
	var userType string
	if utype == "1" {
		userType = "1"
	} else {
		userType = "0"
	}

	var dobTime *time.Time
	if dob != "" {
		parsedDOB, err := time.Parse("2006-01-02", dob)
		if err != nil {
			fmt.Print("Invalid date of birth")
		}
		dobTime = &parsedDOB
	}
	// Check if email already exists
	existingUser := controller.AuthService.FindByEmail(email)
	if existingUser.Id != 0 {
		helper.ResponseHandler(ctx, http.StatusBadRequest, "Email already exists.", nil)
		return
	}

	// //Required Username
	// if username == "" {
	// 	ctx.Set("UsernameError", "Username is required.")
	// 	helper.ResponseHandler(ctx, http.StatusBadRequest, "Username is required.", nil)
	// 	return
	// }

	// //Required Email
	// if email == "" {
	// 	ctx.Set("EmailError", "Email is required.")
	// 	helper.ResponseHandler(ctx, http.StatusBadRequest, "Email is required.", nil)
	// 	return
	// }

	// //Required Password
	// if password == "" || len(password) < 6 {
	// 	ctx.Set("PasswordError", "Password must be at least 6 characters long.")
	// 	helper.ResponseHandler(ctx, http.StatusBadRequest, "Password must be at least 6 characters long.", nil)
	// 	return
	// }
	photoFile, err := ctx.FormFile("photo")
	if err != nil && err != http.ErrMissingFile {
		helper.ErrorPanic(err)
	}

	var photoPath string
	if photoFile != nil {
		// Generate a unique file name for the photo
		photoFileName := fmt.Sprintf("%d_%s", time.Now().Unix(), photoFile.Filename)
		photoPath = filepath.Join("static", "images", photoFileName)

		// Save the uploaded file to the desired location
		err := ctx.SaveUploadedFile(photoFile, photoPath)
		if err != nil {
			helper.ErrorPanic(err)
		}

		// Convert backslashes to forward slashes
		photoPath = filepath.ToSlash(photoPath)
	}
	createUserRequest := request.CreateUserRequest{
		Username:      username,
		Email:         email,
		Password:      password,
		Phone:         phone,
		Address:       address,
		Date_Of_Birth: dobTime,
		Type:          userType,
		Profile_Photo: photoPath,
	}
	fmt.Println(createUserRequest)
	aerr := controller.AuthService.Register(createUserRequest)
	fmt.Println(aerr, "ererrrrrrr")
	if aerr != nil {
		if validationErr, ok := aerr.(validator.ValidationErrors); ok {
			errorMessages := make(map[string]string)
			for _, fieldErr := range validationErr {
				fieldName := fieldErr.Field()
				errorMessage := ""
				switch fieldErr.Tag() {
				case "required":
					errorMessage = fieldName + " field is required"
				case "min":
					errorMessage = fieldName + " must be at least " + fieldErr.Param() + " characters long"
				case "max":
					errorMessage = fieldName + " must not exceed " + fieldErr.Param() + " characters"
				default:
					errorMessage = "Field validation failed"
				}
				errorMessages[fieldName] = errorMessage
				fmt.Println(errorMessages[fieldName])
			}

			ctx.HTML(http.StatusBadRequest, "usercreateform.html", gin.H{
				"Errors": errorMessages,
			})
			return
		} else {

			// Handle other types of errors (if needed)
			// ...
			ctx.HTML(http.StatusBadRequest, "usercreateform.html", gin.H{
				"Errors": map[string]string{
					"general": "An error occurred during registration",
				},
			})
			return
		}
	} else {
		ctx.Redirect(http.StatusFound, "/login")
	}
}

// Login Controller
func (controller *AuthController) Login(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	loginRequest := request.LoginRequest{
		Email:    email,
		Password: password,
	}
	token, err_token := controller.AuthService.Login(loginRequest)
	if err_token != nil {
		helper.ResponseHandler(ctx, http.StatusBadRequest, "Invalid email or password.", nil)
		return
	}
	// resp := response.LoginResponse{
	// 	TokenType: "Bearer",
	// 	Token:     token,
	// }
	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour), // Set cookie expiration time
		MaxAge:   3600,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // Set to true if using HTTPS
		SameSite: http.SameSiteStrictMode,
	}
	ctx.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
	ctx.Redirect(http.StatusFound, "/posts")
}

// Register Form
func (controller *AuthController) RegisterForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.html", gin.H{})
}

// Login Form
func (controller *AuthController) LoginForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}

// Logout

func (controller *AuthController) Logout(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "", false, true)
	ctx.Redirect(http.StatusFound, "/")
}
