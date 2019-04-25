package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/chaintex/currencies/tomochain"
	raven "github.com/getsentry/raven-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sentry"
	"github.com/gin-gonic/gin"
)

type InfoData struct {
	TokenAPIs []tomochain.TokenAPI `json:"tokens"`
}
type HTTPServer struct {
	host string
	r    *gin.Engine
}

func (httpServer *HTTPServer) GetCurrencies(c *gin.Context) {
	var file []byte
	var err error
	var infoData InfoData

	file, err = ioutil.ReadFile("env/tokens.json")
	if err != nil {
		log.Print(err)
	} else {
		err = json.Unmarshal(file, &infoData)
		if err != nil {
			log.Print(err)
		}
	}

	c.JSON(
		http.StatusOK,
		gin.H{"success": true, "data": infoData.TokenAPIs},
	)
}

func (httpServer *HTTPServer) Run(chainTexENV string) {
	httpServer.r.GET("/currencies", httpServer.GetCurrencies)

	httpServer.r.Run(httpServer.host)
}

func NewHTTPServer(host string) *HTTPServer {
	r := gin.Default()
	r.Use(sentry.Recovery(raven.DefaultClient, false))
	r.Use(cors.Default())

	return &HTTPServer{host, r}
}
