# Script Runner

[![Step changelog](https://shields.io/github/v/release/bitrise-steplib/steps-script-runner?include_prereleases&label=changelog&color=blueviolet)](https://github.com/bitrise-steplib/steps-script-runner/releases)

Run scripts from file


<details>
<summary>Description</summary>

Run scripts from file

</details>

## 🧩 Get started

Add this step directly to your workflow in the [Bitrise Workflow Editor](https://devcenter.bitrise.io/steps-and-workflows/steps-and-workflows-index/).

You can also run this step directly with [Bitrise CLI](https://github.com/bitrise-io/bitrise).

## ⚙️ Configuration

<details>
<summary>Inputs</summary>

| Key | Description | Flags | Default |
| --- | --- | --- | --- |
| `file_path` | The script you want to run. | required |  |
| `runner` | The executor to be used for running the script. You can use any binary which is in the PATH (bash/ruby/etc.), multipart commands (e. g. go run), absolute paths (e. g. /bin/sh) or binaries from env (e. g. /usr/bin/env python). You can specify flags as well (e. g. /bin/bash -l). | required | `bash` |
| `working_dir` | This directory will be set as the current working directory for the script. Any relative path in the Script (file_path) will be relative to this directory. |  | `$BITRISE_SOURCE_DIR` |
| `is_debug` | If debug=yes the step will print debug infos about the working dir, tmp file path, exit code, etc. |  | `no` |
</details>

<details>
<summary>Outputs</summary>
There are no outputs defined in this step
</details>

## 🙋 Contributing

We welcome [pull requests](https://github.com/bitrise-steplib/steps-script-runner/pulls) and [issues](https://github.com/bitrise-steplib/steps-script-runner/issues) against this repository.

For pull requests, work on your changes in a forked repository and use the Bitrise CLI to [run step tests locally](https://devcenter.bitrise.io/bitrise-cli/run-your-first-build/).

Learn more about developing steps:

- [Create your own step](https://devcenter.bitrise.io/contributors/create-your-own-step/)
- [Testing your Step](https://devcenter.bitrise.io/contributors/testing-and-versioning-your-steps/)
