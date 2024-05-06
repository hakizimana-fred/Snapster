package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	_, err := snapShooter(data)

	if err != nil {
		return
	}

	// Change data
	data.Name = "Haki"
	data.Age = 30

	fmt.Println("data", data.Name)
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
