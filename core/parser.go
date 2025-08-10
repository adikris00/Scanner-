package core

import (
	"regexp"
	"strings"
)

func URLDomain(site string) string {
	if strings.HasPrefix(site, "http://") {
		site = strings.Replace(site, "http://", "", 1)
	} else if strings.HasPrefix(site, "https://") {
		site = strings.Replace(site, "https://", "", 1)
	}
	pattern := regexp.MustCompile(`(.*)/`)
	for {
		matches := pattern.FindAllStringSubmatch(site, -1)
		if len(matches) == 0 {
			break
		}
		site = matches[0][1]
	}
	return site
}

func IndexOf(contents string) bool {
	return strings.Contains(contents, "<title>Index of")
}

func ExtractFolders(name string) bool {
	return !strings.Contains(name, ".")
}

func ExtractFiles(name string) bool {
	if strings.Contains(name, ".") {
		return strings.Contains(name, ".php") ||
			strings.Contains(name, ".phtml") ||
			strings.Contains(name, ".php5") ||
			strings.Contains(name, ".php4") ||
			strings.Contains(name, ".phar") ||
			strings.Contains(name, ".shtml") ||
			strings.Contains(name, ".haxor") ||
			strings.Contains(name, ".py") ||
			strings.Contains(name, ".env") ||
			strings.Contains(name, ".alfa") ||
			strings.Contains(name, ".php7")
	}
	return false
}

func Extract(contents, selected string) []string {
	var pathFiles [][]string
	if strings.Contains(contents, `</td><td><a href="`) {
		if strings.Contains(selected, "Files") || strings.Contains(selected, "Folders") {
			re := regexp.MustCompile(`</td><td><a href="(.*?)">`)
			pathFiles = re.FindAllStringSubmatch(contents, -1)
			return extractMatches(pathFiles)
		}
	} else if strings.Contains(contents, `]"> <a href="`) {
		if strings.Contains(selected, "Files") || strings.Contains(selected, "Folders") {
			re := regexp.MustCompile(`]"> <a href="(.*?)">`)
			pathFiles = re.FindAllStringSubmatch(contents, -1)
			return extractMatches(pathFiles)
		}
	} else if strings.Contains(contents, "width=device-width, initial-scale=1.0") ||
		strings.Contains(contents, `<tr><td data-sort=`) {
		if strings.Contains(selected, "Files") || strings.Contains(selected, "Folders") {
			re := regexp.MustCompile(`"><a href="(.*?)"><img class="`)
			pathFiles = re.FindAllStringSubmatch(contents, -1)
			return extractMatches(pathFiles)
		}
	}
	return []string{}
}

func extractMatches(matches [][]string) []string {
	var results []string
	for _, match := range matches {
		if len(match) > 1 {
			results = append(results, match[1])
		}
	}
	return results
}

func CheckBackdoors(response, sign string) string {
	if response != "" && strings.Contains(response, sign) {
		if !strings.Contains(response, "<?php") &&
			!strings.Contains(response, "#!/usr/bin/perl") &&
			!strings.Contains(response, "#!/usr/bin/python") &&
			!strings.Contains(response, "#!/bin/bash") &&
			!strings.Contains(response, "#!/usr/local/bin/perl") {
			return sign
		}
	}
	return ""
}
