package user

import (
	"context"
	"fmt"
	pb "frontend/proto"
	"frontend/service"
	"net/http"
)

type ProductController struct {
	BaseController
}

type SearchProductRequest struct {
	Query string `form:"query"`
	Page  int64  `form:"page"`
}

func (p *ProductController) SearchProduct() {
	var req SearchProductRequest
	if err := p.ParseForm(&req); err != nil {
		p.ApiJsonReturn(fmt.Sprintf("难以解析,出错信息为%s", err.Error()), http.StatusBadRequest, nil)
	}
	rsp, err := service.Svc.ProductService.SearchProducts(context.Background(), &pb.SearchProductsRequest{
		Query: req.Query,
		Page:  int32(req.Page),
	})
	if err != nil {
		p.ApiJsonReturn(fmt.Sprintf("rpc服务出错,出错信息为%s", err.Error()), http.StatusInternalServerError, nil)
	}
	if rsp.Status == "500" {
		p.ApiJsonReturn(fmt.Sprintf("rpc服务出错,出错信息为%s", rsp.Description), http.StatusInternalServerError, nil)
	}
	m := map[string]interface{}{}
	m["productInfos"] = rsp.Results
	m["productImgs"] = rsp.Imgs

	p.ApiJsonReturn("搜索成功", http.StatusOK, m)
}
