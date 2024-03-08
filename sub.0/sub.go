package main

import (
	"fmt"
	"os"
	"os/signal"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var broker = "tcp://localhost:1883"
var topic = "test/topic"

func onMessageReceived(client MQTT.Client, message MQTT.Message) {
	fmt.Printf("Mensagem recebida no tópico: %s\n", message.Topic())
	fmt.Printf("Payload: %s\n", message.Payload())
}

func handleDisconnect(client MQTT.Client, err error) {
	fmt.Printf("Erro ao desconectar do broker MQTT: %v\n", err)
}

func main() {
	opts := MQTT.NewClientOptions().AddBroker(broker)
	opts.SetClientID("go-mqtt-client")
	opts.SetDefaultPublishHandler(onMessageReceived)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Verifique se a conexão foi estabelecida
	if !client.IsConnected() {
		fmt.Println("Falha ao conectar ao broker MQTT")
		return
	}

	fmt.Printf("Conectado ao broker MQTT em %s\n", broker)

	// Inscreva-se no tópico
	if token := client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Canal para sinalizar a conclusão da recepção de mensagens
	done := make(chan struct{})

	// Captura de sinal para lidar com o encerramento
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	// Loop para receber mensagens continuamente
	go func() {
		for {
			select {
			case <-done:
				return
			default:
			}
		}
	}()

	// Aguardar um sinal de interrupção para terminar a recepção de mensagens
	<-signalChannel
	close(done)
	fmt.Println("Programa encerrado")
}
