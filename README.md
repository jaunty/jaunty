# Jaunty

The web interface for Jaunty, a Minecraft server community.

# Configuration
Jaunty uses environment variables for configuration by default.

```
JAUNTY_WEB_ADDR=:5000
JAUNTY_SESSION_KEY=somethingsecret
JAUNTY_DISCORD_PUBLIC_KEY=yourapplicationspublickey
JAUNTY_DB_DSN=postgres://youraddress
JAUNTY_RCON_ADDR=someaddress:25575
JAUNTY_GUILD_ID=1234567890193745
JAUNTY_WHITELIST_CHANNEL_ID=123456789098837658
JAUNTY_UNAPPROVED_ROLE_ID=12345627456437658
JAUNTY_MAX_REQUESTS=2
JAUNTY_BOT_TOKEN=longtoken.here.
JAUNTY_OAUTH2_CLIENT_ID=yourappsclientid
JAUNTY_OAUTH2_CLIENT_SECRET=yourappssecret
JAUNTY_OAUTH2_SCOPES=identify,guilds.join
JAUNTY_OAUTH2_REDIRECT_URI=http://youraddress.gay/callback
```

# Requirements
- PostgreSQL
- Go
- A Discord
- - OAuth2 application
- - Server
