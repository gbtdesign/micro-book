package service

import (
	"context"
	"golang/webook/internal/service/sms/tencent"
	"os"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

func TestTencent(t *testing.T) {
	secretId, ok := os.LookupEnv("SMS_SECRET_ID")
	if !ok {
	}
	secretKey, ok := os.LookupEnv("SMS_SECRET_KEY")
	tencentSms, err := sms.NewClient(common.NewCredential(secretId, secretKey),
		"ap-nanjing",
		profile.NewClientProfile())
	if err != nil {
	}

	s := tencent.NewService(tencentSms, "1400842696", "妙影科技")
	retrySms := NewRetrySmsService(s)
	err = retrySms.Send(context.Background(), "aaaa", []string{"ccccc"}, "10086")
}
