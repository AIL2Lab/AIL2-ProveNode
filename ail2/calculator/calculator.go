package calculator

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"strings"
)

/*
 * https://orion.deeplink.cloud/shortterm
 * https://galaxyrace.deepbrainchain.org/rule
 * https://www.nvidia.cn/geforce/graphics-cards/compare/
 * https://www.nvidia.com/en-us/geforce/graphics-cards/compare/
 */

type GpuInfo struct {
	Name        string
	CudaCore    int32
	MemoryTotal int32
}

var nvidiaGpuInfoList = []GpuInfo{}

func LoadGpuList(filename string) error {
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("read nvidia gpu support list failed: %v", err)
	}

	err = json.Unmarshal(jsonData, &nvidiaGpuInfoList)
	if err != nil {
		return fmt.Errorf("parse nvidia gpu support list failed: %v", err)
	}

	if len(nvidiaGpuInfoList) == 0 {
		return errors.New("empty nvidia gpu support list")
	}
	return nil
}

// gpu model exact match with gpu_memory
func CalculatePointExact(gpuNames []string, gpuMemoryTotals []int32, memoryTotal int32) (float64, error) {
	if len(gpuNames) != len(gpuMemoryTotals) {
		return 0.0, errors.New("gpu name and memory list are inconsistent")
	}
	gpuCount := len(gpuNames)
	// cudaCores := make([]int32, 0, len(gpuNames))
	var (
		minCudaCore int32 = 0
		minIndex    int   = 0
	)
	for i, name := range gpuNames {
		// cudaCores[i] = 0
		// gpuInfo, ok := nvidiaGpuInfoList[name]
		// if !ok {
		// 	return 0.0
		// }
		for _, gpuInfo := range nvidiaGpuInfoList {
			if gpuInfo.Name == name && gpuInfo.MemoryTotal == gpuMemoryTotals[i] {
				// if strings.EqualFold(gpuInfo.Name, name) && gpuInfo.MemoryTotal == gpuMemoryTotals[i] {
				// cudaCores[i] = gpuInfo.CudaCore
				if minCudaCore == 0 || gpuInfo.CudaCore < minCudaCore {
					minCudaCore = gpuInfo.CudaCore
					minIndex = i
				}
				break
			}
		}
	}
	if minCudaCore == 0 {
		return 0.0, errors.New("can not find the number of cuda cores")
	}
	// fmt.Println(minCudaCore)
	// fmt.Println(minIndex)

	point := float64(memoryTotal) / 3.5
	point += float64(gpuCount) * 25
	point += math.Sqrt(float64(minCudaCore)) * math.Sqrt(float64(gpuMemoryTotals[minIndex])/10) * float64(gpuCount)
	// return point, nil
	ratio := math.Pow(10, 2)
	return math.Round(point*ratio) / ratio, nil
}

// gpu model fuzzy query without gpu_memory
func CalculatePointFuzzy(gpuNames []string, memoryTotal int32) (float64, error) {
	gpuCount := len(gpuNames)
	// cudaCores := make([]int32, 0, len(gpuNames))
	var (
		minCudaCore  int32 = 0
		minGpuMemory int32 = 0
	)
	for _, name := range gpuNames {
		// cudaCores[i] = 0
		// gpuInfo, ok := nvidiaGpuInfoList[name]
		// if !ok {
		// 	return 0.0
		// }
		for _, gpuInfo := range nvidiaGpuInfoList {
			if gpuInfo.Name == name {
				// if strings.EqualFold(gpuInfo.Name, name) {
				// cudaCores[i] = gpuInfo.CudaCore
				if minCudaCore == 0 || gpuInfo.CudaCore < minCudaCore {
					minCudaCore = gpuInfo.CudaCore
					minGpuMemory = gpuInfo.MemoryTotal
				}
				break
			}
		}
	}
	if minCudaCore == 0 {
		return 0.0, errors.New("can not find the number of cuda cores")
	}
	// fmt.Println(minCudaCore)
	// fmt.Println(minGpuMemory)

	point := float64(memoryTotal) / 3.5
	point += float64(gpuCount) * 25
	point += math.Sqrt(float64(minCudaCore)) * math.Sqrt(float64(minGpuMemory)/10) * float64(gpuCount)
	// return point, nil
	ratio := math.Pow(10, 2)
	return math.Round(point*ratio) / ratio, nil
}

func CalculatePointExactFromReport(gpuNames []string, gpuMemoryTotals []int32, memoryTotal int64) (float64, error) {
	findFirstDigit := func(s string) int {
		for i, char := range s {
			if char >= '0' && char <= '9' {
				return i
			}
		}
		return -1
	}

	// "NVIDIA GeForce RTX 4060 Ti" => "4060 ti"
	matchedGpuNames := make([]string, 0)
	for _, name := range gpuNames {
		fd := findFirstDigit(name)
		if fd == -1 {
			continue
		}
		// matchedGpuNames = append(matchedGpuNames, strings.ToLower(name[fd:]))
		parts := strings.Split(strings.ToLower(name[fd:]), " ")
		filteredParts := []string{}
		for ii, part := range parts {
			if ii == 0 || part == "ti" || part == "super" || part == "d" || part == "oc" || part == "ultra" {
				filteredParts = append(filteredParts, part)
			}
		}
		matchedGpuNames = append(matchedGpuNames, strings.Join(filteredParts, " "))
	}

	// 8192 MB => 8 GB
	for i, mem := range gpuMemoryTotals {
		gpuMemoryTotals[i] = int32(math.Round(float64(mem) / 1024))
	}

	// 17105440768 Bytes => 16 GB
	// memoryTotal = int64(math.Round(float64(memoryTotal) / (1024 * 1024 * 1024)))
	return CalculatePointExact(matchedGpuNames, gpuMemoryTotals, int32(memoryTotal))
}

func CalculatePointFuzzyFromReport(gpuNames []string, memoryTotal int64) (float64, error) {
	findFirstDigit := func(s string) int {
		for i, char := range s {
			if char >= '0' && char <= '9' {
				return i
			}
		}
		return -1
	}

	// "NVIDIA GeForce RTX 4060 Ti" => "4060 ti"
	matchedGpuNames := make([]string, 0)
	for _, name := range gpuNames {
		fd := findFirstDigit(name)
		if fd == -1 {
			continue
		}
		// matchedGpuNames = append(matchedGpuNames, strings.ToLower(name[fd:]))
		parts := strings.Split(strings.ToLower(name[fd:]), " ")
		filteredParts := []string{}
		for ii, part := range parts {
			if ii == 0 || part == "ti" || part == "super" || part == "d" || part == "oc" || part == "ultra" {
				filteredParts = append(filteredParts, part)
			}
		}
		matchedGpuNames = append(matchedGpuNames, strings.Join(filteredParts, " "))
	}

	// 17105440768 Bytes => 16 GB
	// memoryTotal = int64(math.Round(float64(memoryTotal) / (1024 * 1024 * 1024)))
	return CalculatePointFuzzy(matchedGpuNames, int32(memoryTotal))
}
