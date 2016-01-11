package main

import (
	"fmt"
	"os"
	"strconv"
	"io/ioutil"
	"strings"
	"gopkg.in/yaml.v1" // YAML decoder
	"github.com/robertkrimen/otto" // Javascript interpreter
	"github.com/russross/blackfriday" // Markdown to HTML
)

// get all the files in a given directory and every file in every subdirectory recursively
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

// read the specified file and return its string contents
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

type DateData struct {
	Year int
	Month int
	Day int
}

type PostData struct {
	File string // the path to the file
	Meta MetaData // the metadata of the post (parsed from the YAML)
	RawContent string // the contents of the original file
	Lines []string // the raw file split into lines
	ContentLines []string // the lines of the file without the metadata, but with the inline code
	ParsedContent string // the file without any inline code or metadata (ie. the finished post)
	Location string // the location of the resultant html file
	Date DateData // the Year, Month and Day the post was created (parsed from the file path ./posts/YYYY/MM/DD/blah.md)
}

func decodeYAMLMetaData(raw string) (MetaData, interface{}) {
	var r MetaData
	r = MetaData{Title:"untitled",Tags:[]string{},Categories:[]string{}}
	err := yaml.Unmarshal([]byte(raw), &r)
	return r, err
}

// run the Javascript given in the first argument to generate a page
func runJavascript(script string, currId int, posts []PostData) (string, interface{}) {
	vm := otto.New()
	response := ""

	// make any call to echo() in the JS append the argument to the output html
	vm.Set("echo", func(call otto.FunctionCall) otto.Value {
	    response += call.Argument(0).String()
	    return otto.Value{}
	})

	vm.Set("posts", posts)
	vm.Set("currId", currId)
	_, err := vm.Run(script)
	return response, err
}

func main() {
	postGenerator := readFile("templates/posts.js")
	tagPageGenerator := readFile("templates/tag.js")
	categoryPageGenerator := readFile("templates/category.js")

	// get a lits of all of the files in the ./posts dir, and figure out which are actual posts (*.md) and which are asset files to go with posts (eg. image files for a post)
	postDirFiles := getFilesInDirRecursive("posts")
	postFiles := []string{}
	postAssetFiles := []string{}
	for _, file := range postDirFiles{
		if file[len(file)-3:] == ".md" {
			postFiles = append(postFiles, file)
		}else{
			postAssetFiles = append(postAssetFiles, file)
		}
	}

	fmt.Println(len(postFiles), "posts found, with", len(postAssetFiles), "additional assets.")

	// load all of the raw post content into allPostData
	allPostData := make([]PostData, len(postFiles))
	for index, file := range postFiles {
		raw := readFile(file)
		lines := strings.Split(raw, "\n")

		allPostData[index] = PostData{file, MetaData{"not parsed", []string{"not parsed"}, []string{"not parsed"}, index}, raw, lines, []string{"not parsed"}, "not parsed", "not parsed", DateData{-1, -1, -1}}
	}

	// pick out and decode the YAML parts of each post, and pick out the unparsed Markdown
	for index, entry := range allPostData {
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
		metaData, err := decodeYAMLMetaData(postMeta)

		if err != nil {
			fmt.Println("error decoding YAML metadata:", err)
		}

		metaData.Id = index
		entry.Meta = metaData
		allPostData[index] = entry
	}

	// convert the Markdown to HTML, get the location of the final file, and parse the date from the URL
	for index, value := range allPostData {
		postText := ""
		for _, line := range value.ContentLines {
			postText += line + "\n"
		}

		allPostData[index].ParsedContent = string(blackfriday.MarkdownBasic([]byte(postText)))
		allPostData[index].Location = strings.Replace(value.File, ".md", ".html", -1)

		dateParts := strings.Split(allPostData[index].Location, "/")
		year, _ := strconv.Atoi(dateParts[1])
		month, _ := strconv.Atoi(dateParts[2])
		day, _ := strconv.Atoi(dateParts[3])
		allPostData[index].Date = DateData{year, month, day}
	}

	// write the final posts to ./site/posts
	numFilesWritten := 0
	for index, value := range allPostData {
		html, err := runJavascript(postGenerator, index, allPostData)
		if err != nil {
			panic(err)
			return
		}
		outFile := "site/" + value.Location

		outDir := outFile[0:strings.LastIndex(outFile, "/")]
		err = os.MkdirAll(outDir, 0777)
		if err != nil {
			fmt.Println(err)
		}
		err = ioutil.WriteFile(outFile, []byte(html), 0666)
		if err == nil {
			numFilesWritten ++
		}else{
			fmt.Println(err)
		}
	}
	fmt.Printf("%d post files written successfully.\n", numFilesWritten)

	// copy all of the asset files from ./posts to ./site/posts
	numFilesWritten = 0
	for _, origFile := range postAssetFiles {
		outFile := "site/" + origFile
		outDir := outFile[0:strings.LastIndex(outFile, "/")]

		fileData := readFile(origFile)
		err:= os.MkdirAll(outDir, 0777)
		if err != nil {
			fmt.Println(err)
		}
		err = ioutil.WriteFile(outFile, []byte(fileData), 0666)
		if err == nil {
			numFilesWritten ++
		}else{
			fmt.Println(err)
		}
	}
	fmt.Printf("%d asset files written successfully.\n", numFilesWritten)

	// compile lists of all of the tags and all of the categories
	categories := make(map[string][]PostData)
	tags := make(map[string][]PostData)
	for _, val := range allPostData {
		for _, cat := range val.Meta.Categories {
			_, ok := categories[cat]
			if ok {
				categories[cat] = append(categories[cat], val)
			} else {
				categories[cat] = []PostData{val}
			}
		}
		for _, tag := range val.Meta.Tags {
			_, ok := tags[tag]
			if ok {
				tags[tag] = append(tags[tag], val)
			} else {
				tags[tag] = []PostData{val}
			}
		}
	}

	// create pages for each category
	numFilesWritten = 0
	for category, posts := range categories {
		html, err := runJavascript(categoryPageGenerator, -1, posts)
		if err != nil {
			panic(err)
			return
		}
		outFile := "site/categories/" + category + ".html"

		outDir := outFile[0:strings.LastIndex(outFile, "/")]
		err = os.MkdirAll(outDir, 0777)
		if err != nil {
			fmt.Println(err)
		}
		err = ioutil.WriteFile(outFile, []byte(html), 0666)
		if err == nil {
			numFilesWritten ++
		}else{
			fmt.Println(err)
		}
	}
	fmt.Printf("%d category pages written successfully.\n", numFilesWritten)

	// create pages for each tag
	numFilesWritten = 0
	for tag, posts := range tags {
		html, err := runJavascript(tagPageGenerator, -1, posts)
		if err != nil {
			panic(err)
			return
		}
		outFile := "site/tags/" + tag + ".html"

		outDir := outFile[0:strings.LastIndex(outFile, "/")]
		err = os.MkdirAll(outDir, 0777)
		if err != nil {
			fmt.Println(err)
		}
		err = ioutil.WriteFile(outFile, []byte(html), 0666)
		if err == nil {
			numFilesWritten ++
		}else{
			fmt.Println(err)
		}
	}
	fmt.Printf("%d tag pages written successfully.\n", numFilesWritten)

	indexGenerator := readFile("templates/index.js");
	html, err := runJavascript(indexGenerator, -1, allPostData)
	if err != nil {
		panic(err)
		return
	}
	err = ioutil.WriteFile("site/index.html", []byte(html), 0666)
	if err != nil {
		panic(err)
		return
	}
}