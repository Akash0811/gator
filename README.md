# GATOR - A blog aggregator CLI App

### Features:
- Add RSS feeds from across the internet to be collected
- Follow and unfollow RSS feeds that other users have added
- View summaries of the aggregated posts in the terminal, with a link to the full post

### Requisites
- Go (1.26+)
- Postgresql (15+)

### Installation
```
go install github.com/Akash0811/gator@latest
```
After installation, make sure to run all the up migrations
```
cd repo/sql/schema
goose postgres <connection string> up
```

### Set up configuration
Setup your database url (db_url) preferably in your home directory at the location `~/.gatorconfig.json`
An example is 
```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable"
}
```

### Cli commands supported

| command | description | arguments | example |
| --- | --- | --- | --- |
| register | regsiters a user | 1 | gator register username |
| login | logs a registered user in | 1 | gator login username |
| users | lists all users | 0 | gator users |
| addfeed | Add RSS feed | 2 | gator addfeed feed_name url |
| feeds | List RSS feeds | 0 | gator feeds |
| agg | Aggregate content for RSS feed after a time interval | 1 | gator agg 5s |
| browse | Browse content | 1 or 0 | gator browse 10 |
| follow | Follow a feed | 1 | gator follow url |
| unfollow | Unfollow a feed | 1 | gator unfollow url |
| following | List feeds that are being followed | 0 | gator following |