// +build ignore

package main

import (
	"bytes"
	"fmt"
	"github.com/racerxdl/goboy/cpu/generator"
	"github.com/racerxdl/goboy/cpu/generator/gendata"
	"github.com/racerxdl/goboy/cpu/generator/testdata"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"text/template"
	"time"
)

// We need to generate the instruction set, but this doesnt work from a go-generate
// So run manually: go run geninstructionset.go
//go:generate go run geninstructionset.go

func main() {
	z, err := os.Create("instructions.go")
	if err != nil {
		log.Fatalf(err.Error())
	}

	f := bytes.NewBuffer(nil)

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

	data, err := format.Source(f.Bytes())

	if err != nil {
		panic(err)
	}

	z.Write(data)
	z.Close()

	// CB Instructions
	instructions = gendata.BuildCB()
	f.Reset()
	z, err = os.Create("cbinstructions.go")
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

	data, err = format.Source(f.Bytes())

	if err != nil {
		panic(err)
	}

	z.Write(data)
	z.Close()

	// Generate Tests
	f.Reset()
	z, err = os.Create("instructions_test.go")
	if err != nil {
		log.Fatalf(err.Error())
	}

	f.WriteString(`package cpu

import (
    "math/rand"
    "testing"
)

const RunCycles = 10

`)

	for _, v := range generator.InstructionSet {
		if v.TemplateName == "RSTXX" || v.TemplateName == "CBCall" {
			continue
		}
		tpld, err := ioutil.ReadFile(fmt.Sprintf("generator/testtpl/%s.gotpl", v.TemplateName))

		if err != nil {
			panic(err)
		}

		tpl := template.Must(template.New(v.TemplateName).Parse(string(tpld)))

		arg0 := ""
		arg1 := ""
		arg2 := ""

		if len(v.TemplateArgs) >= 1 {
			arg0 = v.TemplateArgs[0]
		}

		if len(v.TemplateArgs) >= 2 {
			arg1 = v.TemplateArgs[1]
		}

		if len(v.TemplateArgs) >= 3 {
			arg2 = v.TemplateArgs[2]
		}

		flags := testdata.GenFlagTest(v.ZSHC)

		d := struct {
			OpCodeX     string
			OpCode      string
			Instruction string
			Cycles      []int
			Asserts     string
			Flags       string
			Arg0        string
			Arg1        string
			Arg2        string
		}{
			OpCodeX:     fmt.Sprintf("%02X", v.Opcode),
			OpCode:      fmt.Sprintf("%d", v.Opcode),
			Instruction: v.Name,
			Cycles:      v.Cycles,
			Arg0:        arg0,
			Arg1:        arg1,
			Arg2:        arg2,
			Flags:       flags,
		}

		err = tpl.Execute(f, d)

		f.WriteString("\n")
		if err != nil {
			panic(err)
		}
	}

	data, err = format.Source(f.Bytes())

	if err != nil {
		z.Write(f.Bytes())
		fmt.Println("Error formatting code: ")
		panic(err)

	}

	z.Write(data)
	z.Close()
}
