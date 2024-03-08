package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	// Configuração do cliente MQTT
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("Conexão Zona Leste - SP")
	client := MQTT.NewClient(opts)

	// Conecte ao broker
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	defer client.Disconnect(250)

	// Verifique se a conexão foi estabelecida
	if !client.IsConnected() {
		fmt.Println("Falha ao conectar ao broker MQTT")
		return
	}

	// Captura de sinal para lidar com o encerramento
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	// Loop para publicar mensagens continuamente
	go func() {
		for {
			zl1, zl2 := lerSensor()
			message := fmt.Sprintf("Leitura do sensor: PM2.5 %.2f (µg/m³), PM10 %.2f (µg/m³)", zl1, zl2)
			token := client.Publish("test/topic", 0, false, message)
			token.Wait()
			fmt.Printf("Publicado: %s\n", message)
			time.Sleep(5 * time.Second)
		}
	}()

	// Aguarde um sinal de interrupção
	<-signalChannel
	fmt.Println("Publicação encerrada")
}

// lerSensor simula a leitura do sensor
func lerSensor() (float64, float64) {
	// Simulação de leitura de sensor
	return SimularSPS30()
}
