package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"flag"
	"io/ioutil"
	"os"
	"os/user"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)


var api *anaconda.TwitterApi
func main(){
	screennamePtr := flag.String("name", "", "twitter @ screenname")
	idPtr  := flag.Int64("id", 0, "twitter integer Id")
	flag.Parse()
	
	api = connectTwitterApi()
	var uv=url.Values{}

	var id int64
	if *idPtr != 0 {
		id = *idPtr
	} else if *screennamePtr != "" {
		id = name2id(*screennamePtr)
	} else {
		fmt.Fprintln(os.Stderr, "-name=username or -id=99999999")
		os.Exit(1)
	}
	fmt.Println(id)

	uv["count"] = []string{"100"}
	lists, err := api.GetListsOwnedBy(id, uv)
	if err != nil {
		fmt.Println("%#v\n", err)
		os.Exit(1)
	}
	//jsonLists, _ := json.Marshal(lists)
	//fmt.Println(string(jsonLists))

	for i, list := range lists {
		fmt.Printf("%d.\n", i + 1)
		fmt.Println("Name: ", list.Name)
		fmt.Println("URL: ", list.URL)
		fmt.Println("Created at: ", list.CreatedAt)
		fmt.Println("Id: ", list.Id)
		fmt.Println("Count(Member,Subscriber): ", list.MemberCount, list.SubscriberCount)
		fmt.Println("Mode: ", list.Mode)
		//jsonList, _ := json.Marshal(list)
		//fmt.Println(string(jsonList))
		if true {
			//uv["count"] = []string{"5000"} //面倒
			uv.Set("count", "5000")
			//uv.Set("include_entities", "false")　//status中のentityに関係
			uv.Set("skip_status", "true") //データ減少
			
			//users, err := api.GetListMembers("akira0000k", list.Id, uv)
			users, err := api.GetListMembers("", list.Id, uv)
			if err != nil {
				fmt.Println("%#v\n", err)
				os.Exit(1)
			}
			//jsonUsers, _ := json.Marshal(users)
			//fmt.Println(string(jsonUsers))
			for _, user := range users.Users {
				username := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(user.Name,	       "\n", `\n`), "\r", `\r`), "\"", `\"`)
				userdesc := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(user.Description, "\n", `\n`), "\r", `\r`), "\"", `\"`)
				fmt.Printf("%s\t\"%s\"\t\"%s\"\n", user.ScreenName, username, userdesc)
			}
		}
	}
}

func name2id(screen_name string) (id int64) {
	var uv=url.Values{}
	uv.Set("skip_status", "true") //データ減少
	users, err := api.GetUsersLookup(screen_name, uv)
	if err != nil {
		fmt.Println("%#v\n", err)
		os.Exit(1)
	}
	//jsonUser, _ := json.Marshal(users[0])
	//fmt.Println(string(jsonUser))
	//os.Exit(9)
	
	var userinfo anaconda.User = users[0]

	id = userinfo.Id
	return id
}
	

func connectTwitterApi() *anaconda.TwitterApi {
	usr, _ := user.Current()
	raw, error := ioutil.ReadFile(usr.HomeDir + "/twitter/twitterAccount.json")

	if error != nil {
		fmt.Println(error.Error())
		return nil
	}

	var twitterAccount TwitterAccount
	json.Unmarshal(raw, &twitterAccount)

	return anaconda.NewTwitterApiWithCredentials(
		twitterAccount.AccessToken,
		twitterAccount.AccessTokenSecret,
		twitterAccount.ConsumerKey,
		twitterAccount.ConsumerSecret)

}

type TwitterAccount struct {
	AccessToken string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	ConsumerKey string `json:"consumerKey"`
	ConsumerSecret string `json:"consumerSecret"`
}

