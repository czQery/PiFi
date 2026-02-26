package sse

import (
	"bufio"
	"math"
	"os"
	"sync"
	"time"

	"github.com/czQery/PiFi/backend/cmd"
	"github.com/czQery/PiFi/backend/hp"
	"github.com/gofiber/fiber/v2"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
)

type responseStats struct {
	Cpu      float64              `json:"cpu"`
	MemTotal uint64               `json:"mem_total"`
	MemUsed  uint64               `json:"mem_used"`
	Hotspot  responseStatsHotspot `json:"hotspot"`
	GPS      cmd.GPSData          `json:"gps"`
}

type responseStatsHotspot struct {
	SSID   string `json:"ssid"`
	Portal bool   `json:"portal"`
}

type responseLog struct {
	Raw any `json:"raw"`
}

var DashClients sync.Map

func Dash(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")

	c.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		logChan := make(chan []byte, 5)
		DashClients.Store(logChan, struct{}{})
		defer DashClients.Delete(logChan)

		var (
			data    interface{}
			cpuLast *cpu.Stats
			err     error
		)

		data, cpuLast = getStats(cpuLast)
		_ = sendEvent(w, "stats", data)
		data, _ = os.ReadFile(hp.LogFileName)
		_ = sendEvent(w, "log", responseLog{Raw: data})

		for {
			select {
			case data = <-logChan:
				err = sendEvent(w, "log", responseLog{Raw: data})
			case <-ticker.C:
				data, cpuLast = getStats(cpuLast)
				err = sendEvent(w, "stats", data)
			}

			if err != nil {
				return
			}
		}
	})

	return nil
}

func DashLog(line []byte) {
	DashClients.Range(func(key, _ interface{}) bool {
		ch := key.(chan []byte)
		select {
		case ch <- line:
		default:
		}
		return true
	})
}

func getStats(cpuLast *cpu.Stats) (responseStats, *cpu.Stats) {
	var data responseStats

	memNow, err := memory.Get()
	if err != nil {
		return data, nil
	}

	cpuNow, err := cpu.Get()
	if err != nil {
		return data, nil
	}

	data = responseStats{
		MemTotal: memNow.Total,
		MemUsed:  memNow.Used,
		Hotspot: responseStatsHotspot{
			SSID:   cmd.Hotspot,
			Portal: cmd.Portal != "",
		},
		GPS: cmd.GPS,
	}

	if cpuLast != nil {
		total := float64(cpuNow.Total - cpuLast.Total)
		data.Cpu = math.Round((float64(cpuNow.System-cpuLast.System)/total*100)*100) / 100
	}

	return data, cpuNow
}
