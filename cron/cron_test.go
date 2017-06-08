package cron

import (
	"testing"
	"time"
)

func TestCreateCronJob(t *testing.T) {
	go CreateCronJob(5*time.Second, demo)
	go CreateCronJob(5*time.Second, demo2)
	time.Sleep(30 * time.Second)
}

func demo() string {
	return "demo demo"
}

func demo2() string {
	return "demo2 demo2"
}
