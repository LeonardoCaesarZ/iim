package main

import (
	"encoding/json"
	"fmt"
	"iim/model"
	"io/ioutil"
	"net/http"
	"reflect"
)

type Person struct {
	Name string
	age  int
}

func tmp(dest interface{}) {
	if reflect.TypeOf(dest).Kind() == reflect.Slice {
		// value := reflect.ValueOf(dest)
		ele := reflect.TypeOf(dest).Elem()
		for i := 0; i < ele.NumField(); i++ {
			fmt.Println(ele.Field(i).Type)
		}
		// for i := 0; i < value.Len(); i++ {
		// 	fmt.Println(value.Index(i))
		// }
	} else {
		fmt.Println("not slice")
	}
}

func main() {
	xs := []Person{}
	// xs = append(xs, Person{"abc", 123})
	// xs = append(xs, Person{"bcd", 234})
	tmp(xs)

	// http.HandleFunc("/auth", authentication)
	// err := http.ListenAndServe(":9999", nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
	// fmt.Println("auth server is stopped!")

	// r := sql.NewRedis("localhost", "6379")

	// err := r.Set("a", 123)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// reply, err := r.GetInt("a")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(reply)

	// key := auth.GenerateAESKey()
	// fmt.Println("aes key: ", string(key))

	// encrypted, err := auth.AESEncrypt([]byte("asdfasdfasdfasdfasdf"), key)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("ctypted text: ", string(encrypted), len(encrypted))

	// decrypted, err := auth.AESDecrypt(encrypted, key)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(decrypted))

	// auth.GenerateRSAKeyPair(1024)
	// if err := auth.Init(); err != nil {
	// 	log.Fatal(err)
	// }

	// data := []byte("qwodasdfw")
	// a, err := auth.RSAEncrypt(data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(a), len(a))

	// b, err := auth.RSADecrypt(a)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(b), len(b))
}

func authentication(w http.ResponseWriter, r *http.Request) {
	// api layer

	r.ParseForm()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// log it and return
	}

	data := &model.AuthRequest{}
	err = json.Unmarshal(body, data)
	if err != nil {
		// log it and return
	}

	if data.Account == "" || data.Passwd == "" || data.AesKey == "" {
		// wrong params, return 403
	}

	// bll layer

}
