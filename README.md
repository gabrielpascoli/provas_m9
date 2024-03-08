# provas_m9
provas do modulo 9


p1 

## Como Usar:

1. Certifique-se de ter o Go instalado em seu ambiente de desenvolvimento.
2. Instale a biblioteca Paho MQTT com o seguinte comando:
   ```bash
   go get github.com/eclipse/paho.mqtt.golang
   ```
3. Execute os programas individualmente a partir do terminal. Por exemplo, para executar o Publicador MQTT:
   ```bash
   go run main.go
   ```
   Ou para executar o Simulador de Leitura do Sensor:
   ```bash
   go run sensor_simulator.go
   ```
   E para o Assinante MQTT:
   ```bash
   go run mqtt_subscriber.go
   ```
