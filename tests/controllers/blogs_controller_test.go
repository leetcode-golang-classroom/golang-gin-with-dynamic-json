package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/leetcode-golang-classroom/golang-gin-with-dynamic-json/internal/config"
	"github.com/leetcode-golang-classroom/golang-gin-with-dynamic-json/internal/model"
	"github.com/leetcode-golang-classroom/golang-gin-with-dynamic-json/internal/service/blog"
	models_test "github.com/leetcode-golang-classroom/golang-gin-with-dynamic-json/tests/models"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Setup the MySQL test database
	models_test.SetupTestDB(config.AppConfig)
	code := m.Run()
	// cleanup
	models_test.TeardownTestDB()
	os.Exit(code)
}

type Response struct {
	Blogs      []model.Blog `json:"blogs"`
	Page       int          `json:"page"`
	PageSize   int          `json:"pageSize"`
	TotalPages int          `json:"totalPages"`
}

func TestBlogIndex(t *testing.T) {
	gin.SetMode(gin.TestMode)
	req := httptest.NewRequest(http.MethodGet, "/blogs?format=json", nil)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	// Seed the database with some test data
	blog1 := model.Blog{Title: "First Blog", Content: "Content of the first blog"}
	blog2 := model.Blog{Title: "Second Blog", Content: "Content of the second blog"}
	models_test.TRdb.Db.Create(&blog1)
	models_test.TRdb.Db.Create(&blog2)
	store := blog.NewBlogStore(models_test.TRdb)
	hdr := blog.NewHandler(store)
	hdr.BlogIndex(ctx)

	res := w.Result()
	defer res.Body.Close()
	// Check the status code
	assert.Equal(t, res.StatusCode, http.StatusOK, "API should return 200 stauts code")

	// Read data from the body and parse the JSON
	var result Response
	err := json.NewDecoder(res.Body).Decode(&result)
	assert.NoError(t, err)
	// Check the length of the blogs array
	assert.Len(t, result.Blogs, 2)
	// Check content
	assert.Equal(t, result.Blogs[0].Title, blog2.Title)
	assert.Equal(t, result.Blogs[0].Content, blog2.Content)
}

func TestBlogIndexWithEmptyTable(t *testing.T) {
	gin.SetMode(gin.TestMode)
	req := httptest.NewRequest(http.MethodGet, "/blogs?format=json", nil)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	models_test.TRdb.Db.Exec("DELETE * FROM blogs;")

	store := blog.NewBlogStore(models_test.TRdb)
	hdr := blog.NewHandler(store)
	hdr.BlogIndex(ctx)

	res := w.Result()
	defer res.Body.Close()
	// Check the status code
	assert.Equal(t, res.StatusCode, http.StatusOK, "API should return 200 stauts code")

	// Read data from the body and parse the JSON
	var result Response
	err := json.NewDecoder(res.Body).Decode(&result)
	assert.NoError(t, err)
	// Check the length of the blogs array
	assert.Len(t, result.Blogs, 0)
}
