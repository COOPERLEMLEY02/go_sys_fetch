package main

import (
	"fmt"
	"math"
	"net"
	"os"
	"runtime"

	//"strings"

	. "github.com/klauspost/cpuid/v2"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func main() {

	logo := []string{
		"  _____________________",
		"  /                    \\",
		"  |  .-----------.  |   |-----.",
		"  |  |           |  |   |-=---|",
		"  |  | Apple //c |  |   |-----|",
		"  |  |           |  |   |-----|",
		"  |  |           |  |   |-----|",
		"  |  `-----------'  |   |-----'/\\",
		"   \\________________/___     /  \\",
		"      /                      / / /",
		"     / //               //  / / /",
		"    /                      / / /",
		"   / _/_/_/_/_/_/_/_/_/_/ /   /",
		"  / _/_/_/_/_/_/_/_/_/_/ /   /",
		" / _/_/_/_______/_/_/_/ / __/",
		"/______________________/ /    ",
		"\\______________________/\\/",
	}

	user, hostname, operatingSystem, Architecture, v := userInformation()
	CPUBrand, PhysicalCores, ThreadsPerCore := cpuInfo()
	diskUsage := diskInformation()
	diskUsageTotal := round2Places(bytestoGB(int64(diskUsage.Total)))
	diskUsageUsed := round2Places(bytestoGB(int64(diskUsage.Used)))
	networkInfo := networkInformation()

	info := []string{

		fmt.Sprintf("\033[32mUser:\033[0m %s", user),
		fmt.Sprintf("\033[32mHostname:\033[0m %s", hostname),
		fmt.Sprintf("\033[32mOS:\033[0m %s", operatingSystem),
		fmt.Sprintf("\033[32mArchitecture:\033[0m %s", Architecture),
		fmt.Sprintf("\033[31mName:\033[0m %s", CPUBrand),
		fmt.Sprintf("\033[31mPhysical Cores:\033[0m %v", PhysicalCores),
		fmt.Sprintf("\033[31mThreads Per Core:\033[0m %v", ThreadsPerCore),
		fmt.Sprintf("\033[31mTotal:\033[0m %v, \033[31mFree:\033[0m %v, \033[31mUsedPercent:\033[0m %f%%", v.Total, v.Free, v.UsedPercent),
		fmt.Sprintf("\033[33mDisk Total:\033[0m %v", diskUsageTotal),
		fmt.Sprintf("\033[33mDisk Used:\033[0m %v", diskUsageUsed),
		fmt.Sprintf("\033[34mNetwork Address:\033[0m %s", networkInfo.IP.String()),
	}

	maxLines := max(len(logo), len(info))
	for i := 0; i < maxLines; i++ {
		var artLine, infoLine string
		if i < len(logo) {
			artLine = logo[i]
		}
		if i < len(info) {
			infoLine = info[i]
		}
		fmt.Printf("%-35s  %s\n", artLine, infoLine)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func cpuInfo() (string, int, int) {
	// Function to get the CPU information

	CPUBrand := CPU.BrandName
	PhysicalCores := CPU.PhysicalCores
	ThreadsPerCore := CPU.ThreadsPerCore

	return CPUBrand, PhysicalCores, ThreadsPerCore

}

func userInformation() (string, string, string, string, *mem.VirtualMemoryStat) {
	// Using the go modules OS and Runtime to get desired information

	user := os.Getenv("USER")
	if user == "" {
		user = os.Getenv("USERNAME")
	}

	v, _ := mem.VirtualMemory()
	hostname, _ := os.Hostname()

	operatingSystem := runtime.GOOS
	Architecture := runtime.GOARCH

	return user, hostname, operatingSystem, Architecture, v

}

func diskInformation() *disk.UsageStat {

	diskSpace, _ := disk.Usage("/")

	return diskSpace

}

func networkInformation() *net.UDPAddr {

	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr

}

func bytestoGB(bytes int64) float64 {

	const gb = 1024 * 1024 * 1024

	return float64(bytes) / float64(gb)
}

func round2Places(x float64) float64 {

	rounded := math.Round(x*100) / 100

	return rounded
}
