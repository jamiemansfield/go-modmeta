/*
 * Copyright (c) 2024, Jamie Mansfield <jmansfield@cadixdev.org>
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package viutils

func ParseBool(key string) bool {
	return key == "yes" ||
		key == "true" ||
		key == "on" ||
		key == "allow" ||
		key == "grant" ||
		key == "1"
}
