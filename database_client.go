package main

import (
	"time"
)

type DatabaseClient struct {
}

var DBClient = &DatabaseClient{}

func (c *DatabaseClient) FindByValue1(value int32) int32 {
	time.Sleep(10 * time.Millisecond)
	return value
}

func (c *DatabaseClient) FindByValue2(value int32) int32 {
	time.Sleep(100 * time.Millisecond)
	return value
}

func (c *DatabaseClient) FindByValues(value1, value2 int32) int32 {
	time.Sleep(10 * time.Millisecond)
	return value1 + value2
}
