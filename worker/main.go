package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gocelery/gocelery"
)

func main() {
	// intialize gocelery client
	celeryBroker := gocelery.NewRedisCeleryBroker("redis://localhost:6379")
	celeryBackend := gocelery.NewRedisCeleryBackend("redis://localhost:6379")
	celeryClient, _ := gocelery.NewCeleryClient(celeryBroker, celeryBackend, 2)

	// 1) integer addition
	celeryClient.Register("worker.add_int", addInt)

	// 2) integer addition with kwargs
	celeryClient.Register("worker.add_int_kwargs", &addIntTask{})

	// 3) string addition
	celeryClient.Register("worker.add_str", addStr)

	// 4) integer and string concatenation
	celeryClient.Register("worker.add_str_int", addStrInt)

	// 5) float addition
	celeryClient.Register("worker.add_float", addFloat)

	// 6) maximum array length
	celeryClient.Register("worker.max_arr_len", maxArrLength)

	// 6) array addition
	celeryClient.Register("worker.add_arr", addArr)

	// 7) dictionary addition
	celeryClient.Register("worker.add_dict", addDict)

	// 8) boolean and operation
	celeryClient.Register("worker.and_bool", andBool)

	// start workers
	go celeryClient.StartWorker()

	// gracefully terminate after 1 minute
	time.Sleep(60 * time.Second)
	celeryClient.StopWorker()
}

// 1) integer addition
func addInt(x int, y int) int {
	return x + y
}

// 2) integer addition with kwargs
type addIntTask struct {
	x int
	y int
}

func (a *addIntTask) ParseKwargs(kwargs map[string]interface{}) error {
	kwargA, ok := kwargs["x"]
	if !ok {
		return fmt.Errorf("undefined kwarg x")
	}
	kwargAFloat, ok := kwargA.(float64)
	if !ok {
		return fmt.Errorf("malformed kwarg x")
	}
	a.x = int(kwargAFloat)
	kwargB, ok := kwargs["y"]
	if !ok {
		return fmt.Errorf("undefined kwarg y")
	}
	kwargBFloat, ok := kwargB.(float64)
	if !ok {
		return fmt.Errorf("malformed kwarg y")
	}
	a.y = int(kwargBFloat)
	return nil
}

func (a *addIntTask) RunTask() (interface{}, error) {
	result := a.x + a.y
	return result, nil
}

// 3) string addition
func addStr(a, b string) string {
	return a + b
}

// 4) integer and string concatenation
func addStrInt(a string, b int) string {
	return a + strconv.Itoa(b)
}

// 5) float addition
func addFloat(a float64, b float64) float64 {
	return a + b
}

// 6) maximum array length
func maxArrLength(a, b []string) int {
	if len(a) > len(b) {
		return len(a)
	}
	return len(b)
}

// 6) array addition
func addArr(a, b []interface{}) []interface{} {
	return append(a, b...)
}

// 7) dictionary addition
func addDict(a, b map[string]interface{}) map[string]interface{} {
	c := map[string]interface{}{}
	for k, v := range a {
		c[k] = v
	}
	for k, v := range b {
		c[k] = v
	}
	return c
}

// 8) boolean and operation
func andBool(a, b bool) bool {
	return a && b
}
