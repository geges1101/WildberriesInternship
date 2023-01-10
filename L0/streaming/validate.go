package streaming

import (
	"encoding/json"
	"reflect"
	"strings"
)

type Item struct {
	ChrtId      uint   `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Sale        uint8  `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  uint8  `json:"total_price"`
	NmId        uint   `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      uint16 `json:"status"`
}

type Order struct {
	OrderUID    string `json:"order_uid"`
	TrackNumber string `json:"track_number"`
	Entry       string `json:"entry"`
	Delivery    struct {
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Zip     string `json:"zip"`
		City    string `json:"city"`
		Address string `json:"address"`
		Region  string `json:"region"`
		Email   string `json:"email"`
	} `json:"delivery"`
	Payment struct {
		Transaction  string `json:"transaction"`
		RequestId    string `json:"request_id"`
		Currency     int    `json:"currency"`
		Provider     string `json:"provider"`
		Amount       int    `json:"amount"`
		PaymentDt    int    `json:"payment_dt"`
		Bank         string `json:"bank"`
		DeliveryCost int    `json:"delivery_cost"`
		GoodsTotal   uint   `json:"goods_total"`
		CustomFee    uint   `json:"custom_fee"`
	} `json:"payment"`
	Items             []Item
	Locale            string `json:"locale"`
	InternalSignature string `json:"internal_signature"`
	CustomerId        string `json:"customer_id"`
	DeliveryService   string `json:"delivery_service"`
	ShardKey          string `json:"shardkey"`
	SmId              uint8  `json:"sm_id"`
	DateCreated       string `json:"date_created"`
	OofShard          string `json:"oof_shard"`
}

func CompareJSONToStruct(bytes []byte, empty interface{}) bool {
	var mapped map[string]interface{}

	if err := json.Unmarshal(bytes, &mapped); err != nil {
		return false
	}

	emptyValue := reflect.ValueOf(empty).Type()

	// check if number of fields is the same
	if len(mapped) != emptyValue.NumField() {
		return false
	}

	// check if field names are the same
	for key := range mapped {
		if field, found := emptyValue.FieldByName(key); found {
			if !strings.EqualFold(key, strings.Split(field.Tag.Get("json"), ",")[0]) {
				return false
			}
		}
	}

	return true
}
