package logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "ivan").Info("Hello World")

	logger.WithField("username", "ivan").
		WithField("name", "Ivan Adi Saputra").
		Info("Hello World")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "ivan",
		"name":     "Ivan Adi Saputra",
	}).Info("Hello World")
}