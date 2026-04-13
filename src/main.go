package main

import "fmt"
import "time"

#include "cpu.gh"

func main() {

    const (
        BASE  = 0x200
        BLOCK = 32768
        LOOPS = 10000
    )

    mem := new(AddressSpace)
    mem[VEC_RES]     = BASE >> 8
    mem[VEC_RES + 1] = BASE & 0xFF
    
    mem[0xFF00] = RTS;
        
    var i int
    
    // Block of NOPS
    for i = 0; i <= BLOCK; i += 3 {
        mem[BASE + i] = JSR_AB
        mem[BASE + i + 1] = 0x00
        mem[BASE + i + 2] = 0xFF
    }

    mem[BASE + i] = 0xFF // illegal    

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
        "%d loops of %d operations Took %d nanoseconds [%f MIPS]\n",
        LOOPS,
        BLOCK,
        iNanoSeconds,
        fMIPS,
    )
    cpu.ShowStatus();
}
