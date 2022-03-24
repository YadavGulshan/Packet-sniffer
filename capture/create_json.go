package capture

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CreateJson(fileName string) (err error) {
	if fileName != "" && filepath.Ext(strings.TrimSpace(fileName)) == ".pcap" {
		fileObject, err := os.Open(fileName)
		defer fileObject.Close()
		if err != nil {
			log.Println(err)
			return err
		}

		fileData, err := ioutil.ReadAll(fileObject)
		if err != nil {
			log.Println(err)
			return err
		}
		// convert the file content to json
		jsonData, err := json.Marshal(fileData)
		if err != nil {
			log.Println(err)
			return err
		}
		// write the json data to a file
		err = ioutil.WriteFile("output.json", jsonData, 0644)
		if err != nil {
			log.Println(err)
			return err
		}

		return err
	}
	err = fmt.Errorf("please enter a valid filename")

	return
}
