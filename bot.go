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

	daysStr := m.Vars["days"]

	days, err := strconv.Atoi(daysStr)
	if err != nil {
		return
	}
	model := models.AverageModel{
		CompanyName: company,
		OffsetDay:   days,
	}

	average := commands.Average{}
	average.Calc(model)

	m.Reply("company=" + company + "days" + string(days))
}

func main() {
	bot, err := tbot.NewServer("")
	if err != nil {
		log.Fatal(err)
	}
	bot.HandleFunc("/average {company} {days}", averageHandler)
	bot.ListenAndServe()
}
