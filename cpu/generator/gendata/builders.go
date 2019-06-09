package gendata

import (
	"text/template"
)

var InstructionsFileTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// {{ .Timestamp }}
package cpu

{{ .Instructions }}

var GBInstructions = []GBInstruction{
    
}

`))

/*

var GBInstructions = []GBInstruction{
    // region 0x00 Group
    gbNOP,
    gbLDBCnn,
    gbLDBmC,
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
    gbLDDmE,
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
    gbRSTXX0x00,
    gbRETZ,
    gbRET,
    gbJPZnn,
    gbCBCall,
    gbCALLZnn,
    gbCALLnn,
    gbADCn,
    gbRSTXX0x08,
    // endregion
    // region 0xD0 Group
    gbRETNC,
    gbPOPDE,
    gbJPNCnn,
    gbNOPWARN0xD3,
    gbCALLNCnn,
    gbPUSHDE,
    gbSUBn,
    gbRSTXX0x10,
    gbRETC,
    gbRETI,
    gbJPCnn,
    gbNOPWARN0xDB,
    gbCALLCnn,
    gbNOPWARN0xDD,
    gbSBCn,
    gbRSTXX0x18,
    // endregion
    // region 0xE0 Group
    gbLDIOnA,
    gbPOPHL,
    gbLDIOCA,
    gbNOPWARN0xE3,
    gbNOPWARN0xE4,
    gbPUSHHL,
    gbANDn,
    gbRSTXX0x20,
    gbADDSPn,
    gbJPHL,
    gbLDmmA,
    gbNOPWARN0xEB,
    gbNOPWARN0xEC,
    gbNOPWARN0xED,
    gbXORn,
    gbRSTXX0x28,
    // endregion
    // region 0xF0 Group
    gbLDAIOn,
    gbPOPAF,
    gbLDAIOC,
    gbDI,
    gbNOPWARN0xF4,
    gbPUSHAF,
    gbORn,
    gbRSTXX0x30,
    gbLDHLSPn,
    gbLDHLSPr,
    gbLDmmA,
    gbEI,
    gbNOPWARN0xFC,
    gbNOPWARN0xFD,
    gbCPn,
    gbRSTXX0x38,
    // endregion
}
*/
