package main

type Windows struct{}

func (w *Windows) InsertIntoUSBPort() {
	println("USB connector is plugged into windows machine.")
}
