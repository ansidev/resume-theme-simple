# Resume Theme Simple

## Introduction

This is a simple theme for [Resume CLI](https://github.com/ansidev/resume-cli) which was built with [Bootstrap](https://getbootstrap.com).

## Getting started

### How to use?

Please checkout an [example project](https://github.com/ansidev/resume-cli-example) to understand how to use this tool.

### Theme structure

```
|-- theme.yaml: Theme manifest for the current Resume CLI.
|-- resources
|   |-- assets
|   |   |-- css
|   |   |   |-- resume.css: Resume CSS source.
|   |   |   |-- styles.css: Generated file. DO NOT EDIT.
|   |   |-- js
|   |       |-- scripts.js: Optional front-end script bundle.
|   |-- partials: partial files.
|   |-- resume.hbs: Main resume template.
|-- helper.go: helper functions.
|-- theme.go
|-- theme_test.go
```

## Contact

Le Minh Tri [@ansidev](https://ansidev.xyz/about).

## License

This source code is available under the [AGPL-3.0 LICENSE](/LICENSE).
