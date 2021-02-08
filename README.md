# trading212

Provides access to the undocumented Trading 212 website API.

Work in progress.

## Issues

You can only have one session with Trading 212 at a time, e.g if you open the mobile app
while on the website it will kick  you off the website. Same with anything using this library,
logging in on another platform will end your session.

This isn't ideal as I was going to implement a trailing stop loss but would be unable to run
my script and use the website at the same time.

Sharing this package anyway to since most projects seem to be using Selenium and may
give someone a head start.