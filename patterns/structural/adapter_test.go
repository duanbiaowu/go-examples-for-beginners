package structural

import "testing"

func Test_Adapter(t *testing.T) {
	c := &client{}
	mac := &Mac{}

	c.insertLightningConnectorIntoComputer(mac)

	win := &Windows{}
	winAdapter := &windowsAdapter{win}

	c.insertLightningConnectorIntoComputer(winAdapter)
}
