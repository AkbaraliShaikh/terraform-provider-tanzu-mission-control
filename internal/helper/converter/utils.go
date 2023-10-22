/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package converter

import (
	"fmt"
	"strings"
)

func BuildModelPath(modelPathSeparator string, paths ...string) string {
	return strings.Join(paths, modelPathSeparator)
}

func BuildArrayField(field string) string {
	return fmt.Sprintf("%s%s", field, ArrayFieldMarker)
}
