package main

//
//import (
//	"github.com/SwanHtetAungPhyo/swifcode/app/internal/pkg/log"
//	"github.com/SwanHtetAungPhyo/swifcode/cmd"
//	"os"
//	"os/signal"
//	"syscall"
//)
//
//func main() {
//	port := LoadEnv("PORT", "8080")
//	log.Init()
//	log := log.GetLogger()
//	log.Infof("PORT: %s", port)
//
//	log.Info("Starting server...")
//	go func() {
//		cmd.Start(port, log)
//	}()
//
//	osChan := make(chan os.Signal, 2)
//	signal.Notify(osChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
//	<-osChan
//}
//
//func LoadEnv(parameter, defaultValue string) string {
//	var value string
//	value = os.Getenv(parameter)
//	if value == "" {
//		value = defaultValue
//	}
//	return value
//}
//
// Configuration
