package verification

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/service/visual"
	"github.com/volcengine/volc-sdk-golang/service/visual/model"
)

func convert_photov2() {
	testAk := "ak"
	testSk := "sk"

	visual.DefaultInstance.Client.SetAccessKey(testAk)
	visual.DefaultInstance.Client.SetSecretKey(testSk)
	//visual.DefaultInstance.SetRegion("region")
	//visual.DefaultInstance.SetHost("host")

	//请求入参
	reqBody := &model.ConvertPhotoV2Request{
		ReqKey:           "lens_opr", // 固定值
		BinaryDataBase64: []string{"image_base64"},
		//ImageUrls:        []string{"https://xxxx"},
		IfColor: 1,
	}

	resp, status, err := visual.DefaultInstance.ConvertPhotoV2(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
