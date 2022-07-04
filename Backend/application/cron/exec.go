package cron

import (
	"ProjectAnalysis/domain/spider"
	"encoding/csv"
	"fmt"
	"github.com/robfig/cron"
	"log"
	"os"
	"os/exec"
	"time"
)

var (
	ReadChan chan spider.RawData = make(chan spider.RawData, 1000)
)

func init() {
	go startTask()
}

func startTask() {
	task := cron.New()
	spec := "0 0 12 * *"
	if err := task.AddFunc(spec, execSpider); err != nil {
		log.Fatalln(fmt.Sprintf("start cron task failed with: %v", err))
	}
	task.Start()
}

func execSpider() {
	argDate := fmt.Sprintf("%v-%v-%v", time.Now().Month(), time.Now().Day(), time.Now().Year())
	cmd := exec.Command("python", SpiderPath, argDate, DataTempPath, DataResultPath)
	err := cmd.Run()
	if err != nil {
		log.Println(fmt.Sprintf("cron task start spider error with: %v", err))
	}

	err = loadCsv()
	if err != nil {
		log.Println(fmt.Sprintf("cron task load csv-file error with: %v", err))
	}
}

func loadCsv() (err error) {
	file, err := os.Open(DataResultPath)
	if err != nil {
		return err
	}

	reader := csv.NewReader(file)
	preData, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, line := range preData {
		raw := spider.RawData{}
		raw.Format(line)
		ReadChan <- raw
	}
	return nil
}
