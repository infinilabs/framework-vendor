// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This code was translated into a form compatible with 6a from the public
// domain sources in SUPERCOP: https://bench.cr.yp.to/supercop.html

package main

import (
	. "github.com/mmcloughlin/avo/build"
	"github.com/mmcloughlin/avo/ir"
	. "github.com/mmcloughlin/avo/operand"
	. "github.com/mmcloughlin/avo/reg"
	_ "golang.org/x/crypto/salsa20/salsa"
)

//go:generate go run . -out ../salsa20_amd64.s -pkg salsa

func main() {
	Package("golang.org/x/crypto/salsa20/salsa")
	ConstraintExpr("amd64,!purego,gc")
	salsa2020XORKeyStream()
	Generate()
}

func salsa2020XORKeyStream() {
	Implement("salsa2020XORKeyStream")
	Attributes(0)
	AllocLocal(456) // frame = 424 + 32 byte alignment
	Comment("This needs up to 64 bytes at 360(R12); hence the non-obvious frame size.")

	Load(Param("out"), RDI)
	Load(Param("in"), RSI)
	Load(Param("n"), RDX)
	Load(Param("nonce"), RCX)
	Load(Param("key"), R8)

	MOVQ(RSP, R12)
	ADDQ(Imm(31), R12)
	ANDQ(I32(^31), R12)

	MOVQ(RDX, R9)
	MOVQ(RCX, RDX)
	MOVQ(R8, R10)
	CMPQ(R9, Imm(0))
	JBE(LabelRef("DONE"))

	START()
	BYTESATLEAST256()
	MAINLOOP1()
	BYTESBETWEEN1AND255()
	NOCOPY()
	MAINLOOP2()

	Label("BYTESATLEAST64")
	Label("DONE")
	RET()
	Label("BYTESATLEAST65")
	SUBQ(Imm(64), R9)
	ADDQ(Imm(64), RDI)
	ADDQ(Imm(64), RSI)
	JMP(LabelRef("BYTESBETWEEN1AND255"))
}

func START() {
	Label("START")
	MOVL(Mem{Base: R10}.Offset(20), ECX)
	MOVL(Mem{Base: R10}.Offset(0), R8L)
	MOVL(Mem{Base: EDX}.Offset(0), EAX)
	MOVL(Mem{Base: R10}.Offset(16), R11L)
	MOVL(ECX, Mem{Base: R12}.Offset(0))
	MOVL(R8L, Mem{Base: R12}.Offset(4))
	MOVL(EAX, Mem{Base: R12}.Offset(8))
	MOVL(R11L, Mem{Base: R12}.Offset(12))
	MOVL(Mem{Base: EDX}.Offset(8), ECX)
	MOVL(Mem{Base: R10}.Offset(24), R8L)
	MOVL(Mem{Base: R10}.Offset(4), EAX)
	MOVL(Mem{Base: EDX}.Offset(4), R11L)
	MOVL(ECX, Mem{Base: R12}.Offset(16))
	MOVL(R8L, Mem{Base: R12}.Offset(20))
	MOVL(EAX, Mem{Base: R12}.Offset(24))
	MOVL(R11L, Mem{Base: R12}.Offset(28))
	MOVL(Mem{Base: EDX}.Offset(12), ECX)
	MOVL(Mem{Base: R10}.Offset(12), EDX)
	MOVL(Mem{Base: R10}.Offset(28), R8L)
	MOVL(Mem{Base: R10}.Offset(8), EAX)
	MOVL(EDX, Mem{Base: R12}.Offset(32))
	MOVL(ECX, Mem{Base: R12}.Offset(36))
	MOVL(R8L, Mem{Base: R12}.Offset(40))
	MOVL(EAX, Mem{Base: R12}.Offset(44))
	MOVQ(Imm(1634760805), RDX)
	MOVQ(Imm(857760878), RCX)
	MOVQ(Imm(2036477234), R8)
	MOVQ(Imm(1797285236), RAX)
	MOVL(EDX, Mem{Base: R12}.Offset(48))
	MOVL(ECX, Mem{Base: R12}.Offset(52))
	MOVL(R8L, Mem{Base: R12}.Offset(56))
	MOVL(EAX, Mem{Base: R12}.Offset(60))
	CMPQ(R9, U32(256))
	JB(LabelRef("BYTESBETWEEN1AND255"))
	MOVOA(Mem{Base: R12}.Offset(48), X0)
	PSHUFL(Imm(0x55), X0, X1)
	PSHUFL(Imm(0xAA), X0, X2)
	PSHUFL(Imm(0xFF), X0, X3)
	PSHUFL(Imm(0x00), X0, X0)
	MOVOA(X1, Mem{Base: R12}.Offset(64))
	MOVOA(X2, Mem{Base: R12}.Offset(80))
	MOVOA(X3, Mem{Base: R12}.Offset(96))
	MOVOA(X0, Mem{Base: R12}.Offset(112))
	MOVOA(Mem{Base: R12}.Offset(0), X0)
	PSHUFL(Imm(0xAA), X0, X1)
	PSHUFL(Imm(0xFF), X0, X2)
	PSHUFL(Imm(0x00), X0, X3)
	PSHUFL(Imm(0x55), X0, X0)
	MOVOA(X1, Mem{Base: R12}.Offset(128))
	MOVOA(X2, Mem{Base: R12}.Offset(144))
	MOVOA(X3, Mem{Base: R12}.Offset(160))
	MOVOA(X0, Mem{Base: R12}.Offset(176))
	MOVOA(Mem{Base: R12}.Offset(16), X0)
	PSHUFL(Imm(0xFF), X0, X1)
	PSHUFL(Imm(0x55), X0, X2)
	PSHUFL(Imm(0xAA), X0, X0)
	MOVOA(X1, Mem{Base: R12}.Offset(192))
	MOVOA(X2, Mem{Base: R12}.Offset(208))
	MOVOA(X0, Mem{Base: R12}.Offset(224))
	MOVOA(Mem{Base: R12}.Offset(32), X0)
	PSHUFL(Imm(0x00), X0, X1)
	PSHUFL(Imm(0xAA), X0, X2)
	PSHUFL(Imm(0xFF), X0, X0)
	MOVOA(X1, Mem{Base: R12}.Offset(240))
	MOVOA(X2, Mem{Base: R12}.Offset(256))
	MOVOA(X0, Mem{Base: R12}.Offset(272))

}

func BYTESATLEAST256() {
	Label("BYTESATLEAST256")
	MOVL(Mem{Base: R12}.Offset(16), EDX)
	MOVL(Mem{Base: R12}.Offset(36), ECX)
	MOVL(EDX, Mem{Base: R12}.Offset(288))
	MOVL(ECX, Mem{Base: R12}.Offset(304))
	SHLQ(Imm(32), RCX)
	ADDQ(RCX, RDX)
	ADDQ(Imm(1), RDX)
	MOVQ(RDX, RCX)
	SHRQ(Imm(32), RCX)
	MOVL(EDX, Mem{Base: R12}.Offset(292))
	MOVL(ECX, Mem{Base: R12}.Offset(308))
	ADDQ(Imm(1), RDX)
	MOVQ(RDX, RCX)
	SHRQ(Imm(32), RCX)
	MOVL(EDX, Mem{Base: R12}.Offset(296))
	MOVL(ECX, Mem{Base: R12}.Offset(312))
	ADDQ(Imm(1), RDX)
	MOVQ(RDX, RCX)
	SHRQ(Imm(32), RCX)
	MOVL(EDX, Mem{Base: R12}.Offset(300))
	MOVL(ECX, Mem{Base: R12}.Offset(316))
	ADDQ(Imm(1), RDX)
	MOVQ(RDX, RCX)
	SHRQ(Imm(32), RCX)
	MOVL(EDX, Mem{Base: R12}.Offset(16))
	MOVL(ECX, Mem{Base: R12}.Offset(36))
	MOVQ(R9, Mem{Base: R12}.Offset(352))
	MOVQ(U32(20), RDX)
	MOVOA(Mem{Base: R12}.Offset(64), X0)
	MOVOA(Mem{Base: R12}.Offset(80), X1)
	MOVOA(Mem{Base: R12}.Offset(96), X2)
	MOVOA(Mem{Base: R12}.Offset(256), X3)
	MOVOA(Mem{Base: R12}.Offset(272), X4)
	MOVOA(Mem{Base: R12}.Offset(128), X5)
	MOVOA(Mem{Base: R12}.Offset(144), X6)
	MOVOA(Mem{Base: R12}.Offset(176), X7)
	MOVOA(Mem{Base: R12}.Offset(192), X8)
	MOVOA(Mem{Base: R12}.Offset(208), X9)
	MOVOA(Mem{Base: R12}.Offset(224), X10)
	MOVOA(Mem{Base: R12}.Offset(304), X11)
	MOVOA(Mem{Base: R12}.Offset(112), X12)
	MOVOA(Mem{Base: R12}.Offset(160), X13)
	MOVOA(Mem{Base: R12}.Offset(240), X14)
	MOVOA(Mem{Base: R12}.Offset(288), X15)
}

func MAINLOOP1() {
	Label("MAINLOOP1")
	MOVOA(X1, Mem{Base: R12}.Offset(320))
	MOVOA(X2, Mem{Base: R12}.Offset(336))
	MOVOA(X13, X1)
	PADDL(X12, X1)
	MOVOA(X1, X2)
	PSLLL(Imm(7), X1)
	PXOR(X1, X14)
	PSRLL(Imm(25), X2)
	PXOR(X2, X14)
	MOVOA(X7, X1)
	PADDL(X0, X1)
	MOVOA(X1, X2)
	PSLLL(Imm(7), X1)
	PXOR(X1, X11)
	PSRLL(Imm(25), X2)
	PXOR(X2, X11)
	MOVOA(X12, X1)
	PADDL(X14, X1)
	MOVOA(X1, X2)
	PSLLL(Imm(9), X1)
	PXOR(X1, X15)
	PSRLL(Imm(23), X2)
	PXOR(X2, X15)
	MOVOA(X0, X1)
	PADDL(X11, X1)
	MOVOA(X1, X2)
	PSLLL(Imm(9), X1)
	PXOR(X1, X9)
	PSRLL(Imm(23), X2)
	PXOR(X2, X9)
	MOVOA(X14, X1)
	PADDL(X15, X1)
	MOVOA(X1, X2)
	PSLLL(Imm(13), X1)
	PXOR(X1, X13)
	PSRLL(Imm(19), X2)
	PXOR(X2, X13)
	MOVOA(X11, X1)
	PADDL(X9, X1)
	MOVOA(X1, X2)
	PSLLL(Imm(13), X1)
	PXOR(X1, X7)
	PSRLL(Imm(19), X2)
	PXOR(X2, X7)
	MOVOA(X15, X1)
	PADDL(X13, X1)
	MOVOA(X1, X2)
	PSLLL(Imm(18), X1)
	PXOR(X1, X12)
	PSRLL(Imm(14), X2)
	PXOR(X2, X12)
	MOVOA(Mem{Base: R12}.Offset(320), X1)
	MOVOA(X12, Mem{Base: R12}.Offset(320))
	MOVOA(X9, X2)
	PADDL(X7, X2)
	MOVOA(X2, X12)
	PSLLL(Imm(18), X2)
	PXOR(X2, X0)
	PSRLL(Imm(14), X12)
	PXOR(X12, X0)
	MOVOA(X5, X2)
	PADDL(X1, X2)
	MOVOA(X2, X12)
	PSLLL(Imm(7), X2)
	PXOR(X2, X3)
	PSRLL(Imm(25), X12)
	PXOR(X12, X3)
	MOVOA(Mem{Base: R12}.Offset(336), X2)
	MOVOA(X0, Mem{Base: R12}.Offset(336))
	MOVOA(X6, X0)
	PADDL(X2, X0)
	MOVOA(X0, X12)
	PSLLL(Imm(7), X0)
	PXOR(X0, X4)
	PSRLL(Imm(25), X12)
	PXOR(X12, X4)
	MOVOA(X1, X0)
	PADDL(X3, X0)
	MOVOA(X0, X12)
	PSLLL(Imm(9), X0)
	PXOR(X0, X10)
	PSRLL(Imm(23), X12)
	PXOR(X12, X10)
	MOVOA(X2, X0)
	PADDL(X4, X0)
	MOVOA(X0, X12)
	PSLLL(Imm(9), X0)
	PXOR(X0, X8)
	PSRLL(Imm(23), X12)
	PXOR(X12, X8)
	MOVOA(X3, X0)
	PADDL(X10, X0)
	MOVOA(X0, X12)
	PSLLL(Imm(13), X0)
	PXOR(X0, X5)
	PSRLL(Imm(19), X12)
	PXOR(X12, X5)
	MOVOA(X4, X0)
	PADDL(X8, X0)
	MOVOA(X0, X12)
	PSLLL(Imm(13), X0)
	PXOR(X0, X6)
	PSRLL(Imm(19), X12)
	PXOR(X12, X6)
	MOVOA(X10, X0)
	PADDL(X5, X0)
	MOVOA(X0, X12)
	PSLLL(Imm(18), X0)
	PXOR(X0, X1)
	PSRLL(Imm(14), X12)
	PXOR(X12, X1)
	MOVOA(Mem{Base: R12}.Offset(320), X0)
	MOVOA(X1, Mem{Base: R12}.Offset(320))
	MOVOA(X4, X1)
	PADDL(X0, X1)
	MOVOA(X1, X12)
	PSLLL(Imm(7), X1)
	PXOR(X1, X7)
	PSRLL(Imm(25), X12)
	PXOR(X12, X7)
	MOVOA(X8, X1)
	PADDL(X6, X1)
	MOVOA(X1, X12)
	PSLLL(Imm(18), X1)
	PXOR(X1, X2)
	PSRLL(Imm(14), X12)
	PXOR(X12, X2)
	MOVOA(Mem{Base: R12}.Offset(336), X12)
	MOVOA(X2, Mem{Base: R12}.Offset(336))
	MOVOA(X14, X1)
	PADDL(X12, X1)
	MOVOA(X1, X2)
	PSLLL(Imm(7), X1)
	PXOR(X1, X5)
	PSRLL(Imm(25), X2)
	PXOR(X2, X5)
	MOVOA(X0, X1)
	PADDL(X7, X1)
	MOVOA(X1, X2)
	PSLLL(Imm(9), X1)
	PXOR(X1, X10)
	PSRLL(Imm(23), X2)
	PXOR(X2, X10)
	MOVOA(X12, X1)
	PADDL(X5, X1)
	MOVOA(X1, X2)
	PSLLL(Imm(9), X1)
	PXOR(X1, X8)
	PSRLL(Imm(23), X2)
	PXOR(X2, X8)
	MOVOA(X7, X1)
	PADDL(X10, X1)
	MOVOA(X1, X2)
	PSLLL(Imm(13), X1)
	PXOR(X1, X4)
	PSRLL(Imm(19), X2)
	PXOR(X2, X4)
	MOVOA(X5, X1)
	PADDL(X8, X1)
	MOVOA(X1, X2)
	PSLLL(Imm(13), X1)
	PXOR(X1, X14)
	PSRLL(Imm(19), X2)
	PXOR(X2, X14)
	MOVOA(X10, X1)
	PADDL(X4, X1)
	MOVOA(X1, X2)
	PSLLL(Imm(18), X1)
	PXOR(X1, X0)
	PSRLL(Imm(14), X2)
	PXOR(X2, X0)
	MOVOA(Mem{Base: R12}.Offset(320), X1)
	MOVOA(X0, Mem{Base: R12}.Offset(320))
	MOVOA(X8, X0)
	PADDL(X14, X0)
	MOVOA(X0, X2)
	PSLLL(Imm(18), X0)
	PXOR(X0, X12)
	PSRLL(Imm(14), X2)
	PXOR(X2, X12)
	MOVOA(X11, X0)
	PADDL(X1, X0)
	MOVOA(X0, X2)
	PSLLL(Imm(7), X0)
	PXOR(X0, X6)
	PSRLL(Imm(25), X2)
	PXOR(X2, X6)
	MOVOA(Mem{Base: R12}.Offset(336), X2)
	MOVOA(X12, Mem{Base: R12}.Offset(336))
	MOVOA(X3, X0)
	PADDL(X2, X0)
	MOVOA(X0, X12)
	PSLLL(Imm(7), X0)
	PXOR(X0, X13)
	PSRLL(Imm(25), X12)
	PXOR(X12, X13)
	MOVOA(X1, X0)
	PADDL(X6, X0)
	MOVOA(X0, X12)
	PSLLL(Imm(9), X0)
	PXOR(X0, X15)
	PSRLL(Imm(23), X12)
	PXOR(X12, X15)
	MOVOA(X2, X0)
	PADDL(X13, X0)
	MOVOA(X0, X12)
	PSLLL(Imm(9), X0)
	PXOR(X0, X9)
	PSRLL(Imm(23), X12)
	PXOR(X12, X9)
	MOVOA(X6, X0)
	PADDL(X15, X0)
	MOVOA(X0, X12)
	PSLLL(Imm(13), X0)
	PXOR(X0, X11)
	PSRLL(Imm(19), X12)
	PXOR(X12, X11)
	MOVOA(X13, X0)
	PADDL(X9, X0)
	MOVOA(X0, X12)
	PSLLL(Imm(13), X0)
	PXOR(X0, X3)
	PSRLL(Imm(19), X12)
	PXOR(X12, X3)
	MOVOA(X15, X0)
	PADDL(X11, X0)
	MOVOA(X0, X12)
	PSLLL(Imm(18), X0)
	PXOR(X0, X1)
	PSRLL(Imm(14), X12)
	PXOR(X12, X1)
	MOVOA(X9, X0)
	PADDL(X3, X0)
	MOVOA(X0, X12)
	PSLLL(Imm(18), X0)
	PXOR(X0, X2)
	PSRLL(Imm(14), X12)
	PXOR(X12, X2)
	MOVOA(Mem{Base: R12}.Offset(320), X12)
	MOVOA(Mem{Base: R12}.Offset(336), X0)
	SUBQ(Imm(2), RDX)
	JA(LabelRef("MAINLOOP1"))
	PADDL(Mem{Base: R12}.Offset(112), X12)
	PADDL(Mem{Base: R12}.Offset(176), X7)
	PADDL(Mem{Base: R12}.Offset(224), X10)
	PADDL(Mem{Base: R12}.Offset(272), X4)
	MOVD(X12, EDX)
	MOVD(X7, ECX)
	MOVD(X10, R8)
	MOVD(X4, R9)
	PSHUFL(Imm(0x39), X12, X12)
	PSHUFL(Imm(0x39), X7, X7)
	PSHUFL(Imm(0x39), X10, X10)
	PSHUFL(Imm(0x39), X4, X4)
	XORL(Mem{Base: SI}.Offset(0), EDX)
	XORL(Mem{Base: SI}.Offset(4), ECX)
	XORL(Mem{Base: SI}.Offset(8), R8L)
	XORL(Mem{Base: SI}.Offset(12), R9L)
	MOVL(EDX, Mem{Base: DI}.Offset(0))
	MOVL(ECX, Mem{Base: DI}.Offset(4))
	MOVL(R8L, Mem{Base: DI}.Offset(8))
	MOVL(R9L, Mem{Base: DI}.Offset(12))
	MOVD(X12, EDX)
	MOVD(X7, ECX)
	MOVD(X10, R8)
	MOVD(X4, R9)
	PSHUFL(Imm(0x39), X12, X12)
	PSHUFL(Imm(0x39), X7, X7)
	PSHUFL(Imm(0x39), X10, X10)
	PSHUFL(Imm(0x39), X4, X4)
	XORL(Mem{Base: SI}.Offset(64), EDX)
	XORL(Mem{Base: SI}.Offset(68), ECX)
	XORL(Mem{Base: SI}.Offset(72), R8L)
	XORL(Mem{Base: SI}.Offset(76), R9L)
	MOVL(EDX, Mem{Base: DI}.Offset(64))
	MOVL(ECX, Mem{Base: DI}.Offset(68))
	MOVL(R8L, Mem{Base: DI}.Offset(72))
	MOVL(R9L, Mem{Base: DI}.Offset(76))
	MOVD(X12, EDX)
	MOVD(X7, ECX)
	MOVD(X10, R8)
	MOVD(X4, R9)
	PSHUFL(Imm(0x39), X12, X12)
	PSHUFL(Imm(0x39), X7, X7)
	PSHUFL(Imm(0x39), X10, X10)
	PSHUFL(Imm(0x39), X4, X4)
	XORL(Mem{Base: SI}.Offset(128), EDX)
	XORL(Mem{Base: SI}.Offset(132), ECX)
	XORL(Mem{Base: SI}.Offset(136), R8L)
	XORL(Mem{Base: SI}.Offset(140), R9L)
	MOVL(EDX, Mem{Base: DI}.Offset(128))
	MOVL(ECX, Mem{Base: DI}.Offset(132))
	MOVL(R8L, Mem{Base: DI}.Offset(136))
	MOVL(R9L, Mem{Base: DI}.Offset(140))
	MOVD(X12, EDX)
	MOVD(X7, ECX)
	MOVD(X10, R8)
	MOVD(X4, R9)
	XORL(Mem{Base: SI}.Offset(192), EDX)
	XORL(Mem{Base: SI}.Offset(196), ECX)
	XORL(Mem{Base: SI}.Offset(200), R8L)
	XORL(Mem{Base: SI}.Offset(204), R9L)
	MOVL(EDX, Mem{Base: DI}.Offset(192))
	MOVL(ECX, Mem{Base: DI}.Offset(196))
	MOVL(R8L, Mem{Base: DI}.Offset(200))
	MOVL(R9L, Mem{Base: DI}.Offset(204))
	PADDL(Mem{Base: R12}.Offset(240), X14)
	PADDL(Mem{Base: R12}.Offset(64), X0)
	PADDL(Mem{Base: R12}.Offset(128), X5)
	PADDL(Mem{Base: R12}.Offset(192), X8)
	MOVD(X14, EDX)
	MOVD(X0, ECX)
	MOVD(X5, R8)
	MOVD(X8, R9)
	PSHUFL(Imm(0x39), X14, X14)
	PSHUFL(Imm(0x39), X0, X0)
	PSHUFL(Imm(0x39), X5, X5)
	PSHUFL(Imm(0x39), X8, X8)
	XORL(Mem{Base: SI}.Offset(16), EDX)
	XORL(Mem{Base: SI}.Offset(20), ECX)
	XORL(Mem{Base: SI}.Offset(24), R8L)
	XORL(Mem{Base: SI}.Offset(28), R9L)
	MOVL(EDX, Mem{Base: DI}.Offset(16))
	MOVL(ECX, Mem{Base: DI}.Offset(20))
	MOVL(R8L, Mem{Base: DI}.Offset(24))
	MOVL(R9L, Mem{Base: DI}.Offset(28))
	MOVD(X14, EDX)
	MOVD(X0, ECX)
	MOVD(X5, R8)
	MOVD(X8, R9)
	PSHUFL(Imm(0x39), X14, X14)
	PSHUFL(Imm(0x39), X0, X0)
	PSHUFL(Imm(0x39), X5, X5)
	PSHUFL(Imm(0x39), X8, X8)
	XORL(Mem{Base: SI}.Offset(80), EDX)
	XORL(Mem{Base: SI}.Offset(84), ECX)
	XORL(Mem{Base: SI}.Offset(88), R8L)
	XORL(Mem{Base: SI}.Offset(92), R9L)
	MOVL(EDX, Mem{Base: DI}.Offset(80))
	MOVL(ECX, Mem{Base: DI}.Offset(84))
	MOVL(R8L, Mem{Base: DI}.Offset(88))
	MOVL(R9L, Mem{Base: DI}.Offset(92))
	MOVD(X14, EDX)
	MOVD(X0, ECX)
	MOVD(X5, R8)
	MOVD(X8, R9)
	PSHUFL(Imm(0x39), X14, X14)
	PSHUFL(Imm(0x39), X0, X0)
	PSHUFL(Imm(0x39), X5, X5)
	PSHUFL(Imm(0x39), X8, X8)
	XORL(Mem{Base: SI}.Offset(144), EDX)
	XORL(Mem{Base: SI}.Offset(148), ECX)
	XORL(Mem{Base: SI}.Offset(152), R8L)
	XORL(Mem{Base: SI}.Offset(156), R9L)
	MOVL(EDX, Mem{Base: DI}.Offset(144))
	MOVL(ECX, Mem{Base: DI}.Offset(148))
	MOVL(R8L, Mem{Base: DI}.Offset(152))
	MOVL(R9L, Mem{Base: DI}.Offset(156))
	MOVD(X14, EDX)
	MOVD(X0, ECX)
	MOVD(X5, R8)
	MOVD(X8, R9)
	XORL(Mem{Base: SI}.Offset(208), EDX)
	XORL(Mem{Base: SI}.Offset(212), ECX)
	XORL(Mem{Base: SI}.Offset(216), R8L)
	XORL(Mem{Base: SI}.Offset(220), R9L)
	MOVL(EDX, Mem{Base: DI}.Offset(208))
	MOVL(ECX, Mem{Base: DI}.Offset(212))
	MOVL(R8L, Mem{Base: DI}.Offset(216))
	MOVL(R9L, Mem{Base: DI}.Offset(220))
	PADDL(Mem{Base: R12}.Offset(288), X15)
	PADDL(Mem{Base: R12}.Offset(304), X11)
	PADDL(Mem{Base: R12}.Offset(80), X1)
	PADDL(Mem{Base: R12}.Offset(144), X6)
	MOVD(X15, EDX)
	MOVD(X11, ECX)
	MOVD(X1, R8)
	MOVD(X6, R9)
	PSHUFL(Imm(0x39), X15, X15)
	PSHUFL(Imm(0x39), X11, X11)
	PSHUFL(Imm(0x39), X1, X1)
	PSHUFL(Imm(0x39), X6, X6)
	XORL(Mem{Base: SI}.Offset(32), EDX)
	XORL(Mem{Base: SI}.Offset(36), ECX)
	XORL(Mem{Base: SI}.Offset(40), R8L)
	XORL(Mem{Base: SI}.Offset(44), R9L)
	MOVL(EDX, Mem{Base: DI}.Offset(32))
	MOVL(ECX, Mem{Base: DI}.Offset(36))
	MOVL(R8L, Mem{Base: DI}.Offset(40))
	MOVL(R9L, Mem{Base: DI}.Offset(44))
	MOVD(X15, EDX)
	MOVD(X11, ECX)
	MOVD(X1, R8)
	MOVD(X6, R9)
	PSHUFL(Imm(0x39), X15, X15)
	PSHUFL(Imm(0x39), X11, X11)
	PSHUFL(Imm(0x39), X1, X1)
	PSHUFL(Imm(0x39), X6, X6)
	XORL(Mem{Base: SI}.Offset(96), EDX)
	XORL(Mem{Base: SI}.Offset(100), ECX)
	XORL(Mem{Base: SI}.Offset(104), R8L)
	XORL(Mem{Base: SI}.Offset(108), R9L)
	MOVL(EDX, Mem{Base: DI}.Offset(96))
	MOVL(ECX, Mem{Base: DI}.Offset(100))
	MOVL(R8L, Mem{Base: DI}.Offset(104))
	MOVL(R9L, Mem{Base: DI}.Offset(108))
	MOVD(X15, EDX)
	MOVD(X11, ECX)
	MOVD(X1, R8)
	MOVD(X6, R9)
	PSHUFL(Imm(0x39), X15, X15)
	PSHUFL(Imm(0x39), X11, X11)
	PSHUFL(Imm(0x39), X1, X1)
	PSHUFL(Imm(0x39), X6, X6)
	XORL(Mem{Base: SI}.Offset(160), EDX)
	XORL(Mem{Base: SI}.Offset(164), ECX)
	XORL(Mem{Base: SI}.Offset(168), R8L)
	XORL(Mem{Base: SI}.Offset(172), R9L)
	MOVL(EDX, Mem{Base: DI}.Offset(160))
	MOVL(ECX, Mem{Base: DI}.Offset(164))
	MOVL(R8L, Mem{Base: DI}.Offset(168))
	MOVL(R9L, Mem{Base: DI}.Offset(172))
	MOVD(X15, EDX)
	MOVD(X11, ECX)
	MOVD(X1, R8)
	MOVD(X6, R9)
	XORL(Mem{Base: SI}.Offset(224), EDX)
	XORL(Mem{Base: SI}.Offset(228), ECX)
	XORL(Mem{Base: SI}.Offset(232), R8L)
	XORL(Mem{Base: SI}.Offset(236), R9L)
	MOVL(EDX, Mem{Base: DI}.Offset(224))
	MOVL(ECX, Mem{Base: DI}.Offset(228))
	MOVL(R8L, Mem{Base: DI}.Offset(232))
	MOVL(R9L, Mem{Base: DI}.Offset(236))
	PADDL(Mem{Base: R12}.Offset(160), X13)
	PADDL(Mem{Base: R12}.Offset(208), X9)
	PADDL(Mem{Base: R12}.Offset(256), X3)
	PADDL(Mem{Base: R12}.Offset(96), X2)
	MOVD(X13, EDX)
	MOVD(X9, ECX)
	MOVD(X3, R8)
	MOVD(X2, R9)
	PSHUFL(Imm(0x39), X13, X13)
	PSHUFL(Imm(0x39), X9, X9)
	PSHUFL(Imm(0x39), X3, X3)
	PSHUFL(Imm(0x39), X2, X2)
	XORL(Mem{Base: SI}.Offset(48), EDX)
	XORL(Mem{Base: SI}.Offset(52), ECX)
	XORL(Mem{Base: SI}.Offset(56), R8L)
	XORL(Mem{Base: SI}.Offset(60), R9L)
	MOVL(EDX, Mem{Base: DI}.Offset(48))
	MOVL(ECX, Mem{Base: DI}.Offset(52))
	MOVL(R8L, Mem{Base: DI}.Offset(56))
	MOVL(R9L, Mem{Base: DI}.Offset(60))
	MOVD(X13, EDX)
	MOVD(X9, ECX)
	MOVD(X3, R8)
	MOVD(X2, R9)
	PSHUFL(Imm(0x39), X13, X13)
	PSHUFL(Imm(0x39), X9, X9)
	PSHUFL(Imm(0x39), X3, X3)
	PSHUFL(Imm(0x39), X2, X2)
	XORL(Mem{Base: SI}.Offset(112), EDX)
	XORL(Mem{Base: SI}.Offset(116), ECX)
	XORL(Mem{Base: SI}.Offset(120), R8L)
	XORL(Mem{Base: SI}.Offset(124), R9L)
	MOVL(EDX, Mem{Base: DI}.Offset(112))
	MOVL(ECX, Mem{Base: DI}.Offset(116))
	MOVL(R8L, Mem{Base: DI}.Offset(120))
	MOVL(R9L, Mem{Base: DI}.Offset(124))
	MOVD(X13, EDX)
	MOVD(X9, ECX)
	MOVD(X3, R8)
	MOVD(X2, R9)
	PSHUFL(Imm(0x39), X13, X13)
	PSHUFL(Imm(0x39), X9, X9)
	PSHUFL(Imm(0x39), X3, X3)
	PSHUFL(Imm(0x39), X2, X2)
	XORL(Mem{Base: SI}.Offset(176), EDX)
	XORL(Mem{Base: SI}.Offset(180), ECX)
	XORL(Mem{Base: SI}.Offset(184), R8L)
	XORL(Mem{Base: SI}.Offset(188), R9L)
	MOVL(EDX, Mem{Base: DI}.Offset(176))
	MOVL(ECX, Mem{Base: DI}.Offset(180))
	MOVL(R8L, Mem{Base: DI}.Offset(184))
	MOVL(R9L, Mem{Base: DI}.Offset(188))
	MOVD(X13, EDX)
	MOVD(X9, ECX)
	MOVD(X3, R8)
	MOVD(X2, R9)
	XORL(Mem{Base: SI}.Offset(240), EDX)
	XORL(Mem{Base: SI}.Offset(244), ECX)
	XORL(Mem{Base: SI}.Offset(248), R8L)
	XORL(Mem{Base: SI}.Offset(252), R9L)
	MOVL(EDX, Mem{Base: DI}.Offset(240))
	MOVL(ECX, Mem{Base: DI}.Offset(244))
	MOVL(R8L, Mem{Base: DI}.Offset(248))
	MOVL(R9L, Mem{Base: DI}.Offset(252))
	MOVQ(Mem{Base: R12}.Offset(352), R9)
	SUBQ(U32(256), R9)
	ADDQ(U32(256), RSI)
	ADDQ(U32(256), RDI)
	CMPQ(R9, U32(256))
	JAE(LabelRef("BYTESATLEAST256"))
	CMPQ(R9, Imm(0))
	JBE(LabelRef("DONE"))
}

func BYTESBETWEEN1AND255() {
	Label("BYTESBETWEEN1AND255")
	CMPQ(R9, Imm(64))
	JAE(LabelRef("NOCOPY"))
	MOVQ(RDI, RDX)
	LEAQ(Mem{Base: R12}.Offset(360), RDI)
	MOVQ(R9, RCX)
	// Hack to get Avo to emit:
	// 	REP; MOVSB
	Instruction(&ir.Instruction{Opcode: "REP; MOVSB"})
	LEAQ(Mem{Base: R12}.Offset(360), RDI)
	LEAQ(Mem{Base: R12}.Offset(360), RSI)
}

func NOCOPY() {
	Label("NOCOPY")
	MOVQ(R9, Mem{Base: R12}.Offset(352))
	MOVOA(Mem{Base: R12}.Offset(48), X0)
	MOVOA(Mem{Base: R12}.Offset(0), X1)
	MOVOA(Mem{Base: R12}.Offset(16), X2)
	MOVOA(Mem{Base: R12}.Offset(32), X3)
	MOVOA(X1, X4)
	MOVQ(U32(20), RCX)
}

func MAINLOOP2() {
	Label("MAINLOOP2")
	PADDL(X0, X4)
	MOVOA(X0, X5)
	MOVOA(X4, X6)
	PSLLL(Imm(7), X4)
	PSRLL(Imm(25), X6)
	PXOR(X4, X3)
	PXOR(X6, X3)
	PADDL(X3, X5)
	MOVOA(X3, X4)
	MOVOA(X5, X6)
	PSLLL(Imm(9), X5)
	PSRLL(Imm(23), X6)
	PXOR(X5, X2)
	PSHUFL(Imm(0x93), X3, X3)
	PXOR(X6, X2)
	PADDL(X2, X4)
	MOVOA(X2, X5)
	MOVOA(X4, X6)
	PSLLL(Imm(13), X4)
	PSRLL(Imm(19), X6)
	PXOR(X4, X1)
	PSHUFL(Imm(0x4E), X2, X2)
	PXOR(X6, X1)
	PADDL(X1, X5)
	MOVOA(X3, X4)
	MOVOA(X5, X6)
	PSLLL(Imm(18), X5)
	PSRLL(Imm(14), X6)
	PXOR(X5, X0)
	PSHUFL(Imm(0x39), X1, X1)
	PXOR(X6, X0)
	PADDL(X0, X4)
	MOVOA(X0, X5)
	MOVOA(X4, X6)
	PSLLL(Imm(7), X4)
	PSRLL(Imm(25), X6)
	PXOR(X4, X1)
	PXOR(X6, X1)
	PADDL(X1, X5)
	MOVOA(X1, X4)
	MOVOA(X5, X6)
	PSLLL(Imm(9), X5)
	PSRLL(Imm(23), X6)
	PXOR(X5, X2)
	PSHUFL(Imm(0x93), X1, X1)
	PXOR(X6, X2)
	PADDL(X2, X4)
	MOVOA(X2, X5)
	MOVOA(X4, X6)
	PSLLL(Imm(13), X4)
	PSRLL(Imm(19), X6)
	PXOR(X4, X3)
	PSHUFL(Imm(0x4E), X2, X2)
	PXOR(X6, X3)
	PADDL(X3, X5)
	MOVOA(X1, X4)
	MOVOA(X5, X6)
	PSLLL(Imm(18), X5)
	PSRLL(Imm(14), X6)
	PXOR(X5, X0)
	PSHUFL(Imm(0x39), X3, X3)
	PXOR(X6, X0)
	PADDL(X0, X4)
	MOVOA(X0, X5)
	MOVOA(X4, X6)
	PSLLL(Imm(7), X4)
	PSRLL(Imm(25), X6)
	PXOR(X4, X3)
	PXOR(X6, X3)
	PADDL(X3, X5)
	MOVOA(X3, X4)
	MOVOA(X5, X6)
	PSLLL(Imm(9), X5)
	PSRLL(Imm(23), X6)
	PXOR(X5, X2)
	PSHUFL(Imm(0x93), X3, X3)
	PXOR(X6, X2)
	PADDL(X2, X4)
	MOVOA(X2, X5)
	MOVOA(X4, X6)
	PSLLL(Imm(13), X4)
	PSRLL(Imm(19), X6)
	PXOR(X4, X1)
	PSHUFL(Imm(0x4E), X2, X2)
	PXOR(X6, X1)
	PADDL(X1, X5)
	MOVOA(X3, X4)
	MOVOA(X5, X6)
	PSLLL(Imm(18), X5)
	PSRLL(Imm(14), X6)
	PXOR(X5, X0)
	PSHUFL(Imm(0x39), X1, X1)
	PXOR(X6, X0)
	PADDL(X0, X4)
	MOVOA(X0, X5)
	MOVOA(X4, X6)
	PSLLL(Imm(7), X4)
	PSRLL(Imm(25), X6)
	PXOR(X4, X1)
	PXOR(X6, X1)
	PADDL(X1, X5)
	MOVOA(X1, X4)
	MOVOA(X5, X6)
	PSLLL(Imm(9), X5)
	PSRLL(Imm(23), X6)
	PXOR(X5, X2)
	PSHUFL(Imm(0x93), X1, X1)
	PXOR(X6, X2)
	PADDL(X2, X4)
	MOVOA(X2, X5)
	MOVOA(X4, X6)
	PSLLL(Imm(13), X4)
	PSRLL(Imm(19), X6)
	PXOR(X4, X3)
	PSHUFL(Imm(0x4E), X2, X2)
	PXOR(X6, X3)
	SUBQ(Imm(4), RCX)
	PADDL(X3, X5)
	MOVOA(X1, X4)
	MOVOA(X5, X6)
	PSLLL(Imm(18), X5)
	PXOR(X7, X7)
	PSRLL(Imm(14), X6)
	PXOR(X5, X0)
	PSHUFL(Imm(0x39), X3, X3)
	PXOR(X6, X0)
	JA(LabelRef("MAINLOOP2"))
	PADDL(Mem{Base: R12}.Offset(48), X0)
	PADDL(Mem{Base: R12}.Offset(0), X1)
	PADDL(Mem{Base: R12}.Offset(16), X2)
	PADDL(Mem{Base: R12}.Offset(32), X3)
	MOVD(X0, ECX)
	MOVD(X1, R8)
	MOVD(X2, R9)
	MOVD(X3, EAX)
	PSHUFL(Imm(0x39), X0, X0)
	PSHUFL(Imm(0x39), X1, X1)
	PSHUFL(Imm(0x39), X2, X2)
	PSHUFL(Imm(0x39), X3, X3)
	XORL(Mem{Base: SI}.Offset(0), ECX)
	XORL(Mem{Base: SI}.Offset(48), R8L)
	XORL(Mem{Base: SI}.Offset(32), R9L)
	XORL(Mem{Base: SI}.Offset(16), EAX)
	MOVL(ECX, Mem{Base: DI}.Offset(0))
	MOVL(R8L, Mem{Base: DI}.Offset(48))
	MOVL(R9L, Mem{Base: DI}.Offset(32))
	MOVL(EAX, Mem{Base: DI}.Offset(16))
	MOVD(X0, ECX)
	MOVD(X1, R8)
	MOVD(X2, R9)
	MOVD(X3, EAX)
	PSHUFL(Imm(0x39), X0, X0)
	PSHUFL(Imm(0x39), X1, X1)
	PSHUFL(Imm(0x39), X2, X2)
	PSHUFL(Imm(0x39), X3, X3)
	XORL(Mem{Base: SI}.Offset(20), ECX)
	XORL(Mem{Base: SI}.Offset(4), R8L)
	XORL(Mem{Base: SI}.Offset(52), R9L)
	XORL(Mem{Base: SI}.Offset(36), EAX)
	MOVL(ECX, Mem{Base: DI}.Offset(20))
	MOVL(R8L, Mem{Base: DI}.Offset(4))
	MOVL(R9L, Mem{Base: DI}.Offset(52))
	MOVL(EAX, Mem{Base: DI}.Offset(36))
	MOVD(X0, ECX)
	MOVD(X1, R8)
	MOVD(X2, R9)
	MOVD(X3, EAX)
	PSHUFL(Imm(0x39), X0, X0)
	PSHUFL(Imm(0x39), X1, X1)
	PSHUFL(Imm(0x39), X2, X2)
	PSHUFL(Imm(0x39), X3, X3)
	XORL(Mem{Base: SI}.Offset(40), ECX)
	XORL(Mem{Base: SI}.Offset(24), R8L)
	XORL(Mem{Base: SI}.Offset(8), R9L)
	XORL(Mem{Base: SI}.Offset(56), EAX)
	MOVL(ECX, Mem{Base: DI}.Offset(40))
	MOVL(R8L, Mem{Base: DI}.Offset(24))
	MOVL(R9L, Mem{Base: DI}.Offset(8))
	MOVL(EAX, Mem{Base: DI}.Offset(56))
	MOVD(X0, ECX)
	MOVD(X1, R8)
	MOVD(X2, R9)
	MOVD(X3, EAX)
	XORL(Mem{Base: SI}.Offset(60), ECX)
	XORL(Mem{Base: SI}.Offset(44), R8L)
	XORL(Mem{Base: SI}.Offset(28), R9L)
	XORL(Mem{Base: SI}.Offset(12), EAX)
	MOVL(ECX, Mem{Base: DI}.Offset(60))
	MOVL(R8L, Mem{Base: DI}.Offset(44))
	MOVL(R9L, Mem{Base: DI}.Offset(28))
	MOVL(EAX, Mem{Base: DI}.Offset(12))
	MOVQ(Mem{Base: R12}.Offset(352), R9)
	MOVL(Mem{Base: R12}.Offset(16), ECX)
	MOVL(Mem{Base: R12}.Offset(36), R8L)
	ADDQ(Imm(1), RCX)
	SHLQ(Imm(32), R8)
	ADDQ(R8, RCX)
	MOVQ(RCX, R8)
	SHRQ(Imm(32), R8)
	MOVL(ECX, Mem{Base: R12}.Offset(16))
	MOVL(R8L, Mem{Base: R12}.Offset(36))
	CMPQ(R9, Imm(64))
	JA(LabelRef("BYTESATLEAST65"))
	JAE(LabelRef("BYTESATLEAST64"))
	MOVQ(RDI, RSI)
	MOVQ(RDX, RDI)
	MOVQ(R9, RCX)
	// Hack to get Avo to emit:
	// 	REP; MOVSB
	Instruction(&ir.Instruction{Opcode: "REP; MOVSB"})
}