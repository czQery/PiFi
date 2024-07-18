package api

import (
	"github.com/czQery/PiFi/backend/cmd"
	"github.com/czQery/PiFi/backend/hp"
	"github.com/gofiber/fiber/v2"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"math"
)

type StatsResponse struct {
	Cpu      float64              `json:"cpu"`
	MemTotal uint64               `json:"mem_total"`
	MemUsed  uint64               `json:"mem_used"`
	Hotspot  StatsHotspotResponse `json:"hotspot"`
}

type StatsHotspotResponse struct {
	SSID   string `json:"ssid"`
	Portal bool   `json:"portal"`
}

var StatsCPU *cpu.Stats

func Stats(c *fiber.Ctx) error {
	iface := make(map[string]SettingsInterfaceResponse)
	err := hp.Config.Unmarshal("settings.iface", &iface)
	if err != nil {
		return &Error{Code: 500, Func: "api/settings", Err: err}
	}

	var hotspotSSID string
	for _, i := range iface {
		if i.Mode == "hotspot" {
			hotspotSSID = i.SSID
		}
	}

	memNow, err := memory.Get()
	if err != nil {
		return &Error{Code: 500, Func: "api/stats", Err: err}
	}

	cpuNow, err := cpu.Get()
	if err != nil {
		return &Error{Code: 500, Func: "api/stats", Err: err}
	}

	if StatsCPU == nil {
		StatsCPU = cpuNow
	}

	total := float64(cpuNow.Total - StatsCPU.Total)
	data := StatsResponse{
		Cpu:      math.Round((float64(cpuNow.System-StatsCPU.System)/total*100)*100) / 100,
		MemTotal: memNow.Total / 1000000000,
		MemUsed:  memNow.Used / 1000000000,
		Hotspot: StatsHotspotResponse{
			SSID:   hotspotSSID,
			Portal: cmd.Portal != "",
		},
	}

	return c.Status(200).JSON(Response{Message: "success", Data: data})
}
