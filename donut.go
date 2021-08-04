// Port of donut.c, as presented at a1k0n.net/2006/09/15/obfuscated-c-donut.html.
// The animation requires ANSI- or VT100-like terminal emulation.
//
// The math is beautifully explained at a1k0n.net/2011/07/20/donut-math.html.
package main

import (
	"fmt"
	"math"
	"time"
)

const (
	tau    = 2 * math.Pi // τ
	dTheta = 0.07        // θ spacing
	dPhi   = 0.02        // φ spacing
	dstK1  = 2           // Screen distance, controls the scale of the animation.
	dstK2  = 5           // Viewer distance from torus, controls the depth of the animation.

	charset = ".,-~:;=!*#$@" // ASCII characters, from dimmest to brightest.

	frameTime = 32 * time.Millisecond // How long to wait between frame transitions.
)

func main() {
	var (
		rotationX float64       // Rotation about the X-axis.
		rotationZ float64       // Rotation about the Z-axis.
		zBuff     [1760]float64 // Z-coordinate values in 3D space.
		outBuff   [1760]byte    // ASCII output buffer.
	)

	for {
		for i := 0; i < 1760; i++ { // clear output and Z-coordinate buffers.
			zBuff[i] = 0
			outBuff[i] = 32 // U+0020, empty.
		}

		sinX, cosX := math.Sin(rotationX), math.Cos(rotationX)
		sinZ, cosZ := math.Sin(rotationZ), math.Cos(rotationZ)

		for theta := 0.0; theta < tau; theta += dTheta { // cross-sectional circle of torus.
			sinTheta := math.Sin(theta)
			cosTheta := math.Cos(theta)

			for phi := 0.0; phi < tau; phi += dPhi { // center of revolution of torus.
				sinPhi := math.Sin(phi)
				cosPhi := math.Cos(phi)

				// calculate 3D coordinates, using a 3x3 matrix multiplication in disguise.
				h := cosTheta + dstK1
				z := sinPhi*h*sinX + sinTheta*cosX + dstK2
				invZ := 1 / z
				f := sinPhi*h*cosX - sinTheta*sinX
				x := int(40 + 30*invZ*(cosPhi*h*cosZ-f*sinZ))
				y := int(12 + 15*invZ*(cosPhi*h*sinZ+f*cosZ))

				if x <= 0 || x >= 80 || y <= 0 || y >= 22 { // out of our drawing frame range.
					continue
				}

				buffIndex := x + 80*y
				if invZ <= zBuff[buffIndex] { // pixel is closer to the viewer than what is already plotted.
					continue
				}

				// calculate luminance, ranging from -√2 to +√2 - ugly, but correct.
				lmn := (sinTheta*sinX-sinPhi*cosTheta*cosX)*cosZ - sinPhi*cosTheta*sinX - sinTheta*cosX - cosPhi*cosTheta*sinZ
				lmnIndex := int(8 * lmn) // clamp luminance to 0..11 (8*√(2) = 11.3), our ASCII charset range.

				if lmnIndex < 0 { // surface is pointing away from the viewer.
					lmnIndex = 0
				}

				zBuff[buffIndex] = invZ
				outBuff[buffIndex] = charset[lmnIndex]
			}
		}

		fmt.Print("\x1b[23A") // move cursor to beginning of drawing frame.

		for i := 0; i <= 1760; i++ {
			var p byte = 10 // U+000A (LF) if not within drawing frame range.
			if i%80 > 0 {
				p = outBuff[i]
			}

			fmt.Printf("%c", p)
		}

		rotationX += 0.04
		rotationZ += 0.02

		time.Sleep(frameTime)
	}
}
