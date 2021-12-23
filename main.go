package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type transaccion struct {
	Id       int       `json:"id"`
	Codigo   string    `json:"codigo"`
	Moneda   string    `json:"moneda"`
	Monto    float64   `json:"monto"`
	Emisor   string    `json:"emisor"`
	Receptor string    `json:"receptor"`
	Fecha    time.Time `json:"fecha"`
}

func QueryID(id int) (transaction transaccion) {
	transacciones_data, err := os.ReadFile("./transacciones.json")
	if err != nil {
		fmt.Println("No se pudo leer el archivo transacciones.json")
	}
	var transactions []transaccion
	err = json.Unmarshal(transacciones_data, &transactions)
	if err != nil {
		fmt.Println("No se pudo parsear los datos")
	}
	for _, trans := range transactions {
		if trans.Id == id {
			transaction = trans
		}
	}
	return
}
func GetAll(ctx *gin.Context) {
	transacciones_data, err := os.ReadFile("./transacciones.json")
	if err != nil {
		fmt.Println("No se pudo leer el archivo transacciones.json")
	}
	var transacciones transaccion
	err = json.Unmarshal(transacciones_data, &transacciones)
	if err != nil {
		fmt.Println("No se pudo parsear los datos")
	}
	ctx.JSON(200, transacciones)
}

func ByID(ctx *gin.Context) {
	id := ctx.Param("id")
	id_transformed, _ := strconv.Atoi(id)
	transaction := QueryID(id_transformed)
	if transaction.Id == 0 {
		ctx.String(http.StatusNotFound, "transacción no registrada")
	} else {
		ctx.JSON(200, transaction)
	}
}
func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hola antonio",
		})
	})
	transaction_routes := router.Group("/transacciones")
	{
		transaction_routes.GET("/", GetAll)
		transaction_routes.GET("/:id", ByID)
	}

	router.Run()
}
