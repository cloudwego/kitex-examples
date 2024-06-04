#!/bin/bash

function gen_dependabot_config () {
  dependabot_config=".github/dependabot.yml"

  mkdir -p .github
  {
    echo "version: 2"
    echo "updates:"
    echo "  - package-ecosystem: gomod"
    echo "    schedule:"
    echo "      interval: weekly"
    echo "    allow:"
    echo "      - dependency-name: github.com/cloudwego/kitex"
    echo "    groups:"
    echo "      kitex-dependencies:"
    echo "        patterns:"
    echo "          - github.com/cloudwego/kitex"
    echo "    directories:"
  } > $dependabot_config

  find . -type d -not -path '*/\.*' | sort | while read -r dir; do
    if [ -f "$dir/go.mod" ]; then
      dir="${dir/#.//}"
      dir="${dir/#\/\///}"
      echo "      - $dir" >> $dependabot_config
    fi
  done
}

gen_dependabot_config
