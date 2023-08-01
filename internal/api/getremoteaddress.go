package api

import (
	"log"
	"net"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetRemoteAddress(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"address": "http://" + net.JoinHostPort(GetOutboundIP().String(), "8788"),
	})
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
