# Concet

Concet is an open source note-taking application.

## Configuration

Configuration stored in `configs` folder in `config.json` file.

### Configuration Example

```json
{
  "port": 3000,
  "production": true,
  "db_config": {
    "host": "example.com",
    "port": "5432",
    "username": "name",
    "db_name": "password",
    "ssl_mode": "disable"
  }
}
```

### Environment variables

Concet also uses environment variables for configuration.

List of all environment variables:

```env
DB_PASSWORD
PASSWORD_SALT
JWT_SECRET
```

### Development

To start a local postgres database run:

```bash
sudo docker run -e POSTGRES_PASSWORD=password -p 5432:5432 postgres
```
