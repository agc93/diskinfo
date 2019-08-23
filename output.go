package main

import (
	"fmt"
	"os"

	"github.com/dustin/go-humanize"
	"github.com/olekukonko/tablewriter"
)

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

func printSimple(disks []Disk) {
	for _, disk := range disks {
		fmt.Println(fmt.Sprintf("%s: %s (%s)	%s (%s)	%s", disk.devicePath, disk.model, disk.serial, disk.deviceNumber, disk.speed, humanize.Bytes(uint64(disk.size))))
	}
}

func printCsv(disks []Disk) {
	fmt.Print("Disk,Device,Model,Serial Number,Link Speed,Size\r\n")
	for _, disk := range disks {
		fmt.Printf("%s,%s,%s,%s,%s,%d\r\n", disk.devicePath, disk.deviceNumber, disk.model, disk.serial, disk.speed, disk.size)
	}
}
