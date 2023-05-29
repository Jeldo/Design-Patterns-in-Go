package house

type House struct {
	windowType string
	doorType   string
	floor      int
}

func (h *House) WindowType() string {
	return h.windowType
}

func (h *House) DoorType() string {
	return h.doorType
}

func (h *House) Floor() int {
	return h.floor
}

type Builder interface {
	SetWindowType()
	SetDoorType()
	SetFloor()
	GetHouse() House
}

func GetBuilder(builderType string) Builder {
	if builderType == "normal" {
		return &NormalBuilder{}
	}
	if builderType == "igloo" {
		return &IglooBuilder{}
	}
	return nil
}

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func NewNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) SetWindowType() {
	b.windowType = "Wooden Window"
}

func (b *NormalBuilder) SetDoorType() {
	b.doorType = "Wooden Door"
}

func (b *NormalBuilder) SetFloor() {
	b.floor = 2
}

func (b *NormalBuilder) GetHouse() House {
	return House{
		windowType: b.windowType,
		doorType:   b.doorType,
		floor:      b.floor,
	}
}

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func NewIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) SetWindowType() {
	b.windowType = "Snow Window"
}

func (b *IglooBuilder) SetDoorType() {
	b.doorType = "Snow Door"
}

func (b *IglooBuilder) SetFloor() {
	b.floor = 1
}

func (b *IglooBuilder) GetHouse() House {
	return House{
		windowType: b.windowType,
		doorType:   b.doorType,
		floor:      b.floor,
	}
}
