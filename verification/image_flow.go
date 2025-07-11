package verification

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/volcengine/volc-sdk-golang/service/visual"
)

func image_flow() {
	testAk := "ak"
	testSk := "sk"

	visual.DefaultInstance.Client.SetAccessKey(testAk)
	visual.DefaultInstance.Client.SetSecretKey(testSk)
	//visual.DefaultInstance.SetRegion("region")
	//visual.DefaultInstance.SetHost("host")

	form := url.Values{}
	form.Add("image_base64", "")
	form.Add("motion_ration", "")

	resp, status, err := visual.DefaultInstance.ImageFlow(form)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
