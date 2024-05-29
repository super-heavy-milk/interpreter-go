test:
    #!/usr/bin/env bash
    # from https://www.strangeleaflet.com/blog/better-go-test-output
    go test '{{ justfile_directory() }}/...' -v -json |
        jq -r '.Output // ""' |
        rg -v "=== RUN" |
        sed -e '/^$/d' -e "s/--- PASS/✅/" -e "s/--- FAIL/❌/"
