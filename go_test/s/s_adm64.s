//func Swap(a,b int)(int,int)
TEXT ·Swap(SB), $0
    MOVQ a+0(FP), AX     // AX = a
    MOVQ b+8(FP), BX     // BX = b
    MOVQ BX, ret0+16(FP) // ret0 = BX
    MOVQ AX, ret1+24(FP) // ret1 = AX
    RET
