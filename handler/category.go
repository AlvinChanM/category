package handler

import (
	"category/common"
	"category/domain/model"
	"category/domain/service"
	proto "category/proto/category"
	"context"
	"github.com/micro/go-micro/v2/util/log"
)

type Category struct {
	CategoryDataService service.ICategoryDataService
}

func (c *Category) CreateCategory(ctx context.Context, request *proto.CategoryRequest,
	response *proto.CreateCategoryResponse) error {
	cate := &model.Category{}
	err := common.SwapTo(request, cate)
	if err != nil {
		return err
	}
	cid, err := c.CategoryDataService.AddCategory(cate)
	if err != nil {
		return err
	}
	response.Message = "创建成功"
	response.CategoryId = cid
	return err

}

func (c *Category) UpdateCategory(ctx context.Context, request *proto.CategoryRequest, response *proto.UpdateCategoryResponse) error {
	cate := &model.Category{}
	err := common.SwapTo(request, cate)
	if err != nil {
		return err
	}
	err = c.CategoryDataService.UpdateCategory(cate)
	if err != nil {
		return err
	}
	response.Message = "更新成功！"
	return err
}

func (c *Category) FindCategoryByName(ctx context.Context, request *proto.FindByNameRequest,
	response *proto.CategoryResponse) error {
	cate, err := c.CategoryDataService.FindCategoryByName(request.CategoryName)
	err = common.SwapTo(cate, response)
	return err
}

func (c *Category) FindCategoryByID(ctx context.Context, request *proto.FindByIdRequest,
	response *proto.CategoryResponse) error {
	cate, err := c.CategoryDataService.FindCategoryByID(request.CategoryId)
	if err != nil {
		return err
	}
	err = common.SwapTo(cate, response)
	return err
}

func (c *Category) FindCategoryByLevel(ctx context.Context, request *proto.FindByLevelRequest,
	response *proto.FindAllResponse) error {
	cates, err := c.CategoryDataService.FindCategoryByLevel(request.Level)
	if err != nil {
		return err
	}
	for _, cate := range cates {
		res := &proto.CategoryResponse{}
		err := common.SwapTo(cate, res)
		if err != nil {
			return err
		}
		response.Category = append(response.Category, res)

	}
	return nil
}

func (c *Category) FindCategoryByParent(ctx context.Context, request *proto.FindByParentRequest,
	response *proto.FindAllResponse) error {
	cates, err := c.CategoryDataService.FindCategoryByParent(request.ParentId)
	if err != nil {
		return err
	}
	response = CategoryToResponse(cates)
	return nil
}

func (c *Category) FindAllCategory(ctx context.Context, request *proto.FindAllRequest, response *proto.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindAllCategory()
	if err != nil {
		return err
	}
	response = CategoryToResponse(categorySlice)
	return nil
}

func (c *Category) DeleteCategory(ctx context.Context,
	request *proto.DeleteCategoryRequest, response *proto.DeleteCategoryResponse) error {
	cid := request.CategoryId
	err := c.CategoryDataService.DeleteCategory(cid)
	if err != nil {
		return err
	}
	response.Message = "删除成功"
	return nil
}


func CategoryToResponse(categorySlice []model.Category)(response *proto.FindAllResponse){
	for _, cate := range categorySlice{
		res := proto.CategoryResponse{}
		err := common.SwapTo(cate,&res)
		if err != nil {
			log.Error(err)
			return
		}
		response.Category = append(response.Category, &res)
	}
	return
}