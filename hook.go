/*
 * SPDX-License-Identifier: BSD-2-Clause
 *
 * Copyright (c) 2021, Lewis Cook <lcook@FreeBSD.org>
 * All rights reserved.
 */
package hookrelay

import "net/http"

type Hook interface {
	Response(i interface{}) func(w http.ResponseWriter, r *http.Request)
	LoadConfig(config string) error
	Endpoint() string
	Options() byte
}
