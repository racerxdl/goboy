package models

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const (
	FlagNotAffect = "-"
	FlagSetToZero = "0"
	FlagSetToOne  = "1"
)

type Instruction struct {
	Opcode        uint8
	Name          string
	Instruction   string
	Cycles        []int
	ZSHC          string
	NumberOfBytes int
	Zero          string
	Sub           string
	HalfCarry     string
	Carry         string
	TemplateName  string
	TemplateArgs  []string
}

var tplRgx = regexp.MustCompile(`(.*)\[(.*)\]`)

func InstructionFromTXT(opcode uint8, txtData string) Instruction {
	o := strings.Split(txtData, "|")
	//Name|Instruction|Cycles|ZSHC|Template[arguments]|NumberOfBytes
	cyclesS := strings.Split(o[2], "/")
	cycles := make([]int, 0)
	for _, v := range cyclesS {
		s, _ := strconv.ParseInt(v, 10, 32)
		cycles = append(cycles, int(s))
	}
	numberOfBytes, _ := strconv.ParseInt(o[5], 10, 32)
	ins := Instruction{
		Opcode:        opcode,
		Name:          o[0],
		Instruction:   o[1],
		Cycles:        cycles,
		NumberOfBytes: int(numberOfBytes),
		ZSHC:          o[3],
		Zero:          string(o[3][0]),
		Sub:           string(o[3][1]),
		HalfCarry:     string(o[3][2]),
		Carry:         string(o[3][3]),
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

func ParseInstructionsFile(filename string) (error, []Instruction) {
	ins := make([]Instruction, 256) // Should have 256

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
		ins[i] = InstructionFromTXT(i, l)
		i++
	}

	return nil, ins
}
