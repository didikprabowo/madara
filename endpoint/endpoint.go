package endpoint

import (
    "github.com/widnyana/madara"
    "github.com/gin-gonic/gin"
)


var groupRegFuncs []func(*gin.RouterGroup)

func RegisterGroup(regFunc func(*gin.RouterGroup)) {
    groupRegFuncs = append(groupRegFuncs, regFunc)
}

func RealMain(engine *gin.Engine) {
    group := engine.Group("/")
    for _, regFunc := range groupRegFuncs {
        regFunc(group)
    }
}

func init() {
    madara.RegisterHandler(RealMain)
}
