package api

import (
	"medev21/crypto-dashboard-backend/api/router"
)

func StartApplication() {

	web := router.Setup()

	web.Run(":8080")

}
