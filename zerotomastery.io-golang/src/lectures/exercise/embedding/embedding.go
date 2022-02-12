//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import "fmt"

type Bytes int
type Celcius float32

type Averager interface {
	average() 
}

type BandwidthUsage struct {
	amount []Bytes
}

func (b *BandwidthUsage) AverageBandwidth() int {
	sum := 0
	for _, byte := range b.amount {
		sum += int(byte)
	}
	return sum / len(b.amount)
}

type CpuTemp struct {
	temp []Celcius
}

func (c *CpuTemp) AverageCpuTemp() int {
	sum := 0
	for _, temp := range c.temp {
		sum += int(temp)
	}
	return sum / len(c.temp)
}

type MemoryUsage struct {
	amount []Bytes
}


func (m *MemoryUsage) AverageMemoryUsage() int {
	sum := 0
	for _, memory := range m.amount {
		sum += int(memory)
	}
	return sum / len(m.amount)
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dashboard := Dashboard{
		BandwidthUsage: bandwidth,
		CpuTemp: temp,
		MemoryUsage: memory,
	}

	fmt.Printf("Average bandwidth is: %v\n", dashboard.AverageBandwidth())
	fmt.Printf("Average cpu temperature is: %v\n", dashboard.AverageCpuTemp())
	fmt.Printf("Average memory usage is: %v\n", dashboard.AverageMemoryUsage())
}
