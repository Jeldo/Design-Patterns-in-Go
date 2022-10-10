package compress

import (
	datasource "DesignPatterns/structural/decorator/datasource/data_source"
	"fmt"
	"strings"
)

const compressed = "[COMPRESSED]"

type CompressDecorator struct {
	ds datasource.DataSource
}

func NewCompressDecorator(ds datasource.DataSource) *CompressDecorator {
	return &CompressDecorator{ds: ds}
}

func (d *CompressDecorator) Write(data string) {
	fmt.Println("Before Write Operation of CompressDecorator: ", data)
	compressedData := fmt.Sprintf(compressed + data)
	d.ds.Write(compressedData)
}

func (d *CompressDecorator) Read() string {
	decompressedData := strings.Replace(d.ds.Read(), compressed, "", -1)
	fmt.Println("After Read Operation of CompressDecorator: ", decompressedData)
	return decompressedData
}
