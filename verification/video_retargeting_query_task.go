package verification

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/volcengine/volc-sdk-golang/service/visual"
)

func video_retargeting_query_task() {
	testAk := "ak"
	testSk := "sk"

	visual.DefaultInstance.Client.SetAccessKey(testAk)
	visual.DefaultInstance.Client.SetSecretKey(testSk)
	//visual.DefaultInstance.SetRegion("region")
	//visual.DefaultInstance.SetHost("host")

	query := url.Values{}
	query.Add("task_id", "")

	resp, status, err := visual.DefaultInstance.VideoRetargetingQueryTask(query)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
