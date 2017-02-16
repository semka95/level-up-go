package main

import "fmt"

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB
	PB = 1024 * TB
)

// ByteSize aaa
type ByteSize float64

func (b ByteSize) String() string {
	switch {
	case b >= PB:
		return "Very Big"
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%fB", b)
}

func main() {
	fmt.Println(ByteSize(2048))
	fmt.Println(ByteSize(3292528.64))
}
