package middleware

import (
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"

	"github.com/gin-gonic/gin"
)

// Authorize determines if current user has been authorized to take an action on an object.
func Authorize(sub string, dom string, obj string, act string, enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get current user/subject
		sub, existed := c.Get("username")
		// Get Dom user/domain
		if !existed {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not Valid Token"})
			return
		}

		// Load policy from Database
		err := enforcer.LoadPolicy()
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"msg": "Failed to load policy from DB"})
			return
		}

		// Casbin enforces policy
		ok, err := enforcer.Enforce(fmt.Sprint(sub), dom, obj, act)

		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"msg": "Error occurred when authorizing user"})
			return
		}

		if !ok {
			//c.AbortWithStatusJSON(403, gin.H{"msg": "You are not authorized / Kamu tidak memiliki akses", "Sub": sub, "Dom": dom, "Obj": obj, "Act": act})
			c.AbortWithStatusJSON(http.StatusForbidden, "You are not authorized")
			return
		}
		c.Next()
	}
}
