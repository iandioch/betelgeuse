package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"gopkg.in/yaml.v1"
	"github.com/robertkrimen/otto"
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
	Id int `yaml: ",omitempty"` // automagically generated
}

type PostData struct {
	File string // the path to the file
	Meta MetaData // the metadata of the post (parsed from the YAML)
	RawContent string // the contents of the original file
	Lines []string // the raw file split into lines
	ContentLines []string // the lines of the file without the metadata, but with the inline code
	ParsedContent string // the file without any inline code or metadata (ie. the finished post)
}

func decodeYAMLMetaData(raw string) (MetaData, interface{}) {
	var r MetaData
	r = MetaData{Title:"untitled",Tags:[]string{},Categories:[]string{}}
	err := yaml.Unmarshal([]byte(raw), &r)
	return r, err
}

func runJavascript(script string, currId int, posts []PostData) (string, interface{}) {
	vm := otto.New()

	response := ""

	vm.Set("echo", func(call otto.FunctionCall) otto.Value {
	    fmt.Printf("Script echo: %s\n", call.Argument(0).String())
	    response += call.Argument(0).String()
	    return otto.Value{}
	})

	vm.Set("posts", posts)
	vm.Set("currId", currId)

	_, err := vm.Run(script)

	return response, err
}

func main() {
	files := getFilesInDirRecursive("posts")

	fmt.Println(files)

	//postTemplate := readFile("templates/posts.html")

	postGenerator := readFile("templates/posts.js")

	allPostData := make([]PostData, len(files))

	for index, file := range files {
		raw := readFile(file)
		lines := strings.Split(raw, "\n")

		allPostData[index] = PostData{file, MetaData{"not parsed", []string{"not parsed"}, []string{"not parsed"}, index}, raw, lines, []string{"not parsed"}, "not parsed"}
	}

	for index, entry := range allPostData {
		// parse out the metadata
		postMeta := ""
		numMeta := 0
		prevLineMetaTag := true
		unParsedLines := []string{}

		for _, line := range entry.Lines {
			if trimString(line) == "---" {
				numMeta ++
				prevLineMetaTag = true
				continue
			}
			if numMeta % 2 == 0 {
				if len(line) == 0 && prevLineMetaTag {
					continue
				}
				unParsedLines = append(unParsedLines, line)
				prevLineMetaTag = false
			}else{
				postMeta += line + "\n"
			}
		}

		entry.ContentLines = unParsedLines

		fmt.Println(postMeta)
		metaData, err := decodeYAMLMetaData(postMeta)

		if err != nil {
			fmt.Println("error decoding YAML metadata:", err)
		}

		//entry.Meta = metaData
		metaData.Id = index

		entry.Meta = metaData
		fmt.Println(index, entry.Meta)

		allPostData[index] = entry
	}

	for index, value := range allPostData {
		numCode := 0
		currCode := ""
		postText := ""

		/*variables := make(map[string]interface{})
		variables["id"] = index
		variables["curr"] = value*/


		for _, line := range value.ContentLines {
			if trimString(line) == "```" {
				numCode ++
				if numCode % 2 == 1 {
					currCode = "" // new bit of code
				} else {
					// code bit just ended, run it and insert its result to the document
					response, err := runJavascript(currCode, index, allPostData)
					if err != nil {
						fmt.Println(err)
						postText += err.(string)
					} else {
						fmt.Println(response)
						postText += response
					}
				}
				continue
			}

			if numCode % 2 == 1 {
				currCode += line + "\n"
				continue
			}else{
				postText += line + "<br />"
			}
		}

		allPostData[index].ParsedContent = postText
	}

	for index, value := range allPostData {
		file := value.File
		html, err := runJavascript(postGenerator, index, allPostData)
		if err != nil {
			panic(err)
			return
		}
		//html := value.ParsedContent
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
	/*for _, file := range files {
		template := postTemplate

		postData := readFile(file)
		lines := strings.Split(postData, "\n")

		postText := ""
		postMeta := ""
		numMeta := 0 // the current number of --- lines passed
		prevLineMetaTag := true // allow for some empty lines to be after the --- tags
		numCode := 0
		currCode := ""

		for _, line := range lines {
			if trimString(line) == "---" {
				numMeta ++
				prevLineMetaTag = true
				continue
			}else if trimString(line) == "```" {
				numCode ++
				if numCode % 2 == 1{
					currCode = "" // new bit of code
				}else {
					// code bit just ended, run it and insert its result to the document

					response, err := runJavascript(currCode, map[string]interface{}{"title":"FILLER TITLE"})
					if err != nil {
						fmt.Println(err)
						postText += err.(string)
					} else {
						fmt.Println(response)
						postText += response
					}

				}
				continue
			}

			if numCode % 2 == 1 {
				currCode += line
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
		//fmt.Println("---\n", postMeta, "\n", metaData, "\n---")

		for i := range metaData.Tags {
			fmt.Println("tag:", metaData.Tags[i])
		}

		html := strings.Replace(template, "~~~text~~~", postText, -1)
		html = strings.Replace(html, "~~~title~~~", metaData.Title, -1)
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
	}*/

	//_ = otto.New()
}