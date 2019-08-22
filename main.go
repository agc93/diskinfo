package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	. "github.com/ahmetb/go-linq"
	"github.com/dustin/go-humanize"
	"github.com/jaypipes/ghw"
	"github.com/olekukonko/tablewriter"
)

func main() {
	// outputPtr := flag.String("o", "simple", "Output format {simple|lines|table};.")
	flag.Parse()
	// input := flag.Args()

	// if len(input) != 1 {
	// 	fmt.Println("No input text detected! Pass the full comma-separated CHAR() text as an argument")
	// 	flag.PrintDefaults()
	// 	os.Exit(1)
	// }

	disks, err := GetDiskInfo("/")
	if err != nil {
		return
	}
	printTable(disks)

}

func printTable(disks []Disk) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Disk", "Device", "Model", "Serial", "Speed", "Size"})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Normal},
		tablewriter.Colors{tablewriter.Normal},
		tablewriter.Colors{tablewriter.Normal},
		tablewriter.Colors{tablewriter.Normal},
		tablewriter.Colors{tablewriter.Bold})
	for _, v := range disks {
		table.Append([]string{v.devicePath, v.deviceNumber, v.model, v.serial, v.speed, humanize.Bytes(uint64(v.size))})
	}
	table.Render()
}

// GetDiskInfo returns info about the host disks
func GetDiskInfo(fsRoot string) ([]Disk, error) {
	if fsRoot == "" {
		fsRoot = "/"
	}
	// block, err := ghw.Block(ghw.WithChroot(fsRoot))
	block, err := ghw.Block()
	if err != nil {
		return nil, err
	}

	var info []Disk
	var availableDisks []*ghw.Disk

	From(block.Disks).WhereT(
		func(d *ghw.Disk) bool {
			return d.SerialNumber != "unknown"
		}).ToSlice(&availableDisks)

	for _, disk := range availableDisks {
		devNum := GetDeviceNumber(disk.Name)
		info = append(info, Disk{
			devicePath:   fmt.Sprintf("/dev/%s", disk.Name),
			serial:       disk.SerialNumber,
			size:         int64(disk.SizeBytes),
			model:        disk.Model,
			deviceNumber: devNum,
			speed:        GetLinkSpeed(devNum),
		})
	}
	return info, nil
}

func GetLinkSpeed(deviceNumber string) string {
	linkNumber := strings.TrimPrefix(deviceNumber, "ata")
	path := fmt.Sprintf("/sys/class/ata_link/link%s/sata_spd", linkNumber)
	spd, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(spd))
}

type Disk struct {
	devicePath   string
	serial       string
	deviceNumber string
	size         int64
	model        string
	speed        string
}
