package profileManager

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// ProfileConfig define the configuration for the profile manager. The local and global has the same root & name.
type ProfileConfig struct {
	// RootPath is the path where the profile will be saved. eg: .profile | .toolName/profile
	RootPath string
	// DisableSuffix is a flag to disable the suffix for the profile file. _local.json | _global.json
	DisableSuffix bool
}

type ProfileManager struct {
	LocalProfiles  map[string]interface{}
	GlobalProfiles map[string]interface{}
	Config         ProfileConfig
}

func NewProfileManager(config ProfileConfig) *ProfileManager {
	return &ProfileManager{
		LocalProfiles:  make(map[string]interface{}),
		GlobalProfiles: make(map[string]interface{}),
		Config:         config,
	}
}

func (pm *ProfileManager) SaveProfile(name string, profile interface{}, global bool) error {
	if global {
		pm.GlobalProfiles[name] = profile
		return pm.saveGlobally(name, profile)
	} else {
		pm.LocalProfiles[name] = profile
		return pm.saveLocally(name, profile)
	}
}

func (pm *ProfileManager) saveLocally(name string, profile interface{}) error {
	data, err := json.Marshal(profile)
	if err != nil {
		return err
	}
	fileName := name + ".json"
	if !pm.Config.DisableSuffix {
		fileName = fmt.Sprintf("%s_local.json", name)
	}
	filePath := filepath.Join(pm.Config.RootPath, fileName)
	if err := os.MkdirAll(pm.Config.RootPath, os.ModePerm); err != nil {
		return err
	}
	pm.LocalProfiles[name] = profile
	return os.WriteFile(filePath, data, 0644)
}

func (pm *ProfileManager) saveGlobally(name string, profile interface{}) error {
	data, err := json.Marshal(profile)
	if err != nil {
		return err
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	fileName := name + ".json"
	if !pm.Config.DisableSuffix {
		fileName = fmt.Sprintf("%s_global.json", name)
	}
	filePath := filepath.Join(homeDir, fileName)
	if err := os.MkdirAll(homeDir, os.ModePerm); err != nil {
		return err
	}
	pm.GlobalProfiles[name] = profile
	return os.WriteFile(filePath, data, 0644)
}

func (pm *ProfileManager) LoadProfile(name string, global bool) (interface{}, error) {
	if global {
		if profile, exists := pm.GlobalProfiles[name]; exists {
			return profile, nil
		}
		return pm.loadGlobally(name)
	} else {
		if profile, exists := pm.LocalProfiles[name]; exists {
			return profile, nil
		}
		return pm.loadLocally(name)
	}
}

func (pm *ProfileManager) loadLocally(name string) (interface{}, error) {
	fileName := name
	if !pm.Config.DisableSuffix {
		fileName = fmt.Sprintf("%s_local.json", name)
	}
	filePath := filepath.Join(pm.Config.RootPath, fileName)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var profile interface{}
	err = json.Unmarshal(data, &profile)
	if err != nil {
		return nil, err
	}
	pm.LocalProfiles[name] = profile
	return profile, nil
}

func (pm *ProfileManager) loadGlobally(name string) (interface{}, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	fileName := name + ".json"
	if !pm.Config.DisableSuffix {
		fileName = fmt.Sprintf("%s_global.json", name)
	}
	filePath := filepath.Join(homeDir, fileName)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var profile interface{}
	err = json.Unmarshal(data, &profile)
	if err != nil {
		return nil, err
	}
	pm.GlobalProfiles[name] = profile
	return profile, nil
}

func (pm *ProfileManager) DeleteProfile(name string, global bool) error {
	if global {
		delete(pm.GlobalProfiles, name)
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		fileName := name + ".json"
		if !pm.Config.DisableSuffix {
			fileName = fmt.Sprintf("%s_global.json", name)
		}
		filePath := filepath.Join(homeDir, fileName)
		return os.Remove(filePath)
	} else {
		delete(pm.LocalProfiles, name)
		fileName := name + ".json"
		if !pm.Config.DisableSuffix {
			fileName = fmt.Sprintf("%s_local.json", name)
		}
		filePath := filepath.Join(pm.Config.RootPath, fileName)
		return os.Remove(filePath)
	}
}

// JsonFileToProfile reads a json file and returns the profile
func (pm *ProfileManager) JsonFileToProfile(filepath string) (interface{}, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var profile interface{}
	err = json.Unmarshal(file, &profile)
	if err != nil {
		return nil, err
	}
	return profile, nil
}
