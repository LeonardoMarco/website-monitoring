package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const delay = 60

func main() {
	fmt.Println(("Init monitoring.."))

	for {
		initMonitoring()
	}
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Error:", err)
	}

	if resp.StatusCode != 200 {
		fmt.Println("Site", site, "downed")
		writerLog(site, false)
		// ... send alert to alerts platform
	} else {
		writerLog(site, true)
	}
}

func readerFileSites() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Deu erro")
	}

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		line := strings.TrimSpace(str)
		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func initMonitoring() {
	sites := readerFileSites()

	for _, site := range sites {
		fmt.Println("Testando site:", site)
		testSite(site)
	}

	time.Sleep(delay * time.Second)
}

func writerLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- Online: " + strconv.FormatBool(status) + "\n")
}
