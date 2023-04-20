package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type OrderMap struct {
	m    map[string][]string
	keys []string
}

func newOrderMap() *OrderMap {
	return &OrderMap{
		m:    make(map[string][]string),
		keys: make([]string, 0),
	}
}

func (om *OrderMap) set(key string, value string) {
	if _, ok := om.m[key]; !ok {
		om.m[key] = make([]string, 0)
		om.keys = append(om.keys, key)
	}
	om.m[key] = append(om.m[key], value)
}

func (om *OrderMap) get(key string) []string {
	return om.m[key]
}

func main() {
	var isVersion bool
	var isHelp bool
	flag.BoolVar(&isVersion, "v", false, "version 1.0.0")
	flag.BoolVar(&isHelp, "h", false, "run the program in the same directory as images and pointclouds to produce result.json")
	flag.Parse()
	if isVersion {
		fmt.Println("version 1.0.0")
	}
	if isHelp {
		fmt.Println("run the program in the same directory as images and pointclouds to produce result.json")
	}

	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("Error while reading the directory:", err)
		os.Exit(1)
	}

	myOrderMap := newOrderMap()

	frameNumberCamNames := make(map[string][]string, 0)
	fileInfos := make([]map[string]interface{}, 0)
	for _, file := range files {
		if file.IsDir() {
			if strings.ToLower(file.Name()) == "images" {
				images, err := ioutil.ReadDir("./images")
				if err != nil {
					fmt.Println("Error while reading the directory:", err)
					os.Exit(1)
				}
				for _, image := range images {
					if image.IsDir() {
						imageinfos, err := ioutil.ReadDir(filepath.Join("./images", image.Name()))
						if err != nil {
							fmt.Println("Error while reading the directory:", err)
							os.Exit(1)
						}
						for _, imageinfo := range imageinfos {
							_, ok := frameNumberCamNames[strings.Split(imageinfo.Name(), ".")[0]]
							if !ok {
								frameNumberCamNames[strings.Split(imageinfo.Name(), ".")[0]] = make([]string, 0)
							}
							frameNumberCamNames[strings.Split(imageinfo.Name(), ".")[0]] = append(frameNumberCamNames[strings.Split(imageinfo.Name(), ".")[0]], strings.SplitN(image.Name(), "_", 2)[1])
						}
					}
				}
			} else if strings.ToLower(file.Name()) == "pointclouds" {
				pcds, err := ioutil.ReadDir("./pointclouds")
				if err != nil {
					fmt.Println("Error while reading the directory:", err)
					os.Exit(1)
				}
				for _, pcd := range pcds {
					if pcd.IsDir() {
						pointClouds, err := ioutil.ReadDir(filepath.Join("./pointclouds", pcd.Name()))
						if err != nil {
							fmt.Println("Error while reading the directory:", err)
							os.Exit(1)
						}
						for _, pointcloud := range pointClouds {
							myOrderMap.set(strings.Split(pointcloud.Name(), ".")[0], strings.SplitN(pcd.Name(), "_", 2)[1])
						}
					}
				}
			}
		}
	}

	sort.Strings(myOrderMap.keys)

	for _, framenumber := range myOrderMap.keys {
		fileInfo := make(map[string]interface{}, 0)
		fileInfo["frame_name"] = framenumber

		imagesAnnoInfo := make([]map[string]interface{}, 0)
		camNames, ok := frameNumberCamNames[framenumber]
		if !ok {
			continue
		}
		for _, camname := range camNames {
			curImageInfo := make(map[string]interface{}, 0)
			annotationsObj := make([]map[string]interface{}, 0)
			annotationObj := make(map[string]interface{})
			annotationObj["anno"] = make([]interface{}, 0)
			annotationObj["lidar_name"] = ""
			annotationsObj = append(annotationsObj, annotationObj)
			curImageInfo["annotations"] = annotationsObj
			curImageInfo["cam_name"] = camname
			imagesAnnoInfo = append(imagesAnnoInfo, curImageInfo)
		}

		pointCloudsAnnoInfo := make([]map[string]interface{}, 0)
		for _, lidarName := range myOrderMap.get(framenumber) {
			curPointCloudInfo := make(map[string]interface{}, 0)
			annotationsObj := make([]map[string]interface{}, 0)
			annotationObj := make(map[string]interface{})
			annotationObj["anno"] = make([]interface{}, 0)
			annotationsObj = append(annotationsObj, annotationObj)
			curPointCloudInfo["annotations"] = annotationsObj
			curPointCloudInfo["lidar_name"] = lidarName
			pointCloudsAnnoInfo = append(pointCloudsAnnoInfo, curPointCloudInfo)
		}

		fileInfo["images"] = imagesAnnoInfo
		fileInfo["pointclouds"] = pointCloudsAnnoInfo

		fileInfos = append(fileInfos, fileInfo)
	}

	jsonBytes, err := json.MarshalIndent(fileInfos, "", " ")
	if err != nil {
		fmt.Println("Error while marshaling to JSON:", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile("result.json", jsonBytes, 0644)
	if err != nil {
		fmt.Println("Error while writing JSON to file:", err)
		os.Exit(1)
	}

}
