package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/mainlycricket/CricKendra/backend/internal/dbutils"
)

type info_parse_response struct {
	infoInput *matchInfoInput
	err       error
}

type match_init_response struct {
	matchInfo *matchInfo
	err       error
}

type channelWrapper[T any] struct {
	channel             chan T
	tasks               sync.WaitGroup
	dependencyCondition sync.Cond
	isDependencyFinish  bool
}

func newChannelWrapper[T any](capacity int) *channelWrapper[T] {
	var wrapperChannel channelWrapper[T]

	wrapperChannel.channel = make(chan T, capacity)
	wrapperChannel.dependencyCondition = *sync.NewCond(&sync.Mutex{})

	return &wrapperChannel
}

func (ch *channelWrapper[T]) close(mainWg *sync.WaitGroup) {
	defer mainWg.Done()

	ch.dependencyCondition.L.Lock()
	defer ch.dependencyCondition.L.Unlock()
	for !ch.isDependencyFinish {
		ch.dependencyCondition.Wait()
	}

	ch.tasks.Wait()
	close(ch.channel)
}

func triggerParseInfo(directories map[string]string, parseChannel *channelWrapper[info_parse_response]) {
	parseChannel.dependencyCondition.L.Lock()

	defer func() {
		fmt.Println("all parsing jobs triggered")
		parseChannel.isDependencyFinish = true
		parseChannel.dependencyCondition.L.Unlock()
		parseChannel.dependencyCondition.Broadcast()
	}()

	for basePath, playingFormat := range directories {
		dirEntries, err := os.ReadDir(basePath)
		if err != nil {
			log.Fatalf("error while reading directory: %v", basePath)
		}

		for _, dirEntry := range dirEntries {
			fileName := dirEntry.Name()
			if strings.HasSuffix(fileName, "_info.csv") {
				matchCricsheetId := strings.TrimSuffix(fileName, "_info.csv")

				match, err := dbutils.ReadMatchByCricsheetId(context.Background(), DB_POOL, matchCricsheetId)
				if err != nil && err.Error() != "no rows in result set" {
					log.Printf(`error while reading cricsheet match: %v`, err)
					continue
				}
				if match.IsBBBDone.Bool {
					continue
				}

				matchInfoPath := filepath.Join(basePath, fileName)

				parseChannel.tasks.Add(1)
				go parseMatchInfoFile(matchInfoPath, playingFormat, parseChannel.channel)
			}
		}
	}
}

func triggerMatchInit(parseChannel *channelWrapper[info_parse_response], matchInitChannel *channelWrapper[match_init_response]) {
	matchInitChannel.dependencyCondition.L.Lock()

	defer func() {
		fmt.Println("all match init jobs triggered")
		matchInitChannel.isDependencyFinish = true
		matchInitChannel.dependencyCondition.L.Unlock()
		matchInitChannel.dependencyCondition.Broadcast()
	}()

	for response := range parseChannel.channel {
		if response.err != nil {
			log.Printf("error while parsing match info file: %v", response.err)
		} else {
			matchInitChannel.tasks.Add(1)
			go response.infoInput.initalizeMatch(matchInitChannel.channel)
		}

		parseChannel.tasks.Done()
	}
}

func triggerMatchBbb(matchInitChannel *channelWrapper[match_init_response], bbbChannel *channelWrapper[error]) {
	bbbChannel.dependencyCondition.L.Lock()

	defer func() {
		fmt.Println("all match bbb jobs triggered")
		bbbChannel.isDependencyFinish = true
		bbbChannel.dependencyCondition.L.Unlock()
		bbbChannel.dependencyCondition.Broadcast()
	}()

	for response := range matchInitChannel.channel {
		if response.err != nil {
			log.Printf("error while initializing match: %v", response.err)
		} else {
			basePath := strings.TrimSuffix(response.matchInfo.infoFilePath, "_info.csv")
			match_bbb_path := basePath + ".csv"

			bbbChannel.tasks.Add(1)
			go insertBBB(match_bbb_path, response.matchInfo, bbbChannel.channel)
		}

		matchInitChannel.tasks.Done()
	}
}

func receiveBbb(bbbChannel *channelWrapper[error]) {
	for err := range bbbChannel.channel {
		if err != nil {
			log.Printf("error while inserting BBB data: %v", err)
		}
		bbbChannel.tasks.Done()
	}

	fmt.Println("bbb job finished")
}
