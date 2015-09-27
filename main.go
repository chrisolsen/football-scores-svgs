package main
import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"strings"
	"net/http"
)

type team struct {
	Id int
	Name string
	ShortName string
	CrestUrl string
}

func (t *team) String() string {
	return t.Name + ": " + t.CrestUrl
}

func main() {

	input, err := ioutil.ReadFile("./teams.json")
	if (err != nil) {
		panic(err)
	}

	var teams []team
	err = json.Unmarshal(input, &teams)
	if (err != nil) {
		panic(err)
	}

	for _, team := range teams {
		parts := strings.Split(team.CrestUrl, "/")
		filename := parts[len(parts) - 1]
		parts = strings.Split(filename, ".")
		filenamePngExt := strings.Join(parts[:len(parts) - 1], ".") + ".png"

		resp, err := http.Get(team.CrestUrl)
		if (err != nil) {
			fmt.Println("Failed:", team.CrestUrl)
			continue
		}

		data, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		err = ioutil.WriteFile("./svgs/" + filename, data, 0644)
		if (err != nil) {
			fmt.Print(err)
			continue
		}

		// err = exec.Command(fmt.Sprintf("inkscape --file ./svgs/%v --export-png ./pngs/%v", filename, filenamePngExt)).Run()
		fmt.Printf("inkscape --file ./svgs/%v --export-width 200 --export-height 200 --export-png ./pngs/%v\n", filename, filenamePngExt)
	}
}
