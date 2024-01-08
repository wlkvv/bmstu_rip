package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const (
	expectedToken = "SomeToken123456"
	updateURL     = "http://176.57.215.76:8000/api/orders/update_deadline/"
	token         = "SomeToken123456"
)

type TimeResult struct {
	OrderID   string `json:"order_id"`
	Fact_time int    `json:"fact_time"`
	Token     string `json:"token"`
}

func main() {
	http.HandleFunc("/async_task", handleProcess)
	fmt.Println("Server running at port :8088")
	http.ListenAndServe(":8088", nil)
}

func handleProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	orderid := r.FormValue("order_id")
	token := r.FormValue("token")
	fmt.Println(orderid, token)

	// if token == "" || token != expectedToken {
	// 	http.Error(w, "Токен не сответствует ожидаемому", http.StatusForbidden)
	// 	fmt.Println("Токен не сответствует ожидаемому")
	// 	fmt.Println(token, expectedToken)
	// 	return
	// }

	w.WriteHeader(http.StatusOK)

	go func() {
		delay := 10
		time.Sleep(time.Duration(delay) * time.Second)

		result := rand.Intn(56) + 5

		expResult := TimeResult{
			OrderID:   orderid,
			Fact_time: result,
			Token:     token,
		}
		fmt.Println("json", expResult)
		jsonValue, err := json.Marshal(expResult)
		if err != nil {
			fmt.Println("Ошибка при маршализации JSON:", err)
			return
		}
		req, err := http.NewRequest(http.MethodPut, updateURL, bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println("Ошибка при формировании запроса на обновление:", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		order, err := client.Do(req)
		if err != nil {
			fmt.Println("Ошибка при отправке запроса на обновление:", err)
			return
		}
		defer order.Body.Close()

		fmt.Println("Ответ от сервера обновления:", order.Status)
	}()
}
