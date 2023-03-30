package controller

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/weimob-tech/go-project-base/pkg/x"
	"github.com/weimob-tech/go-project-boot/pkg/wcontext"
	"gorm.io/gorm"
)

// User
/*
CREATE TABLE `user`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_name` longtext CHARACTER SET utf8 COLLATE utf8_bin NULL,
  `password` longtext CHARACTER SET utf8 COLLATE utf8_bin NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_users_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;
*/
type User struct {
	gorm.Model
	UserName string
	Password string
}

func MysqlInsertController() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		user := User{UserName: "jack", Password: "123456"}
		result := wcontext.Global().Mysql().Create(&user)
		if result.Error != nil {
			ctx.JSON(consts.StatusOK, x.Fail("90500", result.Error.Error()))
			return
		}

		ctx.JSON(consts.StatusOK, x.OkAnd("insert success"))
	}
}

func MysqlQueryController() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var users []User
		result := wcontext.Global().Mysql().Where("user_name=?", "jack").Find(&users)
		if result.Error != nil {
			ctx.JSON(consts.StatusOK, x.Fail("90500", result.Error.Error()))
			return
		}

		ctx.JSON(consts.StatusOK, x.OkAnd(users))
	}
}
