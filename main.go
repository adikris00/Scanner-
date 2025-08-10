package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"mawXscanner/core"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if _, err := os.Stat(core.Maw); os.IsNotExist(err) {
		if err := os.Mkdir(core.Maw, 0755); err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}

	core.Banners()

	fmt.Print("\n" + core.FR + "[+] " + core.FC + "IP/DOMAIN LIST: " + core.FW)
	var filename string
	fmt.Scanln(&filename)

	target, err := core.ReadLines(filename)
	if err != nil {
		fmt.Println("[!] File Ga Ada Bang!: " + filename)
		os.Exit(1)
	}

	core.Signs, _ = core.ReadLines("lib-maw/SHELL-STRINGS.txt")
	core.Strings_Shells, _ = core.ReadLines("lib-maw/SHELL-STRINGS.txt")
	core.Locations, _ = core.ReadLines("lib-maw/COMBINE-PATH.txt")
	core.TrustedFiles, _ = core.ReadLines("lib-maw/TRUSTED-FILES.txt")

	uaLines, _ := core.ReadLines("lib-maw/User-Agents.txt")
	for _, line := range uaLines {
		core.UserAgents = append(core.UserAgents, strings.TrimSpace(line))
	}

	var ua string
	if len(core.UserAgents) > 0 {
		ua = core.UserAgents[rand.Intn(len(core.UserAgents))]
	}
	core.Headers = map[string]string{
		"User-Agent":      ua,
		"Content-type":    "*/*",
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
		"Accept-Language": "en-US,en;q=0.5",
		"Connection":      "keep-alive",
	}

	var wg sync.WaitGroup
	results := make(chan string)
	for _, site := range target {
		wg.Add(1)
		go core.CmsCheckers(site, results, &wg)
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	for result := range results {
		fmt.Println(result)
	}
}
