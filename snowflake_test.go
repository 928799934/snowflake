package snowflake

import (
	"log"
	"testing"
)

func TestGen(t *testing.T) {
	conf := Config{1, 2, int64(1288834974657)}
	i := NewIdWorker(conf)
	//id := i.Generate()
	//log.Printf("id: %d \n", id)

	//id := int64(759885747922141184)
	id := int64(851217241810141184)
	machine := i.Machine(id)
	datacenter := i.Datacenter(id)
	log.Printf("time: %d \n", i.Millisecond(id, int64(1288834974657)))
	log.Printf("machine: %d \n", machine)
	log.Printf("datacenter: %d \n", datacenter)
}
