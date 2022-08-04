package delivery

import (
	"fmt"
	"log"
	"net/http"
	"project3/group3/domain"
	_middleware "project3/group3/feature/common"
	_helper "project3/group3/helper"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type productHandler struct {
	productUseCase domain.ProductUseCase
}

func New(ps domain.ProductUseCase) domain.ProductHandler {
	return &productHandler{
		productUseCase: ps,
	}
}

func (ph *productHandler) InsertProductHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp InsertProductFormat
		idFromToken, _ := _middleware.ExtractData(c)
		if idFromToken != 1 && idFromToken != 2 {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("your account not admin"))
		}

		err := c.Bind(&tmp)
		// =================================================================================

		fileData, fileInfo, fileErr := c.Request().FormFile("product_image")

		// return err jika missing file
		if fileErr == http.ErrMissingFile || fileErr != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get file eee"))
		}

		// cek ekstension file upload
		extension, err_check_extension := _helper.CheckFileExtension(fileInfo.Filename)
		if err_check_extension != nil {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("file extension error"))
		}

		// check file size
		err_check_size := _helper.CheckFileSize(fileInfo.Size)
		if err_check_size != nil {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("file size error"))
		}

		// memberikan nama file
		fileName := time.Now().Format("2006-01-02 15:04:05") + "." + extension
		url, errUploadImg := _helper.UploadImageToS3(fileName, fileData)
		if errUploadImg != nil {
			fmt.Println(errUploadImg)
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to upload file"))
		}

		// =================================================================================
		tmp.UserID = idFromToken
		tmp.ProductImage = url
		if err != nil {
			log.Println("cannot parse data", err)
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data"))
		}

		dataProduct := tmp.ToModel()
		result, errCreate := ph.productUseCase.InsertProduct(dataProduct)
		if errCreate != nil {
			log.Println("ini data product : ", dataProduct)
			log.Println("err : ", errCreate)
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to post product"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", FromModel(result)))
	}
}

func (ph *productHandler) DeleteProductHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idproduct := c.Param("id")
		idFromToken, _ := _middleware.ExtractData(c)
		if idFromToken != 1 && idFromToken != 2 {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("your account not admin"))
		}
		id, _ := strconv.Atoi(idproduct)
		if id == 0 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("you dont have access"))
		}
		row, errDel := ph.productUseCase.DeleteProduct(id)

		if errDel != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to delete data product"))
		}

		if row != 1 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to delete data product"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
	}
}

func (ph *productHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit := c.QueryParam("limit")
		offset := c.QueryParam("offset")
		limitint, _ := strconv.Atoi(limit)
		offsetint, _ := strconv.Atoi(offset)
		res, err := ph.productUseCase.GetAllData(limitint, offsetint)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("canot to read all data"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", FromModelList(res)))

	}
}
func (ph *productHandler) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idProduct, _ := strconv.Atoi(id)
		res, err := ph.productUseCase.GetProductById(idProduct)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get detail roduct"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success ", FromModel(res)))
	}

}

func (ph *productHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idProduct, _ := strconv.Atoi(id)
		idFromToken, _ := _middleware.ExtractData(c)
		if idFromToken != 1 && idFromToken != 2 {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("your account not admin"))
		}
		product_name := c.FormValue("product_name")
		product_image := c.FormValue("product_image")
		stock := c.FormValue("stock")
		price := c.FormValue("price")
		stockint, _ := strconv.Atoi(stock)
		priceint, _ := strconv.Atoi(price)

		postReq := InsertProductFormat{
			ProductName:  product_name,
			ProductImage: product_image,
			Stock:        stockint,
			Price:        priceint,
		}

		dataPost := postReq.ToModel()
		row, errUpd := ph.productUseCase.UpdateData(dataPost, idProduct, idFromToken)
		if errUpd != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("you dont have access"))
		}
		if row == 0 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to update data"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
	}
}
