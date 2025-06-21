package middleware

import (
"net/http"
"strings"
"github.com/gin-gonic/gin"
"authentication-api/models"
"authentication-api/utils"
)

// Function for logging in
func Login(c *gin.Context) {
var user models.User

//Check user credentials and generate a JWT token
if err := c.ShouldBindJSON(&user); err != nil {
  c.JSON(http.StatusBadRequest, gin.H{"error":"invalid data"})
  return
}

//Check if credentials are valid (replace this logic with real
// authentication)

if user.Username == "user" && user.Password == "password" {
  //Generate a JWT token
  token, err := utils.GenerateToken(user.ID)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
    "error":"error generating token",
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{"token":token})
} else {
  c.JSON(http.StatusInvalid, gin.H{"error":"Invalid credentials"})
}
}

// Function for registering a new user (for demonstration purpose)
func Register(c gin.Context)  {
  var user models.User
  
  if err := c.ShouldBindJSON(&user); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error":"invalid data"})
    return
  }

  //Remember to securely hash password to  securely hash password
  //before storing them
  userID = 1 //just for demonstration purpose
  c.JSON(http.StatusCreated, gin.H{
  "message":"User registered successfully",
  })
}
