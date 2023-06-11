// Copyright Tharsis Labs Ltd.(Black)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/black/black/blob/main/LICENSE)
package upgrade

import (
	"fmt"
)

// CreateModuleQueryExec creates a Black module query
func (m *Manager) CreateModuleQueryExec(moduleName, subCommand, chainID string) (string, error) {
	cmd := []string{
		"black",
		"q",
		moduleName,
		subCommand,
		fmt.Sprintf("--chain-id=%s", chainID),
		"--keyring-backend=test",
		"--log_format=json",
	}
	return m.CreateExec(cmd, m.ContainerID())
}
