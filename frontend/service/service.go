package service

import pb "frontend/proto"

type FrontendServer struct {
	EmailService     pb.EmailService
	UserService      pb.UserService
	OrderService     pb.OrderService
	InventoryService pb.InventoryService
	ProductService   pb.ProductCatalogService
	CartService      pb.CartService
}

var Svc *FrontendServer
