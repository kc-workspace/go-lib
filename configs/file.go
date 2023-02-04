package configs

import (
	"errors"
	"path/filepath"

	"github.com/kc-workspace/go-lib/fs"
	"github.com/kc-workspace/go-lib/mapper"
	"github.com/kc-workspace/go-lib/xtemplates"
)

func LoadConfigFromFileSystem(input []fs.FileSystem, data mapper.Mapper, strategy mapper.Mapper) (mapper.Mapper, error) {
	var result = mapper.New()
	var files, err = fs.ToFiles(input)
	if err != nil {
		return result, err
	}

	for _, file := range files {
		var ext = filepath.Ext(file.Basename())
		// read content
		var content, err = file.Read()
		if err != nil {
			return result, err
		}

		// compile template data only if data is not empty
		// If data is empty, then no point to parse templates
		if !data.IsEmpty() {
			str, err := xtemplates.Text(string(content), data)
			if err != nil {
				return result, err
			}
			content = []byte(str)
		}

		// convert content to mapper
		var output mapper.Mapper
		if ext == ".yaml" || ext == ".yml" {
			output, err = mapper.FromYaml(content)
			if err != nil {
				return result, err
			}
		} else if ext == ".json" || ext == ".json5" {
			output, err = mapper.FromJson(content)
			if err != nil {
				return result, err
			}
		} else {
			return result, errors.New("only yaml,yml,json,json5 are supported to convert")
		}

		// merge result together
		result = mapper.Merger(result).Add(output).SetConfig(strategy).Merge()
	}

	return result, nil
}
