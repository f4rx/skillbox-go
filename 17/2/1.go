package main

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.WarnLevel)
}

func main (){
	log.Info("ololo")
	log.Error("пыщ-пыщ")

	log.WithFields(log.Fields{
		"custom_field": "моё кастомное сообщение",
	}).Warn("А это поле msg")
}