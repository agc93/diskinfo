# diskinfo

## Introduction

A simple Go CLI used to show a quick and dirty summary of the installed disks in a Linux system. In particular, this is very useful for servers with large disk counts.

Never be confused as to whether `/dev/sdb` is `ata2` or `ata3` again.

## Usage

To run, simply run `./diskinfo` and you'll get a simple summary shown:

```bash
> ./diskinfo
+----------+--------+-------------------------+--------------+----------+--------+
|   DISK   | DEVICE |          MODEL          |    SERIAL    |  SPEED   |  SIZE  |
+----------+--------+-------------------------+--------------+----------+--------+
| /dev/sda | ata2   | TOSHIBA_THNSNJ512GDNU_A | X61S10DXXXXX | 6.0 Gbps | 512 GB |
+----------+--------+-------------------------+--------------+----------+--------+
```

`diskinfo` also supports a single `-o` option to control the output. Available outputs are `table` (default, shown above), `simple` (more compact and copy-paste friendly output) and `csv` (outputs CSV format, best combined with `tee`).

```bash
> ./diskinfo -o table
> ./diskinfo -o simple
> ./diskinfo -o csv
```
