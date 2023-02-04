package fs

import "github.com/kc-workspace/go-lib/mapper"

// Build single 'file or directory' full path map
func BuildSFpMap(fp string) mapper.Mapper {
	return mapper.New().
		Set("mode", string(SINGLE)).
		Set("type", string(AUTO)).
		Set("fullpath", fp)
}

// Build single 'file or directory' path list map
func BuildSPMap(paths ...string) mapper.Mapper {
	return mapper.New().
		Set("mode", string(SINGLE)).
		Set("type", string(AUTO)).
		Set("paths", paths)
}

// Build single file full path map
func BuildSfFpMap(fp string) mapper.Mapper {
	return mapper.New().
		Set("mode", string(SINGLE)).
		Set("type", string(FILE)).
		Set("fullpath", fp)
}

// Build single file path list map
func BuildSfPMap(paths ...string) mapper.Mapper {
	return mapper.New().
		Set("mode", string(SINGLE)).
		Set("type", string(FILE)).
		Set("paths", paths)
}

// Build single file full path map
func BuildSdFpMap(fp string) mapper.Mapper {
	return mapper.New().
		Set("mode", string(SINGLE)).
		Set("type", string(DIRECTORY)).
		Set("fullpath", fp)
}

// Build single file path list map
func BuildSdPMap(paths ...string) mapper.Mapper {
	return mapper.New().
		Set("mode", string(SINGLE)).
		Set("type", string(DIRECTORY)).
		Set("paths", paths)
}

// multiple mode

// Build multiple 'file or directory' full path map
func BuildMFpMap(fp string) mapper.Mapper {
	return mapper.New().
		Set("mode", string(MULTIPLE)).
		Set("type", string(AUTO)).
		Set("fullpath", fp)
}

// Build multiple 'file or directory' path list map
func BuildMPMap(paths ...[]string) mapper.Mapper {
	return mapper.New().
		Set("mode", string(MULTIPLE)).
		Set("type", string(AUTO)).
		Set("paths", paths)
}

// Build multiple file full path map
func BuildMfFpMap(fp string) mapper.Mapper {
	return mapper.New().
		Set("mode", string(MULTIPLE)).
		Set("type", string(FILE)).
		Set("fullpath", fp)
}

// Build multiple file path list map
func BuildMfPMap(paths ...[]string) mapper.Mapper {
	return mapper.New().
		Set("mode", string(MULTIPLE)).
		Set("type", string(FILE)).
		Set("paths", paths)
}

// Build multiple directory full path map
func BuildMdFpMap(fp string) mapper.Mapper {
	return mapper.New().
		Set("mode", string(MULTIPLE)).
		Set("type", string(DIRECTORY)).
		Set("fullpath", fp)
}

// Build multiple directory path list map
func BuildMdPMap(paths ...[]string) mapper.Mapper {
	return mapper.New().
		Set("mode", string(MULTIPLE)).
		Set("type", string(DIRECTORY)).
		Set("paths", paths)
}
