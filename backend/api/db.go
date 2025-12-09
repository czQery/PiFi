package api

import (
	"github.com/czQery/PiFi/backend/db"
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
	data := db.SelectAP()

	var (
		aps      = int64(len(data))
		wigle    int64
		beacondb int64
		dwpa     int64
	)

	for _, ap := range data {
		if ap.Wigle {
			wigle++
		}

		if ap.BeaconDB {
			beacondb++
		}

		if ap.DWPA {
			dwpa++
		}
	}

	return c.Status(200).JSON(Response{Message: "success", Data: DBResponse{
		APs: aps,
		Wigle: DBNetResponse{
			New: aps - wigle,
			Net: wigle,
		},
		BeaconDB: DBNetResponse{
			New: aps - beacondb,
			Net: beacondb,
		},
		DWPA: DBNetResponse{
			New: aps - dwpa,
			Net: dwpa,
		},
	}})
}
