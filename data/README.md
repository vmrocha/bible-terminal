# Translation data

This directory contains source manifests and, later, deterministic import
tooling. A manifest records what may be downloaded and imported; it does not
contain Bible text.

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
metadata may be normalized separately.
