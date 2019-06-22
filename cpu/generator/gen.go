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
    instructions += gendata.BuildBitManipulation()
    // endregion
    // region Interrupt calls
    instructions += gendata.BuildInterruptCalls()
    // endregion
    // region StackManagement calls
    instructions += gendata.BuildStackManagement()
    // endregion
    // region Flow Control
    instructions += gendata.BuildFlowControl()
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

    gendata.CBInstructionsFileTemplate.Execute(f, struct{
        Timestamp    time.Time
        Instructions string
    }{
        Timestamp:    time.Now(),
        Instructions: instructions,
    })
    f.Close()
}
