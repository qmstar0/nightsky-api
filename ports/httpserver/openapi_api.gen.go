// Package httpserver provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package httpserver

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	BearerScopes = "bearer.Scopes"
)

// AdminGetPostListParams defines parameters for AdminGetPostList.
type AdminGetPostListParams struct {
	Page *int `form:"page,omitempty" json:"page,omitempty"`
}

// CreateCategoryJSONBody defines parameters for CreateCategory.
type CreateCategoryJSONBody struct {
	Desc string `json:"desc"`

	// Name 名称
	Name string `json:"name"`
}

// ModifyCategoryDescJSONBody defines parameters for ModifyCategoryDesc.
type ModifyCategoryDescJSONBody struct {
	Desc string `json:"desc"`
}

// SetPostContentMultipartBody defines parameters for SetPostContent.
type SetPostContentMultipartBody struct {
	Content openapi_types.File `json:"content"`
}

// ModifyPostsJSONBody defines parameters for ModifyPosts.
type ModifyPostsJSONBody struct {
	// Category 分组
	Category *uint32   `json:"category,omitempty"`
	Desc     *string   `json:"desc,omitempty"`
	Tags     *[]string `json:"tags,omitempty"`
	Title    *string   `json:"title,omitempty"`
	Visible  *bool     `json:"visible,omitempty"`
}

// GetPostListParams defines parameters for GetPostList.
type GetPostListParams struct {
	Page     *int    `form:"page,omitempty" json:"page,omitempty"`
	Tag      *string `form:"tag,omitempty" json:"tag,omitempty"`
	Category *uint32 `form:"category,omitempty" json:"category,omitempty"`
	Limit    *int    `form:"limit,omitempty" json:"limit,omitempty"`
}

// CreatePostMultipartBody defines parameters for CreatePost.
type CreatePostMultipartBody struct {
	Content openapi_types.File `json:"content"`

	// Uri ID 编号
	Uri string `json:"uri"`
}

// GetAdminTokenParams defines parameters for GetAdminToken.
type GetAdminTokenParams struct {
	Key *string `form:"key,omitempty" json:"key,omitempty"`
}

// CreateCategoryJSONRequestBody defines body for CreateCategory for apps/json ContentType.
type CreateCategoryJSONRequestBody CreateCategoryJSONBody

// ModifyCategoryDescJSONRequestBody defines body for ModifyCategoryDesc for apps/json ContentType.
type ModifyCategoryDescJSONRequestBody ModifyCategoryDescJSONBody

// SetPostContentMultipartRequestBody defines body for SetPostContent for multipart/form-data ContentType.
type SetPostContentMultipartRequestBody SetPostContentMultipartBody

// ModifyPostsJSONRequestBody defines body for ModifyPosts for apps/json ContentType.
type ModifyPostsJSONRequestBody ModifyPostsJSONBody

// CreatePostMultipartRequestBody defines body for CreatePost for multipart/form-data ContentType.
type CreatePostMultipartRequestBody CreatePostMultipartBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Admin Get PostList
	// (GET /admin/posts)
	AdminGetPostList(c *gin.Context, params AdminGetPostListParams)
	// 全部分类
	// (GET /categories)
	GetAllCategorys(c *gin.Context)
	// 创建分类
	// (POST /categories)
	CreateCategory(c *gin.Context)
	// 删除分类
	// (DELETE /category/{id})
	DeleteCategory(c *gin.Context, id uint32)
	// 修改简介
	// (PATCH /category/{id})
	ModifyCategoryDesc(c *gin.Context, id uint32)
	// 删除
	// (DELETE /post/{id})
	DeletePost(c *gin.Context, id uint32)
	// 设置PostContent
	// (POST /post/{id})
	SetPostContent(c *gin.Context, id uint32)
	// 修改Post
	// (PUT /post/{id})
	ModifyPosts(c *gin.Context, id uint32)
	// 帖子
	// (GET /post/{uri})
	GetPostByUri(c *gin.Context, uri string)
	// 帖子列表
	// (GET /posts)
	GetPostList(c *gin.Context, params GetPostListParams)
	// 创建Post
	// (POST /posts)
	CreatePost(c *gin.Context)
	// 最近帖子
	// (GET /posts/recent)
	GetRecentPostList(c *gin.Context)
	// 全部标签
	// (GET /tags)
	GetAllTags(c *gin.Context)
	// 获取管理员token
	// (GET /token)
	GetAdminToken(c *gin.Context, params GetAdminTokenParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// AdminGetPostList operation middleware
func (siw *ServerInterfaceWrapper) AdminGetPostList(c *gin.Context) {

	var err error

	c.Set(BearerScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params AdminGetPostListParams

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", c.Request.URL.Query(), &params.Page)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter page: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.AdminGetPostList(c, params)
}

// GetAllCategorys operation middleware
func (siw *ServerInterfaceWrapper) GetAllCategorys(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetAllCategorys(c)
}

// CreateCategory operation middleware
func (siw *ServerInterfaceWrapper) CreateCategory(c *gin.Context) {

	c.Set(BearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateCategory(c)
}

// DeleteCategory operation middleware
func (siw *ServerInterfaceWrapper) DeleteCategory(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id uint32

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteCategory(c, id)
}

// ModifyCategoryDesc operation middleware
func (siw *ServerInterfaceWrapper) ModifyCategoryDesc(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id uint32

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ModifyCategoryDesc(c, id)
}

// DeletePost operation middleware
func (siw *ServerInterfaceWrapper) DeletePost(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id uint32

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeletePost(c, id)
}

// SetPostContent operation middleware
func (siw *ServerInterfaceWrapper) SetPostContent(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id uint32

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.SetPostContent(c, id)
}

// ModifyPosts operation middleware
func (siw *ServerInterfaceWrapper) ModifyPosts(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id uint32

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ModifyPosts(c, id)
}

// GetPostByUri operation middleware
func (siw *ServerInterfaceWrapper) GetPostByUri(c *gin.Context) {

	var err error

	// ------------- Path parameter "uri" -------------
	var uri string

	err = runtime.BindStyledParameterWithOptions("simple", "uri", c.Param("uri"), &uri, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter uri: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetPostByUri(c, uri)
}

// GetPostList operation middleware
func (siw *ServerInterfaceWrapper) GetPostList(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPostListParams

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", c.Request.URL.Query(), &params.Page)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter page: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "tag" -------------

	err = runtime.BindQueryParameter("form", true, false, "tag", c.Request.URL.Query(), &params.Tag)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter tag: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "category" -------------

	err = runtime.BindQueryParameter("form", true, false, "category", c.Request.URL.Query(), &params.Category)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter category: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", c.Request.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter limit: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetPostList(c, params)
}

// CreatePost operation middleware
func (siw *ServerInterfaceWrapper) CreatePost(c *gin.Context) {

	c.Set(BearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreatePost(c)
}

// GetRecentPostList operation middleware
func (siw *ServerInterfaceWrapper) GetRecentPostList(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetRecentPostList(c)
}

// GetAllTags operation middleware
func (siw *ServerInterfaceWrapper) GetAllTags(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetAllTags(c)
}

// GetAdminToken operation middleware
func (siw *ServerInterfaceWrapper) GetAdminToken(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetAdminTokenParams

	// ------------- Optional query parameter "key" -------------

	err = runtime.BindQueryParameter("form", true, false, "key", c.Request.URL.Query(), &params.Key)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter key: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetAdminToken(c, params)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/admin/posts", wrapper.AdminGetPostList)
	router.GET(options.BaseURL+"/categories", wrapper.GetAllCategorys)
	router.POST(options.BaseURL+"/categories", wrapper.CreateCategory)
	router.DELETE(options.BaseURL+"/category/:id", wrapper.DeleteCategory)
	router.PATCH(options.BaseURL+"/category/:id", wrapper.ModifyCategoryDesc)
	router.DELETE(options.BaseURL+"/post/:id", wrapper.DeletePost)
	router.POST(options.BaseURL+"/post/:id", wrapper.SetPostContent)
	router.PUT(options.BaseURL+"/post/:id", wrapper.ModifyPosts)
	router.GET(options.BaseURL+"/post/:uri", wrapper.GetPostByUri)
	router.GET(options.BaseURL+"/posts", wrapper.GetPostList)
	router.POST(options.BaseURL+"/posts", wrapper.CreatePost)
	router.GET(options.BaseURL+"/posts/recent", wrapper.GetRecentPostList)
	router.GET(options.BaseURL+"/tags", wrapper.GetAllTags)
	router.GET(options.BaseURL+"/token", wrapper.GetAdminToken)
}
