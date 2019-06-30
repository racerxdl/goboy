package models

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type CBInstruction struct {
	Opcode       uint8
	Name         string
	Instruction  string
	Cycles       []int
	ZSHC         string
	Zero         string
	Sub          string
	HalfCarry    string
	Carry        string
	TemplateName string
	TemplateArgs []string
}

func CBInstructionFromTXT(opcode uint8, txtData string) CBInstruction {
	if txtData[0] == '|' {
		txtData = txtData[1:]
	}

	o := strings.Split(txtData, "|")
	//|Name|Instruction|Cycles|ZSHC|Template[arguments]
	cyclesS := strings.Split(o[2], "/")
	cycles := make([]int, 0)
	for _, v := range cyclesS {
		s, _ := strconv.ParseInt(v, 10, 32)
		cycles = append(cycles, int(s))
	}
	ins := CBInstruction{
		Opcode:      opcode,
		Name:        o[0],
		Instruction: o[1],
		Cycles:      cycles,
		ZSHC:        o[3],
		Zero:        string(o[3][0]),
		Sub:         string(o[3][1]),
		HalfCarry:   string(o[3][2]),
		Carry:       string(o[3][3]),
	}

	m := tplRgx.FindAllStringSubmatch(o[4], -1)[0]

	ins.TemplateName = m[1]

	args := strings.Split(m[2], ",")
	for i, v := range args {
		args[i] = strings.Trim(strings.ReplaceAll(v, "\"", ""), " \n\r")
		if len(args[i]) > 2 && args[i][:2] == "0x" {
			// Parse to int
			n, _ := strconv.ParseInt(args[i][2:], 16, 32)
			args[i] = fmt.Sprintf("%d", n)
		}
	}

	ins.TemplateArgs = args

	return ins
}

func ParseCBInstructionsFile(filename string) (error, []CBInstruction) {
	ins := make([]CBInstruction, 256) // Should have 256

	d, err := ioutil.ReadFile(filename)

	if err != nil {
		return err, nil
	}

	data := string(d)
	lines := strings.Split(data, "\n")

	i := uint8(0)

	for _, l := range lines {
		l = strings.Trim(l, " \r\n")
		if len(l) == 0 || l[:2] == "//" {
			continue
		}
		ins[i] = CBInstructionFromTXT(i, l)
		i++
	}

	return nil, ins
}
