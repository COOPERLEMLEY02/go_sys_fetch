package main

import (
	"fmt"

	. "github.com/klauspost/cpuid/v2"
)

func main() {

	CPUBrand, PhysicalCores, ThreadsPerCore := cpuInfo()

	fmt.Println("\033[31mName:\033[0m", CPUBrand)
	fmt.Println("\033[31mPhysical Cores:\033[0m", PhysicalCores)
	fmt.Println("\033[31mThreads Per Core:\033[0m", ThreadsPerCore)

}

func cpuInfo() (string, int, int) {

	CPUBrand := CPU.BrandName
	PhysicalCores := CPU.PhysicalCores
	ThreadsPerCore := CPU.ThreadsPerCore

	return CPUBrand, PhysicalCores, ThreadsPerCore

}
