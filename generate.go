package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"gopkg.in/yaml.v1"
)

func getFilesInDirRecursive(dirPath string) []string {
	fileItems, _ := ioutil.ReadDir(dirPath)

	names := []string{}

	for _, item := range fileItems {
		if item.Name()[0] == '.' {
			continue
		}
		if item.IsDir() {
			subFiles := getFilesInDirRecursive(dirPath + "/" + item.Name())
			
			for _, f := range subFiles {
				names = append(names, f)
			}
		}else {
			names = append(names, dirPath + "/" + item.Name())
		}
	}

	return names
}

func readFile(filePath string) string {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(bytes)
}

func trimString(input string) string {
	return strings.Trim(input, " ")
}

type MetaData struct {
	Title string
	Tags []string `yaml:",flow"`
	Categories []string `yaml:",flow"`
}

func decodeYAMLMetaData(raw string) (MetaData, interface{}) {
	var r MetaData
	r = MetaData{Title:"untitled",Tags:[]string{},Categories:[]string{}}
	err := yaml.Unmarshal([]byte(raw), &r)
	return r, err
}

func main() {
	files := getFilesInDirRecursive("posts")

	fmt.Println(files)

	postTemplate := readFile("templates/posts.html")

	for _, file := range files {
		template := postTemplate

		postData := readFile(file)
		lines := strings.Split(postData, "\n")

		postText := ""
		postMeta := ""
		numMeta := 0 // the current number of --- lines passed
		prevLineMetaTag := true // allow for some empty lines to be after the --- tags

		for _, line := range lines {
			if trimString(line) == "---" {
				numMeta ++
				prevLineMetaTag = true
				continue
			}
			if numMeta % 2 == 0 {
				if len(line) == 0 && prevLineMetaTag {
					continue
				}
				postText += line + "<br/>"
				prevLineMetaTag = false
			}else {
				postMeta += line + "\n"
			}
		}

		metaData, err := decodeYAMLMetaData(postMeta)

		if err != nil {
			fmt.Println("error decoding YAML metadata:", err)

		}
		fmt.Println("---\n", postMeta, "\n", metaData, "\n---")

		for i := range metaData.Tags {
			fmt.Println("tag:", metaData.Tags[i])
		}

		html := strings.Replace(template, "~~~text~~~", postText, -1)
		//fmt.Println(html)

		outFile := "site/" + strings.Replace(file, ".md", ".html", -1)

		outDir := outFile[0:strings.LastIndex(outFile, "/")]
		//fmt.Println(outDir)
		err = os.MkdirAll(outDir, 0777)
		if err != nil {
			fmt.Println(err)
		}
		err = ioutil.WriteFile(outFile, []byte(html), 0666)
		if err != nil {
			fmt.Println(err)
		}else{
			fmt.Println("File '" + outFile + "' written.")
		}
	}
}