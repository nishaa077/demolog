package functionn

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/lf-edge/eve/api/go/info"
	"github.com/tidwall/gjson"
	"google.golang.org/protobuf/proto"
	"main.go/model"
)

func Readappdata(uuid string, dir string, logpath string) ([]model.NodeData, error) {
	todayFiles, path := getPath(uuid, dir, logpath)
	//todayFiles, path := functions.getPath("8433f2c4-6446-4994-b66d-86e67c81e056", "info")

	var finalFileData []model.NodeData
	// var storageList []utility.NodeStorageList
	for _, file := range todayFiles {
		filePath := filepath.Join(path, file.Name())
		jsonString := protobufConversion(file.Name(), filePath, uuid, todayFiles)
		// Extracting data from the json
		zType := gjson.Get(jsonString, "ztype").Int()

		//zType := gjson.Get(jsonString, "ztype").Int()
		if zType == 3 {
			//	fmt.Println("app jsondata: ", jsonString)
			appname := gjson.Get(jsonString, "InfoContent.Ainfo.AppName").String()
			version := gjson.Get(jsonString, "InfoContent.Ainfo.appVersion").String()
			appid := gjson.Get(jsonString, "InfoContent.Ainfo.AppID").String()
			devid := gjson.Get(jsonString, "devId").String()
			state := gjson.Get(jsonString, "InfoContent.Ainfo.state").Int()
			volumeRefs := gjson.Get(jsonString, "InfoContent.Ainfo.volumeRefs").Array()[0].String()

			if state != 0 {
				nodeData := model.NodeData{
					AppName:    appname,
					Version:    version,
					AppId:      appid,
					Devid:      devid,
					State:      state,
					VolumeRefs: volumeRefs,
				}
				finalFileData = append(finalFileData, nodeData)
				break
			}
		}
	}

	return finalFileData, nil
}

func getPath(uuid string, dir string, commondatabase string) ([]os.DirEntry, string) {
	deviceDirectory := path.Join(commondatabase, "device") // common.DatabasePath ="log", common.DeviceDir ="device"
	//	currentTime := "01-02-2006 15:04:05 Monday"
	startOfDay := time.Date(2023, time.October, 9, 0, 0, 0, 0, time.UTC) // here need to change for the current time and date
	endOfDay := startOfDay.Add(24 * time.Hour)

	var path string = ""
	switch dir {
	case "info", "metrics": //common.InfoDir ="info", common.MetricsDir="metrics"
		path = deviceDirectory + "/" + uuid + "/" + dir
	}
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("files : ", files) data is getting fetched
	todayFiles := filterTodayFiles(files, startOfDay, endOfDay)
	return todayFiles, path
}

// Filteration for the particular date
func filterTodayFiles(files []os.DirEntry, startOfDay, endOfDay time.Time) []os.DirEntry {
	var todayFiles []os.DirEntry
	for _, file := range files {
		fileName := file.Name()
		fileName = strings.Replace(fileName, "_", ":", -1)
		fileTimestamp, err := time.Parse("2006-01-02T15:04:05", fileName[:len("2006-01-02T15:04:05")])
		//fmt.Println("filestamp: ", fileTimestamp)
		if err != nil {
			log.Fatal("err1: ", err)
			continue // Skip files with invalid timestamps.
		}
		if fileTimestamp.After(startOfDay) && fileTimestamp.Before(endOfDay) {
			todayFiles = append(todayFiles, file)
		}
	}
	return todayFiles
}
func protobufConversion(fileName string, filePath string, deviceName string, todayFiles []os.DirEntry) string {
	var finalJsonData []byte
	for _, file := range todayFiles {
		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error1 reading protobuf file %s for device %s: %v\n", file.Name(), deviceName, err)
			continue // Continue to the next file.
		}
		protoMessage := &info.ZInfoMsg{}
		if err := proto.Unmarshal(data, protoMessage); err != nil {
			fmt.Printf("Error2 reading protobuf file %s for device %s: %v\n", file.Name(), deviceName, err)
			continue // Continue to the next file.
		}
		// Marshal the protobuf message to JSON.
		jsonData, err := json.MarshalIndent(protoMessage, "", "  ")
		if err != nil {
			fmt.Printf("Error3 reading protobuf file %s for device %s: %v\n", file.Name(), deviceName, err)
			continue // Continue to the next file.
		}
		finalJsonData = append(finalJsonData, jsonData...)
	}

	return string(finalJsonData)
}
