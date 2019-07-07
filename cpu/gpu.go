package cpu

import (
	"github.com/faiface/pixel"
	"github.com/racerxdl/goboy/pixhelp"
	"image"
	"image/color"
	"sort"
)

type GPU struct {
	mode                         GPUMode
	cpu                          *Core
	modeClocks                   int
	line                         byte
	scrollX, scrollY, winX, winY int
	switchBg                     bool
	switchLCD                    bool
	objSize                      bool
	switchObj                    bool
	switchWin                    bool
	lineCompare                  byte
	lcdStat                      byte
	interruptsFired              byte
	currentRow                   []byte
	bgTileBase                   uint16
	bgMapBase                    uint16
	winMapBase                   uint16
	oam                          []byte
	tileBuffer                   *pixel.PictureData
	registers                    []byte
	tileSet                      []gpuTile
	vramBuffer                   *pixel.PictureData
	vram                         []byte
	objs                         []gpuObject
	prioObjs                     []gpuObject
	bgPallete                    [4]color.RGBA
	obj0Pallete                  [4]color.RGBA
	obj1Pallete                  [4]color.RGBA
}

func (g *GPU) GetVRAM() *pixel.PictureData {
	return g.vramBuffer
}

func (g *GPU) GetTileBuffer() *pixel.PictureData {
	return g.tileBuffer
}

func (g *GPU) LycLy() bool {
	return (g.lcdStat & FlagLycLy) > 0
}

func (g *GPU) OamMode() bool {
	return (g.lcdStat & FlagOamMode) > 0
}

func (g *GPU) VBlankMode() bool {
	return (g.lcdStat & FlagVblankMode) > 0
}

func (g *GPU) HBlankMode() bool {
	return (g.lcdStat & FlagHblankMode) > 0
}

func MakeGPU(cpu *Core) *GPU {
	gpu := &GPU{
		cpu:        cpu,
		tileBuffer: pixel.PictureDataFromImage(image.NewRGBA(image.Rect(0, 0, 144, 288))),
	}

	for i := 0; i < len(gpu.tileBuffer.Pix); i++ {
		var x = i % gpu.tileBuffer.Stride
		var y = i / gpu.tileBuffer.Stride
		if (x%9 == 8) || (y%9 == 8) {
			gpu.tileBuffer.Pix[i] = pixhelp.ToRGBA(color.Transparent)
		} else {
			gpu.tileBuffer.Pix[i] = pixhelp.ToRGBA(color.White)
		}
	}

	gpu.registers = make([]byte, 0xFF)
	gpu.currentRow = make([]byte, 160)
	gpu.vram = make([]byte, 0x2000)
	for i := 0; i < 160; i++ {
		gpu.currentRow[i] = 0x00
	}
	gpu.Reset()

	return gpu
}

func (g *GPU) Reset() {
	g.modeClocks = 0
	g.scrollX = 0
	g.scrollY = 0
	g.winX = 0
	g.winY = 0
	g.line = 0
	g.mode = OamRead
	g.tileSet = make([]gpuTile, 512)
	for i := 0; i < 512; i++ {
		g.tileSet[i] = makeGPUTile()
	}

	for i := 0; i < 0x2000; i++ {
		g.vram[i] = 0x00
	}

	g.vramBuffer = pixel.PictureDataFromImage(image.NewRGBA(image.Rect(0, 0, 256, 256)))
	pixhelp.ClearPictureData(g.vramBuffer, color.Black)
	g.oam = make([]byte, 160)

	g.switchLCD = true
	g.switchBg = false
	g.switchWin = false
	g.objSize = false

	g.objs = make([]gpuObject, 40)
	g.prioObjs = make([]gpuObject, 40)

	for i := 0; i < 40; i++ {
		g.objs[i] = gpuObject{
			Pos:     i,
			Y:       -16,
			X:       -8,
			Tile:    0,
			Palette: 0,
			YFlip:   false,
			XFlip:   false,
		}
		g.prioObjs[i] = g.objs[i]
	}

	g.switchObj = false
	g.lineCompare = 0
	g.lcdStat = 0
	g.interruptsFired = 0
	g.bgTileBase = 0x0000
	g.bgMapBase = 0x1800
	g.winMapBase = 0x1800

	copy(g.bgPallete[:], defaultBgPallete)
	copy(g.obj0Pallete[:], defaultObj0Pallete)
	copy(g.obj1Pallete[:], defaultObj1Pallete)
}

func (g *GPU) state() uint8 {
	state := uint8(0)

	if g.switchBg {
		state |= 0x01
	}

	if g.switchObj {
		state |= 0x02
	}

	if g.objSize {
		state |= 0x04
	}

	if g.bgMapBase == 0x1C00 {
		state |= 0x08
	}

	if g.bgTileBase == 0x0000 {
		state |= 0x10
	}

	if g.switchWin {
		state |= 0x20
	}

	if g.winMapBase == 0x1C00 {
		state |= 0x40
	}

	if g.switchLCD {
		state |= 0x80
	}

	return state
}

func (g *GPU) Read(addr uint16) byte {
	if addr < 0xFF40 { // GPU Memory
		switch {
		case addr >= 0x8000 && addr <= 0x9FFF:
			return g.vram[addr-0x8000]
		case addr >= 0xFE00 && addr <= 0xFE9F:
			return g.oam[addr-0xFE00]
		}

		return 0x00
	}

	// GPU Registers
	switch addr {
	case 0xFF40:
		return g.state()
	case 0xFF41:
		ift := g.interruptsFired
		g.interruptsFired = 0x00
		res := uint8(0)
		if g.mode&0x3 > 0 {
			res |= 0x01
		}

		if g.line == g.lineCompare {
			res |= 0x04
		}

		res |= ift << 3
		res |= 0x80

		return res
	case 0xFF42:
		return uint8(g.scrollY)
	case 0xFF43:
		return uint8(g.scrollX)
	case 0xFF44:
		return g.line
	case 0xFF45:
		return g.lineCompare
	case 0xFF4A:
		return uint8(g.winY)
	case 0xFF4B:
		return uint8(g.winX)
	default:
		return g.registers[addr-0xFF40]
	}
}

func (g *GPU) Write(addr uint16, val uint8) {
	if addr < 0xFF40 { // GPU Memory
		switch {
		case addr >= 0x8000 && addr <= 0x9FFF: // Video RAM
			g.vram[addr-0x8000] = val
			g.updateTile(addr, val)
		case addr >= 0xFE00 && addr <= 0xFE9F:
			g.oam[addr-0xFE00] = val
			g.UpdateOAM(addr, val)
		}
		return
	}

	// GPU Registers
	g.registers[addr-0xFF40] = val

	switch addr {
	case 0xFF40:
		g.switchBg = (val & 0x01) > 0
		g.switchObj = (val & 0x02) > 0
		g.objSize = (val & 0x04) > 0

		if val&0x08 > 0 {
			g.bgMapBase = 0x1C00
		} else {
			g.bgMapBase = 0x1800
		}

		if val&0x10 > 0 {
			g.bgTileBase = 0x0000
		} else {
			g.bgTileBase = 0x0800
		}

		g.switchWin = (val & 0x20) > 0

		if val&0x40 > 0 {
			g.winMapBase = 0x1C00
		} else {
			g.winMapBase = 0x1800
		}

		g.switchLCD = (val & 0x80) > 0
	case 0xFF41:
		g.lcdStat = val & 0x78
	case 0xFF42:
		g.scrollY = int(val)
	case 0xFF43:
		g.scrollX = int(val)
	case 0xFF45:
		g.lineCompare = val
	case 0xFF46:
		for i := 0; i < 160; i++ { // DMA
			v := g.cpu.Memory.ReadByte((uint16(val) << 8) + uint16(i))
			g.oam[i] = v
			g.updateOAM(uint16(0xFE00+i), v)
		}
		g.sortOAM()
	case 0xFF47:
		for i := uint(0); i < 4; i++ {
			var b = (uint(val) >> (i * 2)) & 3
			switch b {
			case 0:
				g.bgPallete[i] = color.RGBA{R: 255, G: 255, B: 255, A: 255}
			case 1:
				g.bgPallete[i] = color.RGBA{R: 192, G: 192, B: 192, A: 255}
			case 2:
				g.bgPallete[i] = color.RGBA{R: 96, G: 96, B: 96, A: 255}
			case 3:
				g.bgPallete[i] = color.RGBA{A: 255}
			}
		}

		g.refreshTileData(-1)
	case 0xFF48:
		for i := uint(0); i < 4; i++ {
			var b = (uint(val) >> (i * 2)) & 3
			switch b {
			case 0:
				g.obj0Pallete[i] = color.RGBA{R: 255, G: 255, B: 255, A: 255}
			case 1:
				g.obj0Pallete[i] = color.RGBA{R: 192, G: 192, B: 192, A: 255}
			case 2:
				g.obj0Pallete[i] = color.RGBA{R: 96, G: 96, B: 96, A: 255}
			case 3:
				g.obj0Pallete[i] = color.RGBA{A: 255}
			}
		}
	case 0xFF49:
		for i := uint(0); i < 4; i++ {
			var b = (uint(val) >> (i * 2)) & 3
			switch b {
			case 0:
				g.obj1Pallete[i] = color.RGBA{R: 255, G: 255, B: 255, A: 255}
			case 1:
				g.obj1Pallete[i] = color.RGBA{R: 192, G: 192, B: 192, A: 255}
			case 2:
				g.obj1Pallete[i] = color.RGBA{R: 96, G: 96, B: 96, A: 255}
			case 3:
				g.obj1Pallete[i] = color.RGBA{A: 255}
			}
		}
	case 0xFF4A:
		g.winY = int(val)
	case 0xFF4B:
		g.winX = int(val)
	}
}

func (g *GPU) updateOAM(addr uint16, val uint8) {
	relAddr := addr - 0xFE00
	obj := relAddr >> 2
	if obj < 40 {
		switch relAddr & 3 {
		case 0:
			g.objs[obj].Y = int(val) - 16
		case 1:
			g.objs[obj].X = int(val) - 8
		case 2:
			if g.objSize {
				g.objs[obj].Tile = val & 0xFE
			} else {
				g.objs[obj].Tile = val
			}
		case 3:
			if val&0x10 != 0 {
				g.objs[obj].Palette = 1
			} else {
				g.objs[obj].Palette = 0
			}
			g.objs[obj].XFlip = (val & 0x20) != 0
			g.objs[obj].YFlip = (val & 0x40) != 0
			g.objs[obj].Prio = (val & 0x80) != 0
		}
	}
}

func (g *GPU) UpdateOAM(addr uint16, val uint8) {
	g.updateOAM(addr, val)
	g.sortOAM()
}

func (g *GPU) sortOAM() {
	copy(g.prioObjs, g.objs)
	sort.SliceStable(g.prioObjs, func(i, j int) bool {
		a := g.prioObjs[i]
		b := g.prioObjs[j]

		if a.X < b.X {
			return true
		}

		if a.Pos < b.Pos {
			return true
		}

		return false
	})

}

func (g *GPU) refreshTileData(tileNum int) {
	if tileNum == -1 {
		// Refresh ALL
		for i, v := range g.tileSet {
			// 16 x 32 tiles with 1px spacing
			// 16 * 9 x 32 * 9
			// 144 x 288 Buffer
			for y := 0; y < 8; y++ {
				for x := 0; x < 8; x++ {
					px := (i%16)*9 + x
					py := (i/16)*9 + y
					p := py*g.tileBuffer.Stride + px

					g.tileBuffer.Pix[p] = g.bgPallete[v.TileData[y][x]]
				}
			}
		}
	} else {
		// Refresh single Tile
		i := tileNum
		v := g.tileSet[tileNum]

		// 16 x 32 tiles with 1px spacing
		// 16 * 9 x 32 * 9
		// 144 x 288 Buffer
		for y := 0; y < 8; y++ {
			for x := 0; x < 8; x++ {
				px := (i%16)*9 + x
				py := (i/16)*9 + y
				p := py*g.tileBuffer.Stride + px

				g.tileBuffer.Pix[p] = g.bgPallete[v.TileData[y][x]]
			}
		}
	}
}

func (g *GPU) renderScanline() {
	if g.switchLCD {
		// region Background Draw
		vFrame := g.cpu.Memory.GetVideoFrame()
		if g.switchBg || g.switchWin {
			bufferOffset := int(g.line) * vFrame.Stride

			// region Background Offset Compute
			bgVramOffset := VRamBase
			bgVramOffset += int(g.bgMapBase)
			bgVramOffset += (((int(g.line) + g.scrollY) & 0xFF) / 8) * 32

			bgY := (int(g.line) + g.scrollY) % 8
			bgX := g.scrollX % 8
			bgTileOffset := (g.scrollX / 8) % 32
			// endregion
			// region Window Offset Compute
			winVramOffset := VRamBase
			winVramOffset += int(g.winMapBase)
			winVramOffset += (((int(g.line) + g.winY) & 0xFF) / 8) * 32

			wY := (int(g.line) + g.winY) % 8
			wX := g.winX % 8
			wTileOffset := (g.winX / 8) % 32
			// endregion

			x := 0
			y := 0
			tileOffset := 0
			vramOffset := 0

			if g.switchWin {
				x = wX
				y = wY
				tileOffset = wTileOffset
				vramOffset = winVramOffset
			} else {
				x = bgX
				y = bgY
				tileOffset = bgTileOffset
				vramOffset = bgVramOffset
			}

			tile := int(g.cpu.Memory.ReadByte(uint16(vramOffset + tileOffset)))

			if g.bgTileBase != 0x0000 && tile < 128 {
				tile += 256
			}

			tileRow := g.tileSet[tile].TileData[y]

			for i := 0; i < 160; i++ {
				c := g.bgPallete[tileRow[x]]
				g.currentRow[i] = tileRow[x]
				vFrame.Pix[bufferOffset] = c
				bufferOffset++
				x++
				if x != 8 {
					continue
				}

				x = 0
				tileOffset = (tileOffset + 1) % 32
				tile := int(g.cpu.Memory.ReadByte(uint16(vramOffset + tileOffset)))
				if g.bgTileBase != 0x0000 && tile < 128 {
					tile += 256
				}

				tileRow = g.tileSet[tile].TileData[y]
			}
		}
		// endregion
		// region Object Draw
		if g.switchObj {
			spriteCount := 0
			iline := int(g.line)
			for i := 0; i < 40; i++ {
				obj := g.prioObjs[i]

				if spriteCount > 10 {
					break
				}

				if obj.X < 0 || obj.X >= 168 {
					continue
				}

				if obj.Y < 0 || obj.Y >= 160 {
					continue
				}

				if obj.Y <= iline && (obj.Y+8) > iline {
					var tileRow []byte
					tileData := g.tileSet[obj.Tile]

					if !obj.YFlip {
						tileRow = tileData.TileData[iline-obj.Y]
					} else {
						tileRow = tileData.TileData[7-(iline-obj.Y)]
					}

					pallete := g.obj0Pallete

					if obj.Palette != 0 {
						pallete = g.obj1Pallete
					}

					bufferOffset := (iline * vFrame.Stride) + obj.X
					var c color.RGBA
					for x := 0; x < 8; x++ {
						if obj.XFlip {
							c = pallete[tileRow[x]]
						} else {
							c = pallete[tileRow[7-x]]
						}

						if tileRow[x] != 0x00 && obj.X+x >= 0 && obj.X+x < 160 && (obj.Prio || g.currentRow[x] == 0x00) {
							vFrame.Pix[bufferOffset] = c
						}
						bufferOffset++
					}
					spriteCount++
				}
			}
		}
		// endregion
	}
}

func (g *GPU) UpdateVRAM() {
	for i := range g.vramBuffer.Pix {
		py := i / 256
		px := i % 256
		tileNum := (px / 8) + ((py / 8) * 32)
		v := g.cpu.Memory.ReadByte(uint16(VRamBase + int(g.bgMapBase) + tileNum))
		tile := g.tileSet[v]
		x := px % 8
		y := py % 8
		g.vramBuffer.Pix[i] = g.bgPallete[tile.TileData[y][x]]
	}
}

func (g *GPU) updateTile(addr uint16, val uint8) {
	relAddr := addr & 0x1FFF
	if (addr & 1) > 0 {
		addr--
		relAddr--
	}

	tile := (relAddr >> 4) & 511
	y := (relAddr >> 1) & 7

	b0 := g.cpu.Memory.ReadByte(addr)
	b1 := g.cpu.Memory.ReadByte(addr + 1)

	for x := 0; x < 8; x++ {
		sx := uint8(1 << (7 - uint(x)))
		o := uint8(0)

		if b0&sx != 0 {
			o += 1
		}

		if b1&sx != 0 {
			o += 2
		}

		g.tileSet[tile].TileData[y][x] = o
	}
	g.refreshTileData(int(tile))
}

func (g *GPU) Cycle() {
	g.modeClocks += g.cpu.Registers.LastClockM
	switch g.mode {
	case HBlank:
		if g.modeClocks > horizontalBlankCycles {
			g.modeClocks = 0
			g.line++

			if g.line == 144 {
				g.mode = VBlank

				g.cpu.Registers.TriggerInterrupts |= IntVblank
				if g.VBlankMode() && g.cpu.Registers.InterruptEnable {
					g.cpu.Registers.TriggerInterrupts |= IntLcdstat
				}
			} else {
				g.mode = OamRead
				if g.OamMode() && g.cpu.Registers.InterruptEnable {
					g.cpu.Registers.TriggerInterrupts |= IntLcdstat
				}
			}

			if g.line == g.lineCompare && g.LycLy() && g.cpu.Registers.InterruptEnable {
				g.cpu.Registers.TriggerInterrupts |= IntLcdstat
			}
		}
	case VBlank:
		if g.modeClocks >= (verticalBlankCycles / 9) {
			g.modeClocks = 0
			g.line++
			if g.line == g.lineCompare && g.LycLy() {
				g.cpu.Registers.TriggerInterrupts |= IntLcdstat
			}
			if g.line > 153 {
				g.mode = OamRead
				g.line = 0
				if g.OamMode() && g.cpu.Registers.InterruptEnable {
					g.cpu.Registers.TriggerInterrupts |= IntLcdstat
				}
			}
		}
	case OamRead:
		if g.modeClocks >= oamCycles {
			g.modeClocks = 0
			g.mode = VramRead
		}
	case VramRead:
		if g.modeClocks >= vRamCycles {
			g.modeClocks = 0
			g.renderScanline()
			g.mode = HBlank

			// TODO: DMA

			if g.HBlankMode() && g.cpu.Registers.InterruptEnable {
				g.cpu.Registers.TriggerInterrupts |= IntLcdstat
			}
		}
	}
}
