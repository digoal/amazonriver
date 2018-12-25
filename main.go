/*
 * Copyright 2018 Shanghai Junzheng Network Technology Co.,Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *	   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"flag"
	"io/ioutil"

	"github.com/hellobike/amazonriver/conf"
	"github.com/hellobike/amazonriver/river"

	"github.com/json-iterator/go"
)

var configfile = flag.String("config", "", "config")

func main() {
	flag.Parse()

	data, err := ioutil.ReadFile(*configfile)
	if err != nil {
		panic(err)
	}

	var config conf.Conf
	if err := jsoniter.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	amazon := river.New(&config)
	amazon.Start()

	// block forever
	select {}
}
