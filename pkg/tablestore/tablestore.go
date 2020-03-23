package tablestore

import (
	"os"

	"fmt"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/golang/glog"
	"github.com/knative-sample/knative-rocketmq/pkg/orderservice"
	uuid "github.com/satori/go.uuid"
)

type TableClient struct {
	tableName string
	client    *tablestore.TableStoreClient
}

func InitClient() *TableClient {
	endpoint := os.Getenv("OTS_TEST_ENDPOINT")
	tableName := os.Getenv("TABLE_NAME")
	instanceName := os.Getenv("OTS_TEST_INSTANCENAME")
	accessKeyId := os.Getenv("OTS_TEST_KEYID")
	accessKeySecret := os.Getenv("OTS_TEST_SECRET")
	client := tablestore.NewClient(endpoint, instanceName, accessKeyId, accessKeySecret)
	return &TableClient{
		tableName: tableName,
		client:    client,
	}
}

func (tableClient *TableClient) Store(order *orderservice.OrderInfo) error {
	if order == nil {
		return fmt.Errorf("order is nil")
	}
	putRowRequest := new(tablestore.PutRowRequest)
	putRowChange := new(tablestore.PutRowChange)
	putRowChange.TableName = tableClient.tableName
	putPk := &tablestore.PrimaryKey{}
	putPk.AddPrimaryKeyColumn("orderId", order.OrderId)
	putRowChange.PrimaryKey = putPk
	uid, _ := uuid.NewV4()
	putRowChange.AddColumn("id", uid.String())
	putRowChange.AddColumn("orderStatus", order.OrderStatus)
	putRowChange.AddColumn("userPhoneNo", order.UserPhoneNo)
	putRowChange.AddColumn("prodId", order.ProdId)
	putRowChange.AddColumn("prodName", order.ProdName)
	putRowChange.AddColumn("chargeMoney", order.ChargeMoney)
	putRowChange.AddColumn("chargeTime", order.ChargeTime)
	putRowChange.AddColumn("finishTime", order.FinishTime)
	putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	putRowRequest.PutRowChange = putRowChange
	_, err := tableClient.client.PutRow(putRowRequest)
	if err != nil {
		glog.Errorf("PutRow failed with error: %s", err.Error())
		return err
	}

	return nil
}
