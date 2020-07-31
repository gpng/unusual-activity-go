# Unusual Trading Activity for SGX (Go)

> Alerts for unusually high trading volume on SGX

Scans all tickers on SGX, gets their last x months of daily volume, and alerts if any of the last n days of volume exceeds s standard deviations from the mean

## Table of Contents

- [Usage](#usage)
- [Maintainers](#maintainers)
- [License](#license)

## Usage

1. Clone repo

2. `cd build`

3. `./unusual_activity -h` or `unusual_activity.exe -h` for help text on cli arguments

4. `./unusual_activity` or `unusual_acitivity.exe` to run with defaults

## Maintainers

[@gpng](https://github.com/gpng)

## License

MIT © 2020 Gerald Png
