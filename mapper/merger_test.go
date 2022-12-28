package mapper_test

import (
	"testing"

	"github.com/kc-workspace/go-lib/mapper"
	"github.com/kc-workspace/go-lib/xtests"
)

var base = mapper.New().
	Set("a", mapper.New().
		Set("aa", "aa").
		Set("bb", 15).
		Set("cc", []int{1, 2, 3}),
	).
	Set("b", "b")

var next = mapper.New().
	Set("a", mapper.New().
		Set("bb", 100).
		Set("cc", []string{"4"}).
		Set("dd", false),
	).
	Set("b", "c").
	Set("c", "d")

func TestMerger(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("normal merge").
		WithExpected(mapper.New().
			Set("a", map[string]interface{}{
				"aa": "aa",
				"bb": 100,
				"cc": []interface{}{1, 2, 3, "4"},
				"dd": false,
			}).
			Set("b", "c").
			Set("c", "d")).
		WithActual(mapper.Merger(base).Add(next).Merge()).
		MustDeepEqual()

	assertion.NewName("merge with config string").
		WithExpected(mapper.New().
			Set("a", mapper.New().
				Set("bb", 100).
				Set("cc", []string{"4"}).
				Set("dd", false)).
			Set("b", "c").
			Set("c", "d")).
		WithActual(mapper.Merger(base).Add(next).SetConfigValue("a", mapper.MERGER_OVERRIDE).Merge()).
		MustDeepEqual()

	assertion.NewName("merge with config").
		WithExpected(mapper.New().
			Set("a", map[string]interface{}{
				"aa": "aa",
				"bb": 100,
				"cc": []interface{}{"4"},
				"dd": false,
			}).
			Set("b", "c").
			Set("c", "d")).
		WithActual(mapper.Merger(base).Add(next).SetConfig(mapper.New().Set("a.cc", mapper.MERGER_OVERRIDE)).Merge()).
		MustDeepEqual()
}
