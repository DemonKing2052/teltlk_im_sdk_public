package controllers

import (
	"ImSdk/common/e"
	"ImSdk/common/model"
	"ImSdk/common/protos"
	"ImSdk/common/utils"
	"ImSdk/svc"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// @Summary 获取轮播图列表
// @Description 获取轮播图列表
// @Tags dapp管理
// @ID GetManageBannerList
// @Accept json
// @Produce json
// @Param request body protos.GetManageBannerListReq true "请求体"
// @Success 200 {object} protos.GetManageBannerListResp "成功"
// @Router /api/v1/public/dapp/manage/banner/list [post]
func GetManageBannerList(c *gin.Context) {
	var req protos.GetManageBannerListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("获取请求参数：%s\n", err.Error())
		c.JSON(http.StatusOK, e.GetMsg(e.ErrorInvalidParam))
		return
	}
	pageData := utils.PageData{Page: int(req.Page), PageSize: int(req.PageSize)}
	page := utils.GetPageData(pageData)

	list, count, err := svc.Ctx.DappBannerModel.BannerFindListPage(c.Request.Context(), &page)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
		return
	}
	data := make([]protos.GetManageBannerListResp, 0)
	for _, v := range *list {
		r := protos.GetManageBannerListResp{
			Id:             v.Id,
			Position:       v.Position,
			Title:          v.Title,
			Sort:           v.Sort,
			SkipUrl:        v.SkipUrl,
			ImgHref:        v.ImgHref,
			Status:         v.Status,
			NeedLogin:      v.NeedLogin,
			SkipTarget:     v.SkipTarget,
			StartAt:        v.StartAt,
			ExpirationTime: v.ExpirationTime,
			CreatedTime:    v.CreatedTime,
			UpdatedTime:    v.UpdatedTime,
		}
		data = append(data, r)
	}
	res := protos.CommonListResp{
		List:     data,
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    count,
	}
	c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": e.GetMsg(e.SUCCESS), "data": res})
	return
}

// @Summary 轮播图操作管理
// @Description 轮播图操作管理
// @Tags dapp管理
// @ID ManageBannerOperation
// @Accept json
// @Produce json
// @Param request body protos.ManageBannerOperationReq true "请求体"
// @Success 200 {object} protos.ManageBannerOperationResp "成功"
// @Router /api/v1/public/dapp/manage/banner/operation [post]
func ManageBannerOperation(c *gin.Context) {
	var req protos.ManageBannerOperationReq
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("获取请求参数：%s\n", err.Error())
		c.JSON(http.StatusOK, e.GetMsg(e.ErrorInvalidParam))
		return
	}
	now := time.Now()
	const (
		BannerOperationTypeAdd    = 1
		BannerOperationTypeEdit   = 2
		BannerOperationTypeDelete = 3
	)
	switch req.OperationType {
	case BannerOperationTypeAdd:
		pro := model.DappBanner{
			Position:       req.Position,
			Title:          req.Title,
			Sort:           req.Sort,
			SkipUrl:        req.SkipUrl,
			ImgHref:        req.ImgHref,
			Status:         req.Status,
			NeedLogin:      req.NeedLogin,
			SkipTarget:     req.SkipTarget,
			StartAt:        nil,
			ExpirationTime: nil,
			CreatedTime:    now,
			UpdatedTime:    now,
		}
		if req.StartAt != "" {
			startAt, err := utils.StrToTime(req.StartAt)
			if err == nil {
				pro.StartAt = &startAt
			}
		}
		if req.ExpirationTime != "" {
			expirationTimeAt, err := utils.StrToTime(req.ExpirationTime)
			if err == nil {
				pro.ExpirationTime = &expirationTimeAt
			}
		}
		err := svc.Ctx.DappBannerModel.Insert(c.Request.Context(), &pro)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
			return
		}
	case BannerOperationTypeEdit:
		pro, err := svc.Ctx.DappBannerModel.FindOne(c.Request.Context(), req.Id)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
			return
		}
		pro.Position = req.Position
		pro.Title = req.Title
		pro.Sort = req.Sort
		pro.SkipUrl = req.SkipUrl
		pro.ImgHref = req.ImgHref
		pro.Status = req.Status
		pro.NeedLogin = req.NeedLogin
		pro.SkipTarget = req.SkipTarget
		pro.UpdatedTime = now

		if req.StartAt != "" {
			startAt, err := utils.StrToTime(req.StartAt)
			if err == nil {
				pro.StartAt = &startAt
			}
		}
		if req.ExpirationTime != "" {
			expirationTimeAt, err := utils.StrToTime(req.ExpirationTime)
			if err == nil {
				pro.ExpirationTime = &expirationTimeAt
			}
		}

		err = svc.Ctx.DappBannerModel.Update(c.Request.Context(), nil, pro)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": err.Error(), "data": nil})
			return
		}
	case BannerOperationTypeDelete:
		pro, err := svc.Ctx.DappBannerModel.FindOne(c.Request.Context(), req.Id)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
			return
		}
		err = svc.Ctx.DappBannerModel.Delete(c.Request.Context(), nil, pro.Id)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
			return
		}
	}
	data := protos.ManageBannerOperationResp{}
	c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": e.GetMsg(e.SUCCESS), "data": data})
}

// @Summary 工具分类列表管理
// @Description 工具分类列表管理
// @Tags dapp发现页管理
// @ID GetManageDiscoverToolCategoriesList
// @Accept json
// @Produce json
// @Param request body protos.GetManageDiscoverToolCategoriesListReq true "请求体"
// @Success 200 {object} protos.GetManageDiscoverToolCategoriesListResp "成功"
// @Router /api/v1/public/dapp/manage/discover/tool/categories/list [post]
func GetManageDiscoverToolCategoriesList(c *gin.Context) {
	var req protos.GetManageDiscoverToolCategoriesListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("获取请求参数：%s\n", err.Error())
		c.JSON(http.StatusOK, e.GetMsg(e.ErrorInvalidParam))
		return
	}
	pageData := utils.PageData{Page: int(req.Page), PageSize: int(req.PageSize)}
	page := utils.GetPageData(pageData)

	list, count, err := svc.Ctx.DappDiscoverToolCategoriesModel.DiscoverToolCategoriesFindListPage(c.Request.Context(), &page)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
		return
	}
	data := make([]protos.GetManageDiscoverToolCategoriesListResp, 0)
	for _, v := range *list {
		r := protos.GetManageDiscoverToolCategoriesListResp{
			Id:          v.Id,
			Title:       v.Title,
			Sort:        v.Sort,
			SkipUrl:     v.SkipUrl,
			ImgHref:     v.ImgHref,
			Description: v.Description,
			Status:      v.Status,
			IsShow:      v.IsShow,
			CreatedTime: v.CreatedTime,
			UpdatedTime: v.UpdatedTime,
		}
		data = append(data, r)
	}
	res := protos.CommonListResp{
		List:     data,
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    count,
	}
	c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": e.GetMsg(e.SUCCESS), "data": res})
	return
}

// @Summary 工具分类操作管理
// @Description 工具分类操作管理
// @Tags dapp发现页管理
// @ID ManageDiscoverToolCategoriesOperation
// @Accept json
// @Produce json
// @Param request body protos.ManageDiscoverToolCategoriesOperationReq true "请求体"
// @Success 200 {object} protos.ManageDiscoverToolCategoriesOperationResp "成功"
// @Router /api/v1/public/dapp/manage/discover/tool/categories/operation [post]
func ManageDiscoverToolCategoriesOperation(c *gin.Context) {
	var req protos.ManageDiscoverToolCategoriesOperationReq
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("获取请求参数：%s\n", err.Error())
		c.JSON(http.StatusOK, e.GetMsg(e.ErrorInvalidParam))
		return
	}
	now := time.Now()
	const (
		DiscoverToolCategoriesOperationTypeAdd    = 1
		DiscoverToolCategoriesOperationTypeEdit   = 2
		DiscoverToolCategoriesOperationTypeDelete = 3
	)
	switch req.OperationType {
	case DiscoverToolCategoriesOperationTypeAdd:
		pro := model.DappDiscoverToolCategories{
			Title:       req.Title,
			Sort:        req.Sort,
			SkipUrl:     req.SkipUrl,
			ImgHref:     req.ImgHref,
			Description: req.Description,
			Status:      req.Status,
			IsShow:      req.IsShow,
			CreatedTime: now,
			UpdatedTime: now,
		}
		err := svc.Ctx.DappDiscoverToolCategoriesModel.Insert(c.Request.Context(), &pro)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
			return
		}
	case DiscoverToolCategoriesOperationTypeEdit:
		pro, err := svc.Ctx.DappDiscoverToolCategoriesModel.FindOne(c.Request.Context(), req.Id)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
			return
		}
		pro.Title = req.Title
		pro.Sort = req.Sort
		pro.SkipUrl = req.SkipUrl
		pro.ImgHref = req.ImgHref
		pro.Description = req.Description
		pro.Status = req.Status
		pro.IsShow = req.IsShow
		pro.UpdatedTime = now
		err = svc.Ctx.DappDiscoverToolCategoriesModel.Update(c.Request.Context(), nil, pro)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": err.Error(), "data": nil})
			return
		}
	case DiscoverToolCategoriesOperationTypeDelete:
		pro, err := svc.Ctx.DappDiscoverToolCategoriesModel.FindOne(c.Request.Context(), req.Id)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
			return
		}
		err = svc.Ctx.DappDiscoverToolCategoriesModel.Delete(c.Request.Context(), nil, pro.Id)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
			return
		}
	}
	data := protos.ManageDiscoverToolCategoriesOperationResp{}
	c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": e.GetMsg(e.SUCCESS), "data": data})
}

// @Summary 区块网络列表管理
// @Description 区块网络列表管理
// @Tags dapp管理
// @ID GetManageNetworkList
// @Accept json
// @Produce json
// @Param request body protos.GetManageNetworkListReq true "请求体"
// @Success 200 {object} protos.GetManageNetworkListResp "成功"
// @Router /api/v1/public/dapp/manage/network/list [post]
func GetManageNetworkList(c *gin.Context) {
	var req protos.GetManageNetworkListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("获取请求参数：%s\n", err.Error())
		c.JSON(http.StatusOK, e.GetMsg(e.ErrorInvalidParam))
		return
	}
	pageData := utils.PageData{Page: int(req.Page), PageSize: int(req.PageSize)}
	page := utils.GetPageData(pageData)

	list, count, err := svc.Ctx.DappNetworkModel.NetWorkFindListPage(c.Request.Context(), &page)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
		return
	}
	data := make([]protos.GetManageNetworkListResp, 0)
	for _, v := range *list {
		r := protos.GetManageNetworkListResp{
			Id:          v.Id,
			Name:        v.Name,
			ChainId:     v.ChainId,
			Symbol:      v.Symbol,
			RpcUrl:      v.RpcUrl,
			ExplorerUrl: v.ExplorerUrl,
			Logo:        v.Logo,
			Status:      v.Status,
			Sort:        v.Sort,
			CreatedAt:   v.CreatedAt,
		}
		data = append(data, r)
	}
	res := protos.CommonListResp{
		List:     data,
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    count,
	}
	c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": e.GetMsg(e.SUCCESS), "data": res})
	return
}

// @Summary 区块网络操作管理
// @Description 区块网络操作管理
// @Tags dapp管理
// @ID ManageNetworkOperation
// @Accept json
// @Produce json
// @Param request body protos.ManageNetworkOperationReq true "请求体"
// @Success 200 {object} protos.ManageNetworkOperationResp "成功"
// @Router /api/v1/public/dapp/manage/network/operation [post]
func ManageNetworkOperation(c *gin.Context) {
	var req protos.ManageNetworkOperationReq
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("获取请求参数：%s\n", err.Error())
		c.JSON(http.StatusOK, e.GetMsg(e.ErrorInvalidParam))
		return
	}
	now := time.Now()
	const (
		NetworkOperationTypeAdd    = 1
		NetworkOperationTypeEdit   = 2
		NetworkOperationTypeDelete = 3
	)
	switch req.OperationType {
	case NetworkOperationTypeAdd:
		pro := model.DappNetwork{
			Name:        req.Name,
			ChainId:     req.ChainId,
			Symbol:      req.Symbol,
			RpcUrl:      req.RpcUrl,
			ExplorerUrl: req.ExplorerUrl,
			Logo:        req.Logo,
			Status:      req.Status,
			Sort:        req.Sort,
			CreatedAt:   now,
		}
		err := svc.Ctx.DappNetworkModel.Insert(c.Request.Context(), &pro)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
			return
		}
	case NetworkOperationTypeEdit:
		pro, err := svc.Ctx.DappNetworkModel.FindOne(c.Request.Context(), req.Id)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
			return
		}
		pro.Name = req.Name
		pro.ChainId = req.ChainId
		pro.Symbol = req.Symbol
		pro.RpcUrl = req.RpcUrl
		pro.ExplorerUrl = req.ExplorerUrl
		pro.Logo = req.Logo
		pro.Status = req.Status
		pro.Sort = req.Sort
		pro.CreatedAt = now
		err = svc.Ctx.DappNetworkModel.Update(c.Request.Context(), nil, pro)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": err.Error(), "data": nil})
			return
		}
	case NetworkOperationTypeDelete:
		pro, err := svc.Ctx.DappNetworkModel.FindOne(c.Request.Context(), req.Id)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
			return
		}
		err = svc.Ctx.DappNetworkModel.Delete(c.Request.Context(), nil, pro.Id)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
			return
		}
	}
	data := protos.ManageNetworkOperationResp{}
	c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": e.GetMsg(e.SUCCESS), "data": data})
}

// @Summary 工具列表管理
// @Description 工具列表管理
// @Tags dapp发现页管理
// @ID GetManageDiscoverToolInfoList
// @Accept json
// @Produce json
// @Param request body protos.GetManageDiscoverToolInfoListReq true "请求体"
// @Success 200 {object} protos.GetManageDiscoverToolInfoListResp "成功"
// @Router /api/v1/public/dapp/manage/discover/tool/info/list [post]
func GetManageDiscoverToolInfoList(c *gin.Context) {
	var req protos.GetManageDiscoverToolInfoListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("获取请求参数：%s\n", err.Error())
		c.JSON(http.StatusOK, e.GetMsg(e.ErrorInvalidParam))
		return
	}
	pageData := utils.PageData{Page: int(req.Page), PageSize: int(req.PageSize)}
	page := utils.GetPageData(pageData)

	list, count, err := svc.Ctx.DappDiscoverToolInfoModel.DiscoverToolInfoFindListPage(c.Request.Context(), "", 0, req.CategoriesId, "", &page)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
		return
	}
	data := make([]protos.GetManageDiscoverToolInfoListResp, 0)
	for _, v := range *list {
		r := protos.GetManageDiscoverToolInfoListResp{
			Id:                    v.Id,
			CategoryInfoIds:       nil,
			Title:                 v.Title,
			ShortDesc:             v.ShortDesc,
			Thumbnail:             v.Thumbnail,
			ImgHref:               v.ImgHref,
			SkipUrl:               v.SkipUrl,
			ServiceSupport:        v.ServiceSupport,
			SupportNetworkInfoIds: nil,
			Token:                 v.Token,
			CommunityTutorial:     v.CommunityTutorial,
			Status:                v.Status,
			Sort:                  v.Sort,
			Tags:                  v.Tags,
			CreatedAt:             v.CreatedAt,
			UpdatedAt:             v.UpdatedAt,
		}
		for _, v := range v.CategoryIds {
			categories, err := svc.Ctx.DappDiscoverToolCategoriesModel.FindOne(c.Request.Context(), v)
			if err == nil {
				r.CategoryInfoIds = append(r.CategoryInfoIds, protos.ManageCategoryInfoItem{
					Id:      categories.Id,
					Title:   categories.Title,
					ImgHref: categories.ImgHref,
				})
			}
		}
		for _, v := range v.SupportNetworkIds {
			network, err := svc.Ctx.DappNetworkModel.FindOne(c.Request.Context(), v)
			if err == nil {
				r.SupportNetworkInfoIds = append(r.SupportNetworkInfoIds, protos.ManageNetWorkInfoItem{
					Id:          network.Id,
					Name:        network.Name,
					ChainId:     network.ChainId,
					Symbol:      network.Symbol,
					RpcUrl:      network.RpcUrl,
					ExplorerUrl: network.ExplorerUrl,
					Logo:        network.Logo,
				})
			}
		}
		data = append(data, r)
	}
	res := protos.CommonListResp{
		List:     data,
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    count,
	}
	c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": e.GetMsg(e.SUCCESS), "data": res})
	return
}

// @Summary 工具操作管理
// @Description 工具操作管理
// @Tags dapp发现页管理
// @ID ManageDiscoverToolInfoOperation
// @Accept json
// @Produce json
// @Param request body protos.ManageDiscoverToolInfoOperationReq true "请求体"
// @Success 200 {object} protos.ManageDiscoverToolInfoOperationResp "成功"
// @Router /api/v1/public/dapp/manage/discover/tool/info/operation [post]
func ManageDiscoverToolInfoOperation(c *gin.Context) {
	var req protos.ManageDiscoverToolInfoOperationReq
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("获取请求参数：%s\n", err.Error())
		c.JSON(http.StatusOK, e.GetMsg(e.ErrorInvalidParam))
		return
	}
	fmt.Printf("获取请求参数：%+v\n", req)
	now := time.Now()
	const (
		DiscoverToolInfoOperationTypeAdd    = 1
		DiscoverToolInfoOperationTypeEdit   = 2
		DiscoverToolInfoOperationTypeDelete = 3
	)
	switch req.OperationType {
	case DiscoverToolInfoOperationTypeAdd:

		pro := model.DappDiscoverToolInfo{
			CategoryIds:       req.CategoryIds,
			Title:             req.Title,
			ShortDesc:         req.ShortDesc,
			Thumbnail:         req.Thumbnail,
			ImgHref:           req.ImgHref,
			SkipUrl:           req.SkipUrl,
			ServiceSupport:    req.ServiceSupport,
			SupportNetworkIds: req.SupportNetworkIds,
			Token:             req.Token,
			CommunityTutorial: req.CommunityTutorial,
			Status:            req.Status,
			Sort:              req.Sort,
			Tags:              req.Tags,
			CreatedAt:         now,
			UpdatedAt:         now,
		}
		err := svc.Ctx.DappDiscoverToolInfoModel.Insert(c.Request.Context(), &pro)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
			return
		}
	case DiscoverToolInfoOperationTypeEdit:
		pro, err := svc.Ctx.DappDiscoverToolInfoModel.FindOne(c.Request.Context(), req.Id)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
			return
		}
		pro.Id = req.Id
		pro.CategoryIds = req.CategoryIds
		pro.Title = req.Title
		pro.ShortDesc = req.ShortDesc
		pro.Thumbnail = req.Thumbnail
		pro.ImgHref = req.ImgHref
		pro.SkipUrl = req.SkipUrl
		pro.ServiceSupport = req.ServiceSupport
		pro.SupportNetworkIds = req.SupportNetworkIds
		pro.Token = req.Token
		pro.CommunityTutorial = req.CommunityTutorial
		pro.Status = req.Status
		pro.Sort = req.Sort
		pro.Tags = req.Tags
		pro.UpdatedAt = now
		err = svc.Ctx.DappDiscoverToolInfoModel.Update(c.Request.Context(), nil, pro)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": err.Error(), "data": nil})
			return
		}
	case DiscoverToolInfoOperationTypeDelete:
		pro, err := svc.Ctx.DappDiscoverToolInfoModel.FindOne(c.Request.Context(), req.Id)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
			return
		}
		err = svc.Ctx.DappDiscoverToolInfoModel.Delete(c.Request.Context(), nil, pro.Id)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
			return
		}
	}
	data := protos.ManageDiscoverToolInfoOperationResp{}
	c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": e.GetMsg(e.SUCCESS), "data": data})
}

// @Summary 发现页工具栏列表管理
// @Description 发现页工具栏列表管理
// @Tags dapp发现页管理
// @ID GetManageDiscovertoolbarList
// @Accept json
// @Produce json
// @Param request body protos.GetManageDiscovertoolbarListReq true "请求体"
// @Success 200 {object} protos.GetManageDiscovertoolbarListResp "成功"
// @Router /api/v1/public/dapp/manage/discover/toolbar/list [post]
func GetManageDiscovertoolbarList(c *gin.Context) {
	var req protos.GetManageDiscovertoolbarListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("获取请求参数：%s\n", err.Error())
		c.JSON(http.StatusOK, e.GetMsg(e.ErrorInvalidParam))
		return
	}
	pageData := utils.PageData{Page: int(req.Page), PageSize: int(req.PageSize)}
	page := utils.GetPageData(pageData)

	list, count, err := svc.Ctx.DappDiscoverToolbarModel.DiscoverToolbarFindListPage(c.Request.Context(), &page)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
		return
	}
	data := make([]protos.GetManageDiscovertoolbarListResp, 0)
	for _, v := range *list {
		r := protos.GetManageDiscovertoolbarListResp{
			Id:      v.Id,
			Title:   v.Title,
			ImgHref: v.ImgHref,
			Tag:     v.Tag,
		}
		data = append(data, r)
	}
	res := protos.CommonListResp{
		List:     data,
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    count,
	}
	c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": e.GetMsg(e.SUCCESS), "data": res})
	return
}

// @Summary 发现页工具栏操作管理
// @Description 发现页工具栏操作管理
// @Tags dapp发现页管理
// @ID ManageDiscovertoolbarOperation
// @Accept json
// @Produce json
// @Param request body protos.ManageDiscovertoolbarOperationReq true "请求体"
// @Success 200 {object} protos.ManageDiscovertoolbarOperationResp "成功"
// @Router /api/v1/public/dapp/manage/discover/toolbar/operation [post]
func ManageDiscovertoolbarOperation(c *gin.Context) {
	var req protos.ManageDiscovertoolbarOperationReq
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("获取请求参数：%s\n", err.Error())
		c.JSON(http.StatusOK, e.GetMsg(e.ErrorInvalidParam))
		return
	}
	now := time.Now()
	const (
		DiscoverToolbarOperationTypeAdd    = 1
		DiscoverToolbarOperationTypeEdit   = 2
		DiscoverToolbarOperationTypeDelete = 3
	)
	switch req.OperationType {
	case DiscoverToolbarOperationTypeAdd:
		pro := model.DappDiscoverToolbar{
			Id:        req.Id,
			Title:     req.Title,
			ImgHref:   req.ImgHref,
			Tag:       req.Tag,
			Sort:      req.Sort,
			CreatedAt: now,
		}
		err := svc.Ctx.DappDiscoverToolbarModel.Insert(c.Request.Context(), &pro)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
			return
		}
	case DiscoverToolbarOperationTypeEdit:
		pro, err := svc.Ctx.DappDiscoverToolbarModel.FindOne(c.Request.Context(), req.Id)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
			return
		}
		pro.Id = req.Id
		pro.Title = req.Title
		pro.ImgHref = req.ImgHref
		pro.Sort = req.Sort
		pro.Tag = req.Tag
		err = svc.Ctx.DappDiscoverToolbarModel.Update(c.Request.Context(), nil, pro)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": err.Error(), "data": nil})
			return
		}
	case DiscoverToolbarOperationTypeDelete:
		pro, err := svc.Ctx.DappDiscoverToolbarModel.FindOne(c.Request.Context(), req.Id)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
			return
		}
		err = svc.Ctx.DappDiscoverToolbarModel.Delete(c.Request.Context(), nil, pro.Id)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
			return
		}
	}
	data := protos.ManageDiscovertoolbarOperationResp{}
	c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": e.GetMsg(e.SUCCESS), "data": data})
}
