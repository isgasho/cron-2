package main

import (
	"log"

	"github.com/lnquy/cron"
)

const expr = "0 9 * * 1-5"

func main() {
	exprDesc := cron.NewDescriptor(
		cron.DayOfWeekStartsAtZero(),
		cron.Use24HourTimeFormat(),
		cron.Verbose(),
		cron.SetLogger(&log.Logger{}),
		cron.SetLocales(cron.Locale_da, cron.Locale_de, cron.Locale_en, cron.Locale_es),
	)

	desc, err := exprDesc.ToDescription(expr, cron.Locale_en)
	if err != nil {
		log.Panicf("failed to convert CRON expression to human readable description: %s", err)
	}
	log.Printf("Expression: %s\nDescription: %s", expr, desc)
}
