# Snowflake
This package is based on the golang implementation of the snowflake algorithm used to generate a unique id
# Usage:
---------
```go
conf := Config{1, 2, 0}
i := NewIdWorker(conf)
id := i.Generate()
log.Printf("id: %d \n", id)

machine := i.Machine(id)
log.Printf("machine: %d \n", machine)

datacenter := i.Datacenter(id)
log.Printf("datacenter: %d \n", datacenter)
```