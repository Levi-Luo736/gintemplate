package verification

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/volcengine/volc-sdk-golang/service/visual"
)

func video_inpaint_submit_task() {
	testAk := "ak"
	testSk := "sk"

	visual.DefaultInstance.Client.SetAccessKey(testAk)
	visual.DefaultInstance.Client.SetSecretKey(testSk)
	// visual.DefaultInstance.SetRegion("region")
	// visual.DefaultInstance.SetHost("host")

	form := url.Values{}
	form.Add("video_url", "")
	form.Add("method", "")

	resp, status, err := visual.DefaultInstance.VideoInpaintSubmitTask(form)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
