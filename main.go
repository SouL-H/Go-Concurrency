package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

//Normal şartlarda 2,5 sn süren işlem go ile 0.9 sn düşüyor. Fakat paralel çalıştığı için wg kullanmamız gerekiyor.
func main() {
	var users []User
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		//Sondaki () ile fonksiyon direk tetikleniyor
		go func() {
			users = append(users, getRandromUser(&wg))
		}()

	}
	wg.Wait()
	//Başlangıç zamanından şu ana kadar ki zamanı hesaplar since fonksiyonu.
	fmt.Println(time.Since(start))
	fmt.Println(users)
}

//User detayı alan bir fonksiyon
func getRandromUser(wg *sync.WaitGroup) User {
	defer wg.Done()
	response, _ := http.Get("https://random-data-api.com/api/users/random_user")
	//Gelen veriyi byte arr çeviriyor
	body, _ := ioutil.ReadAll(response.Body)
	var user User
	//Gelen veriyi json istenilen formata çevirir
	json.Unmarshal(body, &user)
	return user

}

type User struct {
	FirshName string `json:"firsh_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
