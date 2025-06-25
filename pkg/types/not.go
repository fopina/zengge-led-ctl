package types

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"strings"

	c "github.com/fopina/zengge-led-ctl/pkg/constants"
)

// NotificationPayload ...
type NotificationPayload struct {
	Device      byte
	MAC         net.HardwareAddr
	On          bool
	Mode        byte
	Brightness  byte
	RGB         *RGBColor
	Temperature uint8
	LEDCount    uint8
}

func NewNotificationPayload(data []byte) *NotificationPayload {
	if len(data) != 14 {
		return nil
	}
	if data[0] != 0x81 {
		return nil
	}
	d := &NotificationPayload{
		Device:      data[1],
		On:          data[2] == c.PowerOnByte,
		Mode:        data[4],
		Brightness:  data[5],
		RGB:         NewRGBColorBytes(data[6:9]),
		Temperature: data[9],
		LEDCount:    data[11],
	}
	return d
}

func (z *NotificationPayload) String() string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("Device: %X ", z.Device))
	if z.On {
		b.WriteString("ON ")
	} else {
		b.WriteString("OFF ")
	}
	b.WriteString(fmt.Sprintf("Mode: %x Brightness: %d RGB: %v Temperature: %d LEDs: %d", z.Mode, z.Brightness, z.RGB, z.Temperature, z.LEDCount))
	return b.String()
}

// NotificationDetails ...
type NotificationDetails struct {
	Code       int    `json:"code"`
	PayloadRaw string `json:"payload"`
	Payload    *NotificationPayload
}

func NewNotificationDetails(data []byte) *NotificationDetails {
	if len(data) < 10 {
		return nil
	}
	var resp NotificationDetails

	jsonData := data[8:]
	err := json.Unmarshal([]byte(jsonData), &resp)
	if err != nil {
		// likely something else, consumer should just use Notification.Raw
		return nil
	}

	bytes, err := hex.DecodeString(resp.PayloadRaw)
	if err == nil {
		resp.Payload = NewNotificationPayload(bytes)
	}

	return &resp
}

func (z NotificationDetails) String() string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("Code: %d ", z.Code))
	if z.Payload == nil {
		b.WriteString(fmt.Sprintf("[Payload: %s]", z.PayloadRaw))
	} else {
		b.WriteString(fmt.Sprintf("[%s]", z.Payload))
	}
	return b.String()
}

// Notification ...
type Notification struct {
	Raw     []byte
	Details *NotificationDetails
}

func NewNotification(data []byte) Notification {
	return Notification{Raw: data, Details: NewNotificationDetails(data)}
}

func (z Notification) String() string {
	if z.Details == nil {
		return fmt.Sprintf("[%X]", z.Raw)
	} else {
		return z.Details.String()
	}
}
