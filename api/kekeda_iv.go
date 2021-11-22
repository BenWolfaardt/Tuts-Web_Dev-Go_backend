package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Hito struct {
	URI   string
	Title string
	date  time.Time
}

type Response struct {
	Msg    string `json:"text"`
	ChatID int64  `json:"chat_id"`
	Method string `json:"method"`
}

func format_date(diff time.Duration) string {
	// Función para convertir una diferencia de fechas
	// a una cadena del tipo dias horas minutos segundos

	// Definimos constantes para partir la diferencia
	const Decisecond = 100 * time.Millisecond
	const Day = 24 * time.Hour

	// Extraemos la cantidad de dias y las quitamos de la diferencia
	d := diff / Day
	diff = diff % Day
	// Extraemos la cantidad de horas
	h := diff / time.Hour
	diff = diff % time.Hour
	// Extraemos la cantidad de minutos
	m := diff / time.Minute
	diff = diff % time.Minute
	// Extraemos la cantidad de secundos
	s := diff / time.Second
	diff = diff % time.Second
	// Nos quedamos con las partes de segundo
	f := diff / Decisecond
	return fmt.Sprintf("%dd %dh %dm %d.%ds", d, h, m, s, f)

}

var hitos = []Hito{
	Hito{
		URI:   "0.Repositorio",
		Title: "Datos básicos y repo",
		date:  time.Date(2021, time.September, 17, 14, 30, 0, 0, time.UTC),
	},
	Hito{
		URI:   "1.Infraestructura",
		Title: "Historias de usuario",
		date:  time.Date(2021, time.September, 24, 14, 30, 0, 0, time.UTC),
	},
	Hito{
		URI:   "2.Entidad",
		Title: "Una entidad",
		date:  time.Date(2021, time.October, 1, 14, 30, 0, 0, time.UTC),
	},
	Hito{
		URI:   "3.Tests",
		Title: "Creando tests",
		date:  time.Date(2021, time.October, 8, 14, 30, 0, 0, time.UTC),
	},
	Hito{
		URI:   "4.CI",
		Title: "Integración continua",
		date:  time.Date(2020, time.November, 6, 23, 59, 0, 0, time.UTC),
	},
	Hito{
		URI:   "5.Serverless",
		Title: "Trabajando con funciones serverless",
		date:  time.Date(2020, time.November, 24, 23, 59, 0, 0, time.UTC),
	},
	Hito{
		URI:   "6.Microservicio",
		Title: "Diseñando un microservicio",
		date:  time.Date(2020, time.December, 11, 23, 59, 0, 0, time.UTC),
	},
	Hito{
		URI:   "7.PaaS",
		Title: "Desplegando en un PaaS",
		date:  time.Date(2021, time.January, 14, 23, 59, 0, 0, time.UTC),
	},
}

func Handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)

	var update tgbotapi.Update

	err := json.Unmarshal(body, &update)
	if err != nil {
		log.Fatal("Error en el update =>", err)
	}
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	currentTime := time.Now()

	var next int
	var queda time.Duration

	for indice, hito := range hitos {
		if hito.date.After(currentTime) {
			next = indice
			queda = hito.date.Sub(currentTime)
		}
	}

	if update.Message.IsCommand() {
		text := ""

		if next == 0 {
			text = "No upcoming deliveries"
		} else {

			switch update.Message.Command() {
			case "upcoming":
				text = queda.String()
			case "details":
				text = fmt.Sprintf(
					"=> Next milestone %s\n🔗 https://jj.github.io/IV/documentos/proyecto/%s\n📅 %s",
					hitos[next].Title,
					hitos[next].URI,
					hitos[next].date.String(),
				)
			default:
				text = "Use /upcoming for what's left for the next milestone, /details for more detail"
			}
		}

		data := Response{
			Msg:    text,
			Method: "sendMessage",
			ChatID: update.Message.Chat.ID,
		}
		msg, _ := json.Marshal(data)

		log.Printf("Response %s", string(msg))

		w.Header().Add("Content-Type", "application/json")

		fmt.Fprintf(w, string(msg))
	}
}
