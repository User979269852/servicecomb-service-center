/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package servicecenter

import (
	"context"

	pb "github.com/apache/servicecomb-service-center/server/core/proto"
)

// CreateSchemas Create schemas to servicecenter
func (c *Client) CreateSchemas(ctx context.Context, domainProject, serviceId string, schemas []*pb.Schema) error {
	err := c.cli.CreateSchemas(ctx, domainProject, serviceId, schemas)
	if err != nil {
		return err
	}
	return nil
}

// GetSchemasByServiceId Get schemas by serviceId from servicecenter
func (c *Client) GetSchemasByServiceId(ctx context.Context, domainProject, serviceId string) ([]*pb.Schema, error) {
	schemas, err := c.cli.GetSchemasByServiceId(ctx, domainProject, serviceId)
	if err != nil {
		return nil, err
	}
	return schemas, nil
}
