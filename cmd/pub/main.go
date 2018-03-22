//|------------------------------------------------------------------
//| Ali LMQ MQTT
//| All rights reserved: By cellgo.cn CopyRight
//| Author:Tommy.Jin Dtime:2018-3-20
//|-------------------------------------------------------------------

package main

import (
	"fmt"
	"sign"
	"strings"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	topic := "hotlife-micro-service/notice/"                     //ali-topic
	broker := "tcp://mqtt-cn-xxxxxxxxxxx.mqtt.aliyuncs.com:1883" //ali-host
	password := "5WbxxxxxxxxxxxxxxxHefRf9wIOIzS"                 //ali-password
	user := "LTAIxxxxxxxxqupe"                                   //ali-username
	clientid := "GID_ERVICE@@@ClientID_0001"                     //ali-clientid :GID_ORDER_ERVICE@@@ClientID_0001
	cleansess := false                                           //ali-cleansess-session:true关闭 false开启
	qos := 1                                                     //qos action
	num := 10                                                    //publish num
	payload := "message-test"                                    //publish value
	KeepAlive := 90 * time.Second                                //heartbeat
	AutoReconnect := true                                        //retry
	store := ":memory:"                                          //store action

	fmt.Printf("Sample Info:\n")
	fmt.Printf("\tbroker:    %s\n", broker)
	fmt.Printf("\tclientid:  %s\n", clientid)
	fmt.Printf("\tuser:      %s\n", user)
	fmt.Printf("\tpassword:  %s\n", password)
	fmt.Printf("\ttopic:     %s\n", topic)
	fmt.Printf("\tmessage:   %s\n", payload)
	fmt.Printf("\tqos:       %d\n", qos)
	fmt.Printf("\tcleansess: %v\n", cleansess)
	fmt.Printf("\tnum:       %d\n", num)
	fmt.Printf("\tKeepAlive:     %s\n", KeepAlive)
	fmt.Printf("\tAutoReconnect:     %s\n", AutoReconnect)
	fmt.Printf("\tstore:     %s\n", store)

	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientid)
	opts.SetUsername(user)
	opts.SetPassword(sign.SublishSignature(strings.Split(clientid, "@@@")[0], password))
	opts.SetCleanSession(cleansess)
	opts.SetKeepAlive(KeepAlive)
	opts.SetAutoReconnect(AutoReconnect)
	opts.SetStore(MQTT.NewFileStore(store))

	client := MQTT.NewClient(opts)
	var token MQTT.Token
	if token = client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Println("\tSample Publisher Started")
	for i := 0; i < num; i++ {
		fmt.Println("\t---- doing publish ----")
		token := client.Publish(topic, byte(qos), false, payload)
		token.Wait()
	}

	client.Disconnect(250)
	fmt.Println("\tSample Publisher Disconnected")

}
