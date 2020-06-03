package internal

import (
	"fmt"
	"os"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
)

// ServerInformation info
type ServerInformation struct {
	MemoryTotal  uint64  `json:"mem_total"`
	MemoryUsed   uint64  `json:"mem_used"`
	MemoryCached uint64  `json:"mem_cached"`
	MemoryFree   uint64  `json:"mem_free"`
	CPUUser      float64 `json:"cpu_user"`
	CPUSystem    float64 `json:"cpu_system"`
	CPUIdle      float64 `json:"cpu_idle"`
}

// GetStat get stats
// stats
func GetStat() *ServerInformation {
	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return nil
	}

	before, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return nil
	}
	time.Sleep(time.Duration(2) * time.Second)
	after, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return nil
	}
	total := float64(after.Total - before.Total)

	stat := &ServerInformation{
		MemoryTotal:  memory.Total,
		MemoryUsed:   memory.Used,
		MemoryCached: memory.Cached,
		MemoryFree:   memory.Free,
		CPUUser:      float64(after.User-before.User) / total * 100,
		CPUSystem:    float64(after.System-before.System) / total * 100,
		CPUIdle:      float64(after.Idle-before.Idle) / total * 100,
	}
	return stat
}
