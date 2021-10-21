/*
 * SPDX-License-Identifier: BSD-2-Clause
 *
 * Copyright (c) 2021, Lewis Cook <lcook@FreeBSD.org>
 * All rights reserved.
 */
package hookrelay

/*
 * Optional middleware a hook can use for convenience.
 */
const (
	/*
	 * Check whether the incoming method is `POST`.
	 */
	OptionCheckMethod byte = 1 << iota
	/*
	 * Check whether the application type sent is `application/json`.
	 */
	OptionCheckType
	/*
	 * Reasonable defaults for webhook listening.
	 */
	DefaultOptions = (OptionCheckMethod | OptionCheckType)
)
