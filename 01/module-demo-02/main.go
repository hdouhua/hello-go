package main

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Println("hello module, go module mode")
	logrus.Println(uuid.NewString())
}
