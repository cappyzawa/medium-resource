# medium-resource
[![Go Report Card](https://goreportcard.com/badge/cappyzawa/medium-resource)](https://goreportcard.com/report/cappyzawa/medium-resource)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Build Status](https://concourse.ik.am/api/v1/teams/cappyzawa/pipelines/medium-resource/jobs/test-master/badge)](https://concourse.ik.am:14161/teams/cappyzawa/pipelines/medium-resource)
[![Docker Pulls](https://img.shields.io/docker/pulls/cappyzawa/medium-resource.svg)](https://hub.docker.com/r/cappyzawa/medium-resource)

A concourse resource for positing a story to medium.  

This resource can post **one story**, so you should prepare preprocessing task(e.g. [tasks/prepare.yml](https://github.com/cappyzawa/medium/blob/master/ci/tasks/prepare.yml))

## Source Configuration
* `access_token`: Required. You can get the access token from [Self\-issued access tokens](https://github.com/Medium/medium-api-docs#22-self-issued-access-tokens)

## Behavior
### `check` & `in` : none
[Medium API](https://github.com/Medium/medium-api-docs#32-publications) does not get stories.  
[starkandwayne/rss\-resource: A Concourse resource to grab RSS feeds and their contents](https://github.com/starkandwayne/rss-resource) can get stories by using rss.  

### `out`: Post an article.
Posts an article to medium based on parameters.
#### Parameters
* `content_file`: Required. This Resource posts an article based on specified file(Markdown or HTML).
* `tags`: Optional. You can set tags as array.
* `canonical_url`: Optional.
* `status`: Optional. Default `draft`.
* `licence`: Optional.

## Example
```yaml
resource_types:
  - name: medium
    type: registry-image
    source:
      repository: cappyzawa/medium-resource
      tag: latest
resources:
  - name: blog-repo
    type: git
    source:
      uri: https://github.com/cappyzawa/medium
      paths:
      - docs/*
  - name: blog
    type: medium
    source:
      access_token: ((medium-token))
jobs:
  - name: publish-to-medium
    build_logs_to_retain: 10
    plan:
      - get: blog-repo
        trigger: true
      - task: prepare
        file: blog-repo/ci/tasks/prepare.yml
      - put: blog
        params:
          content_file: completed/latest.md
```
