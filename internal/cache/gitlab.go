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

package cache

import (
	"bytes"
	"context"
	"github.com/djcass44/gitlab-goproxy/internal/parser"
	"github.com/go-logr/logr"
	"github.com/xanzy/go-gitlab"
	"io"
	"net/http"
	"os"
	"strings"
)

func NewGitLabCache(client *gitlab.Client, projectId int, hidePackages bool) *GitLabCache {
	return &GitLabCache{
		client:       client,
		projectId:    projectId,
		hidePackages: hidePackages,
	}
}

func (c *GitLabCache) Get(ctx context.Context, name string) (io.ReadCloser, error) {
	log := logr.FromContextOrDiscard(ctx).WithValues("name", name)
	log.Info("checking for cached module")

	pkg, err := parser.NewPackage(name)
	if err != nil {
		return nil, err
	}

	// try to download the module
	data, resp, err := c.client.GenericPackages.DownloadPackageFile(c.projectId, safeName(pkg.Name), pkg.Version, safeName(pkg.String()))
	if err != nil {
		// if it doesn't exist, tell the goproxy to go download it manually
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			log.Info("module has not been cached")
			return nil, os.ErrNotExist
		}
		log.Error(err, "failed to download package")
		return nil, err
	}

	// return the data
	log.Info("located cached module")
	return io.NopCloser(bytes.NewReader(data)), nil
}

func (c *GitLabCache) Put(ctx context.Context, name string, content io.ReadSeeker) error {
	log := logr.FromContextOrDiscard(ctx).WithValues("name", name, "hidden", c.hidePackages)
	log.Info("uploading module to cache")

	pkg, err := parser.NewPackage(name)
	if err != nil {
		return err
	}

	// upload the given data
	status := gitlab.PackageDefault
	if c.hidePackages {
		status = gitlab.PackageHidden
	}
	_, _, err = c.client.GenericPackages.PublishPackageFile(c.projectId, safeName(pkg.Name), pkg.Version, safeName(pkg.String()), content, &gitlab.PublishPackageFileOptions{
		Status: gitlab.GenericPackageStatus(status),
	})
	if err != nil {
		log.Error(err, "failed to publish package")
		return err
	}
	log.Info("successfully published package")
	return nil
}

func safeName(s string) string {
	s = strings.ReplaceAll(s, "/@v/", "/")
	s = strings.ReplaceAll(s, "/", "-")
	s = strings.ReplaceAll(s, "!", "_")
	return s
}
