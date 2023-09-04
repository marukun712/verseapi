package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Item struct {
	Title string `json:"title"`
	Image string `json:"image"`
	Id    string `json:"id"`
	Time  string `json:"time"`
}

var member struct {
	uno      string
	itsuki   string
	nanamona string
	otoha    string
	sera     string
}

var items = []Item{}

var apikey string

func main() {
	Init()

	//メンバーごとにデータをfetch
	unoData := fetchData(member.uno)
	getAllItems(unoData["items"].([]interface{}))

	itsukiData := fetchData(member.itsuki)
	getAllItems(itsukiData["items"].([]interface{}))

	nanamonaData := fetchData(member.nanamona)
	getAllItems(nanamonaData["items"].([]interface{}))

	otohaData := fetchData(member.otoha)
	getAllItems(otohaData["items"].([]interface{}))

	seraData := fetchData(member.sera)
	getAllItems(seraData["items"].([]interface{}))

	//書き込み用のjsonファイルをopen
	file, err := os.Create("res.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//jsonファイルに書き込み
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(items); err != nil {
		log.Fatal(err)
	}
}

func Init() {
	//.envをロード
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//apikey
	apikey = os.Getenv("APIKEY")

	// member構造体の初期化
	member.uno = "UCLfAsY3iMUAF2vvDxvIjymQ"
	member.itsuki = "UCi3aT9cC6YyC0BnOpm2XBEw"
	member.nanamona = "UCOg01LJmZF9UnwFbly73CVw"
	member.otoha = "UCYcnLc0n1ryBDZeGWQTVJ_g"
	member.sera = "UCvXsXmpMKthJuX8XHbsnOjQ"
}

// チャンネルごとに全ての配信枠を取得
func getAllItems(item []interface{}) {
	for _, item := range item {
		data := getStreamInfo(item)
		items = append(items, *data)
	}
}

// 配信情報を取得
func getStreamInfo(item interface{}) *Item {
	videoId := (item.(map[string]interface{})["id"].(map[string]interface{})["videoId"].(string))
	title := (item.(map[string]interface{})["snippet"].(map[string]interface{})["title"].(string))
	time := fetchStreamStartTime(videoId)

	data := new(Item)

	data.Title = title

	data.Image = "https://i.ytimg.com/vi/" + videoId + "/mqdefault_live.jpg"

	data.Id = videoId

	data.Time = time

	return data
}

func fetchData(id string) map[string]any {
	//URL
	url := "https://www.googleapis.com/youtube/v3/search"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	//クエリパラメータ
	params := request.URL.Query()
	params.Add("key", apikey)
	params.Add("part", "snippet")
	params.Add("channelId", id)
	params.Add("eventType", "upcoming")
	params.Add("type", "video")

	request.URL.RawQuery = params.Encode()

	//request
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	//構造体にデコード
	var rawJson map[string]any
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&rawJson); err != nil {
		log.Fatal(err)
	}

	return rawJson
}

// 配信開始時刻を取得
func fetchStreamStartTime(id string) string {
	//URL
	url := "https://www.googleapis.com/youtube/v3/videos"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	//クエリパラメータ
	params := request.URL.Query()
	params.Add("key", apikey)
	params.Add("part", "liveStreamingDetails")
	params.Add("id", id)

	request.URL.RawQuery = params.Encode()

	//request
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	//構造体にデコード
	var rawJson map[string]any
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&rawJson); err != nil {
		log.Fatal(err)
	}

	//scheduledStartTimeを取得
	item := rawJson["items"].([]interface{})

	var time string

	for _, item := range item {
		time = item.(map[string]interface{})["liveStreamingDetails"].(map[string]interface{})["scheduledStartTime"].(string)
	}

	return time
}
