package main

import (
	"math"
	"slices"
	"strconv"
	"strings"

	"ttp.sh/advent-of-code/2023/util"
)

type SparseMap struct {
	start  []int
	target []int
	count  []int
}

func NewSparseMap() *SparseMap {
	return &SparseMap{
		start:  []int{},
		target: []int{},
		count:  []int{},
	}
}

func (m *SparseMap) Set(start_range, start_value, count int) {
	i, _ := slices.BinarySearch(m.start, start_value)

	m.start = slices.Insert(m.start, i, start_value)
	m.target = slices.Insert(m.target, i, start_range)
	m.count = slices.Insert(m.count, i, count)
}

func (m *SparseMap) Get(index int) int {
	i, found := slices.BinarySearch(m.start, index)

	if !found {
		i--
	}
	// we are before all starts, so return the index
	if i == -1 {
		return index
	}

	start := m.start[i]
	count := m.count[i]

	if start+count < index {
		return index
	}

	target := m.target[i]
	offset := index - start
	return target + offset
}

func ParseLine(line string) (int, int, int) {
	ints := util.Map(strings.Split(line, " "), func(i int, v string) int {
		str, _ := strconv.Atoi(v)
		return str
	})
	return ints[0], ints[1], ints[2]
}

func Parse(lines <-chan string) ([]int, *SparseMap, *SparseMap, *SparseMap, *SparseMap, *SparseMap, *SparseMap, *SparseMap) {
	var seeds []int

	seedToSoil := NewSparseMap()
	soilToFertilizer := NewSparseMap()
	fertilizerToWater := NewSparseMap()
	waterToLight := NewSparseMap()
	lightToTemperature := NewSparseMap()
	temperatureToHumidity := NewSparseMap()
	humidityToLocation := NewSparseMap()

	state := "seeds"

	for line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		if strings.HasPrefix(line, "seeds: ") {
			seeds = util.Map(strings.Split(strings.TrimPrefix(line, "seeds: "), " "), func(i int, v string) int {
				str, _ := strconv.Atoi(v)
				return str
			})
			continue
		}

		if strings.HasSuffix(line, " map:") {
			state = strings.TrimSuffix(line, " map:")
			continue
		}

		start_range, start_value, count := ParseLine(line)
		switch state {
		case "seed-to-soil":
			seedToSoil.Set(start_range, start_value, count)
		case "soil-to-fertilizer":
			soilToFertilizer.Set(start_range, start_value, count)
		case "fertilizer-to-water":
			fertilizerToWater.Set(start_range, start_value, count)
		case "water-to-light":
			waterToLight.Set(start_range, start_value, count)
		case "light-to-temperature":
			lightToTemperature.Set(start_range, start_value, count)
		case "temperature-to-humidity":
			temperatureToHumidity.Set(start_range, start_value, count)
		case "humidity-to-location":
			humidityToLocation.Set(start_range, start_value, count)
		}
	}

	return seeds, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation
}

func Run(lines <-chan string) int {
	seeds, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation := Parse(lines)

	lowest := math.MaxInt
	for _, seed := range seeds {
		soil := seedToSoil.Get(seed)
		fert := soilToFertilizer.Get(soil)
		wate := fertilizerToWater.Get(fert)
		ligh := waterToLight.Get(wate)
		temp := lightToTemperature.Get(ligh)
		humi := temperatureToHumidity.Get(temp)
		loca := humidityToLocation.Get(humi)

		// log.Printf("seed: %d, soil %d, fert %d, wate %d, ligh %d, temp %d, humi %d, loca %d", seed, soil, fert, wate, ligh, temp, humi, loca)

		new_lowset := loca
		if new_lowset < lowest {
			lowest = new_lowset
		}
	}

	return lowest
}

func main() {
	println(Run(util.Read("./input.txt")))
}
