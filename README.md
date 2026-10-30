# Resume Theme Simple

## Introduction

This is a simple theme for [Resume CLI](https://github.com/ansidev/resume-cli) which was built with [Bootstrap](https://getbootstrap.com).

## Getting started

### Installation

Download the latest `resume` release from [ansidev/resume-cli](https://github.com/ansidev/resume-cli/releases).

### How to use?

Create a `resume.config.yaml` in your project. Choose `local` if you have cloned this theme, or `remote` to let Resume CLI fetch it automatically:

```yaml
theme:
  type: local                                        # or: remote
  location: path/to/resume-theme-simple               # or: https://github.com/ansidev/resume-theme-simple.git
resume:
  basics:
    name: Your Name
    # ... see demo/resume.config.yaml for the full structure
```

Then generate your resume:

```bash
# HTML
resume generate --format html --output dist/index.html

# PDF
resume generate --format pdf --output dist/static/resume.pdf
```

See [`demo/resume.config.yaml`](demo/resume.config.yaml) for a complete working example.

> **Remote helpers:** when `type: remote` is set, Resume CLI refuses to run `helpers.go` unless you
> review the file and pass `--allow-remote-helpers`:
>
> ```bash
> resume generate --allow-remote-helpers
> ```
>
> `local` themes run helpers automatically without the flag.

The `--theme` option overrides `theme.location`. The template engine and helpers are declared by this theme in [`theme.yaml`](/theme.yaml); they are not set in the Resume CLI config.

### Demo

The demo configuration is a representative example for the current Resume CLI theme contract. Run it from the repository root:

```bash
# HTML
resume generate --config demo/resume.config.yaml --format html --output dist/index.html

# PDF
resume generate --config demo/resume.config.yaml --format pdf --output dist/static/resume.pdf
```

### Helper functions

[`helpers.go`](/helpers.go) is interpreted by Resume CLI at runtime (with [yaegi](https://github.com/traefik/yaegi)).
Its `Funcs() map[string]interface{}` function returns the template helpers: `bootstrapIcon`, `joinArray`,
`markdown`, `displayUrl`, `toLowerCase` and `year`. The file must stay self-contained (single file, standard library,
`github.com/mailgun/raymond/v2` and `github.com/yuin/goldmark` only).

> **Security:** helper code runs with the full privileges of the CLI process. It runs automatically for `local` themes. For `remote` themes, pass `--allow-remote-helpers` as described above.

### Theme structure

```
|-- theme.yaml: theme manifest (name, templateEngine, helpers)
|-- resources: view files
|   |-- assets
|   |   |-- css
|   |   |   |-- resume.css: Resume CSS source.
|   |   |   |-- styles.css: Generated file. DO NOT EDIT.
|   |   |-- js
|   |       |-- scripts.js: Optional front-end script bundle.
|   |-- partials: partial files
|   |-- resume.hbs: Main resume template.
|-- helpers.go: helper functions (entry point: Funcs)
|-- helpers_test.go
|-- go.mod
|-- go.sum
```

## Contact

Le Minh Tri [@ansidev](https://ansidev.xyz/about).

## License

This source code is available under the [AGPL-3.0 LICENSE](/LICENSE).
