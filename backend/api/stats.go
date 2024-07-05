package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"time"
)

type StatsResponse struct {
	Cpu      float64 `json:"cpu"`
	MemTotal uint64  `json:"mem_total"`
	MemUsed  uint64  `json:"mem_used"`
}

func Stats(c *fiber.Ctx) error {
	if !VerifyToken(c) {
		return &Error{Code: 401, Func: "api/stats", Message: "unauthorized"}
	}

	mem, err := memory.Get()
	if err != nil {
		return &Error{Code: 500, Func: "api/stats", Err: err}
	}

	cpuBefore, _ := cpu.Get()
	time.Sleep(time.Duration(50) * time.Millisecond)
	cpuAfter, err := cpu.Get()
	if err != nil {
		return &Error{Code: 500, Func: "api/stats", Err: err}
	}

	data := StatsResponse{
		Cpu:      float64(cpuAfter.Total-cpuBefore.Total) / 10,
		MemTotal: mem.Total / 1000000000,
		MemUsed:  mem.Used / 1000000000,
	}

	return c.Status(200).JSON(Response{Message: "success", Data: data})
}
