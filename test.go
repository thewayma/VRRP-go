package main

import (
	"VRRP/VRRP"
	"flag"
	"fmt"
	"net"
	"time"
)

var (
	VRID     int
	Priority int
)

func init() {
	flag.IntVar(&VRID, "vrid", 233, "virtual router ID")
	flag.IntVar(&Priority, "pri", 100, "router priority")
}

func main() {
	flag.Parse()
	var vr = VRRP.NewVirtualRouter(byte(VRID), "ens34", "ens35", false, VRRP.IPv4)

	vip := net.IPv4(10, 0, 0, 27)
	vr.AddIPvXAddr(vip)

	vr.SetPriorityAndMasterAdvInterval(byte(Priority), time.Millisecond*800)
	vr.Enroll(VRRP.Backup2Master, func() {
		fmt.Println("init to master")
	})
	vr.Enroll(VRRP.Master2Init, func() {
		fmt.Println("master to init")
	})
	vr.Enroll(VRRP.Master2Backup, func() {
		fmt.Println("master to backup")
	})
	go func() {
		time.Sleep(time.Minute * 5)
		vr.Stop()
	}()
	vr.StartWithEventSelector()

}
