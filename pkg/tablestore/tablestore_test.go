package tablestore

import (
	"strings"
	"testing"

	"fmt"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/knative-sample/weather-store/pkg/utils"
	"github.com/knative-sample/weather-store/pkg/weather"
)

func FakeInitClient(tableName string) *TableClient {
	endpoint := "https://xx.cn-beijing.ots.aliyuncs.com" //实例访问地址
	instanceName := "knative-weather"                    // 实例名
	accessKeyId := "xx"                                  // AccessKey ID
	accessKeySecret := "xx"                              //Access Key Secret
	client := tablestore.NewClient(endpoint, instanceName, accessKeyId, accessKeySecret)
	return &TableClient{
		tableName: tableName,
		client:    client,
	}
}
func TestStore(t *testing.T) {
	c := FakeInitClient("weather")
	ct := weather.Cast{
		Date:         "2019-09-24",
		Week:         "2",
		Dayweather:   "晴",
		Nightweather: "晴",
		Daytemp:      "31",
		Nighttemp:    "15",
		Daywind:      "南",
		Nightwind:    "南",
		Daypower:     "≤3",
		Nightpower:   "≤3",
	}

	f := weather.Forecast{
		City:       "北京市",
		Adcode:     "110000",
		Reporttime: "2019-09-24 20:50:56",
		Province:   "北京",
		Casts:      []weather.Cast{ct},
	}
	err := c.Store(f)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestStoreCity(t *testing.T) {
	cityCodes := strings.Split(utils.CityInfo, "\n")
	client := FakeInitClient("city")
	pre := ""
	for _, cityCode := range cityCodes {
		if cityCode == "" {
			continue
		}
		items := strings.Split(cityCode, ",")
		ci := weather.CityInfo{IsCity: "false"}
		ci.Name = items[0]
		ci.Adcode = items[1]
		if len(items) == 3 {
			ci.Citycode = items[2]
			if ci.Citycode != "" && pre != ci.Citycode {
				ci.IsCity = "true"
			}
			pre = ci.Citycode
		}

		err := client.StoreCity(ci)
		if err != nil {
			fmt.Errorf("StoreCity: %s", err.Error())
			continue
		}
		fmt.Println("load: ", ci.Name)
	}
}
