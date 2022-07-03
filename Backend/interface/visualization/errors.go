package visualization

import "github.com/gin-gonic/gin"

// Implement 服务常用实现
type Implement func(ctx *gin.Context, message BaseResponse)

// BadRequestImplement 400错误的实现
func BadRequestImplement() Implement {
	return func(ctx *gin.Context, message BaseResponse) {
		ctx.JSON(400, message)
	}
}

// ForbiddenImplement 403错误的实现
func ForbiddenImplement() Implement {
	return func(ctx *gin.Context, message BaseResponse) {
		ctx.JSON(403, message)
	}
}

// NotFoundImplement 404错误的实现
func NotFoundImplement() Implement {
	return func(ctx *gin.Context, message BaseResponse) {
		ctx.JSON(404, message)
	}
}

// MethodNotAllowedImplement 405错误的实现
func MethodNotAllowedImplement() Implement {
	return func(ctx *gin.Context, message BaseResponse) {
		ctx.JSON(405, message)
	}
}

// InternalServerErrorImplement 500错误的实现
func InternalServerErrorImplement() Implement {
	return func(ctx *gin.Context, message BaseResponse) {
		ctx.JSON(500, message)
	}
}

// BadGatewayImplement 502错误的实现
func BadGatewayImplement() Implement {
	return func(ctx *gin.Context, message BaseResponse) {
		ctx.JSON(502, message)
	}
}
