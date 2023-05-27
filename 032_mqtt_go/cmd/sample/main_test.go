package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
	"testing"
	"time"
)

func Test_NewClient_simple(t *testing.T) {
	ops := mqtt.NewClientOptions().SetClientID("foo").AddBroker("tcp://192.168.0.17:1883")
	client := mqtt.NewClient(ops)

	if client == nil {
		t.Fatalf("ops is nil")
	}

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("testtopic/#", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// Publish a message
	token := client.Publish("testtopic/1", 0, false, "Hello World")
	token.Wait()

	time.Sleep(6 * time.Second)

	// Unscribe
	if token := client.Unsubscribe("testtopic/#"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// Disconnect
	client.Disconnect(250)
	time.Sleep(1 * time.Second)

}
