package wxlib

import (
	"github.com/xen0n/go-workwx"
	"net/http"
)

var TestConfig = map[string]map[string]interface{}{
	"201":{
		"Token":          "yourToken",
		"AppId":          "yourCorpID",
		"AgentId":        "yourAgentId",
		"Secret":         "yourSecret",
		"EncodingAESKey": "yourEncodingAesKey",
		"AppType":        1,
	},
}

var WorkwxClient *workwx.WorkwxApp

func init()  {
	agentID := int64(1000004)
	corpID	:= "wwf8191d9b61b6c18e"
	client := workwx.New(corpID)
	corpSecret:="1yAAA4qwC2iLrXI-78FUkxTkHnkecAQ0bTtD5XFMDqM"

	// there're advanced options
	_ = workwx.New(
		corpID,
		//workwx.WithQYAPIHost("http://localhost:8888"),
		workwx.WithHTTPClient(&http.Client{}),
	)

	// work with individual apps
	WorkwxClient  = client.WithApp(corpSecret, agentID)
	WorkwxClient.SpawnAccessTokenRefresher()
	//WorkwxClient = app
}