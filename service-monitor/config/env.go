package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func ReadEnv() {

	prodErr := godotenv.Load("robot-fleet-monitoring/service-monitor/.env")

	if prodErr != nil {

		localErr := godotenv.Load(".env")
		if localErr != nil {
			logrus.Error("Error loading env local file")
		}

	}
}
