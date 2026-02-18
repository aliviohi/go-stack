---
name: Release Checklist
about: Track tasks for an upcoming release
title: 'Release: vX.Y.Z'
labels: release
assignees: ''
---

## Version

- [ ] Confirm PRs use Conventional Commits (`feat`, `fix`, `chore`, etc.)
- [ ] Confirm milestone or scope for this release window

## Changelog

- [ ] Verify generated `CHANGELOG.md` entries in the Release PR
- [ ] Ensure important PR titles are clear for release notes

## Quality

- [ ] Run tests
- [ ] Build locally
- [ ] Smoke test container (`docker run --rm go-stack`)

## Release

- [ ] Merge the Release PR created by `release-please`
- [ ] Verify git tag and GitHub Release were created
- [ ] Announce release notes
