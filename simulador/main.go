package main

import (
	"log"

	"github.com/codeedu/imersaofsfc2-simulator/infra/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {

	producer := kafka.NewKafkaProducer()
	kafka.Publish("ola", "readTest", producer)

	// route := route.Route{
	// 	ID:        "1",
	// 	ClienteID: "1",
	// }
	// route.LoadPositions()
	// stringjson, _ := route.ExportJsonPositions()
	// fmt.Println(stringjson[1])
}
