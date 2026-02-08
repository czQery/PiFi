package api

import (
	"github.com/czQery/PiFi/backend/db"
	"github.com/czQery/PiFi/backend/net"
	"github.com/gofiber/fiber/v2"
)

type DBResponse struct {
	APs      int64         `json:"aps"`
	Wigle    DBNetResponse `json:"wigle"`
	BeaconDB DBNetResponse `json:"beacondb"`
	DWPA     DBNetResponse `json:"dwpa"`
}

type DBNetResponse struct {
	New int64 `json:"new"`
	Net int64 `json:"net"`
}

func DB(c *fiber.Ctx) error {
	data := db.SelectAP(c.Context())

	var (
		aps = int64(len(data))

		netWigle    int64
		netBeacondb int64
		netDwpa     int64
	)

	_, dbWigle := net.WigleGet(c.Context(), true)

	for _, ap := range data {
		if ap.Wigle {
			netWigle++
		}

		if ap.BeaconDB {
			netBeacondb++
		}

		if ap.DWPA {
			netDwpa++
		}
	}

	return c.Status(200).JSON(Response{Message: "success", Data: DBResponse{
		APs: aps,
		Wigle: DBNetResponse{
			New: int64(len(dbWigle)),
			Net: netWigle,
		},
		BeaconDB: DBNetResponse{
			New: 0,
			Net: netBeacondb,
		},
		DWPA: DBNetResponse{
			New: 0,
			Net: netDwpa,
		},
	}})
}
