package cpu

type gpuTile struct {
	TileData [][]byte
}

func makeGPUTile() gpuTile {
	t := gpuTile{
		TileData: make([][]byte, 8),
	}

	for i := 0; i < 8; i++ {
		t.TileData[i] = make([]byte, 8)
	}

	return t
}
