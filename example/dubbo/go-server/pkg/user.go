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

package pkg

import (
	"context"
	"fmt"
	tripleConstant "github.com/dubbogo/triple/pkg/common/constant"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
)

type GreeterProvider struct {
	*GreeterProviderBase
}

func NewGreeterProvider() *GreeterProvider {
	return &GreeterProvider{
		GreeterProviderBase: &GreeterProviderBase{},
	}
}

func (s *GreeterProvider) SayHelloStream(svr Greeter_SayHelloStreamServer) error {
	c, err := svr.Recv()
	if err != nil {
		return err
	}
	logger.Infof("Dubbo-go3 GreeterProvider recv 1 user, name = %s\n", c.Name)

	err = svr.Send(&User{
		Name: "hello " + c.Name,
		Age:  18,
		Id:   "123456789",
	})

	if err != nil {
		return err
	}
	return nil
}

func (s *GreeterProvider) SayHello(ctx context.Context, in *HelloRequest) (*User, error) {
	logger.Infof("Dubbo3 GreeterProvider get user name = %s\n" + in.Name)
	fmt.Println("get triple header tri-req-id = ", ctx.Value(tripleConstant.TripleCtxKey(tripleConstant.TripleRequestID)))
	fmt.Println("get triple header tri-service-version = ", ctx.Value(tripleConstant.TripleCtxKey(tripleConstant.TripleServiceVersion)))
	return &User{Name: in.Name, Id: "12345", Age: 21}, nil
}

func (g *GreeterProvider) Reference() string {
	return "greeterImpl"
}
