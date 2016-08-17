hekaanom
========

[![GoDoc](https://godoc.org/github.com/berkmancenter/hekaanom?status.png)](https://godoc.org/github.com/berkmancenter/hekaanom)

hekaanom is a Go library implementing anomaly detection in time series data as a filter plugin for the [Heka data processing tool](https://hekad.readthedocs.org).

*Note:* Mozilla [has stated](https://mail.mozilla.org/pipermail/heka/2016-May/001059.html) that they intend to stop maintaining Heka. This filter will continue being a useful first pass to look for anomalies in data, but it should not be used as a long-term production tool.

### Getting started

API documentation is available via [godoc](https://godoc.org/github.com/berkmancenter/hekaanom).

### License

Copyright 2016 President and Fellows of Harvard College

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.