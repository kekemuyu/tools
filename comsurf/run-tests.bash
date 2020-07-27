#!/bin/bash
file=$(mktemp)
trap "rm $file" EXIT

bash tests.bash > "$file"
vimdiff <(cat "$file" \
    | perl -pe 's/^Date: \w{3}, \d{2} \w{3} \d{4} \d\d:\d\d:\d\d UTC/Date: [FILTERED BY TEST SCRIPT]/' \
    ) test-fixtures/results.txt
