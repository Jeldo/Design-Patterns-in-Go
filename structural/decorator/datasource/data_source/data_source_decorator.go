package datasource

type DataSourceDecorator struct {
	wrappee DataSource
}

func NewDataSourceDecorator(source DataSource) *DataSourceDecorator {
	return &DataSourceDecorator{wrappee: source}
}

func (d *DataSourceDecorator) Write(data string) {
	d.wrappee.Write(data)
}

func (d *DataSourceDecorator) Read() string {
	return d.wrappee.Read()
}
