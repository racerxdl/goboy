package gendata

import (
	"text/template"
)

var AllRegisters = []string{
	"A", "B", "C", "D", "E", "H", "L",
}

var CBInstructionsFileTemplate = template.Must(template.New("CBIns").Parse(`package cpu

func gbCBCall(cpu *Core) {
    ins := cpu.Memory.ReadByte(cpu.Registers.PC)
    cpu.Registers.PC++
    CBInstructions[ins](cpu)
}

{{ .Instructions }}

var CBInstructions = []GBInstruction{
   //region CB00 Group
   cbRLCrB,
   cbRLCrC,
   cbRLCrD,
   cbRLCrE,
   cbRLCrH,
   cbRLCrL,
   cbRLCHL,
   cbRLCrA,
   cbRRCrB,
   cbRRCrC,
   cbRRCrD,
   cbRRCrE,
   cbRRCrH,
   cbRRCrL,
   cbRRCHL,
   cbRRCrA,
   //endregion
   //region CB10 Group
   cbRLrB,
   cbRLrC,
   cbRLrD,
   cbRLrE,
   cbRLrH,
   cbRLrL,
   cbRLHL,
   cbRLrA,
   cbRRrB,
   cbRRrC,
   cbRRrD,
   cbRRrE,
   cbRRrH,
   cbRRrL,
   cbRRHL,
   cbRRrA,
   //endregion
   //region CB20 Group
   cbSLArB,
   cbSLArC,
   cbSLArD,
   cbSLArE,
   cbSLArH,
   cbSLArL,
   cbSLAHL,
   cbSLArA,
   cbSRArB,
   cbSRArC,
   cbSRArD,
   cbSRArE,
   cbSRArH,
   cbSRArL,
   cbSRAHL,
   cbSRArA,
   //endregion
   //region CB30 Group
   cbSWAPrB,
   cbSWAPrC,
   cbSWAPrD,
   cbSWAPrE,
   cbSWAPrH,
   cbSWAPrL,
   cbSWAPHL,
   cbSWAPrA,
   cbSRLrB,
   cbSRLrC,
   cbSRLrD,
   cbSRLrE,
   cbSRLrH,
   cbSRLrL,
   cbSRLHL,
   cbSRLrA,
   //endregion
   //region CB40 Group
   cbBIT0B,
   cbBIT0C,
   cbBIT0D,
   cbBIT0E,
   cbBIT0H,
   cbBIT0L,
   cbBITm0,
   cbBIT0A,
   cbBIT1B,
   cbBIT1C,
   cbBIT1D,
   cbBIT1E,
   cbBIT1H,
   cbBIT1L,
   cbBITm1,
   cbBIT1A,
   //endregion
   //region CB50 Group
   cbBIT2B,
   cbBIT2C,
   cbBIT2D,
   cbBIT2E,
   cbBIT2H,
   cbBIT2L,
   cbBITm2,
   cbBIT2A,
   cbBIT3B,
   cbBIT3C,
   cbBIT3D,
   cbBIT3E,
   cbBIT3H,
   cbBIT3L,
   cbBITm3,
   cbBIT3A,
   //endregion
   //region CB60 Group
   cbBIT4B,
   cbBIT4C,
   cbBIT4D,
   cbBIT4E,
   cbBIT4H,
   cbBIT4L,
   cbBITm4,
   cbBIT4A,
   cbBIT5B,
   cbBIT5C,
   cbBIT5D,
   cbBIT5E,
   cbBIT5H,
   cbBIT5L,
   cbBITm5,
   cbBIT5A,
   //endregion
   //region CB70 Group
   cbBIT6B,
   cbBIT6C,
   cbBIT6D,
   cbBIT6E,
   cbBIT6H,
   cbBIT6L,
   cbBITm6,
   cbBIT6A,
   cbBIT7B,
   cbBIT7C,
   cbBIT7D,
   cbBIT7E,
   cbBIT7H,
   cbBIT7L,
   cbBITm7,
   cbBIT7A,
   //endregion
   //region CB80 Group
   cbRES0B,
   cbRES0C,
   cbRES0D,
   cbRES0E,
   cbRES0H,
   cbRES0L,
   cbRESHL0,
   cbRES0A,
   cbRES1B,
   cbRES1C,
   cbRES1D,
   cbRES1E,
   cbRES1H,
   cbRES1L,
   cbRESHL1,
   cbRES1A,
   //endregion
   //region CB90 Group
   cbRES2B,
   cbRES2C,
   cbRES2D,
   cbRES2E,
   cbRES2H,
   cbRES2L,
   cbRESHL2,
   cbRES2A,
   cbRES3B,
   cbRES3C,
   cbRES3D,
   cbRES3E,
   cbRES3H,
   cbRES3L,
   cbRESHL3,
   cbRES3A,
   //endregion
   //region CBA0 Group
   cbRES4B,
   cbRES4C,
   cbRES4D,
   cbRES4E,
   cbRES4H,
   cbRES4L,
   cbRESHL4,
   cbRES4A,
   cbRES6B,
   cbRES6C,
   cbRES6D,
   cbRES6E,
   cbRES6H,
   cbRES6L,
   cbRESHL6,
   cbRES6A,
   //endregion
   //region CBB0 Group
   cbRES6B,
   cbRES6C,
   cbRES6D,
   cbRES6E,
   cbRES6H,
   cbRES6L,
   cbRESHL6,
   cbRES6A,
   cbRES7B,
   cbRES7C,
   cbRES7D,
   cbRES7E,
   cbRES7H,
   cbRES7L,
   cbRESHL7,
   cbRES7A,
   //endregion
   //region CBC0 Group
   cbSET0B,
   cbSET0C,
   cbSET0D,
   cbSET0E,
   cbSET0H,
   cbSET0L,
   cbSETHL0,
   cbSET0A,
   cbSET1B,
   cbSET1C,
   cbSET1D,
   cbSET1E,
   cbSET1H,
   cbSET1L,
   cbSETHL1,
   cbSET1A,
   //endregion
   //region CBD0 Group
   cbSET2B,
   cbSET2C,
   cbSET2D,
   cbSET2E,
   cbSET2H,
   cbSET2L,
   cbSETHL2,
   cbSET2A,
   cbSET3B,
   cbSET3C,
   cbSET3D,
   cbSET3E,
   cbSET3H,
   cbSET3L,
   cbSETHL3,
   cbSET3A,
   //endregion
   //region CBE0 Group
   cbSET4B,
   cbSET4C,
   cbSET4D,
   cbSET4E,
   cbSET4H,
   cbSET4L,
   cbSETHL4,
   cbSET4A,
   cbSET5B,
   cbSET5C,
   cbSET5D,
   cbSET5E,
   cbSET5H,
   cbSET5L,
   cbSETHL5,
   cbSET5A,
   //endregion
   //region CBF0 Group
   cbSET6B,
   cbSET6C,
   cbSET6D,
   cbSET6E,
   cbSET6H,
   cbSET6L,
   cbSETHL6,
   cbSET6A,
   cbSET7B,
   cbSET7C,
   cbSET7D,
   cbSET7E,
   cbSET7H,
   cbSET7L,
   cbSETHL7,
   cbSET7A,
   //endregion
}
`))

var InstructionsFileTemplate = template.Must(template.New("Ins").Parse(`// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// {{ .Timestamp }}
package cpu

{{ .Instructions }}


var GBInstructions = []GBInstruction{
   // region 0x00 Group
   gbNOP,
   gbLDBCnn,
   gbLDBCmA,
   gbINCBC,
   gbINCrB,
   gbDECrB,
   gbLDrnB,
   gbRLCA,
   gbLDmmSP,
   gbADDHLBC,
   gbLDABCm,
   gbDECBC,
   gbINCrC,
   gbDECrC,
   gbLDrnC,
   gbRRCA,
   // endregion
   // region 0x10 Group
   gbStop,
   gbLDDEnn,
   gbLDDEmA,
   gbINCDE,
   gbINCrD,
   gbDECrD,
   gbLDrnD,
   gbRLA,
   gbJRn,
   gbADDHLDE,
   gbLDADEm,
   gbDECDE,
   gbINCrE,
   gbDECrE,
   gbLDrnE,
   gbRRA,
   // endregion
   // region 0x20 Group
   gbJRNZn,
   gbLDHLnn,
   gbLDHLIA,
   gbINCHL,
   gbINCrH,
   gbDECrH,
   gbLDrnH,
   gbDAA,
   gbJRZn,
   gbADDHLHL,
   gbLDAHLI,
   gbDECHL,
   gbINCrL,
   gbDECrL,
   gbLDrnL,
   gbCPL,
   // endregion
   // region 0x30 Group
   gbJRNCn,
   gbLDSPnn,
   gbLDHLDA,
   gbINCSP,
   gbINCHLm,
   gbDECHLm,
   gbLDHLmn,
   gbSCF,
   gbJRCn,
   gbADDHLSP,
   gbLDAHLD,
   gbDECSP,
   gbINCrA,
   gbDECrA,
   gbLDrnA,
   gbCCF,
   // endregion
   // region 0x40 Group
   gbLDrrBB,
   gbLDrrBC,
   gbLDrrBD,
   gbLDrrBE,
   gbLDrrBH,
   gbLDrrBL,
   gbLDrHLmB,
   gbLDrrBA,
   gbLDrrCB,
   gbLDrrCC,
   gbLDrrCD,
   gbLDrrCE,
   gbLDrrCH,
   gbLDrrCL,
   gbLDrHLmC,
   gbLDrrCA,
   // endregion
   // region 0x50 Group
   gbLDrrDB,
   gbLDrrDC,
   gbLDrrDD,
   gbLDrrDE,
   gbLDrrDH,
   gbLDrrDL,
   gbLDrHLmD,
   gbLDrrDA,
   gbLDrrEB,
   gbLDrrEC,
   gbLDrrED,
   gbLDrrEE,
   gbLDrrEH,
   gbLDrrEL,
   gbLDrHLmE,
   gbLDrrEA,
   // endregion
   // region 0x60 Group
   gbLDrrHB,
   gbLDrrHC,
   gbLDrrHD,
   gbLDrrHE,
   gbLDrrHH,
   gbLDrrHL,
   gbLDrHLmH,
   gbLDrrHA,
   gbLDrrLB,
   gbLDrrLC,
   gbLDrrLD,
   gbLDrrLE,
   gbLDrrLH,
   gbLDrrLL,
   gbLDrHLmL,
   gbLDrrLA,
   // endregion
   // region 0x70 Group
   gbLDHLmrB,
   gbLDHLmrC,
   gbLDHLmrD,
   gbLDHLmrE,
   gbLDHLmrH,
   gbLDHLmrL,
   gbHALT,
   gbLDHLmrA,
   gbLDrrAB,
   gbLDrrAC,
   gbLDrrAD,
   gbLDrrAE,
   gbLDrrAH,
   gbLDrrAL,
   gbLDrHLmA,
   gbLDrrAA,
   // endregion
   // region 0x80 Group
   gbADDrB,
   gbADDrC,
   gbADDrD,
   gbADDrE,
   gbADDrH,
   gbADDrL,
   gbADDHL,
   gbADDrA,
   gbADCrB,
   gbADCrC,
   gbADCrD,
   gbADCrE,
   gbADCrH,
   gbADCrL,
   gbADCHL,
   gbADCrA,
   // endregion
   // region 0x90 Group
   gbSUBrB,
   gbSUBrC,
   gbSUBrD,
   gbSUBrE,
   gbSUBrH,
   gbSUBrL,
   gbSUBHL,
   gbSUBrA,
   gbSBCrB,
   gbSBCrC,
   gbSBCrD,
   gbSBCrE,
   gbSBCrH,
   gbSBCrL,
   gbSBCHL,
   gbSBCrA,
   // endregion
   // region 0xA0 Group
   gbANDrB,
   gbANDrC,
   gbANDrD,
   gbANDrE,
   gbANDrH,
   gbANDrL,
   gbANDHL,
   gbANDrA,
   gbXORrB,
   gbXORrC,
   gbXORrD,
   gbXORrE,
   gbXORrH,
   gbXORrL,
   gbXORHL,
   gbXORrA,
   // endregion
   // region 0xB0 Group
   gbORrB,
   gbORrC,
   gbORrD,
   gbORrE,
   gbORrH,
   gbORrL,
   gbORHL,
   gbORrA,
   gbCPrB,
   gbCPrC,
   gbCPrD,
   gbCPrE,
   gbCPrH,
   gbCPrL,
   gbCPHL,
   gbCPrA,
   // endregion
   // region 0xC0 Group
   gbRETNZ,
   gbPOPBC,
   gbJPNZnn,
   gbJPnn,
   gbCALLNZnn,
   gbPUSHBC,
   gbADDn,
   func (cpu *Core) { gbRSTXX(cpu, 0x00) },
   gbRETZ,
   gbRET,
   gbJPZnn,
   gbCBCall,
   gbCALLZnn,
   gbCALLnn,
   gbADCn,
   func (cpu *Core) { gbRSTXX(cpu, 0x08) },
   // endregion
   // region 0xD0 Group
   gbRETNC,
   gbPOPDE,
   gbJPNCnn,
   func (cpu *Core) { gbNOPWARN(cpu, 0xD3) },
   gbCALLNCnn,
   gbPUSHDE,
   gbSUBn,
   func (cpu *Core) { gbRSTXX(cpu, 0x10) },
   gbRETC,
   gbRETI,
   gbJPCnn,
   func (cpu *Core) { gbNOPWARN(cpu, 0xDB) },
   gbCALLCnn,
   func (cpu *Core) { gbNOPWARN(cpu, 0xDD) },
   gbSBCn,
   func (cpu *Core) { gbRSTXX(cpu, 0x18) },
   // endregion
   // region 0xE0 Group
   gbLDIOnA,
   gbPOPHL,
   gbLDIOCA,
   func (cpu *Core) { gbNOPWARN(cpu, 0xE3) },
   func (cpu *Core) { gbNOPWARN(cpu, 0xE4) },
   gbPUSHHL,
   gbANDn,
   func (cpu *Core) { gbRSTXX(cpu, 0x20) },
   gbADDSPn,
   gbJPHL,
   gbLDmmA,
   func (cpu *Core) { gbNOPWARN(cpu, 0xEB) },
   func (cpu *Core) { gbNOPWARN(cpu, 0xEC) },
   func (cpu *Core) { gbNOPWARN(cpu, 0xED) },
   gbXORn,
   func (cpu *Core) { gbRSTXX(cpu, 0x28) },
   // endregion
   // region 0xF0 Group
   gbLDAIOn,
   gbPOPAF,
   gbLDAIOC,
   gbDI,
   func (cpu *Core) { gbNOPWARN(cpu, 0xF4) },
   gbPUSHAF,
   gbORn,
   func (cpu *Core) { gbRSTXX(cpu, 0x30) },
   gbLDHLSPn,
   gbLDHLSPr,
   gbLDAmm,
   gbEI,
   func (cpu *Core) { gbNOPWARN(cpu, 0xFC) },
   func (cpu *Core) { gbNOPWARN(cpu, 0xFD) },
   gbCPn,
   func (cpu *Core) { gbRSTXX(cpu, 0x38) },
   // endregion
}
`))
