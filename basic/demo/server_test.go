package basic

import (
	"flag"
	"testing"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"

	// "github.com/smallnest/go-network-programming/codec"
	"log"
	"net"
)

// https://colobu.com/2022/06/05/use-bpf-to-make-the-go-network-program-8x-faster/
var (
	addr = flag.String("s", "localhost", "server address")
	port = flag.Int("p", 8972, "port")
)
var (
	stat         = make(map[string]int)
	lastStatTime = int64(0)
)

// func main() {
func TestServer(t *testing.T) {
	flag.Parse()
	conn, err := net.ListenPacket("ip4:udp", *addr)
	if err != nil {
		panic(err)
	}
	cc := conn.(*net.IPConn)
	cc.SetReadBuffer(20 * 1024 * 1024)
	cc.SetWriteBuffer(20 * 1024 * 1024)
	handleConn(conn)
}

func handleConn(conn net.PacketConn) {
	for {
		buffer := make([]byte, 1024)
		n, remoteaddr, err := conn.ReadFrom(buffer)
		if err != nil {
			log.Fatal(err)
		}
		buffer = buffer[:n]
		packet := gopacket.NewPacket(buffer, layers.LayerTypeUDP, gopacket.NoCopy)
		// Get the UDP layer from this packet
		if udpLayer := packet.Layer(layers.LayerTypeUDP); udpLayer != nil {
			udp, _ := udpLayer.(*layers.UDP)
			if app := packet.ApplicationLayer(); app != nil {
				// data, err := codec.EncodeUDPPacket(net.ParseIP("127.0.0.1"), net.ParseIP("127.0.0.1"), uint16(udp.DstPort), uint16(udp.SrcPort), app.Payload())
				data, err := EncodeUDPPacket(net.ParseIP("127.0.0.1"), net.ParseIP("127.0.0.1"), uint16(udp.DstPort), uint16(udp.SrcPort), app.Payload())
				if err != nil {
					log.Printf("failed to EncodePacket: %v", err)
					return
				}
				if _, err := conn.WriteTo(data, remoteaddr); err != nil {
					log.Printf("failed to write packet: %v", err)
					conn.Close()
					return
				}
			}
		}
	}
}
