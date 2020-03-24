package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"context"
	"flag"

	cloudevents "github.com/cloudevents/sdk-go"

	"encoding/base64"

	"github.com/knative-sample/knative-rocketmq/pkg/controller"
	"github.com/knative-sample/knative-rocketmq/pkg/kncloudevents"
	"github.com/knative-sample/knative-rocketmq/pkg/orderservice"
	"github.com/knative-sample/knative-rocketmq/pkg/utils/logs"
)

func receive(ctx context.Context, event cloudevents.Event) {
	fmt.Printf(event.String())
	payload := &orderservice.Data{}
	if event.Data == nil {
		log.Printf("receive cloudevents.Event\n  Type:%s\n  Data is empty", event.Context.GetType())
		return
	}

	data, ok := event.Data.([]byte)
	if !ok {
		var err error
		data, err = json.Marshal(event.Data)
		if err != nil {
			data = []byte(err.Error())
		}
	}
	err := json.Unmarshal(data, payload)
	if err != nil {
		log.Printf("receive %s, Unmarshal data error: %s", data, err.Error())
		return
	}
	order := &orderservice.OrderInfo{}
	bt, err := base64.StdEncoding.DecodeString(payload.Body)
	if err != nil {
		log.Printf("receive %s, DecodeString Body error: %s", payload.Body, err.Error())
		return
	}
	err = json.Unmarshal(bt, order)
	if err != nil {
		log.Printf("receive %s, Unmarshal Body error: %s", bt, err.Error())
		return
	}
	controller.StoreOrderService(order)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "ok")
}

const tmp = `{"orderId":"123214342","orderStatus":"completed","userPhoneNo":"152122131323","prodId":"2141412","prodName":"test","chargeMoney":"30.0","chargeTime":"1584932320","finishTime":"1584932320"}`

func main() {
	flag.Parse()
	logs.InitLogs()
	defer logs.FlushLogs()
	go func() {
		http.HandleFunc("/health", handler)
		port := os.Getenv("PORT")
		if port == "" {
			port = "8022"
		}
		http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	}()

	c, err := kncloudevents.NewDefaultClient()
	if err != nil {
		log.Fatal("Failed to create client, ", err)
	}
	log.Fatal(c.StartReceiver(context.Background(), receive))

}
