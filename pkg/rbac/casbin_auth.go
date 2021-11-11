package rbac

import (
	"fmt"
	casbin "github.com/casbin/casbin/v2"
	mda "github.com/casbin/mongodb-adapter/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	c "saiga/config"
)

//func ExampleOfCasAuth() {
//	credential := options.Credential{
//		Username: c.Configure().MongoUserName,
//		Password: c.Configure().MongoPassword,
//	}
//	mongoClientoption := options.Client().ApplyURI(c.Configure().MongoURL).SetAuth(credential)
//	a, err := mda.NewAdapterWithClientOption(mongoClientoption, "casbin_rule")
//	if err != nil {
//		fmt.Println(err)
//	}
//	e, err := casbin.NewEnforcer("/home/kcan/go/src/cci-egde-client/configs/rbac_model.conf", a)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	e.LoadPolicy()
//
//	// Policy User
//	e.AddPolicy("user", "data", "GET")
//
//	// Policy Admin
//	e.AddPolicy("admin", "data", "GET")
//	e.AddPolicy("admin", "data", "POST")
//
//	// AddGroupingPolicy
//	e.AddGroupingPolicy("alice", "admin")
//	e.AddGroupingPolicy("bob", "admin")
//
//	res, _ := e.Enforce("bob", "data", "GET")
//	fmt.Println(res)
//}

// Create object oriented interface example
type CasbinEnforcer struct {
	e *casbin.Enforcer
}

// NewCasbinEnforcer
func NewCasEnforcer() *CasbinEnforcer {
	credential := options.Credential{
		Username: c.Configure().MongoUserName,
		Password: c.Configure().MongoPassword,
	}
	mongoOption := options.Client().ApplyURI(c.Configure().MongoURL).SetAuth(credential)
	a, err := mda.NewAdapterWithClientOption(mongoOption, "casbin_rule")
	if err != nil {
		fmt.Println(err)
	}
	path, _ := os.Getwd()
	e, casbinErr := casbin.NewEnforcer(path+"/configs/rbac_model.conf", a)
	if casbinErr != nil {
		fmt.Println(casbinErr)
	}
	errPolicy := e.LoadPolicy()
	if errPolicy != nil {
		return nil
	}
	return &CasbinEnforcer{e: e}
}

// AddPolicy
func (c *CasbinEnforcer) AddPolicy(sub string, obj string, act string) bool {
	policy, err := c.e.AddPolicy(sub, obj, act)
	if err != nil {
		return false
	}
	return policy
}

// AddGroupingPolicy
func (c *CasbinEnforcer) AddGroupingPolicy(sub string, obj string) bool {
	policy, err := c.e.AddGroupingPolicy(sub, obj)
	if err != nil {
		return false
	}
	return policy
}

// CheckPermission
func (c *CasbinEnforcer) CheckPermission(sub string, obj string, act string) (bool, error) {
	return c.e.Enforce(sub, obj, act)
}
