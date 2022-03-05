package wol

import (
	"errors"
	"net"
)


func Wake(macAddr string, broadcastIP string) error {
	var packet [102]byte
	mac, err := net.ParseMAC(macAddr)
	if err != nil {
		return err
	}

	if len(mac) != 6 {
		return errors.New("invalid MAC address")
	}

	copy(packet[0:], []byte{255, 255, 255, 255, 255, 255})
	offset := 6

	for i := 0; i < 16; i++ {
		copy(packet[offset:], mac)
		offset += 6
	}

	conn, err := net.Dial("udp", broadcastIP+":9")
	if err != nil {
		return err
	}
	defer conn.Close()

	conn.Write(packet[:])
	return nil
}
