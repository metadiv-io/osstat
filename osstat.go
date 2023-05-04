package osstat

import (
	"fmt"
	"log"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/metadiv-io/logger"
	"github.com/metadiv-io/micro"
)

// MonitorMemory monitor memory usage
// alertPercent is >= 0 and <= 1
func MonitorMemory(e *micro.Engine, alertFunc func()) {
	micro.Cron(e, "@every 10s", func() {
		m, err := memory.Get()
		if err != nil {
			log.Println("failed to get os memory info", err)
			return
		}
		if float64(m.Used)/float64(m.Total) > ALERT_PERCENTAGE_MEMORY {
			logger.Prefix(fmt.Sprintf("[system:%s]", micro.SYSTEM_UUID)).Error("memory used percent alert: ", float64(m.Used)/float64(m.Total))
			if !DISABLE_ALERT_MEMORY {
				alertFunc()
			}
		}
	})
}

// MonitorCPU monitor cpu usage
// alertPercent is >= 0 and <= 1
func MonitorCPU(e *micro.Engine, alertFunc func()) {
	micro.Cron(e, "@every 10s", func() {
		c, err := cpu.Get()
		if err != nil {
			log.Println("failed to get os cpu info", err)
			return
		}
		if float64(c.User+c.System)/float64(c.Total) > ALERT_PERCENTAGE_CPU {
			logger.Prefix(fmt.Sprintf("[system:%s]", micro.SYSTEM_UUID)).Error("cpu used percent alert: ", float64(c.User+c.System)/float64(c.Total))
			if !DISABLE_ALERT_CPU {
				alertFunc()
			}
		}
	})
}
