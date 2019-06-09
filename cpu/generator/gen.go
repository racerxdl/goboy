// +build ignore

package main

import (
	"github.com/racerxdl/goboy/cpu/generator/gendata"
	"log"
	"os"
	"time"
)

func main() {

	f, err := os.Create("instructions.go")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer f.Close()

	instructions := ""

	// region Load/Store Instructions
	instructions += "// region Load/Store Instructions\n"
	instructions += gendata.BuildLDRR()
	instructions += gendata.BuildLDrHLm()
	instructions += gendata.BuildLDHLmr()
	instructions += gendata.BuildLDrn()
	instructions += gendata.BuildLSSingles()
	instructions += gendata.BuildLDrrmr()
	instructions += gendata.BuildLDmm()
	instructions += gendata.BuildLDrrrm()
	instructions += gendata.BuildLDrmm()
	instructions += gendata.BuildLDRRnn()
	instructions += gendata.BuildLDHLI()
	instructions += gendata.BuildLDrIOn()
	instructions += "// endregion\n"
	// endregion
	// region Data Processing Instructions
	instructions += "// region Data Processing Instructions\n"
	instructions += gendata.BuildADD()
	instructions += "// endregion \n"
	// endregion

	gendata.InstructionsFileTemplate.Execute(f, struct {
		Timestamp    time.Time
		Instructions string
	}{
		Timestamp:    time.Now(),
		Instructions: instructions,
	})
}
