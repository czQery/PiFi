package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/sirupsen/logrus"
	"time"
)

type StatsResponse struct {
	Cpu      float64 `json:"cpu"`
	MemTotal uint64  `json:"mem_total"`
	MemUsed  uint64  `json:"mem_used"`
}

func Stats(c *fiber.Ctx) error {
	if !VerifyToken(c) {
		return c.Status(401).JSON(Response{Message: "unauthorized"})
	}

	mem, err := memory.Get()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("api - stats mem failed")
		return c.Status(500).JSON(Response{Message: "unexpected error"})
	}

	cpuBefore, errC1 := cpu.Get()
	time.Sleep(time.Duration(50) * time.Millisecond)
	cpuAfter, errC2 := cpu.Get()
	if errC1 != nil || errC2 != nil {
		logrus.WithFields(logrus.Fields{
			"error": errC1.Error() + "," + errC2.Error(),
		}).Error("api - stats cpu failed")
		return c.Status(500).JSON(Response{Message: "unexpected error"})
	}

	data := StatsResponse{
		Cpu:      float64(cpuAfter.Total-cpuBefore.Total) / 10,
		MemTotal: mem.Total / 1000000000,
		MemUsed:  mem.Used / 1000000000,
	}

	return c.Status(200).JSON(Response{Message: "success", Data: data})
}
