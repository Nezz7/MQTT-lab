package main

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	dt := time.Now()
	fmt.Print("[")
	fmt.Print(dt.Format("01-02-2006 15:04:05"))
	fmt.Print("] ")
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v\n", err)
}

func main() {
	var broker = "broker.emqx.io"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_clientMDMS")
	opts.SetUsername("mdms")
	opts.SetPassword("public")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	go sub(client, "/topic/consumption", 2)
	go sub(client, "/topic/production", 2)

	go publish(client, "/topic/price", 60, 1)
	go publish(client, "/topic/reductionVoltage", 1, 2)


	time.Sleep(time.Second * 60 * 15)
	client.Disconnect(250)
}

func publish(client mqtt.Client, topic string, freq int, qos byte) {
	num := 5
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("id %d", i)
		token := client.Publish(topic, qos, false, text)
		token.Wait()
		time.Sleep(time.Second * time.Duration(freq))
	}
}

func sub(client mqtt.Client, topic string, qos byte) {
	token := client.Subscribe(topic, qos, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s\n", topic)
}
