package service

import (
	"context"
	"errors"
	"time"

	"golang/webook/internal/service/sms"
)

type RetrySmsService struct {
	svc sms.Service // 包装一个短信服务
}

func NewRetrySmsService(svc sms.Service) *RetrySmsService {
	return &RetrySmsService{svc: svc}
}
func (r *RetrySmsService) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
	// 循环重试 3 次
	for i := 0; i < 3; i++ {
		err := r.svc.Send(ctx, tpl, args, numbers...)
		if err == nil {
			return nil
		}
		time.Sleep(time.Second)
	}
	return errors.New("重试失败")
}
