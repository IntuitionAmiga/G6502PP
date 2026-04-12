package main

import "fmt"
import "time"

#include "cpu.gh"

func main() {

    const (
        BASE  = 0x200
        BLOCK = 32768
        LOOPS = 1
    )

    mem := new(AddressSpace)
    mem[VEC_RES]     = BASE >> 8
    mem[VEC_RES + 1] = BASE & 0xFF
    
    // Block of NOPS
    for i:= 0; i < BLOCK; i++ {
        mem[BASE + i] = NOP
    }
    mem[BASE + BLOCK] = BEQ
    mem[BASE + BLOCK + 1] = 0x4
    mem[BASE + BLOCK + 2] = 0x0F
    mem[BASE + BLOCK + 3] = 0x1F
    mem[BASE + BLOCK + 4] = 0x2F
    mem[BASE + BLOCK + 5] = 0x3F
    mem[BASE + BLOCK + 6] = 0x4F
    mem[BASE + BLOCK + 7] = 0x5F
    mem[BASE + BLOCK + 8] = 0x6F
    mem[BASE + BLOCK + 9] = 0x8F
    mem[BASE + BLOCK + 10] = 0x9F
    mem[BASE + BLOCK + 11] = 0xAF
    mem[BASE + BLOCK + 12] = 0xBF
    mem[BASE + BLOCK + 13] = 0xCF
    mem[BASE + BLOCK + 14] = 0xDF
    mem[BASE + BLOCK + 15] = 0xEF
    mem[BASE + BLOCK + 16] = 0xBF

    cpu := &MOS6502{}
    cpu.Init(mem).ShowStatus();
        
    tStart := time.Now()
    for i := 0; i < LOOPS; i++ {
        cpu.RunFrom(0x200)
    }
    tElapsed := time.Since(tStart)
    iNanoSeconds := uint64(tElapsed.Nanoseconds())
    fMIPS := 1.0e3 * float64(LOOPS * BLOCK) / float64(iNanoSeconds)

    fmt.Printf(
        "%d loops of %d NOP Took %d nanoseconds [%f MIPS]\n",
        LOOPS,
        BLOCK,
        iNanoSeconds,
        fMIPS,
    )
    cpu.ShowStatus();
}
