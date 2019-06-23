package cpu

import (
	"encoding/binary"
	"fmt"
	"github.com/racerxdl/goboy/cpu/generator"
	"strings"
)

type DisasmInstruction struct {
	Address     int
	Opcode      uint8
	Instruction string
	Argument    []byte
	Cycles      int
	ZSHC        string
	cbInst      bool
}

func (d DisasmInstruction) String() string {
	v := d.Instruction
	if strings.Contains(d.Instruction, "d8") { // Imediated 8 bit unsigned
		v = strings.ReplaceAll(v, "d8", fmt.Sprintf("$%02X", d.Argument[0]))
	} else if strings.Contains(d.Instruction, "d16") { // Imediated 16 bit unsigned
		v = strings.ReplaceAll(v, "d16", fmt.Sprintf("$%04X", binary.LittleEndian.Uint16(d.Argument)))
	} else if strings.Contains(d.Instruction, "a8") { // 8 bit unsigned relative to $FF00
		v = strings.ReplaceAll(v, "a8", fmt.Sprintf("%d", d.Argument[0]))
	} else if strings.Contains(d.Instruction, "a16") { // 16 bit unsigned address
		v = strings.ReplaceAll(v, "a16", fmt.Sprintf("$%04X", binary.LittleEndian.Uint16(d.Argument)))
	} else if strings.Contains(d.Instruction, "r8") { // 8 bit signed
		v = strings.ReplaceAll(v, "r8", fmt.Sprintf("%d", int8(d.Argument[0])))
	}

	return fmt.Sprintf("%04x: %s", d.Address, v)
}

func Disasm(offset int, data []byte) []DisasmInstruction {
	dis := make([]DisasmInstruction, 0)

	i := 0

	for i < len(data) {
		d := DisasmInstruction{
			Address: offset + i,
			cbInst:  false,
		}
		v := data[i]
		i++

		ins := generator.InstructionSet[v]
		args := make([]byte, 0)
		if ins.NumberOfBytes > 1 {
			if len(data) < i+ins.NumberOfBytes-1 {
				// Not enough bytes
				break
			}
			args = data[i : i+ins.NumberOfBytes-1]
		}
		i += ins.NumberOfBytes - 1

		if v == 0xCB { // Extended instructions
			cbi := generator.CBInstructionSet[args[0]]
			d.Instruction = cbi.Instruction
			d.Cycles = cbi.Cycles
			d.ZSHC = cbi.ZSHC
			dis = append(dis, d)
			d.Opcode = args[0]
			d.cbInst = true
			continue
		}

		d.Instruction = ins.Instruction
		d.Cycles = ins.Cycles
		d.ZSHC = ins.ZSHC
		d.Argument = args
		d.Opcode = v

		dis = append(dis, d)
	}

	return dis
}
