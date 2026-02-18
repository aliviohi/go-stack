# Versioning and Release Practice

This repository uses Semantic Versioning (`MAJOR.MINOR.PATCH`):

- `MAJOR`: breaking changes
- `MINOR`: backward-compatible features
- `PATCH`: backward-compatible fixes

## Suggested release workflow

1. Merge changes to `main` using Conventional Commit style (`feat:`, `fix:`, `chore:`).
2. GitHub Action `release-please` opens or updates a Release PR.
3. Review the Release PR:
- Version bump
- Generated `CHANGELOG.md`
4. Merge the Release PR.
5. `release-please` creates the git tag and GitHub Release automatically.

`CHANGELOG.md` is managed by `release-please`; do not edit release sections manually.

## How to choose the next version

- New feature: bump `MINOR` (example: `v0.2.0`).
- Bug fix only: bump `PATCH` (example: `v0.1.1`).
- Breaking change: bump `MAJOR` (example: `v1.0.0`).
