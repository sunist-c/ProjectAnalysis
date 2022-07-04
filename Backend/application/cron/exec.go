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
	if err := task.AddFunc(spec, ExecSpider); err != nil {
		log.Fatalln(fmt.Sprintf("start cron task failed with: %v", err))
	}
	task.Start()
}

func ExecSpider() {
	argDate := time.Now().AddDate(0, 0, -1).Format("2006-01-02MST")
	cmd := exec.Command("python3", SpiderPath, argDate, DataTempPath, DataResultPath)
	log.Println(cmd.String())
	bytes, err := cmd.Output()
	log.Println(string(bytes), err)
	err = cmd.Run()
	if err != nil {
		log.Println(fmt.Sprintf("cron task start spider error with: %v", err))
	}

	err = LoadCsv()
	if err != nil {
		log.Println(fmt.Sprintf("cron task load csv-file error with: %v", err))
	}
}

func LoadCsv() (err error) {
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
