package response

import (
	"github.com/aeone1/rotech-post-comment/internal/protocols/http/errors"
	"github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

// TODO Follow JSON[API] recommendations
type Response struct {
	Message *string      `json:"message,omitempty"`
	Data    *interface{} `json:"data,omitempty"`
}

func Json(c *gin.Context, httpCode int, message string, data interface{}) {
	c.Header("Content-Type", "application/json")
	res := Response{
		Message: &message,
		Data:    &data,
	}
	c.JSON(httpCode, res)
}

func Text(c *gin.Context, httpCode int, message string) {
	c.Header("Content-Type", "text/plain")
	c.String(httpCode, message)
}

// TODO: implement response error
func Err(c *gin.Context, err error) {
	const foreignKeyViolationErrorCode = pq.ErrorCode("23503")
	foreignKeyViolationErrorMessage := "ForeignKey Violation. Please check all external links/ids (e.g. post_id)"
	pqErrorMessage := "Something went wrong at the DB level"
	c.Header("Content-Type", "application/json")
	if pgErr, isPGErr := err.(*pq.Error); isPGErr {
		if pgErr.Code == foreignKeyViolationErrorCode {
			res := Response{
				Message: &foreignKeyViolationErrorMessage,
			}
			c.JSON(403, res)
			return
		}
		res := Response{
			Message: &pqErrorMessage,
		}
		c.JSON(500, res)
		return
	}
	_, ok := err.(*errors.RespError)
	if !ok {
		err = errors.InternalServerError(err.Error())
	}

	er, _ := err.(*errors.RespError)
	
	res := Response{
		Message: &er.Message,
	}
	c.JSON(er.Code, res)
}