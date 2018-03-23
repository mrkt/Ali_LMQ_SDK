//|------------------------------------------------------------------
//| Ali LMQ MQTT
//| Author:Tommy.Jin Dtime:2018-3-20
//|-------------------------------------------------------------------

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sign"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type Off struct {
	MaxPushNum int    `json:"maxPushNum"`
	PushOrder  string `json:"pushOrder"`
}

func main() {

	topic := "hotlife-micro-service/notice/"                     //ali-topic
	offline_topic := "$SYS/getOfflineMsg"                        //ali-offline-topic
	Offline_take := 100                                          //ali-offline-takenum
	Offline_sort := "DESC"                                       //ali-offline-orderby
	broker := "tcp://mqtt-cn-xxxxxxxxxx .mqtt.aliyuncs.com:1883" //ali-host
	password := "5WbXXXXXXXXXXXXXXXXXXXXXXXXXXX"                 //ali-password
	user := "LTAIXXXXXXXXXXXX"                                   //ali-username
	clientid := "GID_ERVICE@@@ClientID_0002"                     //ali-clientid :GID_ORDER_ERVICE@@@ClientID_0001
	cleansess := false                                           //ali-cleansess-session:true关闭 false开启
	qos := 1                                                     //qos action
	payload := "message-test"                                    //publish value
	store := ":memory:"

	fmt.Printf("Sample Info:\n")
	fmt.Printf("\tbroker:    %s\n", broker)
	fmt.Printf("\tclientid:  %s\n", clientid)
	fmt.Printf("\tuser:      %s\n", user)
	fmt.Printf("\tpassword:  %s\n", password)
	fmt.Printf("\ttopic:     %s\n", topic)
	fmt.Printf("\toffline_topic:     %s\n", offline_topic)
	fmt.Printf("\tOffline_take:     %s\n", Offline_take)
	fmt.Printf("\tOffline_sort:     %s\n", Offline_sort)
	fmt.Printf("\tmessage:   %s\n", payload)
	fmt.Printf("\tqos:       %d\n", qos)
	fmt.Printf("\tcleansess: %v\n", cleansess)
	fmt.Printf("\tstore:     %s\n", store)
	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientid)
	opts.SetUsername(user)
	opts.SetPassword(sign.PublishSignature(strings.Split(clientid, "@@@")[0], password))
	opts.SetCleanSession(cleansess)
	opts.SetKeepAlive(90)
	opts.SetAutoReconnect(true)
	if store != ":memory:" {
		opts.SetStore(MQTT.NewFileStore(store))
	}

	choke := make(chan [2]string)

	opts.SetDefaultPublishHandler(func(client MQTT.Client, msg MQTT.Message) {
		choke <- [2]string{msg.Topic(), string(msg.Payload())}
	})

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe(topic, byte(qos), nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	o := &Off{Offline_take, Offline_sort}
	offJson, _ := json.Marshal(o)
	if token := client.Publish(offline_topic, byte(0), false, offJson); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	fmt.Printf("\tGet Offline Msg %d Line", Offline_take)
	for {
		incoming := <-choke
		fmt.Printf("\tRECEIVED TOPIC: %s MESSAGE: %s\n", incoming[0], incoming[1])
	}

	client.Disconnect(250)
	fmt.Println("\tSample Subscriber Disconnected")
}
