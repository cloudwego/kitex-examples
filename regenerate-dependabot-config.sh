#!/bin/bash

function gen_dependabot_config () {
  dependabot_config=".github/dependabot.yml"

  mkdir -p .github
  echo -e "version: 2\nupdates:" > $dependabot_config

  find . -type d -not -path '*/\.*' | while read -r dir; do
    if [ -f "$dir/go.mod" ]; then
      dir="${dir/#.//}"
      dir="${dir/#\/\///}"
      {
        echo -e "  - package-ecosystem: gomod"
        echo -e "    directory: $dir"
        echo -e "    schedule:"
        echo -e "      interval: weekly"
        echo -e "    allow:"
        echo -e "      - dependency-name: github.com/cloudwego/kitex"
      } >> $dependabot_config
    fi
  done
}

gen_dependabot_config