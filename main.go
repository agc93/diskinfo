package main

import (
	"flag"
	"fmt"
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
	table.SetHeader([]string{"Disk", "Device", "Model", "Serial", "Size"})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Normal},
		tablewriter.Colors{tablewriter.Normal},
		tablewriter.Colors{tablewriter.Normal},
		tablewriter.Colors{tablewriter.Bold})
	for _, v := range disks {
		table.Append([]string{v.devicePath, v.deviceNumber, v.model, v.serial, humanize.Bytes(uint64(v.size))})
	}
	table.Render()
}

// GetDiskInfo returns info about the host disks
func GetDiskInfo(fsRoot string) ([]Disk, error) {
	if fsRoot == "" {
		fsRoot = "/"
	}
	block, err := ghw.Block(ghw.WithChroot(fsRoot))
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
		info = append(info, Disk{
			devicePath:   fmt.Sprintf("/dev/%s", disk.Name),
			serial:       disk.SerialNumber,
			size:         int64(disk.SizeBytes),
			model:        disk.Model,
			deviceNumber: GetDeviceNumber(disk.Name),
		})
	}
	return info, nil
}

func GetDeviceNumber(deviceName string) string {
	fi, err := os.Readlink(fmt.Sprintf("/sys/block/%s", deviceName))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// segments := strings.Split(strings.TrimPrefix(fi, "/"), "/")
	for _, v := range strings.Split(strings.TrimPrefix(fi, "/"), "/") {
		if strings.HasPrefix(v, "ata") {
			return v
		}
	}
	return ""
	// linkedName := From(segments).FirstWithT(func(s string) bool { return strings.HasPrefix(s, "ata") })
	// fmt.Println(linkedName)
	// return linkedName.data
	// return ""
}

type Disk struct {
	devicePath   string
	serial       string
	deviceNumber string
	size         int64
	model        string
}
