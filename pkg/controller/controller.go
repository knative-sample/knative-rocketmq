package controller

import (
	"github.com/golang/glog"
	"github.com/knative-sample/knative-rocketmq/pkg/orderservice"
	"github.com/knative-sample/knative-rocketmq/pkg/tablestore"
)

func StoreOrderService(orderInfo *orderservice.OrderInfo) {
	client := tablestore.InitClient()
	err := client.Store(orderInfo)
	if err != nil {
		glog.Errorf("Order Store error: %s", err.Error())
		return
	}
}
