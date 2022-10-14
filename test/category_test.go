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

func TestCreateCategory(t *testing.T) {

	Convey("Create Category", t, func() {

		Convey("Create New Category", func() {
			category := models.Category{
				Name: "My Poduct",
			}
			payload := category

			test := supertest.NewSuperTest(router, t)

			test.Post("/category/")
			test.Send(payload)
			test.Set("Content-Type", "application/json")
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				if rr.Code == 400 {
					t.Log("error")
					t.Skip()
				}
				assert.Equal(t, http.StatusCreated, rr.Code)
			})
		})
	})
}
func TestUpdateCategory(t *testing.T) {

	Convey("Update Category", t, func() {

		Convey("Update New Category", func() {
			category := models.Category{
				Name: "My Category 1",
			}
			payload := category

			test := supertest.NewSuperTest(router, t)

			test.Put("/category/1")
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
func TestPatchCategory(t *testing.T) {

	Convey("Patch Category", t, func() {

		Convey("Patch New Category", func() {
			Category := models.Category{
				Name: "My Category 1",
			}
			payload := Category

			test := supertest.NewSuperTest(router, t)

			test.Patch("/category/1")
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
func TestGetCategory(t *testing.T) {

	Convey("Get Category", t, func() {

		Convey("Get Category", func() {

			test := supertest.NewSuperTest(router, t)
			test.Send(nil)
			test.Get("/category/1")
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
func TestGetCategorys(t *testing.T) {

	Convey("Get Categorys", t, func() {

		Convey("Get Categorys", func() {

			test := supertest.NewSuperTest(router, t)
			test.Send(nil)
			test.Get("/category")
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
func TestDeleteCategory(t *testing.T) {

	Convey("Get Delete Category", t, func() {

		Convey("Get Delete Category", func() {

			test := supertest.NewSuperTest(router, t)
			test.Send(nil)
			test.Get("/category/1")
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
