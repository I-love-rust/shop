package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const InBoxHashSecret = "asdfsfsfojnegounegiuegngeitbe"
const OutBoxHashSecret = "orngiurinuirgniurnortbnrtigm"
const ReturnUrl = "http://web:1111/api/purchase/success"

type Code struct {
	Code        int
	Description string `json:"description"`
}

var HashSumNotEq = Code{Description: "Incorrect hashsum", Code: http.StatusUnauthorized}
var IternalServerError = Code{Description: "Iternal server error", Code: http.StatusInternalServerError}

// just emulation transaction
func register(w http.ResponseWriter, r *http.Request) {
	log.Println("new transaction!")
	user := r.URL.Query().Get("user")
	product := r.URL.Query().Get("product")
	price := r.URL.Query().Get("price")
	transHashSum := r.URL.Query().Get("ts")

	realHashSum := GetMD5Hash(fmt.Sprintf("%s,%s,%s,%s", user, product, price, InBoxHashSecret))
	if realHashSum != transHashSum {
		ReturnErr(w, HashSumNotEq)
		return
	}

	URL, err := url.Parse(ReturnUrl)
	if err != nil {
		ReturnErr(w, IternalServerError)
		return
	}

	urlQuery := URL.Query()

	urlQuery.Add("user", user)
	urlQuery.Add("product", product)
	urlQuery.Add("price", price)
	urlQuery.Add("ts",
		GetMD5Hash(fmt.Sprintf("%s,%s,%s,%s", user, product, price, OutBoxHashSecret)))

	URL.RawQuery = urlQuery.Encode()

	client := &http.Client{}
	req, err := http.NewRequest("GET", URL.String(), nil)

	if err != nil {
		ReturnErr(w, IternalServerError)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func ReturnErr(w http.ResponseWriter, err Code) {
	w.WriteHeader(err.Code)
	json.NewEncoder(w).Encode(err)
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func main() {
	http.HandleFunc("/reg", register)

	http.ListenAndServe("0.0.0.0:7777", nil)
}
