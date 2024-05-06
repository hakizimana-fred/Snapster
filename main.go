package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log/slog"
	"time"
)

type Data struct {
	Name  string
	Age   int
	Email string
}

type SnapShot struct {
	Timestamp time.Time
	Data      *Data
}

func main() {
	data := &Data{
		Name:  "Fred",
		Age:   25,
		Email: "hakifred20@gmail.com",
	}
	snapshot, err := snapShooter(data)

	if err != nil {
		return
	}

	// Change data
	data.Name = "Haki"
	data.Age = 30

	snapShooter(data)
	//restore
	err = restoreSnap(snapshot, data)

	if err != nil {
		return
	}

	restored, err := json.Marshal(data)
	if err != nil {
		slog.Error("Something went wrong", err)
	}
	err = ioutil.WriteFile("restored.json", restored, 0644)
	if err != nil {
		fmt.Println("Could not restore")
	}
	fmt.Printf("Restored data: %+v\n", data)
}

func snapShooter(data *Data) ([]byte, error) {
	snapShot, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = ioutil.WriteFile("snapshot.json", snapShot, 0644)
	if err != nil {
		return nil, err
	}
	return snapShot, nil
}

func restoreSnap(snapshot []byte, data *Data) error {
	err := json.Unmarshal(snapshot, data)
	if err != nil {
		return err
	}
	return nil
}
