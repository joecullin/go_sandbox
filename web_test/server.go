package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"slices"
	"strconv"
)

var port = flag.String("port", "80", "port to listen on")
var appDataFile = flag.String("app-data", "./data/appData.json", "data file")

var templ = template.Must(template.New("qr").Parse(templateStr))

type ReleaseData struct {
	File     string
	Platform string
	Tags     []string
	Version  string
}
type AppData struct {
	Releases []ReleaseData
}

func initAppData(dataFilePath string) AppData {
	dataFileReader, err := os.Open(dataFilePath)
	if err != nil {
		fmt.Printf("Error opening data file %s: %e", dataFilePath, err)
	}
	defer dataFileReader.Close()

	jsonData, err := io.ReadAll(dataFileReader)
	if err != nil {
		fmt.Printf("Error reading data file %s: %e", dataFilePath, err)
	}
	fmt.Println(string(jsonData))
	var apps AppData
	json.Unmarshal(jsonData, &apps)
	fmt.Println(apps)
	fmt.Printf("app: file: %v tags: %v", apps.Releases[2].File, apps.Releases[2].Tags)

	return apps
}

func main() {
	flag.Parse()
	fmt.Println("listening on port", *port)

	appData := initAppData(*appDataFile)
	router := http.NewServeMux()

	router.HandleFunc("GET /api/downloads/{os}/{version}/info", func(w http.ResponseWriter, r *http.Request) {
		osParam := r.PathValue("os")
		versionParam := r.PathValue("version")
		if matched, err := regexp.MatchString("^linux|darwin|windows$", osParam); err != nil || !matched {
			fmt.Fprintf(w, "Validation failed for os=%s. Must be a supported os.\n", osParam)
			return
		}
		if matched, err := regexp.MatchString("^([0-9]+[.][0-9]+)|latest$", versionParam); err != nil || !matched {
			fmt.Fprintf(w, "Validation failed for version=%s. Must be x.x or 'latest'.\n", versionParam)
			return
		}
		fmt.Printf("api/downloads/.../info: ! os=%s version=%s\n", osParam, versionParam)

		i := slices.IndexFunc(appData.Releases, func(r ReleaseData) bool {
			if r.Platform == osParam {
				if versionParam == r.Version || (versionParam == "latest" && slices.Contains(r.Tags, "latest")) {
					return true
				}
			}
			return false
		})
		if i == -1 {
			notFoundPage(w)
			return
		}
		fmt.Printf("api/downloads/.../info: found %v\n", appData.Releases[i])
		jsonString, err := json.Marshal(appData.Releases[i])
		if err != nil {
			fmt.Println("error:", err)
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(jsonString))
	})

	router.HandleFunc("GET /api/testdownload", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("api/testdownload\n")

		data, err := os.ReadFile("./server")
		if err != nil {
			fmt.Printf("Error reading data file!")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Oh no! An error occurred on our end. Maybe try again later?")
			return
		}

		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Length", strconv.Itoa(len(data)))
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})

	// homepage plus catch-all for all other paths and methods:
	router.Handle("/", http.HandlerFunc(infoPage))

	log.Fatal(http.ListenAndServe(":"+*port, router))
}

func notFoundPage(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, `
        <html>
        <head>Not found!</head>
        <body>
        Not found! See <a href="/">docs</a> for help.
        </body>
        </html>`)
}

func infoPage(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		templ.Execute(w, req.FormValue("s"))
	} else {
		notFoundPage(w)
	}
}

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET">
    <input maxLength=1024 size=70 name=s value="" title="Text to QR Encode">
    <input type=submit value="Show QR" name=qr>
</form>
</body>
</html>
`

/*
Some test commands:

curl -v 'http://localhost:5500/'
curl -v 'http://localhost:5500/api/downloads/windows/0.1/info'
curl -v 'http://localhost:5500/api/downloads/windows/1.0/info'
curl -v 'http://localhost:5500/api/downloads/windows/1.1/info'
curl -v 'http://localhost:5500/api/downloads/windows/7.7/info'
curl -v 'http://localhost:5500/api/downloads/windows/latest/info'
curl -v 'http://localhost:5500/api/downloads/windows/latest/infoJOE'
curl -v 'http://localhost:5500/api/downloads/windows/latest1/info'
curl -v 'http://localhost:5500/api/downloads/windowsXXXXX/latest/info'
curl -v 'http://localhost:5500/api/downloads/windowx/latest/info'
curl -v 'http://localhost:5500/api/testdownload' > test && file test
*/
