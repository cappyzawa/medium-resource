# medium-resource
[![Go Report Card](https://goreportcard.com/badge/cappyzawa/medium-resource)](https://goreportcard.com/report/cappyzawa/medium-resource)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://concourse.ik.am:14161/api/v1/teams/cappyzawa/pipelines/medium-resource/jobs/test-master/badge)](https://concourse.ik.am:14161/teams/cappyzawa/pipelines/medium-resource)                                                                                                           

concourse resource for positing an article to medium

## Source Configuration
* `access_token`: Required. You can get the access token from [Self\-issued access tokens](https://github.com/Medium/medium-api-docs#22-self-issued-access-tokens)

## Behavior
### `check`: none
### `in`: none
### `out`: Post an article.
Posts an article to medium based on parameters.
#### Parameters
* `content_file`: Required. This Resource posts an article based on specified file.
* `format`: Optional. Default `markdown`. if you want to use `html`, please set it.
* `title`: Optional. If this parameter does not set, first-line of the `content_file` is used as title.
* `tags`: Optional. You can set tags as array.
* `canonical_url`: Optional.
* `status`: Optional. Default `draft`.
* `licence`: Optional.

## Example
```yaml
resource_types:
- name: medium
  type: docker-image
  source:
    repository: cappyzawa/medium-resource
    tag: latest
resources:
- name: blog-repo
  type: git
  source:
    uri: https://github.com/cappyzawa/medium
- name: blog
  type: medium
  source:
    access_token: <your access token>
jobs:
- name: publish-to-medium
  plan:
    - get: blog-repo
    - put: blog
      params:
        content_file: blog-repo/docs/foo.md
        tags:
        - concourse
        - golang
        status: public
```
