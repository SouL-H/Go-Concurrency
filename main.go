package main

import (
	deneme "Concurrency/chan"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Normal şartlarda 2,5 sn süren işlem go ile 0.9 sn düşüyor. Fakat paralel çalıştığı için wg kullanmamız gerekiyor.

func main() {
	var users []User
	usersCh := make(chan User)

	for i := 0; i < 10; i++ {
		//Sondaki () ile fonksiyon direk tetikleniyor
		go func(ch chan User) {
			getRandromUser(ch)
		}(usersCh)

	}
	//Golang kanallar için tasarlanmış bir durum switch case mantığı
	//Sürekli user chanelini kontrol eden for döngüsü
	go func() {
		for {
			select {
			case user := <-usersCh:
				users = append(users, user)
				fmt.Println(user)
			}
		}
	}()
	//Go routin yanında oop dış class örneği
	fmt.Println(deneme.ChanWrite())
	var input string
	fmt.Scan(&input)

}

//User detayı alan bir fonksiyon
func getRandromUser(ch chan User) User {
	response, _ := http.Get("https://random-data-api.com/api/users/random_user")
	//Gelen veriyi byte arr çeviriyor
	body, _ := ioutil.ReadAll(response.Body)
	var user User
	//Gelen veriyi json istenilen formata çevirir
	json.Unmarshal(body, &user)
	ch <- user
	return user

}

type User struct {
	FirshName string `json:"firsh_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
