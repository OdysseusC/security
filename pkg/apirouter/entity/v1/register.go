/*
Copyright 2021 The tKeel Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package v1

import (
	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/tkeel-io/security/pkg/apiserver/config"
	"github.com/tkeel-io/security/pkg/apiserver/filters"
	"github.com/tkeel-io/security/pkg/constants"
)

func AddToRestContainer(c *restful.Container, conf *config.EntityConfig) error {
	var webservice *restful.WebService
	for _, v := range c.RegisteredWebServices() {
		if v.RootPath() == "v1" {
			webservice = v
			break
		}
	}
	if webservice == nil {
		webservice = &restful.WebService{}
		webservice.Path("v1").
			Produces(restful.MIME_JSON).
			Filter(filters.Auth())

		c.Add(webservice)
	}

	handler := newEntityHandler(conf)

	webservice.Route(webservice.GET("entity/{entity_type}/{entity_id}/token").
		To(handler.Token).
		Doc("get a entity token").
		Param(webservice.PathParameter("entity_type", "EntityType")).
		Param(webservice.PathParameter("entity_id", "Entity's ID")).
		Param(webservice.QueryParameter("expires_in", "invalid period( seconds )")).
		Metadata(restfulspec.KeyOpenAPITags, []string{constants.APITagEntity}))

	webservice.Route(webservice.POST("entity/token/valid").
		To(handler.TokenValid).
		Doc("valid a entity token").
		Reads(TokenValidIn{}).
		Metadata(restfulspec.KeyOpenAPITags, []string{constants.APITagEntity}))

	return nil
}
