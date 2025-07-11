package verification

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/service/visual"
)

func image_scorev2() {
	testAk := "ak"
	testSk := "sk"

	visual.DefaultInstance.Client.SetAccessKey(testAk)
	visual.DefaultInstance.Client.SetSecretKey(testSk)
	//visual.DefaultInstance.SetRegion("region")
	//visual.DefaultInstance.SetHost("host")

	//图片base64编码。此算法可选输入1张图片或多张图片
	imageBase64List := []string{""}

	resp, status, err := visual.DefaultInstance.ImageScoreV2(imageBase64List)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
