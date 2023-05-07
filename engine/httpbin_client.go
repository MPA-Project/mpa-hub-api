package engine

import (
	"context"

	"myponyasia.com/hub-api/dto"
)

type HttpBinClient interface {
	PostMethod(ctx context.Context, requestBody *dto.HttpBinDTO, response *map[string]interface{})
}
