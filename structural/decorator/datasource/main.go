package main

import (
	"DesignPatterns/structural/decorator/datasource/compress"
	datasource "DesignPatterns/structural/decorator/datasource/data_source"
	"DesignPatterns/structural/decorator/datasource/encrypt"
)

func main() {
	data := "[ORIGINAL DATA]"
	source := datasource.NewFileDataSource()
	compressDecoratedDataSource := compress.NewCompressDecorator(source)
	encryptDecoratedDataSource := encrypt.NewEncryptDecorator(compressDecoratedDataSource)

}
