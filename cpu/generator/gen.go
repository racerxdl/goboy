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
    instructions += gendata.BuildLDRR()
    instructions += gendata.BuildLDrHLm()
    instructions += gendata.BuildLDHLmr()
    instructions += gendata.BuildLDrn()
    instructions += gendata.BuildLDHLmn()
    instructions += gendata.BuildLDrrmr()
    instructions += gendata.BuildLDmm()
    instructions += gendata.BuildLDrrrm()
    instructions += gendata.BuildLDrmm()
    instructions += gendata.BuildLDRRnn()
    // endregion

    gendata.InstructionsFileTemplate.Execute(f, struct {
        Timestamp time.Time
        Instructions string
    }{
        Timestamp: time.Now(),
        Instructions: instructions,
    })
}

