package main

import (
	"fmt"
	"github.com/sethvargo/go-githubactions"
	"strconv"
	"strings"
)

func main() {
	imageTagPrefix := githubactions.GetInput("image_tag_prefix")
	imageTagSuffix := githubactions.GetInput("image_tag_suffix")
	isAddShortCommitHashToImageTag, _ := strconv.ParseBool(githubactions.GetInput("is_add_short_commit_hash"))

	context, _ := githubactions.Context()
	var shortCommitHash = context.SHA[(len(context.SHA) - 8):]
	var imageTag string

	if imageTagPrefix != "" {
		imageTag += imageTagPrefix
	}

	if context.RefType == "tag" {
		imageTag += context.RefName
	} else {
		imageTag += strings.Replace(context.RefName, "/", "_", -1)
	}

	if isAddShortCommitHashToImageTag {
		imageTag += fmt.Sprintf("-%s", shortCommitHash)
	}

	if imageTagSuffix != "" {
		imageTag += imageTagSuffix
	}

	githubactions.SetEnv("SHORT_COMMIT_HASH", shortCommitHash)
	githubactions.SetEnv("IMAGE_TAG", imageTag)
}
