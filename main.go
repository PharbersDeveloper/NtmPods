package main

import (
	"fmt"
	"github.com/alfredyang1986/BmServiceDef/BmConfig"
	"github.com/alfredyang1986/BmServiceDef/BmPodsDefine"
	"os"
)

func main() {
	version := "v0"
	fmt.Println("NTM pods archi begins, version =", version)

	var pod = BmPodsDefine.Pod{Name: "new TMIST"}
	ntmHome := os.Getenv("NTM_HOME")
	pod.RegisterSerFromYAML(ntmHome + "/Resources/service-def.yaml")

	var bmRouter BmConfig.BmRouterConfig
	bmRouter.GenerateConfig()

	addr := bmRouter.Host + ":" + bmRouter.Port
	fmt.Println("Listening on ", addr)
	//api := api2go.NewAPIWithResolver(version, &BmApiResolver.RequestURL{Addr: addr})
	//pod.RegisterAllResource(api)

	//pod.RegisterAllFunctions(version, api)
	//pod.RegisterAllMiddleware(api)
	//
	//handler := api.Handler().(*httprouter.Router)
	//pod.RegisterPanicHandler(handler)
	//http.ListenAndServe(":"+bmRouter.Port, handler)
	//
	fmt.Println("NTM pods archi ends, version =", version)
}
