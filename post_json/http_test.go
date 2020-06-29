package main

import (
	"fmt"
	"testing"
)

func TestHttp(t *testing.T) {
	jps := JsonPostSample{}
	//jps.SamplePost()
	url := "http://localhost:8998/textMessage/paymentInformation"
	body := `{
    "Head": {
        "BatchID": "2006230002", 
        "RequestDate": "2019-04-03 13:25:35", 
        "RequestSource": "P00", 
        "RequestCode": "400001", 
        "RequestDesc": "示例"
 
    }, 
    "Body": [
        {
        }
    ]
}  `
	fmt.Println(body)
	jps.MyPost(url, body)
}
