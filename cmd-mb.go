/*
 * Mini Copy, (C) 2014,2015 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"github.com/minio-io/cli"
	"github.com/minio-io/mc/pkg/client"
	"github.com/minio-io/mc/pkg/console"
	"github.com/minio-io/minio/pkg/iodine"
	"github.com/minio-io/minio/pkg/utils/log"
)

// doMakeBucketCmd creates a new bucket
func doMakeBucketCmd(ctx *cli.Context) {
	if len(ctx.Args()) < 1 {
		cli.ShowCommandHelpAndExit(ctx, "mb", 1) // last argument is exit code
	}
	config, err := getMcConfig()
	if err != nil {
		log.Debug.Println(iodine.New(err, nil))
		console.Fatalln("mc: Unable to read config")
	}
	for _, arg := range ctx.Args() {
		u, err := parseURL(arg, config.GetMapString("Aliases"))
		if err != nil {
			log.Debug.Println(iodine.New(err, nil))
			console.Fatalf("mc: Unable to parse URL[%s]\n", u)
		}
		bucket, _, err := url2Object(u)
		if err != nil {
			log.Debug.Println(iodine.New(err, nil))
			console.Fatalf("mc: Unable to decode bucket and object from URL[%s]\n", u)
		}

		// this is handled differently since http based URLs cannot have
		// nested directories as buckets, buckets are a unique alphanumeric
		// name having subdirectories is only supported for fsClient
		if getURLType(u) != urlFS {
			if !client.IsValidBucketName(bucket) {
				log.Debug.Println(iodine.New(err, nil))
				console.Fatalf("mc: Invalid bucket name: %s", bucket)
			}
		}

		manager := mcClientManager{}
		clnt, err := manager.getNewClient(u, globalDebugFlag)
		if err != nil {
			log.Debug.Println(iodine.New(err, nil))
			console.Fatalf("mc: Unable to create new client to [%s]\n", u)
		}

		// retry - 5 times
		for r := retries.init(); r.retry(); {
			err = clnt.PutBucket(bucket)

			if !isValidRetry(err) {
				break
			}
		}
		if err != nil {
			log.Debug.Println(iodine.New(err, nil))
			// error message returned properly by PutBucket()
			console.Fatalln(err)
		}
	}
}
