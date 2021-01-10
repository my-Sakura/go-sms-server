package Interface

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/my-Sakura/SMS/pkg/utils"
)

type Dispatcher interface {
	Send(mobile, code string) error
}

func (c *client) Send(mobile string) error {

	switch c.Name {
	case "esand":
		code := utils.GetCode(4)
		url := fmt.Sprintf("https://smssend.shumaidata.com/sms/send?mobile=%s&templateID=%s", mobile, code, c.TemplateCode)
		client := &http.Client{}

		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			return err
		}

		req.Header.Add("Authorization", "APPCODE"+" "+c.AppCode)
		resp, _ := client.Do(req)
		b, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("response is %v\n", string(b))

		return nil
	case "tianyan":
		code := utils.GetCode(6)
		url := fmt.Sprintf("https://smssend.shumaidata.com/sms/send?receive=%s&tag=%s&templateId=%s", mobile, code, c.TemplateCode)

		client := &http.Client{}

		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			return err
		}

		req.Header.Add("Authorization", "APPCODE"+" "+c.AppCode)
		resp, _ := client.Do(req)
		b, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("response is %v\n", string(b))

		return nil
	}

	return nil
}
