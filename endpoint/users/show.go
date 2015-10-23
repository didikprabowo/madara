package users
import (
    "github.com/gin-gonic/gin"
    "github.com/widnyana/madara/endpoint"
)

type User struct {
    ID int `json:"id"`
    Name string `json:"name"`
    Email string `json:"email_address"`
}


func RealMain(routerGroup *gin.RouterGroup) {
    routerGroup.GET("/users", func(c *gin.Context) {
        data := User{
            ID: 1,
            Name: "Widnyana",
            Email: "name@domain.tld",
        }

        c.JSON(200, data)
    })
}

func init() {
    endpoint.RegisterGroup(RealMain)
}
