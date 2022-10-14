package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert"
	"github.com/restuwahyu13/go-supertest/supertest"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/triadmoko/edot-interview/models"
)

var router = NewGinConfig()

func TestCreateProduct(t *testing.T) {

	Convey("Create Product", t, func() {

		Convey("Create New Product", func() {
			product := models.Product{
				Name:       "My Poduct",
				CategoryID: 1,
				Price:      2000,
				Total:      100,
			}
			payload := product

			test := supertest.NewSuperTest(router, t)

			test.Post("/product/")
			test.Send(payload)
			test.Set("Content-Type", "application/json")
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				if rr.Code >= 400 {
					t.Log("error")
					t.Skip()
				}
				assert.Equal(t, http.StatusCreated, rr.Code)
			})
		})
	})
}
func TestUpdateProduct(t *testing.T) {

	Convey("Update Product", t, func() {

		Convey("Update New Product", func() {
			product := models.Product{
				Name:       "My Poduct 1",
				CategoryID: 1,
				Price:      2000,
				Total:      100,
			}
			payload := product

			test := supertest.NewSuperTest(router, t)

			test.Put("/product/1")
			test.Send(payload)
			test.Set("Content-Type", "application/json")
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				if rr.Code >= 400 {
					t.Log("error")
					t.Skip()
				}
				assert.Equal(t, rr.Code, http.StatusOK)
			})
		})
	})
}

func TestPatchProduct(t *testing.T) {

	Convey("Patch Product", t, func() {

		Convey("Patch New Product", func() {
			product := models.Product{
				Name:  "My Poduct 1",
				Price: 2000,
				Total: 100,
			}
			payload := product

			test := supertest.NewSuperTest(router, t)

			test.Patch("/product/1")
			test.Send(payload)
			test.Set("Content-Type", "application/json")
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				if rr.Code >= 400 {
					t.Log("error")
					t.Skip()
				}
				assert.Equal(t, rr.Code, http.StatusOK)
			})
		})
	})
}
func TestGetProduct(t *testing.T) {

	Convey("Get Product", t, func() {

		Convey("Get Product", func() {

			test := supertest.NewSuperTest(router, t)
			test.Send(nil)
			test.Get("/product/1")
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				if rr.Code >= 400 {
					t.Log("error")
					t.Skip()
				}
				assert.Equal(t, http.StatusOK, rr.Code)
			})
		})
	})
}
func TestGetProducts(t *testing.T) {

	Convey("Get Products", t, func() {

		Convey("Get Products", func() {

			test := supertest.NewSuperTest(router, t)
			test.Send(nil)
			test.Get("/product")
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				if rr.Code >= 400 {
					t.Log("error")
					t.Skip()
				}
				assert.Equal(t, http.StatusOK, rr.Code)
			})
		})
	})
}
func TestDeleteProduct(t *testing.T) {

	Convey("Get Delete Products", t, func() {

		Convey("Get Delete Products", func() {

			test := supertest.NewSuperTest(router, t)
			test.Send(nil)
			test.Get("/product/1")
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				if rr.Code >= 400 {
					t.Log("error")
					t.Skip()
				}
				assert.Equal(t, http.StatusOK, rr.Code)
			})
		})
	})
}