# Docker Secrets Simulator

When running locally for development, it's hard to use secrets as they require Docker Swarm. And, it's often hard to use Docker Swarm during development as you frequently want to use a `docker-compose.yml` with `build` directives, volume mounts, etc.  So, how do you get in the practice of using secrets?  Well, simulate them!

There are a few ways to simulate secrets:

- Mount secrets to `/run/secrets` from the host (which requires you to define the files somewhere)
- Use this image!

The "Secrets Simulator" will take all defined environment variables and create secrets files where the variable name is the filename and the variable value is the secret content. This allows you to keep everything within your docker-compose.yml!

## Example

The following `docker-compose.yml` file is in the repo, so you can give it a try. I only add a service using the `mikesir87/secrets-simulator` image. Any declared environment variables are then converted into files in the `/run/secrets` directory. Since that directory is a volume, I can then mount it to other locations that need secrets!

```yaml
version: '3.1'

services:
  secrets-simulator:
    image: mikesir87/secrets-simulator
    environment:
      DB_USERNAME: admin
      DB_PASSWORD: password1234!
    volumes:
      - secrets:/run/secrets:rw

  viewer:
    image: mikesir87/secrets-viewer
    volumes:
      - secrets:/run/secrets:ro

volumes:
  secrets:
    driver: local
```

## FAQ

**How can I pick different secrets for different apps?**

You can run as many simulator services as you wish! You'll just need to create a new volume for each distinct set of secrets.


**Why mount secrets as `ro`?**

To help simulate a "real" scenario, I bind the secrets as read-only in the app. That way, I can't make any accidental changes to them.



