package webserver

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
// )

// func Start(config *BlogConf) {
// 	fmt.Println("[i] Starting Local WebServer...")
// 	if config == nil {
// 		config = DefaultConf()
// 	}

// 	setRoutesWithConfig := func(w http.ResponseWriter, r *http.Request) {
// 		SetRoutes(w, r, *config)
// 	}

// 	funcframework.RegisterHTTPFunction("/", setRoutesWithConfig)

// 	if err := funcframework.Start(config.port); err != nil {
// 		log.Fatalf("funcframework.Start: %v\n", err)
// 	}
// }
