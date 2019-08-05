package cpu

type gpuTile struct {
	TileData [][]byte
}

type gpuTileAttr struct {
	BackgroundPallete int
	TileVRAMBank      int
	HorizontalFlip    bool
	VerticalFlip      bool
	BGtoOamPriority   bool
}

func tileAttr(attr uint8) gpuTileAttr {
	g := gpuTileAttr{}

	g.BackgroundPallete = int(attr & 0x7)
	g.TileVRAMBank = int(attr&8) >> 3

	g.HorizontalFlip = attr&0x20 > 0
	g.VerticalFlip = attr&0x40 > 0
	g.BGtoOamPriority = attr&0x80 > 0

	return g
}

func makeGPUTile() gpuTile {
	t := gpuTile{
		TileData: make([][]byte, 8),
	}

	for i := 0; i < 8; i++ {
		t.TileData[i] = make([]byte, 16)
	}

	return t
}
