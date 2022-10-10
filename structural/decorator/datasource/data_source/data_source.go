package datasource

type DataSource interface {
	Write(data string)
	Read() string
}

type FileDataSource struct {
	storage string
}

func NewFileDataSource() *FileDataSource {
	return &FileDataSource{}
}

func (f *FileDataSource) Write(data string) {
	f.storage = data
}

func (f *FileDataSource) Read() string {
	return f.storage
}
