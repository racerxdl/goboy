// +build ignore

package main

import (
	"github.com/racerxdl/goboy/cpu/generator/gendata"
	"log"
	"os"
	"time"
)

// We need to generate the instruction set, but this doesnt work from a go-generate
// So run manually: go run geninstructionset.go
//go:generate go run geninstructionset.go

func main() {
	f, err := os.Create("instructions.go")
	if err != nil {
		log.Fatalf(err.Error())
	}

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
	instructions += gendata.BuildSUB()
	instructions += gendata.BuildCP()
	instructions += gendata.BuildOperators()
	instructions += gendata.BuildIncDec()
	instructions += "// endregion \n"
	// endregion
	// region Bit Manipulation
	instructions += "// region Bit Manipulation Instructions\n"
	instructions += gendata.BuildBitManipulation()
	instructions += "// endregion \n"
	// endregion
	// region Interrupt calls
	instructions += "// region Interrupt Calls Instructions\n"
	instructions += gendata.BuildInterruptCalls()
	instructions += "// endregion \n"
	// endregion
	// region StackManagement calls
	instructions += "// region Stack Management Instructions\n"
	instructions += gendata.BuildStackManagement()
	instructions += "// endregion \n"
	// endregion
	// region Flow Control
	instructions += "// region Flow Control Instructions\n"
	instructions += gendata.BuildFlowControl()
	instructions += "// endregion \n"
	// endregion

	gendata.InstructionsFileTemplate.Execute(f, struct {
		Timestamp    time.Time
		Instructions string
	}{
		Timestamp:    time.Now(),
		Instructions: instructions,
	})
	f.Close()

	// CB Instructions
	instructions = gendata.BuildCB()

	f, err = os.Create("cbinstructions.go")
	if err != nil {
		log.Fatalf(err.Error())
	}

	gendata.CBInstructionsFileTemplate.Execute(f, struct {
		Timestamp    time.Time
		Instructions string
	}{
		Timestamp:    time.Now(),
		Instructions: instructions,
	})
	f.Close()
}
