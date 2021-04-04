# git-fork
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->
[![Maintainability](https://api.codeclimate.com/v1/badges/badf947964bf48e30fc8/maintainability)](https://codeclimate.com/github/angeliski/git-fork/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/badf947964bf48e30fc8/test_coverage)](https://codeclimate.com/github/angeliski/git-fork/test_coverage)

[WIP] That is a work in progress. Be carefully to use in your projects 

A simple way to manage your forks. You can find some tutorials how to do that, but how about just run a `git-fork sync` and get your repository synchronized?

The `git-fork` is based in a simple idea: You do all your work outside from the main branch, and we sync the main branch when needed 


## Installation
Not done yet. If you would like to try, clone the repo and run `make build` to get a binary

## Develop

You can see the commands run `make help`, here some tips:
- `make quality` -  that command will run all quality checks in your code
- `make mocks` - that command will generate the mocks for all interfaces

## Build Binaries

```shell
docker run --rm --privileged \
  --user=$$(id -u):$$(id -g)
  -v $PWD:/git-fork \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -w /git-fork \
  goreleaser/goreleaser build

```

## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://github.com/angeliski"><img src="https://avatars.githubusercontent.com/u/1574240?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Rogerio Angeliski</b></sub></a><br /><a href="https://github.com/angeliski/git-fork/commits?author=angeliski" title="Documentation">📖</a></td>
  </tr>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!