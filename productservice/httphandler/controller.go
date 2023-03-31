package httphandler

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"go-micro.dev/v4/logger"
	"io/ioutil"
	"net/http"
	db "productservice/db/sqlc"
	"productservice/utils"
	"strconv"
	"time"
)

type UploadProductImageRequest struct {
	ProductID int64  `form:"productId"`
	Order     int64  `form:"order"`
	Creator   string `form:"creator"`
}

type UploadController struct {
	BaseController
}

type JsonReturn struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data"` //Data字段需要设置为interface类型以便接收任意数据
	//json标签意义是定义此结构体解析为json或序列化输出json时value字段对应的key值,如不想此字段被解析可将标签设为`json:"-"`
}

func (c *BaseController) ApiJsonReturn(msg string, code int, data interface{}) {
	var JsonReturn JsonReturn
	JsonReturn.Msg = msg
	JsonReturn.Code = code
	JsonReturn.Data = data
	c.Data["json"] = JsonReturn //将结构体数组根据tag解析为json
	c.ServeJSON()               //对json进行序列化输出
	c.StopRun()                 //终止执行逻辑
}

type BaseController struct {
	web.Controller
}

type DownloadProductImageController struct {
	BaseController
}

type DownloadProductImageRequest struct {
	ProductID int64 `form:"productId"`
}

func (this *DownloadProductImageController) Post() {
	var req DownloadProductImageRequest
	err := this.ParseForm(&req)
	fmt.Println(err)
	imgs, err := db.DBStore.GetProductInfoImg(context.Background(), req.ProductID)
	if len(imgs) == 0 || err != nil {
		this.ApiJsonReturn("这个商品不存在图片", http.StatusBadRequest, "")
		this.Abort(strconv.Itoa(http.StatusBadRequest))
		return
	}
	m := make(map[string]interface{}, 4)
	m["productID"] = req.ProductID
	m["pictures"] = make([][]byte, 0)
	for _, val := range imgs {
		fileBytes, err := ioutil.ReadFile("./img/" + val.Source.String)
		if err != nil {
			continue
		}
		m["pictures"] = append((m["pictures"]).([][]byte), fileBytes)
	}
	result, err := json.MarshalIndent(m, "", "    ")
	this.ApiJsonReturn("商品图片", http.StatusOK, result)
}

func (this *UploadController) Post() {
	var req UploadProductImageRequest
	this.ParseForm(&req)
	f, _, err := this.GetFile("product.png")

	if f != nil {
		defer f.Close()
	}
	fmt.Println(req)
	if err != nil {
		logger.Errorf("上传图片出错,错误是：%s", err.Error())
		return
	}
	fileName := utils.GetName(req.ProductID, req.Order)
	err = this.SaveToFile("product.png", "./img/"+fileName)

	if err != nil {
		this.Ctx.WriteString(fmt.Errorf("出错信息：%s", err).Error())
		this.Abort(strconv.Itoa(http.StatusInternalServerError))
	}
	_, err = db.DBStore.InsertProductImage(context.Background(), db.InsertProductImageParams{
		Title:     sql.NullString{},
		ProductID: req.ProductID,
		Source: sql.NullString{
			String: fileName,
			Valid:  true,
		},
		Thumbnail: sql.NullString{},
		CreateBy: sql.NullString{
			String: req.Creator,
			Valid:  true,
		},
		CreationDate: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		Orders: sql.NullInt32{
			Int32: int32(req.Order),
			Valid: true,
		},
	})

	if err != nil {
		this.Ctx.WriteString(fmt.Errorf("出错信息：%s", err).Error())
		this.Abort(strconv.Itoa(http.StatusInternalServerError))
	}
	this.ApiJsonReturn("操作成功", http.StatusOK, "")
}
