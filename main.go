// This file is auto-generated, don't edit it. Thanks.
package main

import (
	"net/http"
	"os"
	"strconv"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dm20151123 "github.com/alibabacloud-go/dm-20151123/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gin-gonic/gin"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateSmsClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	// _result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

type SmsParam struct {
	Access_key_id     string `json:"access_key_id" binding:"required"`
	Access_key_secret string `json:"access_key_secret" binding:"required"`
	Phone_number      string `json:"phone_number" binding:"required"`
	Sign_name         string `json:"sign_name" binding:"required"`
	Template_code     string `json:"template_code" binding:"required"`
	Template_param    string `json:"template_param" binding:"required"`
}

type EmailParam struct {
	AccessKeyId     string `json:"access_key_id" binding:"required"`
	AccessKeySecret string `json:"access_key_secret" binding:"required"`
	AccountName     string `json:"account_name" binding:"required"`
	ToAddress       string `json:"to_address" binding:"required"`
	Subject         string `json:"subject" binding:"required"`
	HtmlBody        string `json:"html_body" binding:"required"`
	FromAlias       string `json:"from_alias"`
	ReplyAddress    string `json:"reply_address"`
}

func CreateEmailClient(accessKeyId *string, accessKeySecret *string) (_result *dm20151123.Client, _err error) {
	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dm.aliyuncs.com")
	// _result = &dm20151123.Client{}
	_result, _err = dm20151123.NewClient(config)
	return _result, _err
}

func main() {
	var port int = 80
	if len(os.Args) > 1 {
		v, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic("端口格式不正确")
		}
		port = v
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/sms", func(ctx *gin.Context) {
		var sms SmsParam
		if err := ctx.BindJSON(&sms); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		client, _err := CreateSmsClient(&sms.Access_key_id, &sms.Access_key_secret)
		if _err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": _err.Error()})
			return
		}

		sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
			PhoneNumbers:  tea.String(sms.Phone_number),
			SignName:      tea.String(sms.Sign_name),
			TemplateCode:  tea.String(sms.Template_code),
			TemplateParam: tea.String(sms.Template_param),
		}
		runtime := &util.RuntimeOptions{}
		resp, err1 := client.SendSmsWithOptions(sendSmsRequest, runtime)
		if err1 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
			return
		}

		ctx.JSON(http.StatusOK, tea.ToMap(resp))
	})

	r.POST("/email", func(ctx *gin.Context) {
		var email EmailParam
		if err := ctx.BindJSON(&email); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		client, _err := CreateEmailClient(&email.AccessKeyId, &email.AccessKeySecret)
		if _err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": _err.Error()})
			return
		}

		singleSendMailRequest := &dm20151123.SingleSendMailRequest{
			AccountName:    &email.AccountName,
			AddressType:    tea.Int32(1),
			ReplyToAddress: tea.Bool(true),
			ToAddress:      tea.String(email.ToAddress),
			Subject:        tea.String(email.Subject),
			HtmlBody:       tea.String(email.HtmlBody),
			FromAlias:      tea.String(email.FromAlias),
			ReplyAddress:   tea.String(email.ReplyAddress),
		}
		runtime := &util.RuntimeOptions{}

		resp, _err := client.SingleSendMailWithOptions(singleSendMailRequest, runtime)
		if _err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": _err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, tea.ToMap(resp))
	})
	r.Run("0.0.0.0:" + strconv.Itoa(port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
