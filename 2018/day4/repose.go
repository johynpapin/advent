package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type LogEntry struct {
	Id     int
	Asleep bool
	Time   time.Time
	Minute int
}

type GuardEntry struct {
	Asleep  bool
	Since   int
	Total   int
	Minutes map[int]int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var entries []*LogEntry

	for scanner.Scan() {
		rawTime := scanner.Text()[1:17]
		rawText := scanner.Text()[19:]

		time, err := time.Parse("2006-01-02 15:04", rawTime)
		if err != nil {
			fmt.Println(err)
		}

		entry := &LogEntry{
			Time:   time,
			Minute: time.Minute(),
		}

		entry.Id = -1

		if strings.Contains(rawText, "up") {
			entry.Asleep = false
		} else if strings.Contains(rawText, "asleep") {
			entry.Asleep = true
		} else {
			fmt.Sscanf(rawText, "Guard #%d begins shift", &entry.Id)
		}

		entries = append(entries, entry)
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Time.Before(entries[j].Time)
	})

	guards := make(map[int]*GuardEntry)

	var currentGuardId int

	for _, entry := range entries {
		if entry.Id != -1 {
			currentGuardId = entry.Id
			continue
		}

		guard, exist := guards[currentGuardId]
		if !exist {
			guard = &GuardEntry{
				Minutes: make(map[int]int),
			}
			guards[currentGuardId] = guard
		}

		if entry.Asleep {
			guard.Since = entry.Minute
			continue
		}

		guard.Total += entry.Minute - guard.Since

		for i := guard.Since; i < entry.Minute; i++ {
			guard.Minutes[i]++
		}
	}

	var bestGuardId, bestGuardTime int
	var freqGuardId, freqGuardMinute, freqGuardValue int
	for id, guard := range guards {
		if guard.Total > bestGuardTime {
			bestGuardId = id
			bestGuardTime = guard.Total
		}

		for minute, total := range guard.Minutes {
			if total > freqGuardValue {
				freqGuardId = id
				freqGuardMinute = minute
				freqGuardValue = total
			}
		}
	}

	var bestMinute, bestTotal int
	for minute, total := range guards[bestGuardId].Minutes {
		if total > bestTotal {
			bestMinute = minute
			bestTotal = total
		}
	}

	fmt.Println("part1", bestGuardId*bestMinute)
	fmt.Println("part2", freqGuardId*freqGuardMinute)
}
