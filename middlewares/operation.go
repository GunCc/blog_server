package middlewares

import (
	"blog_server/global"
	"blog_server/model/system"
	"blog_server/service"
	"blog_server/utils"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 创建记录api
var operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService

// 内存池 **
var respPool sync.Pool

// 操作记录 中间件
func OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte
		var userId int
		// 如果不是 Get 请求
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = io.ReadAll(c.Request.Body)
			if err != nil {
				global.BLOG_LOG.Error("read body from request error:", zap.Error(err))
			} else {
				c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		} else {
			query := c.Request.URL.RawQuery
			// 解析 query
			query, _ = url.QueryUnescape(query)
			// 拆分数据
			split := strings.Split(query, "&")
			m := make(map[string]string)
			for _, v := range split {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					m[kv[0]] = kv[1]
				}
			}
			body, _ = json.Marshal(&m)
		}
		cc, _ := utils.GetClaims(c)
		if cc.ID != 0 {
			userId = int(cc.ID)
		} else {
			// 字符串转为int
			i, err := strconv.Atoi(c.Request.Header.Get("x-user-id"))
			if err != nil {
				userId = 0
			}
			userId = i
		}

		record := system.SysOperationRecord{
			Ip:     c.ClientIP(),
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
			Agent:  c.Request.UserAgent(),
			Body:   string(body),
			UserId: userId,
		}
		if strings.Contains(c.GetHeader("Content-Type"), "multipart/form-data") {
			if len(record.Body) > 1024 {
				// 截断
				newBody := respPool.Get().([]byte)
				copy(newBody, []byte(record.Body))
				record.Body = string(newBody)
				defer respPool.Put(newBody[:0])
			}
		}
		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}

		c.Writer = writer
		now := time.Now()
		c.Next()
		latency := time.Since(now)
		record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = c.Writer.Status()
		record.Latency = latency
		record.Resp = writer.body.String()

		if getHeaderRecordLen(c) {
			// 截断
			newBody := respPool.Get().([]byte)
			copy(newBody, []byte(record.Body))
			record.Body = string(newBody)
			defer respPool.Put(newBody[:0])
		}
		if err := operationRecordService.CreateSysOperationRecord(record); err != nil {
			global.BLOG_LOG.Error("操作记录创建失败：", zap.Error(err))
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func getHeaderRecordLen(c *gin.Context) bool {
	writerHeader := c.Writer.Header()
	if strings.Contains(writerHeader.Get("Pragma"), "public") ||
		strings.Contains(writerHeader.Get("Expires"), "0") ||
		strings.Contains(writerHeader.Get("Cache-Control"), "must-revalidate, post-check=0, pre-check=0") ||
		strings.Contains(writerHeader.Get("Content-Type"), "application/force-download") ||
		strings.Contains(writerHeader.Get("Content-Type"), "application/octet-stream") ||
		strings.Contains(writerHeader.Get("Content-Type"), "application/vnd.ms-excel") ||
		strings.Contains(writerHeader.Get("Content-Type"), "application/download") ||
		strings.Contains(writerHeader.Get("Content-Disposition"), "attachment") ||
		strings.Contains(writerHeader.Get("Content-Transfer-Encoding"), "binary") {
		return true
	}
	return false
}
