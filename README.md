# Word Search System
## Introduction
This repository is the source for the word_search_system component of the word_search stack. This component manages the words list and exposes 3 methods for manipulating and quering this words list:
1. SearchWord - takes a keyword as a parameter and searches through the words list, returning possible matches.
2. AddWords - takes a list of words as a parameter and adds them to the words list.
3. GetTop5KeyWords - returns the top 5 most searched keywords

The methods can be called over gRPC and the gRPC proto buffers are defined at https://github.com/chrisjpalmer/word_search_system_grpc, which is a dependency of this repo.

## Docker Building
### Clone CLI Tool and Install it
```sh
git clone https://github.com/chrisjpalmer/word_search_cli && cd word_search_cli && npm link
cd /my/blank/proj/dir #specify a blank project directory
word_search_cli init #initializes a new word_search_proj in your current directory
```

### Prepare Docker Machine
```sh
eval $(docker-machine env default)
```

### Build source via the CLI
```sh
word_search_cli build --build-repo-tag=1.0.0 --src-repo-tag=1.0.0 word_search_system #see https://github.com/chrisjpalmer/word_search_system for more tags
```

The correct build-repo-tag needs to be used for building this source. You can find the compatible build-repo-tag at https://github.com/chrisjpalmer/word_search_system_builder

## Debugging / Local Run
```sh
cp config.template.json config.json
go run github.com/chrisjpalmer/word_search_system
```

## GRPC Protocol Buffers
This table shows which version of `word_search_system_grpc` is used in each version of `word_search_system`.
This can be used to percieve incompatibilities between versions of `word_search_system` and `word_search_api`.

Where `word_search_system` and `word_search_api` use different major versions of the `word_search_system_grpc`, they are not compatible.
| word_search_system version | word_search_system_grpc version |
| -------------- | ------------------------ |
| 1.0.0 | 1.0.0 |
