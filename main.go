// Ниже реализован сервис бронирования номеров в отеле. В предметной области
// выделены два понятия: Order — заказ, который включает в себя даты бронирования
// и контакты пользователя, и RoomAvailability — количество свободных номеров на
// конкретный день.
//
// Задание:
// - провести рефакторинг кода с выделением слоев и абстракций
// - применить best-practices там где это имеет смысл
// - исправить имеющиеся в реализации логические и технические ошибки и неточности
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Order struct {
	HotelID   string    `json:"hotel_id"` // int | кастомный тип ID завести справочник отелей
	RoomID    string    `json:"room_id"`  // int | кастомный тип ID завести справочник отелей
	UserEmail string    `json:"email"`    // пользователя вынести в справочник и поле сделать user_id(int)
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`

	// CreatedAt
	// Price -
	// Status NEW|NEED_CONFIRM
}

// (mutex lock unlock)
var Orders = []Order{}

type RoomAvailability struct {
	HotelID string    `json:"hotel_id"` // int | кастомный тип ID надо завести справочник отелей
	RoomID  string    `json:"room_id"`  // int | кастомный тип ID надо завести справочник комнат
	Date    time.Time `json:"date"`
	Quota   int       `json:"quota"` // кажется тут bool подойдет

}

// (mutex lock unlock)
var Availability = []RoomAvailability{
	{"reddison", "lux", date(2024, 1, 1), 1},
	{"reddison", "lux", date(2024, 1, 2), 1},
	{"reddison", "lux", date(2024, 1, 3), 1},
	{"reddison", "lux", date(2024, 1, 4), 1},
	{"reddison", "lux", date(2024, 1, 5), 0},
}

func main() {
	// у приложения нет конфига
	mux := http.NewServeMux()
	mux.HandleFunc("/orders", createOrder) // ручка обрабатывает все методы (POST|GET)

	LogInfo("Server listening on localhost:8080")
	err := http.ListenAndServe(":8080", mux) // захардкожен порт
	if errors.Is(err, http.ErrServerClosed) {
		LogInfo("Server closed")
	} else if err != nil {
		LogErrorf("Server failed: %s", err)
		os.Exit(1) // избыточно
	}
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder Order                        // создать модель OrderRequest
	json.NewDecoder(r.Body).Decode(&newOrder) // не обработана ошибка + ее лог и 400

	// валидация данных + 422 + log
	// вернуть ошибку {"code":"validation_error","message": "Ошибка входных данных"}

	daysToBook := daysBetween(newOrder.From, newOrder.To)

	// логику с проверкой свободных дат вынести в отдельный сервис
	unavailableDays := make(map[time.Time]struct{})

	for _, day := range daysToBook {
		unavailableDays[day] = struct{}{}
	}

	for _, dayToBook := range daysToBook {
		for i, availability := range Availability {
			if !availability.Date.Equal(dayToBook) || availability.Quota < 1 {
				continue
			}

			availability.Quota -= 1 // номера бронируются даже если не удалось создать заказ
			Availability[i] = availability
			delete(unavailableDays, dayToBook)
		}
	}

	if len(unavailableDays) != 0 {
		http.Error(w, "Hotel room is not available for selected dates", http.StatusInternalServerError)
		LogErrorf("Hotel room is not available for selected dates:\n%v\n%v", newOrder, unavailableDays)
		return
	}

	//  вызов сервиса запуска заказа
	Orders = append(Orders, newOrder)

	w.Header().Set("Content-Type", "application/json") // в мидлварь
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOrder) // обработка ошибки + создать модель OrderResponse

	LogInfo("Order successfully created: %v", newOrder)
}

func daysBetween(from time.Time, to time.Time) []time.Time {
	if from.After(to) {
		return nil
	}

	days := make([]time.Time, 0)
	for d := toDay(from); !d.After(toDay(to)); d = d.AddDate(0, 0, 1) {
		days = append(days, d)
	}

	return days
}

func toDay(timestamp time.Time) time.Time {
	return time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day(), 0, 0, 0, 0, time.UTC)
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

var logger = log.Default() // не делать глобальным объектом а прокинуть как зависимость

func LogErrorf(format string, v ...any) { // LogError(err error)
	msg := fmt.Sprintf(format, v...)
	logger.Printf("[Error]: %s\n", msg)
}

func LogInfo(format string, v ...any) { // нейминг LogInfof
	msg := fmt.Sprintf(format, v...)
	logger.Printf("[Info]: %s\n", msg)
}
