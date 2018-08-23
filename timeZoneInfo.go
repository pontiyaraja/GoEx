package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
	"time"
)

var zoneDirs = []string{
	// Update path according to your OS
	"/usr/share/zoneinfo/",
	"/usr/share/lib/zoneinfo/",
	"/usr/lib/locale/TZ/",
}

var zoneDir string

func main() {
	// for _, zoneDir = range zoneDirs {
	//     ReadFile("")
	// }
	date := time.Now()
	fmt.Println("Time ------> ", date)
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	fmt.Println("Time ------> ", startOfDay)
	fmt.Println(runtime.NumGoroutine())
	// for i := 1; i < 10; i++ {
	// 	go printData(i)
	// }
	// time.Sleep(5 * time.Second)
	// reader := bufio.NewReader(os.Stdin)
	// reader.ReadString('\n')
}
func printData(i int) {
	if i%2 == 0 {
		fmt.Println("Exiting go routine ", i, runtime.NumGoroutine())
		runtime.Goexit()
	} else {
		fmt.Println("In goroutien ", i)
	}
	time.Sleep(3 * time.Second)
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println(t)
			ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(50*time.Millisecond))
			print(ctx, cancel)
		}
	}()
	time.Sleep(2 * time.Second)
}

func print(ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept====================")
	case <-ctx.Done():
		fmt.Println("error =====================  ", ctx.Err())
	}
	time.Sleep(2 * time.Second)
	fmt.Println("    Go  === ", runtime.NumGoroutine(), "  CPU USAGE  =====  ", runtime.NumCPU())
}

func ReadFile(path string) {
	files, _ := ioutil.ReadDir(zoneDir + path)
	for _, f := range files {
		if f.Name() != strings.ToUpper(f.Name()[:1])+f.Name()[1:] {
			continue
		}
		if f.IsDir() {
			ReadFile(path + "/" + f.Name())
		} else {
			fmt.Println((path + "/" + f.Name())[1:])
		}
	}
}
