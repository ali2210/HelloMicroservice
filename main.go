package main


import(
	"github.com/gin-gonic/gin"
	 service "github.com/pusher/pusher-http-go"
	"net/http"
	"fmt"
)

const(
	app_id string = "1208250"
	key  string = "cc7e790e1675802bcca8"
	secret string = "550664a82fe91017b83c"
	cluster string = "ap2"
)

var(
	eventsMicroservice map[string]string
)

func InitService() (map[string]string) {
	eventType := "message"
	eventsMicroservice = make(map[string]string)
	eventsMicroservice[eventType] = "Hello-Microservice"
	return eventsMicroservice
}

func main(){
	route := gin.Default()
	route.LoadHTMLGlob("templates/*")
	serviceClient := service.Client{
		AppID : app_id,
		Key: key,
		Secret: secret,
		Cluster: cluster,
		Secure: true,
	}
	pushEvent:= InitService()
	fmt.Println(pushEvent)
	route.GET("/",func (c *gin.Context){
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title" : "Hello-Microservice",
				"payload": serviceClient.Trigger("hellomicro", "my-event", pushEvent),			
			},
		)
	})
	route.Run()
}