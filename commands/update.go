package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/astaxie/beego"
	"github.com/lifei6671/godoc/conf"
)

//检查最新版本.
func CheckUpdate() {

	resp, err := http.Get("https://api.github.com/repos/lifei6671/godoc/tags")

	if err != nil {
		beego.Error("CheckUpdate => ", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Error("CheckUpdate => ", err)
		os.Exit(1)
	}

	var result []*struct {
		Name string `json:"name"`
	}

	err = json.Unmarshal(body, &result)
	fmt.Println("MinDoc current version => ", conf.VERSION)
	if err != nil {
		beego.Error("CheckUpdate => ", err)
		os.Exit(0)
	}

	fmt.Println("MinDoc last version => ", result[0].Name)
	os.Exit(0)

}
