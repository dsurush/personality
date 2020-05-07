package settings

import (
	"encoding/json"
	//"fmt"
//	log "github.com/sirupsen/logrus"
	"io/ioutil"
	//"net/http"
	//"strings"
)

var (
	// AppSettings app settnigs
	AppSettings Settings
	HamsoyaSettings Settings
)

// ReadSettings to init app settings
func ReadSettings(filepath string) Settings {
	var appParams Settings
	doc, err := ioutil.ReadFile(filepath)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(doc, &appParams)
	if err != nil {
		panic(err)
	}

	return appParams
}
