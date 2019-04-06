package main

import (
	"github.com/yanzay/tbot"
	"log"
	"miniature-parakeet/commands"
	"miniature-parakeet/models"
	"strconv"
)

func averageHandler(m *tbot.Message) {
	company := m.Vars["company"]

	daysStr := m.Vars["day"]

	days, err := strconv.Atoi(daysStr)
	if err != nil {
		return
	}

	month := m.Vars["month"]

	model := models.AverageModel{
		CompanyName: company,
		Day:         days,
		Month:       month,
	}

	average := commands.Average{}
	result := average.Calc(model)
	m.Replyf("average=%.6f", result)
}

func main() {
	bot, err := tbot.NewServer("")
	if err != nil {
		log.Fatal(err)
	}
	bot.HandleFunc("/average {company} {month} {day}", averageHandler)
	log.Println(bot.ListenAndServe())
}
