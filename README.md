# diskinfo

## Introduction

A simple Go CLI used to show a quick and dirty summary of the installed disks in a Linux system. In particular, this is very useful for servers with large disk counts.

Never be confused as to whether `/dev/sdb` is `ata2` or `ata3` again.

## Usage

To run, simply run `./diskinfo` and you'll get a simple summary shown:

```bash
> ./diskinfo 
+----------+--------+-------------------------+--------------+--------+
|   DISK   | DEVICE |          MODEL          |    SERIAL    |  SIZE  |
+----------+--------+-------------------------+--------------+--------+
| /dev/sda | ata2   | TOSHIBA_THNSNJ512GDNU_A | X61S10DXXXXX | 512 GB |
+----------+--------+-------------------------+--------------+--------+
```
