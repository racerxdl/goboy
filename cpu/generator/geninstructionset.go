// +build ignore

package main

import (
	"bytes"
	"github.com/racerxdl/goboy/cpu/generator/models"
	"go/format"
	"os"
	"text/template"
	"time"
)

var instTpl = template.Must(template.New("InstSet").Parse(`// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// {{ .Timestamp }}
package generator

import "github.com/racerxdl/goboy/cpu/generator/models"

var InstructionSet = []models.Instruction{ {{.INSTS}} }
var CBInstructionSet = []models.CBInstruction{ {{.CBINSTS}} }
`))

var instructionTemplate = template.Must(template.New("Instruction Template").Parse(`
    {
        Opcode:        {{.Opcode}},
        Name:          "{{.Name}}",
        Instruction:   "{{.Instruction}}",
        Cycles:        []int{ {{range .Cycles}} {{.}}, {{end}} },
        ZSHC:          "{{.ZSHC}}",
        NumberOfBytes: {{.NumberOfBytes}},
        Zero:          "{{.Zero}}",
        Sub:           "{{.Sub}}",
        HalfCarry:     "{{.HalfCarry}}",
        Carry:         "{{.Carry}}",
        TemplateName:  "{{.TemplateName}}",
        TemplateArgs:  []string{ {{range .TemplateArgs}}"{{.}}", {{end}} },
    },`))

var cbInstructionTemplate = template.Must(template.New("CB Instruction Template").Parse(`
    {
        Opcode:        {{.Opcode}},
        Name:          "{{.Name}}",
        Instruction:   "{{.Instruction}}",
        Cycles:        []int{ {{range .Cycles}} {{.}}, {{end}} },
        ZSHC:          "{{.ZSHC}}",
        Zero:          "{{.Zero}}",
        Sub:           "{{.Sub}}",
        HalfCarry:     "{{.HalfCarry}}",
        Carry:         "{{.Carry}}",
        TemplateName:  "{{.TemplateName}}",
        TemplateArgs:  []string{ {{range .TemplateArgs}}"{{.}}", {{end}} },
    },`))

func main() {
	err, insts := models.ParseInstructionsFile("instructions.txt")
	if err != nil {
		panic(err)
	}

	err, cbinsts := models.ParseCBInstructionsFile("instructions_cb.txt")
	if err != nil {
		panic(err)
	}

	f, err := os.Create("instructionset.go")
	if err != nil {
		panic(err)
	}

	// Normal
	b := bytes.NewBuffer(nil)
	for _, v := range insts {
		instructionTemplate.Execute(b, v)
	}

	built := b.String()

	// CB
	b = bytes.NewBuffer(nil)
	for _, v := range cbinsts {
		cbInstructionTemplate.Execute(b, v)
	}

	builtcb := b.String()

	o := bytes.NewBuffer(nil)

	instTpl.Execute(o, struct {
		INSTS     string
		CBINSTS   string
		Timestamp time.Time
	}{
		INSTS:     built,
		CBINSTS:   builtcb,
		Timestamp: time.Now(),
	})

	// Format the Code
	data, err := format.Source(o.Bytes())

	if err != nil {
		panic(err)
	}

	f.Write(data)

	f.Close()
}
