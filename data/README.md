# Translation data

This directory contains source manifests for the deterministic import tooling.
A manifest records what may be downloaded and imported; it does not contain
Bible text.

## Provenance contract

Every translation manifest must include:

- a stable application identifier and publisher identifier;
- language, edition, and canon metadata;
- publisher and rights-notice links;
- the exact source archive URL and member name;
- retrieval date and SHA-256 checksum;
- expected book and verse totals; and
- the first and last expected references.

Import tooling must reject a source when its checksum, structure, identifiers,
or expected totals differ from the manifest. Updating a source is an intentional
repository change: review the publisher's current rights notice, update the
checksum and expectations, regenerate the database, and describe textual
changes in the pull request.

Verse text must remain unchanged after UTF-8 decoding and removal of the source
record separator. Reference identifiers, book ordering, and other application
metadata may be normalized separately. A source may contain an addressable verse
with an empty text payload; the importer preserves both the reference and the
empty payload.

## Importing WEBP

Download the official archive linked by the manifest, then run:

```console
go run ./cmd/bible-import \
  --manifest data/translations/engwebp/manifest.json \
  --archive /path/to/engwebp_vpl.zip \
  --output data/generated/engwebp.db
```

The command refuses to overwrite an existing database. Before parsing any
content it verifies the complete archive against the manifest checksum. It then
validates book order, contiguous references, duplicate references, canonical
boundaries, and expected totals before atomically publishing the SQLite file.
The generated database also contains an FTS5 token index over the unchanged
verse text. The index is rebuilt in canonical verse order so repeated imports of
the same pinned archive remain byte-for-byte reproducible.
