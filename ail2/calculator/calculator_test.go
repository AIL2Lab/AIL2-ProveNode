package calculator

import (
	"encoding/json"
	"log"
	"math"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/fsnotify/fsnotify"
)

// go test -v -timeout 30s -count=1 -run TestCalcPoint AIL2-ProveNode/ail2/calculator
func TestCalcPoint(t *testing.T) {
	gpuNames := []string{"NVIDIA GeForce RTX 4060 Ti"}
	// gpuNames := []string{"NVIDIA GeForce RTX 3070 Laptop GPU"}
	// gpuMemoryTotals := []int32{8192}
	gpuMemoryTotals := []int32{16384}
	// gpuNames := []string{
	// 	"NVIDIA GeForce RTX 4060 Ti",
	// 	"NVIDIA GeForce RTX 4060 Ti",
	// 	"NVIDIA GeForce RTX 4060 Ti",
	// 	"NVIDIA GeForce RTX 4060 Ti",
	// }
	// gpuMemoryTotals := []int32{8192, 8192, 8192, 8192}
	memoryTotal := 17105440768
	// memoryTotal := 34256556032

	findFirstDigit := func(s string) int {
		for i, char := range s {
			if char >= '0' && char <= '9' {
				return i
			}
		}
		return -1
	}

	for i, name := range gpuNames {
		fd := findFirstDigit(name)
		if fd == -1 {
			log.Fatalf("invalid gpu type: %v", name)
		}
		// gpuNames[i] = strings.ToLower(name[fd:])
		parts := strings.Split(strings.ToLower(name[fd:]), " ")
		filteredParts := []string{}
		for ii, part := range parts {
			if ii == 0 || part == "ti" || part == "super" || part == "d" || part == "oc" || part == "ultra" {
				filteredParts = append(filteredParts, part)
			}
		}
		gpuNames[i] = strings.Join(filteredParts, " ")
	}
	log.Printf("match gpu type: %v", gpuNames)

	for i, mem := range gpuMemoryTotals {
		gpuMemoryTotals[i] = mem / 1024
	}
	log.Printf("format gpu memory unit to GB: %v", gpuMemoryTotals)

	memoryTotal = int(math.Round(float64(memoryTotal) / (1024 * 1024 * 1024)))
	log.Printf("physical memory total: %vGB", memoryTotal)

	calcPoint, err := CalculatePointExact(gpuNames, gpuMemoryTotals, int32(memoryTotal))
	log.Printf("%v %v %v %v", gpuNames, gpuMemoryTotals, calcPoint, err)
}

func TestCalcPoint2(t *testing.T) {
	gpuNames := []string{"NVIDIA GeForce RTX 4060 Ti"}
	// gpuMemoryTotals := []int32{8192}
	gpuMemoryTotals := []int32{16384}
	memoryTotal := int64(17105440768)
	// memoryTotal := int64(34256556032)
	memoryTotal = int64(math.Round(float64(memoryTotal) / (1024 * 1024 * 1024)))

	calcPoint, err := CalculatePointExactFromReport(gpuNames, gpuMemoryTotals, memoryTotal)
	log.Printf("%v %v %v %v", gpuNames, gpuMemoryTotals, calcPoint, err)

	gpuNames = []string{
		"NVIDIA GeForce RTX 4080",
		"NVIDIA GeForce RTX 4060 Ti",
		"NVIDIA GeForce RTX 4060 Ti",
		"NVIDIA GeForce RTX 4060 Ti",
	}
	gpuMemoryTotals = []int32{16384, 8192, 8192, 8192}
	memoryTotal = int64(128) // int64(137438953472)

	calcPoint, err = CalculatePointExactFromReport(gpuNames, gpuMemoryTotals, memoryTotal)
	log.Printf("%v %v %v %v", gpuNames, gpuMemoryTotals, calcPoint, err)
}

func TestCalcPoint3(t *testing.T) {
	gpuNames := []string{"NVIDIA GeForce RTX 4070"}
	// gpuMemoryTotals := []int32{8192}
	gpuMemoryTotals := []int32{12282}
	memoryTotal := int64(32)

	calcPoint, err := CalculatePointExactFromReport(gpuNames, gpuMemoryTotals, memoryTotal)
	log.Printf("%v %v %v %v", gpuNames, gpuMemoryTotals, calcPoint, err)
}

func TestCalcPoint4(t *testing.T) {
	gpuNames := []string{"NVIDIA GeForce RTX 3070 Laptop GPU"}
	// gpuMemoryTotals := []int32{8192}
	gpuMemoryTotals := []int32{16384}
	memoryTotal := int64(32)

	calcPoint, err := CalculatePointExactFromReport(gpuNames, gpuMemoryTotals, memoryTotal)
	log.Printf("%v %v %v %v", gpuNames, gpuMemoryTotals, calcPoint, err)

	calcPoint, err = CalculatePointFuzzyFromReport(gpuNames, memoryTotal)
	log.Printf("%v %v %v %v", gpuNames, gpuMemoryTotals, calcPoint, err)
}

// go test -v -timeout 30s -count=1 -run TestFsNotify AIL2-ProveNode/ail2/calculator
func TestFsNotify(t *testing.T) {
	wg := sync.WaitGroup{}
	gil := []GpuInfo{
		{
			Name:        "2060",
			CudaCore:    1920,
			MemoryTotal: 6,
		},
		{
			Name:        "2060",
			CudaCore:    2176,
			MemoryTotal: 12,
		},
		{
			Name:        "2060 super",
			CudaCore:    2176,
			MemoryTotal: 8,
		},
	}
	jsondata, err := json.Marshal(gil)
	if err != nil {
		t.Fatalf("Marshal json failed: %v", err)
	}
	if err := os.WriteFile("aaa.json", jsondata, 0644); err != nil {
		t.Fatalf("Write json file failed: %v", err)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		t.Fatalf("Create watcher failed: %v", err)
	}
	defer watcher.Close()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				t.Logf("event: %v", event)
				if event.Has(fsnotify.Write) {
					t.Logf("modified file: %v", event.Name)
					if strings.HasSuffix(event.Name, "aaa.json") {
						return
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				t.Logf("watcher error: %v", err)
			}
		}
	}()

	// if err := watcher.Add("/Volumes/data/code/AIL2-ProveNode/ail2/calculator"); err != nil {
	if err := watcher.Add("."); err != nil {
		t.Fatalf("Add watcher failed: %v", err)
	}

	gil = append(gil, GpuInfo{
		Name:        "3070",
		CudaCore:    5888,
		MemoryTotal: 8,
	})
	jsondata2, err := json.Marshal(gil)
	if err != nil {
		t.Fatalf("Marshal json failed: %v", err)
	}
	if err := os.WriteFile("aaa.json", jsondata2, 0644); err != nil {
		t.Fatalf("Write json file failed: %v", err)
	}

	wg.Wait()
	os.Remove("aaa.json")

	// jsondata3, err := json.Marshal(nvidiaGpuInfoList)
	// if err != nil {
	// 	t.Fatalf("Marshal json failed: %v", err)
	// }
	// if err := os.WriteFile("gpus.json", jsondata3, 0644); err != nil {
	// 	t.Fatalf("Write json file failed: %v", err)
	// }
}
