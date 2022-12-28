package configs

import (
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
		output, err := mapper.FromJson(content)
		if err != nil {
			return result, err
		}

		// merge result together
		result = mapper.Merger(result).Add(output).SetConfig(strategy).Merge()
	}

	return result, nil
}
