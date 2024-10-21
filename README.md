
# CDC Data Exchange (DEX) CSV/TSV Structural Validator and Transformer

## Overview
The CSV/TSV Validator and Transformer is a high-performance tool designed to validate and transform large CSV/TSV files in near real-time. Its currently in the proof-of-concent phase, and this repository reflects the initial development efforts.

## WHY
The intent behind providing these two features in DEX is to ensure CSV files are parseable by downstream big-data processing systems elsewhere in the CDC, such as those built on top of DataBricks. By checking structural validation of CSV files, DEX can rapidly (within seconds in many cases) respond to a sending organization with CSV validation errors and warnings. Doing so gives that sender rapid feedback about why their CSV file isn't parseable. Faster resolution of certain data quality issues related to the structure of the files being sent is ideal, as opposed to waiting hours or days for a batch processing job.


## Key Features

- **CSV/TSV validation** - The validator, which checks the files for their adherence to the rules set forth in [RFC 4180](https://www.rfc-editor.org/rfc/rfc4180). The CSV validator does not validate field content, header fields or data types at this point. Each validated row is assigned:
    - **File UUID**: Each processed file is assigned a unique `uuid` that is used to link rows to its originating file. This will enable downstream consumers to track back rows to its source file.
    - **Row UUID**: Each row is assigned unique `uuid` during validation. This ensure that every row can be uniquely referenced.
    - **SHA-256 Row Hashing**: The content of each row is hashed using  `SHA-256` algorithm to provide ability of detecting duplicates and ensure data entegrity throughout the data pipeline. 

- **CSV/TSV transformer** - The  transformer, which generates a JSON object for each row in the file. If the optional header is present, then the header specification is used to populate the JSON property names.

- **Encoding Detection** - If encoding is not specified in the optional command line argument `config.json` , the tool will sample the file for upto 1024 bytes, and attempt to `auto-detect` the file's encoding(with best effor accuracy). Suppported encodings are: 
    - `UTF-8`
    - `UTF-8 with BOM`
    - `USASCII`
    - `ISO-8859-1`
    - `Windows1252`

- **Delimiter Detection** - If the delimiter is not provided in optional `config.json` command file flag, the tool will sample the file upto 1024 bytes of data, and will attempt to `auto-detect` delimiter. Supported delimiters include: 
    - `,` (comma)
    - `\t` (tab) 

## Installation
1. Clone the repository:
    ```bash
    git clone https://github.com/CDCgov/data-exchange-csv.git
    cd data-exchange-csv
    go build
    ```

## Usage
The DEX CSV Validator and Transformer accepts following command-line flags:
- `-fileURL:` [Required] The path to the file that will be validated.
- `-destination:` [Required] The path to the folder where validation/transformation results will be stored.
- `-debug:` [Optional] If true, `debug-level` logs will be generated.
- `-log-file:` [Optional] If true, logs will be written to logs/validation.json, default is stdout.
- `-transform:` [Optional] If true, the valid rows will be transformed to a `JSON` object.
- `-config:` [Optional] The path to the `config.json`. If provided overrides auto-detection of encoding, and separator.
    
    ```json
    {
    "encoding": "UTF-8",
    "separator": ",",
    "hasHeader": true
    }
      
## Future Enhancements
- **Non-blocking Validation/transformation**: Currently, the validation process is performed synchronously, which may introduce delays when processing large files. To address this, we are exploring the use of Go routines to parallelize the validation and transformation process. By leveraging concurrency, we aim to significantly improve performance and reduce processing time.

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
## Additional Standard Notices
Please refer to [CDC's Template Repository](https://github.com/CDCgov/template)
for more information about [contributing to this repository](https://github.com/CDCgov/template/blob/master/CONTRIBUTING.md),
[public domain notices and disclaimers](https://github.com/CDCgov/template/blob/master/DISCLAIMER.md),
and [code of conduct](https://github.com/CDCgov/template/blob/master/code-of-conduct.md).
