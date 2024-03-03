package main

import (
	"time"
)

type DatabaseClient struct {
}

var DBClient = &DatabaseClient{}

func (c *DatabaseClient) GetValue1(value int32) int32 {
	time.Sleep(10 * time.Millisecond)
	return 1
}

func (c *DatabaseClient) GetSlowValue2(value int32) int32 {
	time.Sleep(100 * time.Millisecond)
	return 1
}

func (c *DatabaseClient) Sum(value1, value2 int32) int32 {
	time.Sleep(10 * time.Millisecond)
	return value1 + value2
}
