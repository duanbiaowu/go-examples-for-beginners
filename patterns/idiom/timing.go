package idiom

import (
	"log"
	"time"
)

func Duration(invocation time.Time, name string) {
	elapsed := time.Since(invocation)
	log.Printf("%s lasted %s", name, elapsed)
}
