package controllers

import (
	"ImSdk/common/e"
	"ImSdk/common/model"
	"ImSdk/common/protos"
	"ImSdk/common/utils"
	"ImSdk/interfaces"
	"ImSdk/svc"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// @Summary 获取轮播图列表
// @Description 获取轮播图列表
// @Tags 发现页
// @ID GetBannerList
// @Accept json
// @Produce json
// @Param request body protos.GetBannerListReq true "请求体"
// @Success 200 {object} protos.GetBannerListResp "成功"
// @Router /api/v1/public/dapp/banner/list [post]
func GetBannerList(c *gin.Context) {
	var req protos.GetBannerListReq
	if err := c.ShouldBind(&req); err != nil {
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
	data := make([]protos.GetBannerListResp, 0)
	for _, v := range *list {
		r := protos.GetBannerListResp{
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

// @Summary 获取工具分类列表
// @Description 获取工具分类列表
// @Tags 发现页
// @ID GetDiscoverToolCategoriesList
// @Accept json
// @Produce json
// @Param request body protos.GetDiscoverToolCategoriesListReq true "请求体"
// @Success 200 {object} protos.GetDiscoverToolCategoriesListResp "成功"
// @Router /api/v1/public/dapp/tool/categories/list [post]
func GetDiscoverToolCategoriesList(c *gin.Context) {
	var req protos.GetDiscoverToolCategoriesListReq
	if err := c.ShouldBind(&req); err != nil {
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
	data := make([]protos.GetDiscoverToolCategoriesListResp, 0)
	for _, v := range *list {
		r := protos.GetDiscoverToolCategoriesListResp{
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

// @Summary 获取网络列表
// @Description 获取网络列表
// @Tags 发现页
// @ID GetNetworkList
// @Accept json
// @Produce json
// @Param request body protos.GetNetworkListReq true "请求体"
// @Success 200 {object} protos.GetNetworkListResp "成功"
// @Router /api/v1/public/dapp/network/list [post]
func GetNetworkList(c *gin.Context) {
	var req protos.GetNetworkListReq
	if err := c.ShouldBind(&req); err != nil {
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
	data := make([]protos.GetNetworkListResp, 0)
	for _, v := range *list {
		r := protos.GetNetworkListResp{
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

// @Summary 获取工具列表
// @Description 获取工具列表
// @Tags 发现页
// @ID GetDiscoverToolInfoList
// @Accept json
// @Produce json
// @Param request body protos.GetDiscoverToolInfoListReq true "请求体"
// @Success 200 {object} protos.GetDiscoverToolInfoListResp "成功"
// @Router /api/v1/public/dapp/tool/info/list [post]
func GetDiscoverToolInfoList(c *gin.Context) {
	var req protos.GetDiscoverToolInfoListReq
	if err := c.ShouldBind(&req); err != nil {
		fmt.Printf("获取请求参数：%s\n", err.Error())
		c.JSON(http.StatusOK, e.GetMsg(e.ErrorInvalidParam))
		return
	}
	pageData := utils.PageData{Page: int(req.Page), PageSize: int(req.PageSize)}
	page := utils.GetPageData(pageData)

	list, count, err := svc.Ctx.DappDiscoverToolInfoModel.DiscoverToolInfoFindListPage(c.Request.Context(), req.ToolName, req.NetWorkId, req.CategoriesId, req.Tag, &page)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
		return
	}
	data := make([]protos.GetDiscoverToolInfoListResp, 0)
	for _, v := range *list {
		r := protos.GetDiscoverToolInfoListResp{
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
			CreatedAt:             v.CreatedAt,
			UpdatedAt:             v.UpdatedAt,
		}
		for _, v := range v.CategoryIds {
			categories, err := svc.Ctx.DappDiscoverToolCategoriesModel.FindOne(c.Request.Context(), v)
			if err == nil {
				r.CategoryInfoIds = append(r.CategoryInfoIds, protos.CategoryInfoItem{
					Id:      categories.Id,
					Title:   categories.Title,
					ImgHref: categories.ImgHref,
				})
			}
		}
		for _, v := range v.SupportNetworkIds {
			network, err := svc.Ctx.DappNetworkModel.FindOne(c.Request.Context(), v)
			if err == nil {
				r.SupportNetworkInfoIds = append(r.SupportNetworkInfoIds, protos.NetWorkInfoItem{
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

// @Summary 获取发现页工具栏列表
// @Description 获取发现页工具栏列表
// @Tags 发现页
// @ID GetDiscoverToolbarList
// @Accept json
// @Produce json
// @Param request body protos.GetDiscoverToolbarListReq true "请求体"
// @Success 200 {object} protos.GetDiscoverToolbarListResp "成功"
// @Router /api/v1/public/dapp/toolbar/list [post]
func GetDiscoverToolbarList(c *gin.Context) {
	var req protos.GetDiscoverToolbarListReq
	if err := c.ShouldBind(&req); err != nil {
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
	data := make([]protos.GetDiscoverToolbarListResp, 0)
	for _, v := range *list {
		r := protos.GetDiscoverToolbarListResp{
			Id:        v.Id,
			Title:     v.Title,
			Tag:       v.Tag,
			ImgHref:   v.ImgHref,
			Sort:      v.Sort,
			CreatedAt: v.CreatedAt,
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

// @Summary 获取发现页工具收藏列表
// @Description 获取发现页工具收藏列表
// @Tags 发现页
// @ID GetDiscoverToolFavoritesList
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param request body protos.GetDiscoverToolFavoritesListReq true "请求体"
// @Success 200 {object} protos.GetDiscoverToolFavoritesListResp "成功"
// @Router /api/v1/public/dapp/tool/favorites/list [post]
func GetDiscoverToolFavoritesList(c *gin.Context) {
	var req protos.GetDiscoverToolFavoritesListReq
	if err := c.ShouldBind(&req); err != nil {
		fmt.Printf("获取请求参数：%s\n", err.Error())
		c.JSON(http.StatusOK, e.GetMsg(e.ErrorInvalidParam))
		return
	}
	token := c.GetHeader("token")
	usInfo, err := interfaces.GetUserInfo(token, svc.Ctx.Config.Mode)
	if err != nil {
		fmt.Errorf("get user failed ! :%s ", err.Error())
		c.JSON(http.StatusOK, e.GetMsg(e.ErrorTokenNotExist))
		return
	}
	uid := ""
	if _, ok := usInfo.Data["userID"]; ok {
		uid = usInfo.Data["userID"].(string)
	}

	pageData := utils.PageData{Page: int(req.Page), PageSize: int(req.PageSize)}
	page := utils.GetPageData(pageData)

	list, count, err := svc.Ctx.DappDiscoverToolFavoritesModel.DiscoverToolFavoritesFindListPage(c.Request.Context(), uid, req.NetWorkId, &page)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
		return
	}
	data := make([]protos.GetDiscoverToolFavoritesListResp, 0)
	for _, v := range *list {
		toolInfo, _ := svc.Ctx.DappDiscoverToolInfoModel.FindOne(c.Request.Context(), v.ToolId)
		if toolInfo != nil {
			r := protos.GetDiscoverToolFavoritesListResp{
				Id:                    v.Id,
				ToolId:                v.ToolId,
				CategoryInfoIds:       nil,
				Title:                 toolInfo.Title,
				ShortDesc:             toolInfo.ShortDesc,
				Thumbnail:             toolInfo.Thumbnail,
				ImgHref:               toolInfo.ImgHref,
				SkipUrl:               toolInfo.SkipUrl,
				ServiceSupport:        toolInfo.ServiceSupport,
				SupportNetworkInfoIds: nil,
				Token:                 toolInfo.Token,
				CommunityTutorial:     toolInfo.CommunityTutorial,
				Status:                toolInfo.Status,
				Sort:                  toolInfo.Sort,
				CreatedAt:             v.CreatedAt,
			}
			for _, v := range toolInfo.CategoryIds {
				categories, err := svc.Ctx.DappDiscoverToolCategoriesModel.FindOne(c.Request.Context(), v)
				if err == nil {
					r.CategoryInfoIds = append(r.CategoryInfoIds, protos.FavoritesCategoryInfoItem{
						Id:      categories.Id,
						Title:   categories.Title,
						ImgHref: categories.ImgHref,
					})
				}
			}
			for _, v := range toolInfo.SupportNetworkIds {
				network, err := svc.Ctx.DappNetworkModel.FindOne(c.Request.Context(), v)
				if err == nil {
					r.SupportNetworkInfoIds = append(r.SupportNetworkInfoIds, protos.FavoritesNetWorkInfoItem{
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

// @Summary 发现页工具收藏操作
// @Description 发现页工具收藏操作
// @Tags 发现页
// @ID DiscoverToolFavoritesOperation
// @Accept json
// @Produce json
// @Param request body protos.DiscoverToolFavoritesOperationReq true "请求体"
// @Success 200 {object} protos.DiscoverToolFavoritesOperationResp "成功"
// @Router /api/v1/public/dapp/tool/favorites/operation [post]
func DiscoverToolFavoritesOperation(c *gin.Context) {
	var req protos.DiscoverToolFavoritesOperationReq
	if err := c.ShouldBind(&req); err != nil {
		fmt.Printf("获取请求参数：%s\n", err.Error())
		c.JSON(http.StatusOK, e.GetMsg(e.ErrorInvalidParam))
		return
	}
	token := c.GetHeader("token")
	usInfo, err := interfaces.GetUserInfo(token, svc.Ctx.Config.Mode)
	if err != nil {
		fmt.Errorf("get user failed ! :%s ", err.Error())
		c.JSON(http.StatusOK, e.GetMsg(e.ErrorTokenNotExist))
		return
	}
	uid := ""
	if _, ok := usInfo.Data["userID"]; ok {
		uid = usInfo.Data["userID"].(string)
	}
	toolInfo, err := svc.Ctx.DappDiscoverToolInfoModel.FindOne(c.Request.Context(), req.ToolId)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
		return
	}
	pro, err := svc.Ctx.DappDiscoverToolFavoritesModel.FindOneByUserIdToolId(c.Request.Context(), uid, req.ToolId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = svc.Ctx.DappDiscoverToolFavoritesModel.Insert(c.Request.Context(), &model.DappDiscoverToolFavorites{
				UserId:            uid,
				ToolId:            req.ToolId,
				SupportNetworkIds: toolInfo.SupportNetworkIds,
				Status:            req.Status,
			})
		}
	} else {
		pro.Status = req.Status
		err = svc.Ctx.DappDiscoverToolFavoritesModel.Update(c.Request.Context(), nil, pro)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
			return
		}
	}
	res := protos.DiscoverToolFavoritesOperationResp{}
	c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": e.GetMsg(e.SUCCESS), "data": res})
}

// @Summary 发现页工具访问记录列表
// @Description 发现页工具访问记录列表
// @Tags 发现页
// @ID GetDiscoverToolEventList
// @Accept json
// @Produce json
// @Param request body protos.GetDiscoverToolEventListReq true "请求体"
// @Success 200 {object} protos.GetDiscoverToolEventListResp "成功"
// @Router /api/v1/public/dapp/tool/event/list [post]
func GetDiscoverToolEventList(c *gin.Context) {
	var req protos.GetDiscoverToolEventListReq
	if err := c.ShouldBind(&req); err != nil {
		fmt.Printf("获取请求参数：%s\n", err.Error())
		c.JSON(http.StatusOK, e.GetMsg(e.ErrorInvalidParam))
		return
	}
	if req.EventType == "" {
		req.EventType = "view"
	}
	token := c.GetHeader("token")
	usInfo, err := interfaces.GetUserInfo(token, svc.Ctx.Config.Mode)
	if err != nil {
		fmt.Errorf("get user failed ! :%s ", err.Error())
		c.JSON(http.StatusOK, e.GetMsg(e.ErrorTokenNotExist))
		return
	}
	uid := ""
	if _, ok := usInfo.Data["userID"]; ok {
		uid = usInfo.Data["userID"].(string)
	}
	pageData := utils.PageData{Page: int(req.Page), PageSize: int(req.PageSize)}
	page := utils.GetPageData(pageData)

	list, count, err := svc.Ctx.DappDiscoverToolEventsModel.DiscoverToolEventsFindListPage(c.Request.Context(), uid, req.EventType, &page)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
		return
	}
	data := make([]protos.GetDiscoverToolEventListResp, 0)
	for _, v := range *list {
		toolInfo, _ := svc.Ctx.DappDiscoverToolInfoModel.FindOne(c.Request.Context(), v.ItemId)
		if toolInfo != nil {
			r := protos.GetDiscoverToolEventListResp{
				Id:                    v.Id,
				ToolId:                v.ItemId,
				CategoryInfoIds:       nil,
				Title:                 toolInfo.Title,
				ShortDesc:             toolInfo.ShortDesc,
				Thumbnail:             toolInfo.Thumbnail,
				ImgHref:               toolInfo.ImgHref,
				SkipUrl:               toolInfo.SkipUrl,
				ServiceSupport:        toolInfo.ServiceSupport,
				SupportNetworkInfoIds: nil,
				Token:                 toolInfo.Token,
				CommunityTutorial:     toolInfo.CommunityTutorial,
				Status:                toolInfo.Status,
				Sort:                  toolInfo.Sort,
				CreatedAt:             v.CreatedAt,
			}
			for _, v := range toolInfo.CategoryIds {
				categories, err := svc.Ctx.DappDiscoverToolCategoriesModel.FindOne(c.Request.Context(), v)
				if err == nil {
					r.CategoryInfoIds = append(r.CategoryInfoIds, protos.EventCategoryInfoItem{
						Id:      categories.Id,
						Title:   categories.Title,
						ImgHref: categories.ImgHref,
					})
				}
			}
			for _, v := range toolInfo.SupportNetworkIds {
				network, err := svc.Ctx.DappNetworkModel.FindOne(c.Request.Context(), v)
				if err == nil {
					r.SupportNetworkInfoIds = append(r.SupportNetworkInfoIds, protos.EventNetWorkInfoItem{
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

// @Summary 发现页工具访问记录操作
// @Description 发现页工具访问记录操作
// @Tags 发现页
// @ID DiscoverToolEventOperation
// @Accept json
// @Produce json
// @Param request body protos.DiscoverToolEventOperationReq true "请求体"
// @Success 200 {object} protos.DiscoverToolEventOperationResp "成功"
// @Router /api/v1/public/dapp/tool/event/operation [post]
func DiscoverToolEventOperation(c *gin.Context) {
	var req protos.DiscoverToolEventOperationReq
	if err := c.ShouldBind(&req); err != nil {
		fmt.Printf("获取请求参数：%s\n", err.Error())
		c.JSON(http.StatusOK, e.GetMsg(e.ErrorInvalidParam))
		return
	}
	token := c.GetHeader("token")
	usInfo, err := interfaces.GetUserInfo(token, svc.Ctx.Config.Mode)
	if err != nil {
		fmt.Errorf("get user failed ! :%s ", err.Error())
		c.JSON(http.StatusOK, e.GetMsg(e.ErrorTokenNotExist))
		return
	}
	uid := ""
	if _, ok := usInfo.Data["userID"]; ok {
		uid = usInfo.Data["userID"].(string)
	}
	ip := c.GetHeader("X-Real-IP")
	if ip == "" {
		ip = c.ClientIP()
	}
	ua := c.GetHeader("User-Agent")

	meta := c.GetHeader("meta")
	now := time.Now()
	_, err = svc.Ctx.DappDiscoverToolInfoModel.FindOne(c.Request.Context(), req.ToolId)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
		return
	}
	pro, err := svc.Ctx.DappDiscoverToolEventsModel.FindOneByUserIdEventTypeToolId(c.Request.Context(), uid, req.EventType, req.ToolId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = svc.Ctx.DappDiscoverToolEventsModel.Insert(c.Request.Context(), &model.DappDiscoverToolEvents{
				UserId:    uid,
				ItemId:    req.ToolId,
				EventType: req.EventType,
				Ip:        ip,
				Ua:        ua,
				Meta:      meta,
				CreatedAt: now,
			})
		}
	} else {
		pro.CreatedAt = now
		err = svc.Ctx.DappDiscoverToolEventsModel.Update(c.Request.Context(), nil, pro)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{"code": e.ERROR, "message": err.Error(), "data": nil})
			return
		}
	}
	res := protos.DiscoverToolEventOperationResp{}
	c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": e.GetMsg(e.SUCCESS), "data": res})
}
