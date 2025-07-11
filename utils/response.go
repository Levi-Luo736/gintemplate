package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, message string, data interface{}, status int) {
	c.JSON(http.StatusOK, gin.H{
		"code":    status,
		"message": message,
		"data":    data,
	})
}

func ErrorResponse(c *gin.Context, err error, status int) {
	var code int
	var message string

	switch err.(type) {
	case *gin.Error:
		code = 500
		message = "Internal Server Error"
	default:
		code = 400
		message = err.Error()
	}

	c.JSON(status, gin.H{
		"code":    code,
		"message": message,
	})
}

// 成功响应的封装
// modules.success = function (res, message, data = {}, status = 200) {
//   res.status(status).json({
//     code: 0,
//     message,
//     data,
//   });
// };
// // 失败响应的封装
// modules.failure = function (res, error, code ) {
//   // 初始化状态码和错误信息
//   let statusCode = 500;
//   // 错误消息
//   let errors = [error.message || '服务器错误，请联系管理员~ '];
//   // 错误代码
//   let message = '服务器错误，请联系管理员~ ';
//   if (error.name === 'UnauthorizedError') {
//     statusCode = 401;
//     code = 11;
//     errors = [error.message];
//     message = '认证失败';
//   }
//   if (error.name === 'SequelizeValidationError') {
//     const errors = error.errors.map((e) => e.message);
//     statusCode = 400;
//     code = 13;
//     errors = errors;
//     message = '请求参数错误';
//   }
//   if (error.name === 'JsonWebTokenError') {
//     statusCode = 401;
//     code = 11;
//     errors = [error.message];
//     message = '认证失败';
//   }
//   if (error.name === 'SequelizeValidationError') {
//     const errors = error.errors.map((e) => e.message);
//     statusCode = 400;
//     code = 13;
//     errors = errors;
//     message = '请求参数错误';
//   }
//   if (error.name === 'JsonWebTokenError') {
//     statusCode = 401;
//     code = 11;
//     errors = [error.message];
//     message = '认证失败';
//   }
//   if (error.name === 'SequelizeValidationError') {
//     const errors = error.errors.map((e) => e.message);
//     statusCode = 400;
//     code = 13;
//     errors = errors;
//   }
//   if (error instanceof multer.MulterError) {
//     // multer 上传错误
//     if (error.code === 'LIMIT_FILE_SIZE') {
//       statusCode = 413;
//       errors = '文件大小超出限制。';
//     } else {
//       statusCode = 400;
//       errors = error.message;
//     }
//     return;
//   }
//   if (error.name === 'NotFoundError') {
//     // 资源未找到错误
//     statusCode = 404;
//     code = 12;
//     errors = [error.message];
//     message = error.name || '资源不存在';
//   }
//   res.status(statusCode).json({
//     code: code,
//     message: message,
//     errors: errors,
//   });
// };
