# osstat

## Installation

```bash
go get -u github.com/metadiv-io/osstat
```

## Highlights

* MonitorMemory(e *micro.Engine, alertFunc func()) 

* MonitorCPU(e *micro.Engine, alertFunc func())

## Environment Variables

* ALERT_PERCENTAGE_CPU - from 0 to 1

* ALERT_PERCENTAGE_MEMORY - from 0 to 1

* DISABLE_ALERT_CPU - true/false

* DISABLE_ALERT_MEMORY - true/false
