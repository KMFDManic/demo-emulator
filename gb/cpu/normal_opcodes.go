package cpu

func opcode0x00() {
	// NOP
}

func opcode0x01() {
	// LD BC,nn
	opcodesLDAddrToReg(bc.GetLowReg(), pc.GetValue())
	pc.Increment()
	opcodesLDAddrToReg(bc.GetHighReg(), pc.GetValue())
	pc.Increment()
}

func opcode0x02() {
	// LD (BC),A
	opcodesLDValueToAddr(bc.GetValue(), af.GetHigh())
}

func opcode0x03() {
	// INC BC
	bc.Increment()
}

func opcode0x04() {
	// INC B
	opcodesINC(bc.GetHighReg())
}

func opcode0x05() {
	// DEC B
	opcodesDEC(bc.GetHighReg())
}

func opcode0x06() {
	// LD B,n
	opcodesLDAddrToReg(bc.GetHighReg(), pc.GetValue())
	pc.Increment()
}

func opcode0x07() {
	// RLCA
	opcodesRLCA(af.GetHighReg())
}

func opcode0x08() {
	// LD (nn),SP
	l := uint16(mem.Read(pc.GetValue()))
	pc.Increment()
	h := uint16(mem.Read(pc.GetValue()))
	pc.Increment()
	address := ((h << 8) + l)
	mem.Write(address, sp.GetLow())
	mem.Write(address+1, sp.GetHigh())
}

func opcode0x09() {
	// ADD HL,BC
	opcodesADDHL(bc.GetValue())
}

func opcode0x0A() {
	// LD A,(BC)
	opcodesLDAddrToReg(af.GetHighReg(), bc.GetValue())
}

func opcode0x0B() {
	// DEC BC
	bc.Decrement()
}

func opcode0x0C() {
	// INC C
	opcodesINC(bc.GetLowReg())
}

func opcode0x0D() {
	// DEC C
	opcodesDEC(bc.GetLowReg())
}

func opcode0x0E() {
	// LD C,n
	opcodesLDAddrToReg(bc.GetLowReg(), pc.GetValue())
	pc.Increment()
}

func opcode0x0F() {
	// RRCA
	opcodesRRCA(af.GetHighReg())
}

func opcode0x10() {
	// STOP
	pc.Increment()
}

func opcode0x11() {
	// LD DE,nn
	opcodesLDAddrToReg(de.GetLowReg(), pc.GetValue())
	pc.Increment()
	opcodesLDAddrToReg(de.GetHighReg(), pc.GetValue())
	pc.Increment()
}

func opcode0x12() {
	// LD (DE),A
	opcodesLDValueToAddr(de.GetValue(), af.GetHigh())
}

func opcode0x13() {
	// INC DE
	de.Increment()
}

func opcode0x14() {
	// INC D
	opcodesINC(de.GetHighReg())
}

func opcode0x15() {
	// DEC D
	opcodesDEC(de.GetHighReg())
}

func opcode0x16() {
	// LD D,n
	opcodesLDAddrToReg(de.GetHighReg(), pc.GetValue())
	pc.Increment()
}

func opcode0x17() {
	// RLA
	opcodesRLA(af.GetHighReg())
}

func opcode0x18() {
	// JR n
	offset := int8(mem.Read(pc.GetValue()))
	value := int(pc.GetValue()) + 1 + int(offset)
	pc.SetValue(uint16(value))
}

func opcode0x19() {
	// ADD HL,DE
	opcodesADDHL(de.GetValue())
}

func opcode0x1A() {
	// LD A,(DE)
	opcodesLDAddrToReg(af.GetHighReg(), de.GetValue())
}

func opcode0x1B() {
	// DEC DE
	de.Decrement()
}

func opcode0x1C() {
	// INC E
	opcodesINC(de.GetLowReg())
}

func opcode0x1D() {
	// DEC E
	opcodesDEC(de.GetLowReg())
}

func opcode0x1E() {
	// LD E,n
	opcodesLDAddrToReg(de.GetLowReg(), pc.GetValue())
	pc.Increment()
}

func opcode0x1F() {
	// RRA
	opcodesRRA(af.GetHighReg())
}

func opcode0x20() {
	// JR NZ,n
	if !isSetFlag(flagZero) {
		offset := int8(mem.Read(pc.GetValue()))
		value := int(pc.GetValue()) + 1 + int(offset)
		pc.SetValue(uint16(value))
		branchTaken = true
	} else {
		pc.Increment()
	}
}

func opcode0x21() {
	// LD HL,nn
	opcodesLDAddrToReg(hl.GetLowReg(), pc.GetValue())
	pc.Increment()
	opcodesLDAddrToReg(hl.GetHighReg(), pc.GetValue())
	pc.Increment()
}

func opcode0x22() {
	// LD (HLI),A
	opcodesLDValueToAddr(hl.GetValue(), af.GetHigh())
	hl.Increment()
}

func opcode0x23() {
	// INC HL
	hl.Increment()
}

func opcode0x24() {
	// INC H
	opcodesINC(hl.GetHighReg())
}

func opcode0x25() {
	// DEC H
	opcodesDEC(hl.GetHighReg())
}

func opcode0x26() {
	// LD H,n
	opcodesLDAddrToReg(hl.GetHighReg(), pc.GetValue())
	pc.Increment()
}

func opcode0x27() {
	// DAA
	a := uint16(af.GetHigh())

	if !isSetFlag(flagNegative) {
		if isSetFlag(flagHalf) || ((a & 0xF) > 9) {
			a += 0x06
		}
		if isSetFlag(flagCarry) || (a > 0x9F) {
			a += 0x60
		}
	} else {
		if isSetFlag(flagHalf) {
			a = (a - 6) & 0xFF
		}
		if isSetFlag(flagCarry) {
			a -= 0x60
		}
	}

	resetFlag(flagHalf)
	resetFlag(flagZero)

	if (a & 0x100) == 0x100 {
		setFlag(flagCarry)
	}

	a &= 0xFF

	setZeroFlagFromResult(uint8(a))

	af.SetHigh(uint8(a))
}

func opcode0x28() {
	// JR Z,n
	if isSetFlag(flagZero) {
		offset := int8(mem.Read(pc.GetValue()))
		value := int(pc.GetValue()) + 1 + int(offset)
		pc.SetValue(uint16(value))
		branchTaken = true
	} else {
		pc.Increment()
	}
}

func opcode0x29() {
	// ADD HL,HL
	opcodesADDHL(hl.GetValue())
}

func opcode0x2A() {
	// LD A,(HLI)
	opcodesLDAddrToReg(af.GetHighReg(), hl.GetValue())
	hl.Increment()
}

func opcode0x2B() {
	// DEC HL
	hl.Decrement()
}

func opcode0x2C() {
	// INC L
	opcodesINC(hl.GetLowReg())
}

func opcode0x2D() {
	// DEC L
	opcodesDEC(hl.GetLowReg())
}

func opcode0x2E() {
	// LD L,n
	opcodesLDAddrToReg(hl.GetLowReg(), pc.GetValue())
	pc.Increment()
}

func opcode0x2F() {
	// CPL
	af.SetHigh(^af.GetHigh())
	setFlag(flagHalf)
	setFlag(flagNegative)
}

func opcode0x30() {
	// JR NC,n
	if !isSetFlag(flagCarry) {
		offset := int8(mem.Read(pc.GetValue()))
		value := int(pc.GetValue()) + 1 + int(offset)
		pc.SetValue(uint16(value))
		branchTaken = true
	} else {
		pc.Increment()
	}
}

func opcode0x31() {
	// LD SP,nn
	sp.SetLow(mem.Read(pc.GetValue()))
	pc.Increment()
	sp.SetHigh(mem.Read(pc.GetValue()))
	pc.Increment()
}

func opcode0x32() {
	// LD (HLD), A
	opcodesLDValueToAddr(hl.GetValue(), af.GetHigh())
	hl.Decrement()
}

func opcode0x33() {
	// INC SP
	sp.Increment()
}

func opcode0x34() {
	// INC (HL)
	opcodesINCHL()
}

func opcode0x35() {
	// DEC (HL)
	opcodesDECHL()
}

func opcode0x36() {
	// LD (HL),n
	mem.Write(hl.GetValue(), mem.Read(pc.GetValue()))
	pc.Increment()
}

func opcode0x37() {
	// SCF
	setFlag(flagCarry)
	resetFlag(flagHalf)
	resetFlag(flagNegative)
}

func opcode0x38() {
	// JR C,n
	if isSetFlag(flagCarry) {
		offset := int8(mem.Read(pc.GetValue()))
		value := int(pc.GetValue()) + 1 + int(offset)
		pc.SetValue(uint16(value))
		branchTaken = true
	} else {
		pc.Increment()
	}
}

func opcode0x39() {
	// ADD HL,SP
	opcodesADDHL(sp.GetValue())
}

func opcode0x3A() {
	// LD A,(HLD)
	opcodesLDAddrToReg(af.GetHighReg(), hl.GetValue())
	hl.Decrement()
}

func opcode0x3B() {
	// DEC SP
	sp.Decrement()
}

func opcode0x3C() {
	// INC A
	opcodesINC(af.GetHighReg())
}

func opcode0x3D() {
	// DEC A
	opcodesDEC(af.GetHighReg())

}

func opcode0x3E() {
	// LD A,n
	opcodesLDAddrToReg(af.GetHighReg(), pc.GetValue())
	pc.Increment()
}

func opcode0x3F() {
	// CCF
	flipFlag(flagCarry)
	resetFlag(flagHalf)
	resetFlag(flagNegative)
}

func opcode0x40() {
	// LD B,B
	opcodesLDValueToReg(bc.GetHighReg(), bc.GetHigh())
}

func opcode0x41() {
	// LD B,C
	opcodesLDValueToReg(bc.GetHighReg(), bc.GetLow())
}

func opcode0x42() {
	// LD B,D
	opcodesLDValueToReg(bc.GetHighReg(), de.GetHigh())
}

func opcode0x43() {
	// LD B,E
	opcodesLDValueToReg(bc.GetHighReg(), de.GetLow())
}

func opcode0x44() {
	// LD B,H
	opcodesLDValueToReg(bc.GetHighReg(), hl.GetHigh())
}

func opcode0x45() {
	// LD B,L
	opcodesLDValueToReg(bc.GetHighReg(), hl.GetLow())
}

func opcode0x46() {
	// LD B,(HL)
	opcodesLDAddrToReg(bc.GetHighReg(), hl.GetValue())
}

func opcode0x47() {
	// LD B,A
	opcodesLDValueToReg(bc.GetHighReg(), af.GetHigh())
}

func opcode0x48() {
	// LD C,B
	opcodesLDValueToReg(bc.GetLowReg(), bc.GetHigh())
}

func opcode0x49() {
	// LD C,C
	opcodesLDValueToReg(bc.GetLowReg(), bc.GetLow())
}

func opcode0x4A() {
	// LD C,D
	opcodesLDValueToReg(bc.GetLowReg(), de.GetHigh())
}

func opcode0x4B() {
	// LD C,E
	opcodesLDValueToReg(bc.GetLowReg(), de.GetLow())
}

func opcode0x4C() {
	// LD C,H
	opcodesLDValueToReg(bc.GetLowReg(), hl.GetHigh())
}

func opcode0x4D() {
	// LD C,L
	opcodesLDValueToReg(bc.GetLowReg(), hl.GetLow())
}

func opcode0x4E() {
	// LD C,(HL)
	opcodesLDAddrToReg(bc.GetLowReg(), hl.GetValue())
}

func opcode0x4F() {
	// LD C,A
	opcodesLDValueToReg(bc.GetLowReg(), af.GetHigh())
}

func opcode0x50() {
	// LD D,B
	opcodesLDValueToReg(de.GetHighReg(), bc.GetHigh())
}

func opcode0x51() {
	// LD D,C
	opcodesLDValueToReg(de.GetHighReg(), bc.GetLow())
}

func opcode0x52() {
	// LD D,D
	opcodesLDValueToReg(de.GetHighReg(), de.GetHigh())
}

func opcode0x53() {
	// LD D,E
	opcodesLDValueToReg(de.GetHighReg(), de.GetLow())
}

func opcode0x54() {
	// LD D,H
	opcodesLDValueToReg(de.GetHighReg(), hl.GetHigh())
}

func opcode0x55() {
	// LD D,L
	opcodesLDValueToReg(de.GetHighReg(), hl.GetLow())
}

func opcode0x56() {
	// LD D,(HL)
	opcodesLDAddrToReg(de.GetHighReg(), hl.GetValue())
}

func opcode0x57() {
	// LD D,A
	opcodesLDValueToReg(de.GetHighReg(), af.GetHigh())
}

func opcode0x58() {
	// LD E,B
	opcodesLDValueToReg(de.GetLowReg(), bc.GetHigh())
}

func opcode0x59() {
	// LD E,C
	opcodesLDValueToReg(de.GetLowReg(), bc.GetLow())
}

func opcode0x5A() {
	// LD E,D
	opcodesLDValueToReg(de.GetLowReg(), de.GetHigh())
}

func opcode0x5B() {
	// LD E,E
	opcodesLDValueToReg(de.GetLowReg(), de.GetLow())
}

func opcode0x5C() {
	// LD E,H
	opcodesLDValueToReg(de.GetLowReg(), hl.GetHigh())
}

func opcode0x5D() {
	// LD E,L
	opcodesLDValueToReg(de.GetLowReg(), hl.GetLow())
}

func opcode0x5E() {
	// LD E,(HL)
	opcodesLDAddrToReg(de.GetLowReg(), hl.GetValue())
}

func opcode0x5F() {
	// LD E,A
	opcodesLDValueToReg(de.GetLowReg(), af.GetHigh())
}

func opcode0x60() {
	// LD H,B
	opcodesLDValueToReg(hl.GetHighReg(), bc.GetHigh())
}

func opcode0x61() {
	// LD H,C
	opcodesLDValueToReg(hl.GetHighReg(), bc.GetLow())
}

func opcode0x62() {
	// LD H,D
	opcodesLDValueToReg(hl.GetHighReg(), de.GetHigh())
}

func opcode0x63() {
	// LD H,E
	opcodesLDValueToReg(hl.GetHighReg(), de.GetLow())
}

func opcode0x64() {
	// LD H,H
	opcodesLDValueToReg(hl.GetHighReg(), hl.GetHigh())
}

func opcode0x65() {
	// LD H,L
	opcodesLDValueToReg(hl.GetHighReg(), hl.GetLow())
}

func opcode0x66() {
	// LD H,(HL)
	opcodesLDAddrToReg(hl.GetHighReg(), hl.GetValue())
}

func opcode0x67() {
	// LD H,A
	opcodesLDValueToReg(hl.GetHighReg(), af.GetHigh())
}

func opcode0x68() {
	// LD L,B
	opcodesLDValueToReg(hl.GetLowReg(), bc.GetHigh())
}

func opcode0x69() {
	// LD L,C
	opcodesLDValueToReg(hl.GetLowReg(), bc.GetLow())
}

func opcode0x6A() {
	// LD L,D
	opcodesLDValueToReg(hl.GetLowReg(), de.GetHigh())
}

func opcode0x6B() {
	// LD L,E
	opcodesLDValueToReg(hl.GetLowReg(), de.GetLow())
}

func opcode0x6C() {
	// LD L,H
	opcodesLDValueToReg(hl.GetLowReg(), hl.GetHigh())
}

func opcode0x6D() {
	// LD L,L
	opcodesLDValueToReg(hl.GetLowReg(), hl.GetLow())
}

func opcode0x6E() {
	// LD L,(HL)
	opcodesLDAddrToReg(hl.GetLowReg(), hl.GetValue())
}

func opcode0x6F() {
	// LD L,A
	opcodesLDValueToReg(hl.GetLowReg(), af.GetHigh())
}

func opcode0x70() {
	// LD (HL),B
	opcodesLDValueToAddr(hl.GetValue(), bc.GetHigh())
}

func opcode0x71() {
	// LD (HL),C
	opcodesLDValueToAddr(hl.GetValue(), bc.GetLow())
}

func opcode0x72() {
	// LD (HL),D
	opcodesLDValueToAddr(hl.GetValue(), de.GetHigh())
}

func opcode0x73() {
	// LD (HL),E
	opcodesLDValueToAddr(hl.GetValue(), de.GetLow())
}

func opcode0x74() {
	// LD (HL),H
	opcodesLDValueToAddr(hl.GetValue(), hl.GetHigh())
}

func opcode0x75() {
	// LD (HL),L
	opcodesLDValueToAddr(hl.GetValue(), hl.GetLow())
}

func opcode0x76() {
	// HALT
	if imeCycles > 0 {
		// If EI is pending interrupts are triggered before Halt
		imeCycles = 0
		ime = true
		pc.Decrement()
	} else {
		ifreg := mem.Read(0xFF0F)
		iereg := mem.Read(0xFFFF)

		halted = true

		// if there is an interrupt pending and
		// the cpu is halted it fails to advance the PC register
		// once the cpu resumes operation
		// this bug is present in all the original DMGs
		if !ime && ((ifreg & iereg & 0x1F) != 0) {
			skipPCBug = true
		}
	}
}

func opcode0x77() {
	// LD (HL),A
	opcodesLDValueToAddr(hl.GetValue(), af.GetHigh())
}

func opcode0x78() {
	// LD A,B
	opcodesLDValueToReg(af.GetHighReg(), bc.GetHigh())
}

func opcode0x79() {
	// LD A,C
	opcodesLDValueToReg(af.GetHighReg(), bc.GetLow())
}

func opcode0x7A() {
	// LD A,D
	opcodesLDValueToReg(af.GetHighReg(), de.GetHigh())
}

func opcode0x7B() {
	// LD A,E
	opcodesLDValueToReg(af.GetHighReg(), de.GetLow())
}

func opcode0x7C() {
	// LD A,H
	opcodesLDValueToReg(af.GetHighReg(), hl.GetHigh())
}

func opcode0x7D() {
	// LD A,L
	opcodesLDValueToReg(af.GetHighReg(), hl.GetLow())
}

func opcode0x7E() {
	// LD A,(HL)
	opcodesLDAddrToReg(af.GetHighReg(), hl.GetValue())
}

func opcode0x7F() {
	// LD A,A
	opcodesLDValueToReg(af.GetHighReg(), af.GetHigh())
}

func opcode0x80() {
	// ADD A,B
	opcodesADD(bc.GetHigh())
}

func opcode0x81() {
	// ADD A,C
	opcodesADD(bc.GetLow())
}

func opcode0x82() {
	// ADD A,D
	opcodesADD(de.GetHigh())
}

func opcode0x83() {
	// ADD A,E
	opcodesADD(de.GetLow())
}

func opcode0x84() {
	// ADD A,H
	opcodesADD(hl.GetHigh())
}

func opcode0x85() {
	// ADD A,L
	opcodesADD(hl.GetLow())
}

func opcode0x86() {
	// ADD A,(HL)
	opcodesADD(mem.Read(hl.GetValue()))
}

func opcode0x87() {
	// ADD A,A
	opcodesADD(af.GetHigh())
}

func opcode0x88() {
	// ADC A,B
	opcodesADC(bc.GetHigh())
}

func opcode0x89() {
	// ADC A,C
	opcodesADC(bc.GetLow())
}

func opcode0x8A() {
	// ADC A,D
	opcodesADC(de.GetHigh())
}

func opcode0x8B() {
	// ADC A,E
	opcodesADC(de.GetLow())
}

func opcode0x8C() {
	// ADC A,H
	opcodesADC(hl.GetHigh())
}

func opcode0x8D() {
	// ADC A,L
	opcodesADC(hl.GetLow())
}

func opcode0x8E() {
	// ADC A,(HL)
	opcodesADC(mem.Read(hl.GetValue()))
}

func opcode0x8F() {
	// ADC A,A
	opcodesADC(af.GetHigh())
}

func opcode0x90() {
	// SUB B
	opcodesSUB(bc.GetHigh())
}

func opcode0x91() {
	// SUB C
	opcodesSUB(bc.GetLow())
}

func opcode0x92() {
	// SUB D
	opcodesSUB(de.GetHigh())
}

func opcode0x93() {
	// SUB E
	opcodesSUB(de.GetLow())
}

func opcode0x94() {
	// SUB H
	opcodesSUB(hl.GetHigh())
}

func opcode0x95() {
	// SUB L
	opcodesSUB(hl.GetLow())
}

func opcode0x96() {
	// SUB (HL)
	opcodesSUB(mem.Read(hl.GetValue()))
}

func opcode0x97() {
	// SUB A
	opcodesSUB(af.GetHigh())
}

func opcode0x98() {
	// SBC B
	opcodesSBC(bc.GetHigh())
}

func opcode0x99() {
	// SBC C
	opcodesSBC(bc.GetLow())
}

func opcode0x9A() {
	// SBC D
	opcodesSBC(de.GetHigh())
}

func opcode0x9B() {
	// SBC E
	opcodesSBC(de.GetLow())
}

func opcode0x9C() {
	// SBC H
	opcodesSBC(hl.GetHigh())
}

func opcode0x9D() {
	// SBC L
	opcodesSBC(hl.GetLow())
}

func opcode0x9E() {
	// SBC (HL)
	opcodesSBC(mem.Read(hl.GetValue()))
}

func opcode0x9F() {
	// SBC A
	opcodesSBC(af.GetHigh())
}

func opcode0xA0() {
	// AND B
	opcodesAND(bc.GetHigh())
}

func opcode0xA1() {
	// AND C
	opcodesAND(bc.GetLow())
}

func opcode0xA2() {
	// AND D
	opcodesAND(de.GetHigh())
}

func opcode0xA3() {
	// AND E
	opcodesAND(de.GetLow())
}

func opcode0xA4() {
	// AND H
	opcodesAND(hl.GetHigh())
}

func opcode0xA5() {
	// AND L
	opcodesAND(hl.GetLow())
}

func opcode0xA6() {
	// AND (HL)
	opcodesAND(mem.Read(hl.GetValue()))
}

func opcode0xA7() {
	// AND A
	opcodesAND(af.GetHigh())
}

func opcode0xA8() {
	// XOR B
	opcodesXOR(bc.GetHigh())
}

func opcode0xA9() {
	// XOR C
	opcodesXOR(bc.GetLow())
}

func opcode0xAA() {
	// XOR D
	opcodesXOR(de.GetHigh())
}

func opcode0xAB() {
	// XOR E
	opcodesXOR(de.GetLow())
}

func opcode0xAC() {
	// XOR H
	opcodesXOR(hl.GetHigh())
}

func opcode0xAD() {
	// XOR L
	opcodesXOR(hl.GetLow())
}

func opcode0xAE() {
	// XOR (HL)
	opcodesXOR(mem.Read(hl.GetValue()))
}

func opcode0xAF() {
	// XOR A
	opcodesXOR(af.GetHigh())
}

func opcode0xB0() {
	// OR B
	opcodesOR(bc.GetHigh())
}

func opcode0xB1() {
	// OR C
	opcodesOR(bc.GetLow())
}

func opcode0xB2() {
	// OR D
	opcodesOR(de.GetHigh())
}

func opcode0xB3() {
	// OR E
	opcodesOR(de.GetLow())
}

func opcode0xB4() {
	// OR H
	opcodesOR(hl.GetHigh())
}

func opcode0xB5() {
	// OR L
	opcodesOR(hl.GetLow())
}

func opcode0xB6() {
	// OR (HL)
	opcodesOR(mem.Read(hl.GetValue()))
}

func opcode0xB7() {
	// OR A
	opcodesOR(af.GetHigh())
}

func opcode0xB8() {
	// CP B
	opcodesCP(bc.GetHigh())
}

func opcode0xB9() {
	// CP C
	opcodesCP(bc.GetLow())
}

func opcode0xBA() {
	// CP D
	opcodesCP(de.GetHigh())
}

func opcode0xBB() {
	// CP E
	opcodesCP(de.GetLow())
}

func opcode0xBC() {
	// CP H
	opcodesCP(hl.GetHigh())
}

func opcode0xBD() {
	// CP L
	opcodesCP(hl.GetLow())
}

func opcode0xBE() {
	// CP (HL)
	opcodesCP(mem.Read(hl.GetValue()))
}

func opcode0xBF() {
	// CP A
	opcodesCP(af.GetHigh())
}

func opcode0xC0() {
	// RET NZ
	if !isSetFlag(flagZero) {
		stackPop(&pc)
		branchTaken = true
	}
}

func opcode0xC1() {
	// POP BC
	stackPop(&bc)
}

func opcode0xC2() {
	// JP NZ,nn
	if !isSetFlag(flagZero) {
		l := mem.Read(pc.GetValue())
		pc.Increment()
		h := mem.Read(pc.GetValue())
		pc.SetHigh(h)
		pc.SetLow(l)
		branchTaken = true
	} else {
		pc.Increment()
		pc.Increment()
	}
}

func opcode0xC3() {
	// JP nn
	l := mem.Read(pc.GetValue())
	pc.Increment()
	h := mem.Read(pc.GetValue())
	pc.SetHigh(h)
	pc.SetLow(l)
}

func opcode0xC4() {
	// CALL NZ,nn
	if !isSetFlag(flagZero) {
		l := mem.Read(pc.GetValue())
		pc.Increment()
		h := mem.Read(pc.GetValue())
		pc.Increment()
		stackPush(&pc)
		pc.SetHigh(h)
		pc.SetLow(l)
		branchTaken = true
	} else {
		pc.Increment()
		pc.Increment()
	}
}

func opcode0xC5() {
	// PUSH BC
	stackPush(&bc)
}

func opcode0xC6() {
	// ADD A,n
	opcodesADD(mem.Read(pc.GetValue()))
	pc.Increment()
}

func opcode0xC7() {
	// RST 00H
	stackPush(&pc)
	pc.SetValue(0x0000)
}

func opcode0xC8() {
	// RET Z
	if isSetFlag(flagZero) {
		stackPop(&pc)
		branchTaken = true
	}
}

func opcode0xC9() {
	// RET
	stackPop(&pc)
}

func opcode0xCA() {
	// JP Z,nn
	if isSetFlag(flagZero) {
		l := mem.Read(pc.GetValue())
		pc.Increment()
		h := mem.Read(pc.GetValue())
		pc.SetHigh(h)
		pc.SetLow(l)
		branchTaken = true
	} else {
		pc.Increment()
		pc.Increment()
	}
}

func opcode0xCB() {
	// CB prefixed instruction
}

func opcode0xCC() {
	// CALL Z,nn
	if isSetFlag(flagZero) {
		l := mem.Read(pc.GetValue())
		pc.Increment()
		h := mem.Read(pc.GetValue())
		pc.Increment()
		stackPush(&pc)
		pc.SetHigh(h)
		pc.SetLow(l)
		branchTaken = true
	} else {
		pc.Increment()
		pc.Increment()
	}
}

func opcode0xCD() {
	// CALL nn
	l := mem.Read(pc.GetValue())
	pc.Increment()
	h := mem.Read(pc.GetValue())
	pc.Increment()
	stackPush(&pc)
	pc.SetHigh(h)
	pc.SetLow(l)
}

func opcode0xCE() {
	// ADC A,n
	opcodesADC(mem.Read(pc.GetValue()))
	pc.Increment()
}

func opcode0xCF() {
	// RST 08H
	stackPush(&pc)
	pc.SetValue(0x0008)
}

func opcode0xD0() {
	// RET NC
	if !isSetFlag(flagCarry) {
		stackPop(&pc)
		branchTaken = true
	}
}

func opcode0xD1() {
	// POP DE
	stackPop(&de)
}

func opcode0xD2() {
	// JP NC,nn
	if !isSetFlag(flagCarry) {
		l := mem.Read(pc.GetValue())
		pc.Increment()
		h := mem.Read(pc.GetValue())
		pc.SetHigh(h)
		pc.SetLow(l)
		branchTaken = true
	} else {
		pc.Increment()
		pc.Increment()
	}
}

func opcode0xD3() {
	invalidOPCode()
}

func opcode0xD4() {
	// CALL NC,nn
	if !isSetFlag(flagCarry) {
		l := mem.Read(pc.GetValue())
		pc.Increment()
		h := mem.Read(pc.GetValue())
		pc.Increment()
		stackPush(&pc)
		pc.SetHigh(h)
		pc.SetLow(l)
		branchTaken = true
	} else {
		pc.Increment()
		pc.Increment()
	}
}

func opcode0xD5() {
	// PUSH DE
	stackPush(&de)
}

func opcode0xD6() {
	// SUB n
	opcodesSUB(mem.Read(pc.GetValue()))
	pc.Increment()
}

func opcode0xD7() {
	// RST 10H
	stackPush(&pc)
	pc.SetValue(0x0010)
}

func opcode0xD8() {
	// RET C
	if isSetFlag(flagCarry) {
		stackPop(&pc)
		branchTaken = true
	}
}

func opcode0xD9() {
	// RETI
	stackPop(&pc)
	ime = true
}

func opcode0xDA() {
	// JP C,nn
	if isSetFlag(flagCarry) {
		l := mem.Read(pc.GetValue())
		pc.Increment()
		h := mem.Read(pc.GetValue())
		pc.SetHigh(h)
		pc.SetLow(l)
		branchTaken = true
	} else {
		pc.Increment()
		pc.Increment()
	}
}

func opcode0xDB() {
	invalidOPCode()
}

func opcode0xDC() {
	// CALL C,nn
	if isSetFlag(flagCarry) {
		l := mem.Read(pc.GetValue())
		pc.Increment()
		h := mem.Read(pc.GetValue())
		pc.Increment()
		stackPush(&pc)
		pc.SetHigh(h)
		pc.SetLow(l)
		branchTaken = true
	} else {
		pc.Increment()
		pc.Increment()
	}
}

func opcode0xDD() {
	invalidOPCode()
}

func opcode0xDE() {
	// SBC n
	opcodesSBC(mem.Read(pc.GetValue()))
	pc.Increment()
}

func opcode0xDF() {
	// RST 18H
	stackPush(&pc)
	pc.SetValue(0x0018)
}

func opcode0xE0() {
	// LD (0xFF00+n),A
	opcodesLDValueToAddr(0xFF00+uint16(mem.Read(pc.GetValue())), af.GetHigh())
	pc.Increment()
}

func opcode0xE1() {
	// POP HL
	stackPop(&hl)
}

func opcode0xE2() {
	// LD (0xFF00+C),A
	opcodesLDValueToAddr(0xFF00+uint16(bc.GetLow()), af.GetHigh())
}

func opcode0xE3() {
	invalidOPCode()
}

func opcode0xE4() {
	invalidOPCode()
}

func opcode0xE5() {
	// PUSH HL
	stackPush(&hl)
}

func opcode0xE6() {
	// AND n
	opcodesAND(mem.Read(pc.GetValue()))
	pc.Increment()
}

func opcode0xE7() {
	// RST 20H
	stackPush(&pc)
	pc.SetValue(0x0020)
}

func opcode0xE8() {
	// ADD SP,n
	opcodesADDSP(int8(mem.Read(pc.GetValue())))
	pc.Increment()
}

func opcode0xE9() {
	// JP (HL)
	pc.SetValue(hl.GetValue())
}

func opcode0xEA() {
	// LD (nn),A
	var tmp SixteenBitReg
	tmp.SetLow(mem.Read(pc.GetValue()))
	pc.Increment()
	tmp.SetHigh(mem.Read(pc.GetValue()))
	pc.Increment()
	opcodesLDValueToAddr(tmp.GetValue(), af.GetHigh())
}

func opcode0xEB() {
	invalidOPCode()
}

func opcode0xEC() {
	invalidOPCode()
}

func opcode0xED() {
	invalidOPCode()
}

func opcode0xEE() {
	// XOR n
	opcodesXOR(mem.Read(pc.GetValue()))
	pc.Increment()
}

func opcode0xEF() {
	// RST 28H
	stackPush(&pc)
	pc.SetValue(0x28)
}

func opcode0xF0() {
	// LD A,(0xFF00+n)
	opcodesLDAddrToReg(af.GetHighReg(), 0xFF00+uint16(mem.Read(pc.GetValue())))
	pc.Increment()
}

func opcode0xF1() {
	// POP AF
	stackPop(&af)
	af.SetLow(af.GetLow() & 0xF0)
}

func opcode0xF2() {
	// LD A,(C)
	opcodesLDAddrToReg(af.GetHighReg(), 0xFF00+uint16(bc.GetLow()))
}

func opcode0xF3() {
	// DI
	ime = false
	imeCycles = 0
}

func opcode0xF4() {
	invalidOPCode()
}

func opcode0xF5() {
	// PUSH AF
	stackPush(&af)
}

func opcode0xF6() {
	// OR n
	opcodesOR(mem.Read(pc.GetValue()))
	pc.Increment()
}

func opcode0xF7() {
	// RST 30H
	stackPush(&pc)
	pc.SetValue(0x0030)
}

func opcode0xF8() {
	// LD HL,SP+n
	offset := int8(mem.Read(pc.GetValue()))
	value := int(sp.GetValue()) + int(offset)
	result := uint16(value)
	clearAllFlags()
	if ((sp.GetValue() ^ uint16(offset) ^ result) & 0x100) == 0x100 {
		setFlag(flagCarry)
	}
	if ((sp.GetValue() ^ uint16(offset) ^ result) & 0x10) == 0x10 {
		setFlag(flagHalf)
	}
	hl.SetValue(result)
	pc.Increment()
}

func opcode0xF9() {
	// LD SP,HL
	sp.SetValue(hl.GetValue())
}

func opcode0xFA() {
	// LD A,(nn)
	var tmp SixteenBitReg
	tmp.SetLow(mem.Read(pc.GetValue()))
	pc.Increment()
	tmp.SetHigh(mem.Read(pc.GetValue()))
	pc.Increment()
	opcodesLDAddrToReg(af.GetHighReg(), tmp.GetValue())
}

func opcode0xFB() {
	// EI
	imeCycles = (int(machineCycles[0xFB]) * 4) + 1
}

func opcode0xFC() {
	invalidOPCode()
}

func opcode0xFD() {
	invalidOPCode()
}

func opcode0xFE() {
	// CP n
	opcodesCP(mem.Read(pc.GetValue()))
	pc.Increment()
}

func opcode0xFF() {
	// RST 38H
	stackPush(&pc)
	pc.SetValue(0x0038)
}
