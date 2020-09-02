//the task is not up to the mark, due to lack of time as i am currently working as an intern in another software company, all the things i have done are done by me, within a time frame of 4-5 hours, i request you to have a look at it, Thanking you Regards, Paras Dandwani.

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"encoding/json"
	// "io/ioutil"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



// import "go.mongodb.org/mongo-driver/mongo"

// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// defer cancel()

// if err != nil { log.Fatal(err) }


//mongodb+srv://user:<password>@meeting.a5tc7.mongodb.net/<dbname>?retryWrites=true&w=majority

type meeting struct{
	Id int `json:"id" bson:"_id,omitempty"`
	Title string `json:"title" bson:"title"`
	Participants []string `json:"participants" bson:"participants"`
	StartTime int `json:"startTime" bson:"startTime"`
	EndTime int `json:"endTime" bson:"endTime"`
	TimeStamp int `json:"timeStamp" bson:"timeStamp"`
}

// type user struct{
// 	Name string `json:"name"`
// 	Email string `json:"email"`
// 	Rsvp string `json:"rsvp"`
// }

func connectToMongo()  (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://user:12345678@meeting.a5tc7.mongodb.net/<dbname>?retryWrites=true&w=majority"))

	if err != nil {
		return nil, fmt.Errorf("error connecting to mongodb: %v", err)
	}
	defer client.Disconnect(ctx)

	database := client.Database("quickstart")
	meetings := database.Collection("meetings")
	// episodesCollection := database.Collection("episodes")
	return client, nil
}

func meetingRoute(w http.ResponseWriter, r *http.Request) {
	// options := options.Client().ApplyURI(
	// 	"mongodb+srv://user:<password>@meeting.a5tc7.mongodb.net/<dbname>?retryWrites=true&w=majority",
	//  )
	// client, err := mongo.Connect(context.Background(), options)
	//b, err := ioutil.ReadAll(r.Body)
	// defer r.Body.Close()

	// if err != nil {
    //     http.Error(w, err.Error(), http.StatusBadRequest)
    //     return
	// }

	m := meeting{}
	err := json.NewDecoder(r.Body).Decode(&m)
	// u, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	// err = json.Unmarshal(u, &u)
	
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	// res, _ := json.Marshal(&u)
	userJson , err := json.Marshal(m)

	// fmt.Fprintln(w,string(res))
	w.Write(userJson)
	//w.Write(res)

}



func main() {
	connectToMongo()
    http.HandleFunc("/meeting", meetingRoute)
	log.Fatal(http.ListenAndServe(":8080", nil))
	
}

