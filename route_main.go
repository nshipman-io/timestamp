package main

import(
	json2 "encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"strconv"

	//"encoding/json"
	"time"
	"net/http"
	"log"
)

type TimeStamp struct {
	Unix int64
	UTC string
}

type ErrorResp struct {
	Error string
}

//Fri, 25 Dec 2015 00:00:00 GMT
var layoutISO = "2006-01-02"
var layoutUTC = "Mon, 02 Jan 2006 03:04:05 MST"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the timestamp api!\n" +
		"Here are some example calls you can make:\n" +
		"0.0.0.0:8080/api/timestamp/2015-12-25\n" +
		"0.0.0.0:8080/api/timestamp/1606874322")
}

func TimeStampGenerator(w http.ResponseWriter, r *http.Request) {
	timestamp := mux.Vars(r)["timestamp"]
	if timestamp == "" {
		output := GetTimeNow()
		json, err := json2.Marshal(output)
		if err != nil {
			log.Println(err)
		}
		w.Write(json)
		return
	}

	output,err := ConvertTimeUTC(timestamp)
	if err != nil {
		log.Println("Notice: Attempting Machine Time conversion")
		output, err = MachineTime(timestamp)
		if err != nil {
			e := ErrorResp{
				Error: "Invalid Date",
			}

			json, err := json2.Marshal(e)
			if err != nil {
				log.Println(err)
			}
			w.Write(json)
			return
		}

	}

	json, err := json2.Marshal(output)
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(json)
}
func GetTimeNow() (*TimeStamp){
	var ts TimeStamp
	t := time.Now()

	ts.Unix = t.Unix()
	ts.UTC = t.Format(layoutUTC)

	return &ts
}

func ConvertTimeUTC(timestamp string) (*TimeStamp, error) {
	var ts TimeStamp
	t,err := time.Parse(layoutISO, timestamp)
	if err != nil {
		log.Println(err)
	}
	ts.Unix = t.Unix()
	ts.UTC = t.Format(layoutUTC)

	return &ts,err

}

func MachineTime(timestamp string) (*TimeStamp, error) {
	var ts TimeStamp
	i, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		log.Println(err)
		return &ts,err
	}
	t := time.Unix(i, 0)
	ts.Unix = t.Unix()
	ts.UTC = t.Format(layoutUTC)
	return &ts, err
}
