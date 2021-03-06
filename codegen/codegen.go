/*
 *  Copyright (c) 2017, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	mtproto_parser "github.com/nebulaim/mtprotoc/codegen/parser"
	"github.com/golang/glog"
	"flag"
	// "github.com/nebulaim/mtprotoc/codegen/gen/golang"
	// "github.com/nebulaim/mtprotoc/codegen/gen/proto"
	// "github.com/nebulaim/mtprotoc/codegen/gen/golang"
	"github.com/nebulaim/mtprotoc/codegen/gen/proto"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

func main() {
	flag.Parse()

	schemas, err := mtproto_parser.Parse("./schemas/scheme.tl")
	if err != nil {
		glog.Fatal(err)
	}

	// glog.Info(schemas.Layer)
	// glog.Info(schemas.TypeMap)
	genproto.GenProto(schemas, "./")
	// gengolang.GenGolang(schemas, "./")
}
