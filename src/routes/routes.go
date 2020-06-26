package routes

import (
	"net/http"

	"github.com/arifseft/go-auth/src/controllers"
	"github.com/arifseft/go-auth/src/validations"
	"github.com/gin-gonic/gin"
)

// Router middleware to handler routes
func Router(g *gin.RouterGroup) {
	// // the jwt middleware
	// authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
	//     Realm:       "test zone",
	//     Key:         []byte("secret key"),
	//     Timeout:     time.Hour,
	//     MaxRefresh:  time.Hour,
	//     IdentityKey: identityKey,
	//     PayloadFunc: func(data interface{}) jwt.MapClaims {
	//         if v, ok := data.(*User); ok {
	//             return jwt.MapClaims{
	//                 identityKey: v.UserName,
	//             }
	//         }
	//         return jwt.MapClaims{}
	//     },
	//     IdentityHandler: func(c *gin.Context) interface{} {
	//         claims := jwt.ExtractClaims(c)
	//         return &User{
	//             UserName: claims[identityKey].(string),
	//         }
	//     },
	//     Authenticator: func(c *gin.Context) (interface{}, error) {
	//         var loginVals login
	//         if err := c.ShouldBind(&loginVals); err != nil {
	//             return "", jwt.ErrMissingLoginValues
	//         }
	//         userID := loginVals.Username
	//         password := loginVals.Password
	//
	//         if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
	//             return &User{
	//                 UserName:  userID,
	//                 LastName:  "Bo-Yi",
	//                 FirstName: "Wu",
	//             }, nil
	//         }
	//
	//         return nil, jwt.ErrFailedAuthentication
	//     },
	//     Authorizator: func(data interface{}, c *gin.Context) bool {
	//         if v, ok := data.(*User); ok && v.UserName == "admin" {
	//             return true
	//         }
	//
	//         return false
	//     },
	//     Unauthorized: func(c *gin.Context, code int, message string) {
	//         c.JSON(code, gin.H{
	//             "code":    code,
	//             "message": message,
	//         })
	//     },
	//     // TokenLookup is a string in the form of "<source>:<name>" that is used
	//     // to extract token from the request.
	//     // Optional. Default value "header:Authorization".
	//     // Possible values:
	//     // - "header:<name>"
	//     // - "query:<name>"
	//     // - "cookie:<name>"
	//     // - "param:<name>"
	//     TokenLookup: "header: Authorization, query: token, cookie: jwt",
	//     // TokenLookup: "query:token",
	//     // TokenLookup: "cookie:token",
	//
	//     // TokenHeadName is a string in the header. Default value is "Bearer"
	//     TokenHeadName: "Bearer",
	//
	//     // TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
	//     TimeFunc: time.Now,
	// })
	//
	// if err != nil {
	//     log.Fatal("JWT Error:" + err.Error())
	// }
	//
	// r.POST("/login", authMiddleware.LoginHandler)
	//
	// r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
	//     claims := jwt.ExtractClaims(c)
	//     log.Printf("NoRoute claims: %#v\n", claims)
	//     c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	// })
	//
	// auth := r.Group("/auth")
	// // Refresh time can be longer than token timeout
	// auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	// auth.Use(authMiddleware.MiddlewareFunc())
	// {
	//     auth.GET("/hello", helloHandler)
	// }

	controller := controllers.UController()
	{
		g.GET("/users", controller.GetUsers)
		g.GET("/user/:id", validations.GetUser, controller.GetUser)
		g.POST("/user", validations.CreateUser, controller.CreateUser)
		g.PATCH("/user/:id", validations.UpdateUser, controller.UpdateUser)
		g.DELETE("/user/:id", validations.DeleteUser, controller.DeleteUser)
		// g.POST("/user/login", validations.LoginUser, controller.LoginUser)

		g.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":   http.StatusOK,
				"message":  "ready",
				"database": "error",
			})
		})
	}
}
