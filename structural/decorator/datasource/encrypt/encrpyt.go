package encrypt

import (
	datasource "DesignPatterns/structural/decorator/datasource/data_source"
	"fmt"
	"strings"
)

const encrypted = "[ENCRYPTED]"

type EncryptDecorator struct {
	ds datasource.DataSource
}

func NewEncryptDecorator(ds datasource.DataSource) *EncryptDecorator {
	return &EncryptDecorator{ds: ds}
}

func (d *EncryptDecorator) Write(data string) {
	fmt.Println("Before Write Operation of EncryptDecorator: ", data)
	encryptedData := fmt.Sprintf(encrypted + data)
	d.ds.Write(encryptedData)
}

func (d *EncryptDecorator) Read() string {
	decryptedData := strings.Replace(d.ds.Read(), encrypted, "", -1)
	fmt.Println("After Read Operation of EncryptDecorator: ", decryptedData)
	return decryptedData
}
