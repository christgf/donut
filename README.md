# Have a donut

Go port of the original [donut.c](https://www.a1k0n.net/2006/09/15/obfuscated-c-donut.html) with some minor tweaks.

```
                    for{fori:=0;i
                <1760;i++{zb[i]=0b[i]=
              32}sx,cx:=math.Sin(rx),mat
            h.Cos(rx)sz,cz:=math.Sin(rz),m
          ath.Cos(rz)fort:=0.0;t<6.48;t+=0.0
         7{st:=math.Sin(t)ct:=math.Cos(t)forp
        :=0.0;p<6.48;p+=0.02{sp:=math.Sin(p)cp
        :=math.Cos(p)h:=ct+2zz:=1/(sp*h*sx+st*
        cx+5)f:=sp*h*cx        -st*sxx:=int(40
       +30*zz*(cp*h*cz          -f*sz))y:=int(1
       2+15*zz*(cp*h*s          z+f*cz))ifx<=0|
        |x>=80||y<=0||y        >=22{continue}b
        i:=x+80*yifzz<=zb[bi]{continue}lm:=(st
        *sx-sp*ct*cx)*cz-sp*ct*sx-st*cx-cp*ct*
         szl:=int(8*lm)ifl<0{l=0}zb[bi]=zzb[b
          i]=".,-~:;=!*#$@"[l]}}}fmt.Print("
            \x1b[23A")fori:=0;i<=1760;i++{
              varpbyte=10ifi%80>0{p=b[i]
                }fmt.Printf("%c",p)}rx
                    +=0.04rz+=0.02
```

Run with `go run donut.go` - requires ANSI- or VT100-like terminal emulation.

The math is beautifully explained [here](https://www.a1k0n.net/2011/07/20/donut-math.html).