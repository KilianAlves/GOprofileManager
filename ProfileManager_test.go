package profileManager

import (
	"fmt"
	_ "github.com/stretchr/testify/assert"
	"testing"
)

type Profile struct {
	Name string
}

var ProfileConfigClassic = ProfileConfig{
	RootPath:      ".FolderTest",
	DisableSuffix: false,
}

func TestProfileManager_SaveProfile_Local(t *testing.T) {
	// Create a new profile manager
	profileManager := NewProfileManager(ProfileConfigClassic)

	// Create a new profile
	profile := Profile{
		Name: "Test",
	}
	fmt.Println("hezsbezfbhj")
	// Add the profile to the profile manager
	err := profileManager.SaveProfile("Profile1", profile, false)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	// Check if the profile has been added to the local profiles
	if _, ok := profileManager.LocalProfiles["Profile1"]; !ok {
		t.Errorf("The profile has not been added to the local profiles")
	}
}

func TestProfileManager_SaveProfile_Global(t *testing.T) {
	// Create a new profile manager
	profileManager := NewProfileManager(ProfileConfigClassic)

	// Create a new profile
	profile := Profile{
		Name: "Test",
	}

	// Add the profile to the profile manager
	err := profileManager.SaveProfile("Profile1", profile, true)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	// Check if the profile has been added to the local profiles
	if _, ok := profileManager.GlobalProfiles["Profile1"]; !ok {
		t.Errorf("The profile has not been added to the local profiles")
	}
}

func TestProfileManager_SaveProfile_Name(t *testing.T) {
	//suffix is disabled
	NewProfileManager(ProfileConfig{
		RootPath:      ".FolderTest",
		DisableSuffix: true,
	})
}

func TestProfileManager_LoadProfile(t *testing.T) {

}

// TODO: Add test for : getProfile(), deleteProfile() for local and global profiles
// Add test for testing the ProfileConfig (the folder generation and the file creation)
