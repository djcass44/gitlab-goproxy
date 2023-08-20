/*
 *    Copyright 2023 Django Cass
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 *
 */

package main

import (
	"context"
	"github.com/djcass44/gitlab-goproxy/internal/cache"
	"github.com/djcass44/go-utils/logging"
	"github.com/goproxy/goproxy"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"github.com/xanzy/go-gitlab"
	"gitlab.com/autokubeops/serverless"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type environment struct {
	Port     int `envconfig:"PORT" default:"8080"`
	LogLevel int `split_words:"true"`
	GitLab   struct {
		URL       string `split_words:"true" required:"true"`
		Token     string `split_words:"true" required:"true"`
		ProjectID int    `split_words:"true" required:"true"`
	}
}

func main() {
	var e environment
	envconfig.MustProcess("app", &e)

	// configure logging
	zc := zap.NewProductionConfig()
	zc.Level = zap.NewAtomicLevelAt(zapcore.Level(e.LogLevel * -1))

	log, _ := logging.NewZap(context.Background(), zc)

	git, err := gitlab.NewClient(e.GitLab.Token, gitlab.WithBaseURL(e.GitLab.URL))
	if err != nil {
		log.Error(err, "failed to create gitlab client")
		os.Exit(1)
	}

	router := mux.NewRouter()
	router.Use(logging.Middleware(log))

	router.PathPrefix("/").Handler(&goproxy.Goproxy{
		Cacher:        cache.NewGitLabCache(git, e.GitLab.ProjectID),
		ProxiedSUMDBs: nil,
	})

	serverless.NewBuilder(router).
		WithPort(e.Port).
		WithLogger(log).
		Run()
}
