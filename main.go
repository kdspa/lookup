package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Config struct {
	token string
	port  int
}

var userCache = make(map[string]string)
var guildCache = make(map[string]string)

func userRoute(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	cachedUser, cached := userCache[id]
	if cached {
		fmt.Fprintf(w, cachedUser)
	} else {
		fmt.Printf("Fetching user: %s\n", id)
		client := &http.Client{}
		req, getErr := http.NewRequest("GET", "https://discord.com/api/v8/users/"+id, nil)
		if getErr != nil {
			log.Fatal(getErr)
		}
		req.Header.Add("Authorization", "Bot "+os.Getenv("TOKEN"))
		res, getErr := client.Do(req)
		fmt.Printf("HTTP: %s\n", res.Status)
		if getErr != nil {
			log.Fatal(getErr)
		}
		if res.Status != "200 OK" {
			http.Error(w, "An error occurred while fetching this user", http.StatusInternalServerError)
		} else {
			if res.Body != nil {
				defer res.Body.Close()
			}
			body, readErr := ioutil.ReadAll(res.Body)
			if readErr != nil {
				log.Fatal(readErr)
			}
			userCache[id] = string(body)
			fmt.Fprintf(w, string(body))
		}
	}
}

func guildRoute(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	cachedGuild, cached := guildCache[id]
	if cached {
		fmt.Fprintf(w, cachedGuild)
	} else {
		fmt.Printf("Fectching guild: %s\n", id)
		client := &http.Client{}
		req, getErr := http.NewRequest("GET", "https://discord.com/api/v8/users/"+id, nil)
		if getErr != nil {
			log.Fatal(getErr)
		}
		req.Header.Add("Authorization", "Bot "+os.Getenv("TOKEN"))
		res, getErr := client.Do(req)
		fmt.Printf("HTTP: %s\n", res.Status)
		if getErr != nil {
			log.Fatal(getErr)
		}
		if res.Status != "200 OK" {
			http.Error(w, "An occurred while fetching this guild", http.StatusInternalServerError)
		} else {
			if res.Body != nil {
				defer res.Body.Close()
			}
			body, readErr := ioutil.ReadAll(res.Body)
			if readErr != nil {
				log.Fatal(readErr)
			}
			guildCache[id] = string(body)
			fmt.Fprintf(w, string(body))
		}
	}
}

func inviteRoute(w http.ResponseWriter, r *http.Request) {

	code := mux.Vars(r)["code"]

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	client := &http.Client{}
	req, getErr := http.NewRequest("GET", "https://discord.com/api/v8/invites/"+code, nil)
	if getErr != nil {
		log.Fatal(getErr)
	}
	req.Header.Add("Authorization", "Bot "+os.Getenv("TOKEN"))
	res, getErr := client.Do(req)
	fmt.Printf("HTTP: %s\n", res.Status)
	if getErr != nil {
		log.Fatal(getErr)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Fprintf(w, string(body))

}
