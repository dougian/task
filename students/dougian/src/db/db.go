package db

import (
	"github.com/boltdb/bolt"
	"os/user"
	"log"
	"encoding/binary"
	"path/filepath"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Id uint64
	Value string
	Status bool
}

var db *bolt.DB
var bucketName []byte = []byte("taskBucket")

func InitBolt() (error) {
	usr, err := user.Current()
	dbFile := filepath.Join(usr.HomeDir, ".task_cli.db")
	db, err = bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		return err
	})
}

func CompleteTask(id uint64) (error) {

	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, id)

		task := bucket.Get(b)
		if task == nil {
			fmt.Println("Key not found!")
			os.Exit(-1)
		}
		var storedTask Task
		json.Unmarshal(task, &storedTask)
		fmt.Println("Found: ", storedTask)

		storedTask.Status = true

		jsonRepr, _  := json.Marshal(storedTask)
		return bucket.Put(b, jsonRepr)
	})
}


func UpdateTask(taskValue Task) (error) {

	return db.Update(func(tx *bolt.Tx) error {

		bucket := tx.Bucket(bucketName)
		jsonRepr, err  := json.Marshal(taskValue)

		if err != nil {
			return err
		}

		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, taskValue.Id)

		return bucket.Put(b, jsonRepr)
	})
}

func AddTask(taskValue Task) (error) {

	return db.Update(func(tx *bolt.Tx) error {

		bucket := tx.Bucket(bucketName)
		id, _ := bucket.NextSequence()
		taskValue.Id = id
		jsonRepr, err  := json.Marshal(taskValue)

		if err != nil {
			return err
		}

		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, id)

		return bucket.Put(b, jsonRepr)
	})
}

func ListTasks() ([]Task) {
	var results []Task
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		c := bucket.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var storedTask Task
			json.Unmarshal(v, &storedTask)

			results = append(results, storedTask)
		}
		return nil
	})
	return results
}
