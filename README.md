
# CDC Data Exchange (DEX) CSV Structural Validator

## Overview
This is a code repository for a validator and transformer specific to character-separated values (CSV) files. The intent behind DEX's CSV product is to provide a high-speed, ultra-scalable CSV validator that can be run in near real-time even at high load and with huge file sizes.

The two key functions of the CSV product are:
- The CSV validator, which checks CSV files for their adherence to the rules set forth in [RFC 4180](https://www.rfc-editor.org/rfc/rfc4180) as well as against an optional header specification. If a optional header specification is provided, the CSV validator will ensure the header row in the CSV matches the header specification. The CSV validator does not validate field content or data types.
- The CSV transformer, which generates a JSON file for each record in the CSV file. If the optional header specification is present, then the header specification is used to populate the JSON property names.

> The DEX CSV product is under development in a 'proof-of-concept' phase only. It is subject to significant changes and updates.

The intent behind providing these two features in DEX is to ensure CSV files are parseable by downstream big-data processing systems elsewhere in the CDC, such as those built on top of DataBricks. By checking structural validation of CSV files, DEX can rapidly (within seconds in many cases) respond to a sending organization with CSV validation errors and warnings. Doing so gives that sender rapid feedback about why their CSV file isn't parseable. Faster resolution of certain data quality issues related to the structure of the files being sent is ideal, as opposed to waiting hours or days for a batch processing job.

Transforming CSV records into JSON files is an optional feature intended for cases where a consumer of the DEX CSV product would rather work with individual JSON records than CSV files.

Each transformed record is assigned a UUIDv4 and is hashed. The intent behind hashing each record is to aid downstream CDC data systems in detecting duplicate records. The file hash and UUIDv4, among other properties, are included in the file metadata when it's written to disk.

The CSV product, when built, will be fully integrated with the rest of DEX and will be optional for any teams wishing to process CSV files through DEX. More information will be forthcoming as the CSV products moves from proof-of-concept into the Alpha and Beta phases of product maturity.

## Character encoding
The CSV code modules will optionally accept one of several character encodings, such as `UTF-8` or `ISO-8859-1`, and can auto-detect character encodings (albeit with no guarantee of 100% accuracy) if one is not provided by the calling function. 

## Public Domain Standard Notice
This repository constitutes a work of the United States Government and is not
subject to domestic copyright protection under 17 USC § 105. This repository is in
the public domain within the United States, and copyright and related rights in
the work worldwide are waived through the [CC0 1.0 Universal public domain dedication](https://creativecommons.org/publicdomain/zero/1.0/).
All contributions to this repository will be released under the CC0 dedication. By
submitting a pull request you are agreeing to comply with this waiver of
copyright interest.

## License Standard Notice
The repository utilizes code licensed under the terms of the Apache Software
License and therefore is licensed under ASL v2 or later.

This source code in this repository is free: you can redistribute it and/or modify it under
the terms of the Apache Software License version 2, or (at your option) any
later version.

This source code in this repository is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
PARTICULAR PURPOSE. See the Apache Software License for more details.

You should have received a copy of the Apache Software License along with this
program. If not, see http://www.apache.org/licenses/LICENSE-2.0.html

The source code forked from other open source projects will inherit its license.

## Privacy Standard Notice
This repository contains only non-sensitive, publicly available data and
information. All material and community participation is covered by the
[Disclaimer](https://github.com/CDCgov/template/blob/master/DISCLAIMER.md)
and [Code of Conduct](https://github.com/CDCgov/template/blob/master/code-of-conduct.md).
For more information about CDC's privacy policy, please visit [http://www.cdc.gov/other/privacy.html](https://www.cdc.gov/other/privacy.html).

## Contributing Standard Notice
Anyone is encouraged to contribute to the repository by [forking](https://help.github.com/articles/fork-a-repo)
and submitting a pull request. (If you are new to GitHub, you might start with a
[basic tutorial](https://help.github.com/articles/set-up-git).) By contributing
to this project, you grant a world-wide, royalty-free, perpetual, irrevocable,
non-exclusive, transferable license to all users under the terms of the
[Apache Software License v2](http://www.apache.org/licenses/LICENSE-2.0.html) or
later.

All comments, messages, pull requests, and other submissions received through
CDC including this GitHub page may be subject to applicable federal law, including but not limited to the Federal Records Act, and may be archived. Learn more at [http://www.cdc.gov/other/privacy.html](http://www.cdc.gov/other/privacy.html).

## Records Management Standard Notice
This repository is not a source of government records, but is a copy to increase
collaboration and collaborative potential. All government records will be
published through the [CDC web site](http://www.cdc.gov).

## Related documents

* [Open Practices](open_practices.md)
* [Rules of Behavior](rules_of_behavior.md)
* [Thanks and Acknowledgements](thanks.md)
* [Disclaimer](DISCLAIMER.md)
* [Contribution Notice](CONTRIBUTING.md)
* [Code of Conduct](code-of-conduct.md)
ls -
## Additional Standard Notices
Please refer to [CDC's Template Repository](https://github.com/CDCgov/template)
for more information about [contributing to this repository](https://github.com/CDCgov/template/blob/master/CONTRIBUTING.md),
[public domain notices and disclaimers](https://github.com/CDCgov/template/blob/master/DISCLAIMER.md),
and [code of conduct](https://github.com/CDCgov/template/blob/master/code-of-conduct.md).
